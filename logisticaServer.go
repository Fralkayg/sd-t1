// Package main implements a server for Greeter service.
package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	pb "github.com/Fralkayg/sd-t1/Service"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	seguimiento     int
	lock            bool
	colaRetail      []paquete
	colaPrioritario []paquete
	colaNormal      []paquete
}

type paquete struct {
	IDPaquete   string
	Seguimiento int
	Tipo        string
	Valor       int
	Intentos    int
	Estado      string
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *server) GenerarOrdenPyme(ctx context.Context, ordenPyme *pb.OrdenPyme) (*pb.SeguimientoPyme, error) {
	//Chequear si el servidor esta ocupado en otro requerimiento
	for s.lock {
	}

	s.lock = true

	log.Printf("Id orden: %v", ordenPyme.GetId())
	s.seguimiento++

	registroOrdenPyme(ordenPyme, s.seguimiento)
	if ordenPyme.GetPrioritario() == 1 {
		enqueue(s.colaPrioritario, paquete{IDPaquete: ordenPyme.GetId(),
			Seguimiento: s.seguimiento,
			Tipo:        "Prioritario",
			Valor:       int(ordenPyme.GetValor()),
			Intentos:    0,
			Estado:      "En bodega"})
	} else {
		enqueue(s.colaNormal, paquete{IDPaquete: ordenPyme.GetId(),
			Seguimiento: s.seguimiento,
			Tipo:        "Normal",
			Valor:       int(ordenPyme.GetValor()),
			Intentos:    0,
			Estado:      "En bodega"})
	}
	fmt.Println("Cola prioritario: ", s.colaPrioritario)
	fmt.Println("Cola normal: ", s.colaNormal)
	log.Printf("Aqui deberia estar generandose la orden de Pyme")

	s.lock = false
	return &pb.SeguimientoPyme{Id: int32(s.seguimiento)}, nil
}

func (s *server) GenerarOrdenRetail(ctx context.Context, ordenRetail *pb.OrdenRetail) (*pb.SeguimientoRetail, error) {
	for s.lock {
	}

	s.lock = true

	log.Printf("Id orden: %v", ordenRetail.GetId())
	s.seguimiento++

	registroOrdenRetail(ordenRetail, s.seguimiento)
	enqueue(s.colaRetail, paquete{IDPaquete: ordenRetail.GetId(),
		Seguimiento: s.seguimiento,
		Tipo:        "Retail",
		Valor:       int(ordenRetail.GetValor()),
		Intentos:    0,
		Estado:      "En bodega"})

	log.Printf("Aqui deberia estar generandose la orden Retail")
	fmt.Println("Cola retail: ", s.colaRetail)

	s.lock = false
	return &pb.SeguimientoRetail{Id: int32(s.seguimiento)}, nil
}

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
	timeString := timestamp.Format("2020-01-01 00:00")

	var fileData [][]string

	log.Printf("Generando linea en archivo registro.csv, Retail")

	fileData = append(fileData, []string{timeString,
		strconv.Itoa(idSeguimiento),
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

	log.Printf("Generando linea en archivo registro.csv, PYME tipo %v", tipoPyme)

	var fileData [][]string
	fileData = append(fileData, []string{timestamp.Format("2020-01-01 00:00"),
		strconv.Itoa(idSeguimiento),
		tipoPyme,
		ordenPyme.GetProducto(),
		strconv.Itoa(int(ordenPyme.GetValor())),
		ordenPyme.GetOrigen(),
		ordenPyme.GetDestino(),
		strconv.Itoa(idSeguimiento)})

	csvWriter := csv.NewWriter(seguimientoFile)
	csvWriter.WriteAll(fileData)
}

func enqueue(queue []paquete, element paquete) []paquete {
	queue = append(queue, element) // Simply append to enqueue.
	// fmt.Println("Enqueued:", element)
	return queue
}

func dequeue(queue []paquete) []paquete {
	element := queue[0] // The first element is the one to be dequeued.
	// fmt.Println("Dequeued:", element)
	return queue[1:] // Slice off the element once it is dequeued.
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
