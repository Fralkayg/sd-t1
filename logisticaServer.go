// Package main implements a server for Greeter service.
package main

import (
	"context"
	"log"
	"net"
	"strconv"

	pb "github.com/Fralkayg/sd-t1/Service"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	seguimientoPyme   int
	seguimientoRetail int
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *server) GenerarOrdenPyme(ctx context.Context, ordenPyme *pb.OrdenPyme) (*pb.SeguimientoPyme, error) {
	log.Printf("Id orden: %v", ordenPyme.GetId())
	idSeguimiento, err := strconv.Atoi(ordenPyme.GetId())
	if err != nil {
		log.Printf("Ocurrio un error al hacer la transformación de datos.")
	}

	log.Printf("Aqui deberia estar generandose la orden de Pyme")
	return &pb.SeguimientoPyme{Id: int32(idSeguimiento)}, nil
}

func (s *server) GenerarOrdenRetail(ctx context.Context, ordenRetail *pb.OrdenRetail) (*pb.SeguimientoRetail, error) {
	log.Printf("Id orden: %v", ordenRetail.GetId())
	idSeguimiento, err := strconv.Atoi(ordenRetail.GetId())
	if err != nil {
		log.Printf("Ocurrio un error al hacer la transformación de datos.")
	}

	log.Printf("Aqui deberia estar generandose la orden Retail")
	return &pb.SeguimientoRetail{Id: int32(idSeguimiento)}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	s := server{}

	//Contador para indicar el ID de seguimiento para cada uno de los servicios.
	s.seguimientoPyme = 0
	s.seguimientoRetail = 0

	pb.RegisterLogisticaServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
