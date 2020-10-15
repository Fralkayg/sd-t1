// Package main implements a server for Greeter service.
package main

import (
	"context"
	"encoding/csv"
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
	seguimiento int
	lock        bool
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

	log.Printf("Aqui deberia estar generandose la orden Retail")

	s.lock = false
	return &pb.SeguimientoRetail{Id: int32(s.seguimiento)}, nil
}

func registroOrdenRetail(ordenRetail *pb.OrdenRetail, idSeguimiento int) {
	seguimientoFile, err := os.OpenFile("./registro.csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Printf("Hubo un error al abrir/crear archivo seguimiento. Tipo: Retail")
	}

	defer seguimientoFile.Close()

	timestamp := time.Now()

	var fileData [][]string

	log.Printf("Generando linea en archivo registro.csv, Retail")

	fileData = append(fileData, []string{timestamp.String(),
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
	fileData = append(fileData, []string{timestamp.String(),
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
