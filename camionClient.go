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
	//defaultName = "world"
)

type Camion struct{
  	Tipo string
 	Id int
  	infoPaquete1 infoPaquete
 	infoPaquete2 infoPaquete
 	cantPaquetes int
}

type infoPaquete struct{
 	Id int
 	Tipo string
 	Valor int
 	Origen string
	Destino string
	Intentos int
	Fecha string
}


func pedirPaquete(conn *grpc.ClientConn, truck Camion){
	//weas raras
	truck.cantPaquetes ++

func cargarCamion(conn *grpc.ClientConn, truck Camion, time int){
	pedirPaquete(conn, truck)
	if truck.cantPaquetes == 1{
		time.Sleep(time.Duration(periodo) * time.Second)
		//pedirPaquete(conn)

	}

}

func main(){
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()


	camion1 := &Camion{Tipo: "retail",Paquetes:0,Estado:0,Id:"1"}
  	camion2 := &Camion{Tipo: "retail",Paquetes:0,Estado:0,Id:"2"}
  	camion3 := &Camion{Tipo: "normal",Paquetes:0,Estado:0,Id:"3"}

  	var waitTime int
  	log.Printf("Ingrese el tiempo de espera de 2do paquete")
  	fmt.Scanln(&waitTime)

  	for{
  		// carga de paquetes
  		if camion1.cantPaquetes == 0{
			cargarCamion(conn, camion1)
  		}
  		if camion2.cantPaquetes == 0{
  			cargarCamion(camion2)
  		}
  		if camion3.cantPaquetes == 0{
  			cargarCamion(camion3)
  		}

  		// entrega de paquetes
  		if camion1.cantPaquetes != 0{
			entregaRetail(camion1)
  		}
  		if camion2.cantPaquetes != 0{
  			entregaRetail(camion2)
  		}
  		if camion3.cantPaquetes != 0{
  			entregaNormal(camion3)
  		}





  	}
}