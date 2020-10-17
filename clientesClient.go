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

func generarOrdenRetail(conn *grpc.ClientConn, lineaALeer int) int {
	c := pb.NewLogisticaServiceClient(conn)

	file, _ := os.Open("./retail.csv")
	fileReader := csv.NewReader(bufio.NewReader(file))
	for i := 0; true; i++ {
		linea, error := fileReader.Read()

		if error == io.EOF {
			return -1
		} else if error != nil {
			log.Fatal(error)
			return -1
		}

		log.Printf("Leyendo archivo retail. Linea: %v", strconv.Itoa(i))

		if lineaALeer == i {
			log.Printf("Encontro la linea correspondiente en Retail. %v", strconv.Itoa(i))
			valorInt, _ := strconv.Atoi(linea[2])
			seguimientoRetail, errorRetail := c.GenerarOrdenRetail(context.Background(), &pb.OrdenRetail{
				Id:       linea[0],
				Producto: linea[1],
				Valor:    int32(valorInt),
				Origen:   linea[3],
				Destino:  linea[4],
			})

			if errorRetail != nil {
				log.Fatalf("Error al enviar orden retail")
			} else {
				log.Printf("Se recibio exitosamente su orden. ")
				//Su ID de seguimiento es: %v", strconv.Itoa(int(seguimientoRetail.Id)
			}
			return int(seguimientoRetail.Id)
		}
	}
	return -1
}

func generarOrdenPyme(conn *grpc.ClientConn, lineaALeer int) int {
	c := pb.NewLogisticaServiceClient(conn)

	file, _ := os.Open("./pymes.csv")
	fileReader := csv.NewReader(bufio.NewReader(file))
	for i := 0; true; i++ {
		linea, error := fileReader.Read()
		if error == io.EOF {
			return -1
		} else if error != nil {
			log.Fatal(error)
			return -1
		}
		log.Printf("Leyendo archivo PYMES. Linea: %v", strconv.Itoa(i))

		if lineaALeer == i {
			log.Printf("Encontro la linea correspondiente en PYME. %v", strconv.Itoa(i))
			valorInt, _ := strconv.Atoi(linea[2])
			PrioriInt, _ := strconv.Atoi(linea[5])
			seguimientoPyme, errorPyme := c.GenerarOrdenPyme(context.Background(), &pb.OrdenPyme{
				Id:          linea[0],
				Producto:    linea[1],
				Valor:       int32(valorInt),
				Origen:      linea[3],
				Destino:     linea[4],
				Prioritario: int32(PrioriInt)})

			if errorPyme != nil {
				log.Fatalf("Error al enviar orden PYME")
			} else {
				log.Printf("Se recibio exitosamente su orden. Su ID de seguimiento es: %v", strconv.Itoa(int(seguimientoPyme.Id)))
				file.Close()
				return int(seguimientoPyme.Id)
			}
		}
	}
	return -1
}

func pymeTest(conn *grpc.ClientConn, codigoSeguimiento int) int {
	c := pb.NewLogisticaServiceClient(conn)

	seguimientoPyme, errorPyme := c.GenerarOrdenPyme(context.Background(), &pb.OrdenPyme{Id: "1", Producto: "Caca", Valor: 1000, Origen: "Camilo", Destino: "Water", Prioritario: 1})
	if errorPyme != nil {
		log.Fatalf("Error al enviar orden PYME")
	} else {
		log.Printf("Se recibio exitosamente su orden. Su ID de seguimiento es: %v", strconv.Itoa(int(seguimientoPyme.Id)))
		return int(seguimientoPyme.Id)
	}
	return -1
}

func hacerSeguimiento(conn *grpc.ClientConn, codigoSeguimiento int) {
	c := pb.NewLogisticaServiceClient(conn)

	infoSeguimiento, errorSeguimiento := c.SolicitarSeguimiento(context.Background(), &pb.SeguimientoPyme{Id: int32(codigoSeguimiento)})

	if errorSeguimiento != nil {
		log.Printf("Ocurrio un error al realizar el seguimiento.")
	} else {
		log.Println("ID Paquete: %v", infoSeguimiento.IDPaquete)
		log.Println("Estado: %v", infoSeguimiento.Estado)
	}
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

	var comp int
	log.Printf("Seleccione el comportamiento. PYME(0) o Retail(1). (0/1):")
	fmt.Scanln(&comp)

	var seguimientos []int

	var cantPedidos int

	var opcion int

	cantPedidos = 1
	opcion = 0

	for cantPedidos < 51 { //while algo pase xd 50 pedidos maybe?
		rand.Seed(time.Now().UnixNano())
		opcion = rand.Intn(2)
		opcionAux := strconv.Itoa(opcion)

		log.Printf("Opcion: %v", opcionAux)
		if opcion == 0 && comp == 0 {
			// orden pyme
			log.Printf("Entro bien en orden PYME")
			var seguimientoPyme int
			seguimientoPyme = generarOrdenPyme(conn, cantPedidos) //entrega el codigo de seguimiento
			if seguimientoPyme != -1 {
				seguimientos = append(seguimientos, seguimientoPyme)

				cantPedidos++
			}

		} else if opcion == 0 && comp == 1 {
			// orden retail
			log.Printf("Entro bien en orden Retail")
			var seguimientoRetail int
			seguimientoRetail = generarOrdenRetail(conn, cantPedidos) //entrega el codigo de seguimiento
			if seguimientoRetail != -1 {
				seguimientos = append(seguimientos, seguimientoRetail)

				cantPedidos++
			}

		} else if opcion == 1 {
			// pedir seguimiento

			log.Printf("Entro bien en Seguimiento")
			if len(seguimientos) > 0 {
				randSeguimiento := rand.Intn(int(len(seguimientos))) + 1

				fmt.Println("Seguimiento random escogido: ", randSeguimiento)

				hacerSeguimiento(conn, randSeguimiento)
			}

		}
		time.Sleep(time.Duration(periodo) * time.Second)
	}

}
