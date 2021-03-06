// Package main implements a server for Greeter service.
package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"errors"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	pb "github.com/Fralkayg/sd-t1/Service"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	seguimiento         int
	lock                bool
	colaRetail          []paquete
	colaPrioritario     []paquete
	colaNormal          []paquete
	seguimientoPaquetes []SeguimientoPaquete
}

type SeguimientoPaquete struct {
	IDPaquete     string
	Estado        string
	IDCamion      int
	IDSeguimiento int
	Intentos      int
	Tipo          string
	Valor         int
}

type paquete struct {
	IDPaquete   string
	Seguimiento int
	Tipo        string
	Valor       int
	Intentos    int
	Estado      string
	Origen      string
	Destino     string
}

type Finanzas struct {
	IDPaquete string
	Tipo      string
	Valor     int
	Intentos  int
	Estado    string
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

//Función: failOnError
//Descripción: Función necesaria para el funcionamiento de RabbitMQ. Importada desde https://www.rabbitmq.com/tutorials/tutorial-one-go.html
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

//Funcion: ActualizarSeguimiento
//Descripción: Actualiza el registro en memoria de seguimiento de paquetes. Recibido - No recibido - Intentos realizados.
//Parametros:
//context: --
//UpdateSeguimiento: Estructura de datos que entrega informacion del estado de un paquete.
//Retorno:
//StatusSeguimiento: Contiene un string con el mensaje del resultado de la operación
//error: Mensaje de error en caso de fallar algo
func (s *server) ActualizarSeguimiento(ctx context.Context, updateSeguimiento *pb.UpdateSeguimiento) (*pb.StatusSeguimiento, error) {
	for s.lock {
	}

	s.lock = true

	// fmt.Println("Seguimiento a actualizar: ", updateSeguimiento.Seguimiento)

	index, _, err := Find(s.seguimientoPaquetes, int(updateSeguimiento.Seguimiento))

	if err != nil {
		log.Printf("Hubo un error al actualizar el paquete solicitado.")
		s.lock = false
		return &pb.StatusSeguimiento{Mensaje: "Error al actualizar estado del paquete"}, errors.New("Error al actualizar estado del paquete")
	}

	if updateSeguimiento.Entregado {
		s.seguimientoPaquetes[index].Estado = "Recibido"
		log.Printf("El paquete %v ha sido recibido.", s.seguimientoPaquetes[index].IDPaquete)
	} else {
		s.seguimientoPaquetes[index].Estado = "No recibido"
		log.Printf("El paquete %v no ha sido recibido.", s.seguimientoPaquetes[index].IDPaquete)
	}

	s.seguimientoPaquetes[index].Intentos = int(updateSeguimiento.Intentos)

	// fmt.Println("Seguimiento: ", s.seguimientoPaquetes)

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

	finanzas := Finanzas{
		IDPaquete: s.seguimientoPaquetes[index].IDPaquete,
		Tipo:      s.seguimientoPaquetes[index].Tipo,
		Valor:     s.seguimientoPaquetes[index].Valor,
		Intentos:  s.seguimientoPaquetes[index].Intentos,
		Estado:    s.seguimientoPaquetes[index].Estado}

	b, _ := json.Marshal(finanzas)

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(b),
		})
	failOnError(err, "Failed to publish a message")

	s.lock = false

	return &pb.StatusSeguimiento{Mensaje: "Paquete actualizado correctamente"}, nil
}

//Funcion: SolicitarSeguimiento
//Descripción: Entrega el estado de un paquete solicitado por cliente
//Parametros:
//context: --
//SeguimientoPyme: Estructura que contiene solo un entero representado el ID de seguimiento de un paquete (tanto para retail como para PYME)
//Retorno:
//SeguimientoPaqueteSolicitado: Estructura que contiene la información necesaria del paquete solicitado.
//error: Mensaje de error en caso de fallar algo
func (s *server) SolicitarSeguimiento(ctx context.Context, seguimientoPyme *pb.SeguimientoPyme) (*pb.SeguimientoPaqueteSolicitado, error) {
	for s.lock {
	}

	s.lock = true

	_, paquete, err := Find(s.seguimientoPaquetes, int(seguimientoPyme.Id))

	if err != nil {
		log.Printf("El paquete solicitado no se encuentra en la lista de seguimiento de paquetes.")
		s.lock = false
		return &pb.SeguimientoPaqueteSolicitado{}, errors.New("El paquete solicitado no se encuentra en la lista de seguimiento de paquetes.")
	}

	s.lock = false

	return &pb.SeguimientoPaqueteSolicitado{IDPaquete: paquete.IDPaquete, Estado: paquete.Estado}, nil
}

//Funcion: GenerarOrdenPyme
//Descripción: Genera registro en memoria de un pedido de PYME realizado por cliente.
//Parametros:
//context: --
//OrdenPyme: Estructura de datos que contiene la información de la orden realizada.
//Retorno:
//SeguimientoPyme: Estructura que contiene un entero con el ID de seguimiento de la orden generada.
//error: Mensaje de error en caso de fallar algo
func (s *server) GenerarOrdenPyme(ctx context.Context, ordenPyme *pb.OrdenPyme) (*pb.SeguimientoPyme, error) {
	//Chequear si el servidor esta ocupado en otro requerimiento
	for s.lock {
	}

	s.lock = true

	log.Printf("Se recibio orden PYME. Id orden: %v", ordenPyme.GetId())
	s.seguimiento++

	registroOrdenPyme(ordenPyme, s.seguimiento)

	// s.seguimientoPaquetes = append(s.seguimientoPaquetes, SeguimientoPaquete{
	// 	IDPaquete:     ordenPyme.GetId(),
	// 	Estado:        "En bodega",
	// 	IDCamion:      0,
	// 	IDSeguimiento: s.seguimiento,
	// 	Intentos:      0})

	if ordenPyme.GetPrioritario() == 1 {
		s.colaPrioritario = enqueue(s.colaPrioritario, paquete{IDPaquete: ordenPyme.GetId(),
			Seguimiento: s.seguimiento,
			Tipo:        "Prioritario",
			Valor:       int(ordenPyme.GetValor()),
			Intentos:    0,
			Origen:      ordenPyme.GetOrigen(),
			Destino:     ordenPyme.GetDestino(),
			Estado:      "En bodega"})

		s.seguimientoPaquetes = append(s.seguimientoPaquetes, SeguimientoPaquete{
			IDPaquete:     ordenPyme.GetId(),
			Estado:        "En bodega",
			IDCamion:      0,
			IDSeguimiento: s.seguimiento,
			Tipo:          "Prioritario",
			Valor:         int(ordenPyme.GetValor()),
			Intentos:      0})
	} else {
		s.colaNormal = enqueue(s.colaNormal, paquete{IDPaquete: ordenPyme.GetId(),
			Seguimiento: s.seguimiento,
			Tipo:        "Normal",
			Valor:       int(ordenPyme.GetValor()),
			Intentos:    0,
			Origen:      ordenPyme.GetOrigen(),
			Destino:     ordenPyme.GetDestino(),
			Estado:      "En bodega"})

		s.seguimientoPaquetes = append(s.seguimientoPaquetes, SeguimientoPaquete{
			IDPaquete:     ordenPyme.GetId(),
			Estado:        "En bodega",
			IDCamion:      0,
			IDSeguimiento: s.seguimiento,
			Tipo:          "Normal",
			Valor:         int(ordenPyme.GetValor()),
			Intentos:      0})
	}
	// fmt.Println("Cola prioritario: ", s.colaPrioritario)
	// fmt.Println("Cola normal: ", s.colaNormal)
	// log.Printf("Aqui deberia estar generandose la orden de Pyme")
	// fmt.Println("Seguimiento orden: ", s.seguimientoPaquetes)

	s.lock = false
	return &pb.SeguimientoPyme{Id: int32(s.seguimiento)}, nil
}

//Funcion: GenerarOrdenRetail
//Descripción: Genera registro en memoria de un pedido de retail realizado por cliente.
//Parametros:
//context: --
//OrdenRetail: Estructura de datos que contiene la información de la orden realizada.
//Retorno:
//SeguimientoRetail: Estructura que contiene un entero con el ID de seguimiento de la orden generada.
//error: Mensaje de error en caso de fallar algo
func (s *server) GenerarOrdenRetail(ctx context.Context, ordenRetail *pb.OrdenRetail) (*pb.SeguimientoRetail, error) {
	for s.lock {
	}

	s.lock = true

	log.Printf("Se recibio orden Retail. Id orden: %v", ordenRetail.GetId())
	s.seguimiento++

	registroOrdenRetail(ordenRetail, s.seguimiento)

	s.seguimientoPaquetes = append(s.seguimientoPaquetes, SeguimientoPaquete{
		IDPaquete:     ordenRetail.GetId(),
		Estado:        "En bodega",
		IDCamion:      0,
		IDSeguimiento: s.seguimiento,
		Tipo:          "Retail",
		Valor:         int(ordenRetail.GetValor()),
		Intentos:      0})

	s.colaRetail = enqueue(s.colaRetail, paquete{IDPaquete: ordenRetail.GetId(),
		Seguimiento: s.seguimiento,
		Tipo:        "Retail",
		Valor:       int(ordenRetail.GetValor()),
		Intentos:    0,
		Origen:      ordenRetail.GetOrigen(),
		Destino:     ordenRetail.GetDestino(),
		Estado:      "En bodega"})

	// log.Printf("Aqui deberia estar generandose la orden Retail")
	// fmt.Println("Cola retail: ", s.colaRetail)
	// fmt.Println("Seguimiento orden: ", s.seguimientoPaquetes)

	s.lock = false
	return &pb.SeguimientoRetail{Id: int32(s.seguimiento)}, nil
}

//Funcion: registroOrdenRetail
//Descripción: Genera registro en un archivo CSV con la información de los pedidos generados.
//Parametros:
//context: --
//OrdenRetail: Estructura de datos que contiene la información de la orden realizada.
//idSeguimiento: ID seguimiento de orden.
//Retorno:
//--
func registroOrdenRetail(ordenRetail *pb.OrdenRetail, idSeguimiento int) {
	seguimientoFile, err := os.OpenFile("./registro.csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		seguimientoAux, errAux := os.Create("./registro.csv")
		if errAux != nil {
			log.Printf("wea")
		}
		seguimientoFile = seguimientoAux
		log.Printf("Hubo un error al abrir/crear archivo seguimiento. Tipo: Retail")
	}

	defer seguimientoFile.Close()

	timestamp := time.Now()
	timeString := timestamp.Format("2006-01-02 15:04:05")

	var fileData [][]string

	// log.Printf("Generando linea en archivo registro.csv, Retail")

	fileData = append(fileData, []string{timeString,
		ordenRetail.GetId(),
		"retail",
		ordenRetail.GetProducto(),
		strconv.Itoa(int(ordenRetail.GetValor())),
		ordenRetail.GetOrigen(),
		ordenRetail.GetDestino(),
		strconv.Itoa(idSeguimiento)})

	csvWriter := csv.NewWriter(seguimientoFile)
	csvWriter.WriteAll(fileData)
	// csvWriter.Flush()
}

//Funcion: registroOrdenPyme
//Descripción: Genera registro en un archivo CSV con la información de los pedidos generados.
//Parametros:
//context: --
//OrdenPyme: Estructura de datos que contiene la información de la orden realizada.
//idSeguimiento: Estructura que contiene un entero con el ID de seguimiento de la orden generada.
//Retorno:
//--
func registroOrdenPyme(ordenPyme *pb.OrdenPyme, idSeguimiento int) {
	seguimientoFile, err := os.OpenFile("./registro.csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		seguimientoAux, errAux := os.Create("./registro.csv")
		seguimientoFile = seguimientoAux
		if errAux != nil {
			log.Printf("wea2")
		}
		log.Printf("Hubo un error al abrir/crear archivo seguimiento. Tipo: Retail")
	}

	defer seguimientoFile.Close()

	timestamp := time.Now()

	var tipoPyme string
	if ordenPyme.GetPrioritario() == 1 {
		tipoPyme = "prioritario"
	} else {
		tipoPyme = "normal"
	}

	// log.Printf("Generando linea en archivo registro.csv, PYME tipo %v", tipoPyme)

	var fileData [][]string
	fileData = append(fileData, []string{timestamp.Format("2006-01-02 15:04:05"),
		ordenPyme.GetId(),
		tipoPyme,
		ordenPyme.GetProducto(),
		strconv.Itoa(int(ordenPyme.GetValor())),
		ordenPyme.GetOrigen(),
		ordenPyme.GetDestino(),
		strconv.Itoa(idSeguimiento)})

	csvWriter := csv.NewWriter(seguimientoFile)
	csvWriter.WriteAll(fileData)
}

//Función extraida de https://www.educative.io/edpresso/how-to-implement-a-queue-in-golang
//Modificada para el funcionamiento con la estructura necesaria
func enqueue(queue []paquete, element paquete) []paquete {
	queue = append(queue, element) // Simply append to enqueue.
	// fmt.Println("Enqueued:", element)
	return queue
}

//Función extraida de https://www.educative.io/edpresso/how-to-implement-a-queue-in-golang
//Modificada para el funcionamiento con la estructura necesaria
func dequeue(queue []paquete) ([]paquete, paquete) {
	element := queue[0] // The first element is the one to be dequeued.
	// fmt.Println("Dequeued:", element)
	return queue[1:], element // Slice off the element once it is dequeued.
}

//Funcion: SolicitarPaquete
//Descripción: Se encarga de entregar los paquetes a cada camión dependiendo del tipo de camión y de los paquetes que se tienen en las colas.
//Parametros:
//context: --
//Camion: Estructura de datos que contiene la información de los camiones.
//Retorno:
//PaqueteCamion: Estructura que contiene la información del paquete que se enviara al camión.
//error: Mensaje de error en caso de fallar algo
func (s *server) SolicitarPaquete(ctx context.Context, camion *pb.Camion) (*pb.PaqueteCamion, error) {
	if camion.GetTipo() == "Retail" {
		if len(s.colaRetail) > 0 {
			colaRetail, paquete := dequeue(s.colaRetail)
			s.colaRetail = colaRetail

			if paquete.IDPaquete != "" {
				paqueteCamion := &pb.PaqueteCamion{
					Id:          paquete.IDPaquete,
					Tipo:        paquete.Tipo,
					Origen:      paquete.Origen,
					Destino:     paquete.Destino,
					Valor:       int32(paquete.Valor),
					Seguimiento: int32(paquete.Seguimiento),
				}

				index, _, err := Find(s.seguimientoPaquetes, paquete.Seguimiento)

				if err != nil {
					log.Printf("El paquete solicitado no se encuentra en la lista de seguimiento de paquetes.")
				} else {
					log.Printf("Se cambio el estado del paquete %v a En camino", paqueteCamion.Id)
					s.seguimientoPaquetes[index].Estado = "En camino"
					s.seguimientoPaquetes[index].IDCamion = int(camion.Id)
				}

				// s.seguimientoPaquetes = append(s.seguimientoPaquetes, SeguimientoPaquete{
				// 	IDPaquete: paquete.IDPaquete,
				// 	Estado: "En camino",
				// 	IDCamion: camion.Id,
				// 	IDSeguimiento: paquete.Seguimiento,
				// 	Intentos: 0})

				// fmt.Println("Cola retail: ", s.colaRetail)
				// fmt.Println("Seguimiento: ", s.seguimientoPaquetes)

				return paqueteCamion, nil
			}
		} else {
			if camion.GetEntregaRetail() {
				if len(s.colaPrioritario) > 0 {
					colaPrioritario, paquete := dequeue(s.colaPrioritario)
					s.colaPrioritario = colaPrioritario

					if paquete.IDPaquete != "" {
						paqueteCamion := &pb.PaqueteCamion{
							Id:          paquete.IDPaquete,
							Tipo:        paquete.Tipo,
							Origen:      paquete.Origen,
							Destino:     paquete.Destino,
							Valor:       int32(paquete.Valor),
							Seguimiento: int32(paquete.Seguimiento),
						}

						index, _, err := Find(s.seguimientoPaquetes, paquete.Seguimiento)

						if err != nil {
							log.Printf("El paquete solicitado no se encuentra en la lista de seguimiento de paquetes.")
						} else {
							log.Printf("Se cambio el estado del paquete %v a En camino", paqueteCamion.Id)
							s.seguimientoPaquetes[index].Estado = "En camino"
							s.seguimientoPaquetes[index].IDCamion = int(camion.Id)
						}

						// s.seguimientoPaquetes = append(s.seguimientoPaquetes, SeguimientoPaquete{
						// 	IDPaquete: paquete.IDPaquete,
						// 	Estado: "En camino",
						// 	IDCamion: camion.Id,
						// 	IDSeguimiento: paquete.Seguimiento,
						// 	Intentos: 0})

						// fmt.Println("Cola prioritario: ", s.colaPrioritario)
						// fmt.Println("Seguimiento: ", s.seguimientoPaquetes)

						return paqueteCamion, nil
					}
				}
			}
		}
	} else if camion.GetTipo() == "Normal" {
		if len(s.colaPrioritario) > 0 {
			colaPrioritario, paquete := dequeue(s.colaPrioritario)
			s.colaPrioritario = colaPrioritario

			if paquete.IDPaquete != "" {
				paqueteCamion := &pb.PaqueteCamion{
					Id:          paquete.IDPaquete,
					Tipo:        paquete.Tipo,
					Origen:      paquete.Origen,
					Destino:     paquete.Destino,
					Valor:       int32(paquete.Valor),
					Seguimiento: int32(paquete.Seguimiento),
				}

				index, _, err := Find(s.seguimientoPaquetes, paquete.Seguimiento)

				if err != nil {
					log.Printf("El paquete solicitado no se encuentra en la lista de seguimiento de paquetes.")
				} else {
					log.Printf("Se cambio el estado del paquete %v a En camino", paqueteCamion.Id)
					s.seguimientoPaquetes[index].Estado = "En camino"
					s.seguimientoPaquetes[index].IDCamion = int(camion.Id)
				}

				// s.seguimientoPaquetes = append(s.seguimientoPaquetes, SeguimientoPaquete{
				// 	IDPaquete: paquete.IDPaquete,
				// 	Estado: "En camino",
				// 	IDCamion: camion.Id,
				// 	IDSeguimiento: paquete.Seguimiento,
				// 	Intentos: 0})

				// fmt.Println("Cola prioritario: ", s.colaPrioritario)
				// fmt.Println("Seguimiento: ", s.seguimientoPaquetes)

				return paqueteCamion, nil

			}

		} else if len(s.colaNormal) > 0 {

			colaNormal, paquete := dequeue(s.colaNormal)
			s.colaNormal = colaNormal

			if paquete.IDPaquete != "" {
				paqueteCamion := &pb.PaqueteCamion{
					Id:          paquete.IDPaquete,
					Tipo:        paquete.Tipo,
					Origen:      paquete.Origen,
					Destino:     paquete.Destino,
					Valor:       int32(paquete.Valor),
					Seguimiento: int32(paquete.Seguimiento),
				}

				index, _, err := Find(s.seguimientoPaquetes, paquete.Seguimiento)

				if err != nil {
					log.Printf("El paquete solicitado no se encuentra en la lista de seguimiento de paquetes.")
				} else {
					log.Printf("Se cambio el estado del paquete %v a En camino", paqueteCamion.Id)
					s.seguimientoPaquetes[index].Estado = "En camino"
					s.seguimientoPaquetes[index].IDCamion = int(camion.Id)
				}

				// s.seguimientoPaquetes = append(s.seguimientoPaquetes, SeguimientoPaquete{
				// 	IDPaquete: paquete.IDPaquete,
				// 	Estado: "En camino",
				// 	IDCamion: camion.Id,
				// 	IDSeguimiento: paquete.Seguimiento,
				// 	Intentos: 0})

				// fmt.Println("Cola normal: ", s.colaNormal)
				// fmt.Println("Seguimiento: ", s.seguimientoPaquetes)

				return paqueteCamion, nil
			}
		}
	}
	// fmt.Println("Seguimiento: ", s.seguimientoPaquetes)
	return &pb.PaqueteCamion{}, errors.New("Error al entregar paquete")
}

//Función extraida de https://yourbasic.org/golang/find-search-contains-slice/
//Modificada para el funcionamiento de las estructura de seguimientoPaquete
func Find(seguimientoPaquetes []SeguimientoPaquete, idSeguimiento int) (int, SeguimientoPaquete, error) {
	for i, element := range seguimientoPaquetes {
		if idSeguimiento == element.IDSeguimiento {
			return i, element, nil
		}
	}
	return len(seguimientoPaquetes), SeguimientoPaquete{}, errors.New("El paquete solicitado no se encuentra para seguimiento")
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	s := server{}

	//Inicializacion variables servidor logistica.
	s.seguimiento = 0
	s.lock = false

	pb.RegisterLogisticaServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
