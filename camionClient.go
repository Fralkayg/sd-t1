package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	pb "github.com/Fralkayg/sd-t1/Service"
	"google.golang.org/grpc"
)

const (
	address = "dist54:50051"
	//defaultName = "world"
)

type Camion struct {
	Tipo          string
	Id            int
	infoPaquete1  infoPaquete
	infoPaquete2  infoPaquete
	cantPaquetes  int
	entregaRetail bool
}

type infoPaquete struct {
	Id           int
	Tipo         string
	Valor        int
	Origen       string
	Destino      string
	Intentos     int
	fechaEntrega string
}

func ochentaPorcientoXD() bool {
	rand.Seed(time.Now().UnixNano())
	probabilidad = rand.Intn(100)

	if probabilidad < 81 {
		return true
	} else {
		return false
	}

}

func entregaRetail() {
	var intentosTotales int
	intentosTotales = 0

	for intentosTotales < 3 {
		if infoPaquete1.Valor > infoPaquete2.Valor {
			if ochentaPorcientoXD() == true {
				timestamp = time.Now()
				// enviar info de paquete 1 enviado
			} else {
				infoPaquete1.Intentos++
			}

			if ochentaPorcientoXD() == true {
				timestamp = time.Now()
				// enviar info de paquete 1 enviado
			} else {
				infoPaquete2.Intentos++
			}
		} else {
			if ochentaPorcientoXD() == true {
				timestamp = time.Now()
				// enviar info de paquete 1 enviado
			} else {
				infoPaquete2.Intentos++
			}

			if ochentaPorcientoXD() == true {
				timestamp = time.Now()
				// enviar info de paquete 1 enviado
			} else {
				infoPaquete1.Intentos++
			}
		}
		intentosTotales++
	}
}

func pedirPaquete(conn *grpc.ClientConn, truck Camion) {
	//weas raras

	truck.cantPaquetes++
}

func camionTest(conn *grpc.ClientConn) infoPaquete {
	c := pb.NewLogisticaServiceClient(conn)

	paquete, errorPaquete := c.SolicitarPaquete(context.Background(), &pb.Camion{Id: "1", Tipo: "retail", EntregaRetail: false})
	if errorPaquete != nil {
		log.Fatalf("Error al recibir paquete desde logistica")
		return nil
	}
	log.Printf("Se recibio exitosamente el paquete. Su ID es: %v", strconv.Itoa(int(paquete.GetId())))
	infoPaquete := &infoPaquete{
		Id:       paquete.GetId(),
		Tipo:     paquete.GetTipo(),
		Valor:    paquete.GetValor(),
		Origen:   paquete.GetOrigen(),
		Destino:  paquete.GetDestino(),
		Intentos: 0}
	return infoPaquete
}

func cargarCamion(conn *grpc.ClientConn, truck Camion, waitTime int) {
	pedirPaquete(conn, truck)
	if truck.cantPaquetes == 1 {
		time.Sleep(time.Duration(waitTime) * time.Second)
		pedirPaquete(conn, truck)
	}

}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	camion1 := &Camion{
		Tipo:          "retail",
		Id:            1,
		infoPaquete1:  infoPaquete{},
		infoPaquete2:  infoPaquete{},
		Paquetes:      0,
		entregaRetail: false}
	camion2 := &Camion{
		Tipo:          "retail",
		Id:            2,
		infoPaquete1:  infoPaquete{},
		infoPaquete2:  infoPaquete{},
		Paquetes:      0,
		entregaRetail: false}
	camion3 := &Camion{
		Tipo:          "normal",
		Id:            3,
		infoPaquete1:  infoPaquete{},
		infoPaquete2:  infoPaquete{},
		Paquetes:      0,
		entregaRetail: false}

	var waitTime int
	log.Printf("Ingrese el tiempo de espera de 2do paquete: ")
	fmt.Scanln(&waitTime)

	camionTest(conn)
	/*
		  	for{
		  		// carga de paquetes
		  		if camion1.cantPaquetes == 0{
					cargarCamion(conn, camion1)
					log.Printf("Camion 1 cargado")

		  		}
		  		if camion2.cantPaquetes == 0{
		  			cargarCamion(camion2)
		  			log.Printf("Camion 2 cargado")
		  		}
		  		if camion3.cantPaquetes == 0{
		  			cargarCamion(camion3)
		  			log.Printf("Camion 3 cargado")
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


	*/

	// }
}
