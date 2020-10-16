package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
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
	Id           string
	Tipo         string
	Valor        int32
	Origen       string
	Destino      string
	Intentos     int32
	fechaEntrega string

	entregado    bool
	penalizacion int32
}

func ochentaPorcientoXD() bool {
	rand.Seed(time.Now().UnixNano())
	probabilidad := rand.Intn(100)

	if probabilidad < 81 {
		return true
	} else {
		return false
	}

}

func generarRegistro(idCamion string, fecha string, paquete infoPaquete) {
	CamionFile, err := os.OpenFile("./camion"+idCamion+".csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		seguimientoAux, errAux := os.Create("./camion" + idCamion + ".csv")
		if errAux != nil {
			log.Printf("ola")
		}
		CamionFile = seguimientoAux
	}

	defer CamionFile.Close()

	var fileData [][]string

	log.Printf("Generando linea en archivo del camion xd")

	fileData = append(fileData, []string{
		paquete.Id,
		paquete.Tipo,
		strconv.Itoa(int(paquete.Valor)),
		paquete.Origen,
		paquete.Destino,
		strconv.Itoa(int(paquete.Intentos)),
		paquete.fechaEntrega})

	csvWriter := csv.NewWriter(CamionFile)
	csvWriter.WriteAll(fileData)
	// csvWriter.Flush()
}

func entregaNormal(conn *grpc.ClientConn, truck *Camion) {
	var intentosTotales int
	intentosTotales = 0

	for intentosTotales < 2 {

		if truck.infoPaquete1.Valor > truck.infoPaquete2.Valor && truck.infoPaquete1.entregado == false {
			//
			if truck.infoPaquete1.Valor > truck.infoPaquete1.penalizacion {
				truck.infoPaquete1.Intentos++
				if ochentaPorcientoXD() == true {
					timestamp := time.Now()
					generarRegistro(strconv.Itoa(truck.Id), timestamp.String(), truck.infoPaquete1)
					truck.cantPaquetes--

				} else {
					truck.infoPaquete1.penalizacion += 10
				}
			}

			if truck.infoPaquete2.Valor > truck.infoPaquete2.penalizacion && truck.infoPaquete2.entregado == false {
				truck.infoPaquete2.Intentos++
				if ochentaPorcientoXD() == true {
					timestamp := time.Now()
					generarRegistro(strconv.Itoa(truck.Id), timestamp.String(), truck.infoPaquete2)
					truck.cantPaquetes--

				} else {
					truck.infoPaquete2.penalizacion += 10
				}
			}
		}

		//pack2 > pack1
		if truck.infoPaquete2.Valor > truck.infoPaquete1.Valor && truck.infoPaquete2.entregado == false {
			//
			if truck.infoPaquete2.Valor > truck.infoPaquete2.penalizacion {
				truck.infoPaquete2.Intentos++
				if ochentaPorcientoXD() == true {
					timestamp := time.Now()
					generarRegistro(strconv.Itoa(truck.Id), timestamp.String(), truck.infoPaquete2)
					truck.cantPaquetes--

				} else {
					truck.infoPaquete2.penalizacion += 10
				}
			}

			if truck.infoPaquete1.Valor > truck.infoPaquete1.penalizacion && truck.infoPaquete1.entregado == false {
				truck.infoPaquete1.Intentos++
				if ochentaPorcientoXD() == true {
					timestamp := time.Now()
					generarRegistro(strconv.Itoa(truck.Id), timestamp.String(), truck.infoPaquete1)
					truck.cantPaquetes--

				} else {
					truck.infoPaquete1.penalizacion += 10
				}
			}

		}
		intentosTotales++
	}

	if truck.infoPaquete1.entregado == false {
		generarRegistro(strconv.Itoa(truck.Id), "0", truck.infoPaquete1)

	}

	if truck.infoPaquete2.entregado == false {
		generarRegistro(strconv.Itoa(truck.Id), "0", truck.infoPaquete2)
	}

	truck.infoPaquete1 = infoPaquete{}
	truck.infoPaquete2 = infoPaquete{}
}

func entregaRetail(conn *grpc.ClientConn, truck *Camion) {
	var intentosTotales int
	intentosTotales = 0

	for intentosTotales < 3 {

		if truck.infoPaquete1.Valor > truck.infoPaquete2.Valor && truck.infoPaquete1.entregado == false {
			//
			if truck.infoPaquete1.Valor > truck.infoPaquete1.penalizacion {
				truck.infoPaquete1.Intentos++
				if ochentaPorcientoXD() == true {
					timestamp := time.Now()
					truck.infoPaquete1.entregado = true
					generarRegistro(strconv.Itoa(truck.Id), timestamp.String(), truck.infoPaquete1)
					truck.cantPaquetes--

				} else {
					truck.infoPaquete1.penalizacion += 10
				}
			}

			if truck.infoPaquete2.Valor > truck.infoPaquete2.penalizacion && truck.infoPaquete2.entregado == false {
				truck.infoPaquete2.Intentos++
				if ochentaPorcientoXD() == true {
					timestamp := time.Now()
					truck.infoPaquete2.entregado = true
					generarRegistro(strconv.Itoa(truck.Id), timestamp.String(), truck.infoPaquete2)
					truck.cantPaquetes--

				} else {
					truck.infoPaquete2.penalizacion += 10
				}
			}
		}

		//pack2 > pack1
		if truck.infoPaquete2.Valor > truck.infoPaquete1.Valor && truck.infoPaquete2.entregado == false {
			//
			if truck.infoPaquete2.Valor > truck.infoPaquete2.penalizacion {
				truck.infoPaquete2.Intentos++
				if ochentaPorcientoXD() == true {
					timestamp := time.Now()
					truck.infoPaquete2.entregado = true
					generarRegistro(strconv.Itoa(truck.Id), timestamp.String(), truck.infoPaquete2)
					truck.cantPaquetes--

				} else {
					truck.infoPaquete2.penalizacion += 10
				}
			}

			if truck.infoPaquete1.Valor > truck.infoPaquete1.penalizacion && truck.infoPaquete1.entregado == false {
				truck.infoPaquete1.Intentos++
				if ochentaPorcientoXD() == true {
					timestamp := time.Now()
					truck.infoPaquete1.entregado = true
					generarRegistro(strconv.Itoa(truck.Id), timestamp.String(), truck.infoPaquete1)
					truck.cantPaquetes--

				} else {
					truck.infoPaquete1.penalizacion += 10
				}
			}

		}
		intentosTotales++
	}
	fmt.Printf("Antes de cambiar. Camion: %v. Entrega retail: %t", strconv.Itoa(truck.Id), truck.entregaRetail)
	if truck.infoPaquete1.Tipo == "Retail" && truck.infoPaquete2.Tipo == "Retail" {
		truck.entregaRetail = true
	}
	fmt.Printf("Despues de cambiar. Camion: %v. Entrega retail: %t", strconv.Itoa(truck.Id), truck.entregaRetail)

	if truck.infoPaquete1.entregado == false {
		generarRegistro(strconv.Itoa(truck.Id), "0", truck.infoPaquete1)
	}

	if truck.infoPaquete2.entregado == false {
		generarRegistro(strconv.Itoa(truck.Id), "0", truck.infoPaquete2)
	}

	truck.infoPaquete1 = infoPaquete{}
	truck.infoPaquete2 = infoPaquete{}
}

func pedirPaquete(conn *grpc.ClientConn, truck *Camion) infoPaquete {
	//weas raras
	c := pb.NewLogisticaServiceClient(conn)

	paquete, errorPaquete := c.SolicitarPaquete(context.Background(), &pb.Camion{
		Id:            int32(truck.Id),
		Tipo:          truck.Tipo,
		EntregaRetail: truck.entregaRetail})
	if errorPaquete != nil {
		log.Fatalf("Error al recibir paquete desde logistica")
		return infoPaquete{}
	}
	log.Printf("Se recibio exitosamente el paquete. Su ID es: %v", paquete.GetId())
	infoPaquete := infoPaquete{
		Id:           paquete.GetId(),
		Tipo:         paquete.GetTipo(),
		Valor:        paquete.GetValor(),
		Origen:       paquete.GetOrigen(),
		Destino:      paquete.GetDestino(),
		Intentos:     0,
		penalizacion: 0}
	truck.cantPaquetes++
	return infoPaquete
}

func cargarCamion(conn *grpc.ClientConn, truck *Camion, waitTime int) {
	truck.infoPaquete1 = pedirPaquete(conn, truck)
	if truck.cantPaquetes == 1 {
		time.Sleep(time.Duration(waitTime) * time.Second)
		truck.infoPaquete2 = pedirPaquete(conn, truck)
	}
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	camion1 := &Camion{
		Tipo:          "Retail",
		Id:            1,
		infoPaquete1:  infoPaquete{},
		infoPaquete2:  infoPaquete{},
		cantPaquetes:  0,
		entregaRetail: false}
	camion2 := &Camion{
		Tipo:          "Retail",
		Id:            2,
		infoPaquete1:  infoPaquete{},
		infoPaquete2:  infoPaquete{},
		cantPaquetes:  0,
		entregaRetail: false}
	camion3 := &Camion{
		Tipo:          "Normal",
		Id:            3,
		infoPaquete1:  infoPaquete{},
		infoPaquete2:  infoPaquete{},
		cantPaquetes:  0,
		entregaRetail: false}

	var waitTime int
	log.Printf("Ingrese el tiempo de espera de 2do paquete: ")
	fmt.Scanln(&waitTime)

	for {
		// carga de paquetes
		if camion1.cantPaquetes == 0 {
			cargarCamion(conn, camion1, waitTime)
			log.Printf("Camion 1 cargado")
		}
		if camion2.cantPaquetes == 0 {
			cargarCamion(conn, camion2, waitTime)
			log.Printf("Camion 2 cargado")
		}
		if camion3.cantPaquetes == 0 {
			cargarCamion(conn, camion3, waitTime)
			log.Printf("Camion 3 cargado")
		}

		// entrega de paquetes
		if camion1.cantPaquetes != 0 {
			entregaRetail(conn, camion1)
		}
		if camion2.cantPaquetes != 0 {
			entregaRetail(conn, camion2)
		}
		if camion3.cantPaquetes != 0 {
			entregaNormal(conn, camion3)
		}
	}
}
