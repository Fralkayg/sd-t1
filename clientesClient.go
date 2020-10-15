// Package main implements a client for Greeter service.
package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"
	"math/rand"
	"fmt"
	"bufio"
	"encoding/csv"
	"io"

	pb "github.com/Fralkayg/sd-t1/Service"
	"google.golang.org/grpc"
)

const (
	address     = "dist54:50051"
	defaultName = "world"
)
/*
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
*/

func generarOrdenRetail(conn *grpc.ClientConn, lineaALeer int){
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
  	for i:=0 ; true ; i++{
  		linea,error := fileReader.Read()
    	if error == io.EOF{
      		return 0
    	}else if error != nil{
        	log.Fatal(error)
        	continue
    	}

    	if lineaALeer == i{
    		valorInt, _ := strconv.Atoi(linea[2])
    		seguimientoRetail, errorRetail := c.GenerarOrdenRetail(ctx, &pb.OrdenRetail{
    			Id: linea[0], 
    			Producto: linea[1],
    			Valor: valorInt,
    			Origen: linea[3],
    			Destino: linea[4],
    		})

			if errorRetail != nil {
				log.Fatalf("Error al enviar orden retail")
			}
			else{
				log.Printf("Se recibio exitosamente su orden. ")
				//Su ID de seguimiento es: %v", strconv.Itoa(int(seguimientoRetail.Id)
			}
    	}

  	}
}

func generarOrdenPyme(conn *grpc.ClientConn, lineaALeer int){
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
  	for i:=0 ; true ; i++{
  		linea,error := fileReader.Read()
    	if error == io.EOF{
      		return 0
    	}else if error != nil{
        	log.Fatal(error)
        	continue
    	}

    	if lineaALeer == i{
    		valorInt, _ := strconv.Atoi(linea[2])
    		PrioriInt, _ := strconv.Atoi(linea[5])
    		seguimientoPyme, errorPyme := c.GenerarOrdenPyme(ctx, &pb.OrdenPyme{
    			Id: linea[0], 
    			Producto: linea[1],
    			Valor: valorInt,
    			Origen: linea[3],
    			Destino: linea[4],
    			Prioritario: PrioriInt
    		})

			if errorPyme != nil {
				log.Fatalf("Error al enviar orden PYME")
			}
			else{
				log.Printf("Se recibio exitosamente su orden. Su ID de seguimiento es: %v", strconv.Itoa(int(seguimientoPyme.Id)))
				file.Close()
				return int(seguimientoPyme.Id)
			}
    	}

  	}




	
}

func hacerSeguimiento(conn *grpc.ClientConn, code_seguimiento int){
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

  	var code_seguimiento[50] int
  	var cant_pedidos_retail int
  	var cant_pedidos_pyme int
  	var opcion int

  	for (cant_pedidos_retail+cant_pedidos_pyme) < 51{ //while algo pase xd 50 pedidos maybe?
      	opcion = rand.Intn(3)

      	if opcion == 0{
        	// orden pyme
        	pyme_realizada = generarOrdenPyme(conn,cant_pedidos_pyme) //entrega el codigo de seguimiento
        	if pyme_realizada != 0{
          		code_seguimiento[cant_pedidos_pyme] = pyme_realizada
          		cant_pedidos_pyme++
       		}
        

      	}
      	else if opcion == 1{
        // orden retail
        	retail_realizado = generarOrdenRetail(conn,cant_pedidos_retail) //algo entregara xd
       		if retail_realizado != 0{
         		cant_pedidos_retail++
       		}

      	}
     	else{
        // pedir seguimiento
        	if cant_pedidos_pyme > 0{
          		rand_seguimiento = rand.Intn(cant_pedidos_pyme)
          		hacerSeguimiento(conn, code_seguimiento[rand_seguimiento])
        	}
      	}
      	time.Sleep(time.Duration(periodo) * time.Second)
  }

}