// Package main implements a client for Greeter service.
package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	pb "github.com/Fralkayg/sd-t1/Service"
	"google.golang.org/grpc"
)

const (
	address     = "dist54:50051"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewLogisticaServiceClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	seguimientoPyme, errorPyme := c.GenerarOrdenPyme(ctx, &pb.OrdenPyme{Id: "1", Producto: "Caca", Valor: 1000, Origen: "Camilo", Destino: "Water", Prioritario: 1})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	if errorPyme != nil {
		log.Fatalf("Error al enviar orden PYME")
	}

	log.Printf("Se recibio exitosamente su orden. Su ID de seguimiento es: %v", strconv.Itoa(int(seguimientoPyme.Id)))

	log.Printf("Greeting: %s", r.GetMessage())
}
