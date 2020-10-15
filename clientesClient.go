package main

import (
	"bufio"
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
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

func generarOrdenRetail(conn *grpc.ClientConn, lineaALeer int) {
	c := pb.NewLogisticaServiceClient(conn)
	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	file, _ := os.Open("retail.csv")
	fileReader := csv.NewReader(bufio.NewReader(file))
	for i := 0; true; i++ {
		linea, error := fileReader.Read()

		if error == io.EOF {
			return 0
		} else if error != nil {
			log.Fatal(error)
			continue
		}

		if lineaALeer == i {
			valorInt, _ := strconv.Atoi(linea[2])
			seguimientoRetail, errorRetail := c.GenerarOrdenRetail(ctx, &pb.OrdenRetail{
				Id:       linea[0],
				Producto: linea[1],
				Valor:    valorInt,
				Origen:   linea[3],
				Destino:  linea[4],
			})

			if errorRetail != nil {
				log.Fatalf("Error al enviar orden retail")
			} else {
				log.Printf("Se recibio exitosamente su orden. ")
				//Su ID de seguimiento es: %v", strconv.Itoa(int(seguimientoRetail.Id)
			}
		}
	}
}

func generarOrdenPyme(conn *grpc.ClientConn, lineaALeer int) {
	c := pb.NewLogisticaServiceClient(conn)
	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	file, _ := os.Open("pyme.csv")
	fileReader := csv.NewReader(bufio.NewReader(file))
	for i := 0; true; i++ {
		linea, error := fileReader.Read()
		if error == io.EOF {
			return 0
		} else if error != nil {
			log.Fatal(error)
			continue
		}

		if lineaALeer == i {
			valorInt, _ := strconv.Atoi(linea[2])
			PrioriInt, _ := strconv.Atoi(linea[5])
			seguimientoPyme, errorPyme := c.GenerarOrdenPyme(ctx, &pb.OrdenPyme{
				Id:          linea[0],
				Producto:    linea[1],
				Valor:       valorInt,
				Origen:      linea[3],
				Destino:     linea[4],
				Prioritario: PrioriInt})

			if errorPyme != nil {
				log.Fatalf("Error al enviar orden PYME")
			} else {
				log.Printf("Se recibio exitosamente su orden. Su ID de seguimiento es: %v", strconv.Itoa(int(seguimientoPyme.Id)))
				file.Close()
				return int(seguimientoPyme.Id)
			}
		}
	}
}

func hacerSeguimiento(conn *grpc.ClientConn, codigoSeguimiento int) {
	c := pb.NewLogisticaServiceClient(conn)
	//pedir algo xd
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	var periodo int
	log.Printf("Ingrese el tiempo entre ordenes del cliente:")
	fmt.Scanln(&periodo)

	var codigoSeguimiento [50]int
	var cantPedidosRetail int
	var cantPedidosPyme int
	var opcion int

	for (cantPedidosRetail + cantPedidosPyme) < 51 { //while algo pase xd 50 pedidos maybe?
		opcion = rand.Intn(3)

		if opcion == 0 {
			// orden pyme
			var seguimientoOrden int
			seguimientoOrden = generarOrdenPyme(conn, cantPedidosPyme) //entrega el codigo de seguimiento
			if seguimientoOrden != 0 {
				codigoSeguimiento[cantPedidosPyme] = seguimientoOrden
				cantPedidosPyme++
			}

		} else if opcion == 1 {
			// orden retail
			var seguimientoRetail int
			seguimientoRetail = generarOrdenRetail(conn, cantPedidosRetail) //algo entregara xd
			if seguimientoRetail != 0 {
				cantPedidosRetail++
			}

		} else {
			// pedir seguimiento
			if cantPedidosPyme > 0 {
				rand_seguimiento = rand.Intn(cantPedidosPyme)
				hacerSeguimiento(conn, codigoSeguimiento[rand_seguimiento])
			}
		}
		time.Sleep(time.Duration(periodo) * time.Second)
	}

}
