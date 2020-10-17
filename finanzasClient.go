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

type Finanzas struct {
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

func ingresoPaquete(paquete infoPaquete) float32 {
	var ingresos float32
	ingresos = 0
	if paquete.Estado == "Recibido" {
		ingresos += float32(paquete.Valor)
		ingresos += convenioPYME(paquete)
		ingresos -= float32((paquete.Intentos - 1) * 10)
	} else {
		if paquete.Tipo == "Normal" {
			ingresos += 0
			ingresos -= float32((paquete.Intentos - 1) * 10)
		} else if paquete.Tipo == "Prioritario" {
			ingresos += convenioPYME(paquete)
			ingresos -= float32((paquete.Intentos - 1) * 10)
		} else {
			ingresos += float32(paquete.Valor)
			ingresos -= float32((paquete.Intentos - 1) * 10)
		}
	}
	return ingresos
}

func registrarFinanza(paquete infoPaquete) {
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
	ingresos := ingresoPaquete(paquete)
	fileData = append(fileData, []string{
		paquete.IDPaquete,
		paquete.Tipo,
		//strconv.Itoa(int(paquete.Valor)),
		//paquete.Origen,
		//paquete.Destino,
		strconv.Itoa(int(paquete.Intentos)),
		paquete.Estado,
		fmt.Sprintf("%f", ingresos)})

	csvWriter := csv.NewWriter(finanzaFile)
	csvWriter.WriteAll(fileData)
	// csvWriter.Flush()
}

// func main(){
// 	//conn con rabbit xd
// 	// usar colas rabbitMQ
// 	// La representacion de los datos enviados a la cola de mensajes del sistema financiero debe ser mediante JSON.
// 	// se obtiene un paquetito
// 	var balance float32

// 	for{
// 		ingresos = ingresoPaquete(paquetito)
// 		registrarFinanza(paquetito)

// 		balance += ingresos
// 		log.Printf("Balance: %f dignipesos",balance)
// 	}

// }

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
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
			var message Finanzas
			json.Unmarshal(d.body, &message)
			fmt.Println(message)
			// log.Printf("Received a message: %s", d.Body)

			// ingresos = ingresoPaquete(paquetito)
			// registrarFinanza(paquetito)

			// balance += ingresos
			// log.Printf("Balance: %f dignipesos",balance)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
