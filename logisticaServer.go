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
	pb.UnimplementedLogisticaServiceServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *server) GenerarOdenPyme(ctx context.Context, ordenPyme *pb.OrdenPyme) (*pb.SeguimientoPyme, error) {
	log.Printf("Id orden: %v", ordenPyme.GetId())
	idSeguimiento, err := strconv.Atoi(ordenPyme.GetId())
	log.Printf("Aqui deberia estar generandose la orden de Pyme")
	return &pb.SeguimientoPyme{Id: int32(idSeguimiento)}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterLogisticaServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
