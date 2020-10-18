package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/streadway/amqp"
)

type infoPaquete struct {
	IDPaquete string
	Tipo      string
	Valor     int
	Intentos  int
	Estado    string
}

func convenioPYME(paquete infoPaquete) float32 {
	if paquete.Tipo == "Prioritario" {
		var ingreso float32
		ingreso = float32(paquete.Valor) * 0.3
		return ingreso
	} else {
		return 0
	}
}

func ingresoPaquete(paquete infoPaquete) (float32, float32) {
	var ingresos float32
	var perdidas float32
	ingresos = 0
	perdidas = 0
	if paquete.Estado == "Recibido" {
		ingresos += float32(paquete.Valor)
		ingresos += convenioPYME(paquete)
		perdidas -= float32((paquete.Intentos - 1) * 10)
		// ingresos -= float32((paquete.Intentos - 1) * 10)
	} else {
		if paquete.Tipo == "Normal" {
			ingresos += 0
			perdidas -= float32((paquete.Intentos - 1) * 10)
		} else if paquete.Tipo == "Prioritario" {
			ingresos += convenioPYME(paquete)
			perdidas -= float32((paquete.Intentos - 1) * 10)
		} else {
			ingresos += float32(paquete.Valor)
			perdidas -= float32((paquete.Intentos - 1) * 10)
		}
	}
	return ingresos, perdidas
}

func registrarFinanza(paquete infoPaquete, ingresos float32, perdidas float32) {
	finanzaFile, err := os.OpenFile("./registroFinanzas.csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		finanzaAux, errAux := os.Create("./registroFinanzas.csv")
		if errAux != nil {
			log.Printf("ola")
		}
		finanzaFile = finanzaAux
	}

	defer finanzaFile.Close()

	var fileData [][]string

	// log.Printf("Generando linea en archivo del camion xd")
	// ingresos := ingresoPaquete(paquete)
	fileData = append(fileData, []string{
		paquete.IDPaquete,
		paquete.Tipo,
		//strconv.Itoa(int(paquete.Valor)),
		//paquete.Origen,
		//paquete.Destino,
		strconv.Itoa(int(paquete.Intentos)),
		paquete.Estado,
		fmt.Sprintf("%f", ingresos),
		fmt.Sprintf("%f", perdidas)})

	csvWriter := csv.NewWriter(finanzaFile)
	csvWriter.WriteAll(fileData)
	// csvWriter.Flush()
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	var balance float32
	balance = 0

	conn, err := amqp.Dial("amqp://hahngoro:panconpalta@dist54:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			var paquete infoPaquete
			json.Unmarshal(d.Body, &paquete)
			// fmt.Println(paquete)
			// log.Printf("Received a message: %s", d.Body)

			ingresos, perdidas := ingresoPaquete(paquete)
			registrarFinanza(paquete, ingresos, perdidas)

			balance = balance + ingresos - perdidas
			log.Printf("Balance: %f dignipesos", balance)
		}
	}()

	log.Printf(" [*] Esperando la llegada de paquetes. Para salir presione las teclas CTRL+C")
	<-forever
}
