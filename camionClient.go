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
	Seguimiento  int
	entregado    bool
	penalizacion int32
}

/** ochentaPorcientoXD
** Parámetros **
Ninguno
** Retorno **
- bool

** Descripción **
Realiza un random para revisar si la entrega se cumple o no. 
Si el valor está dentro del 80%, se cumple y retorna true. Si no, retorna false.

**/
func ochentaPorcientoXD() bool {
	rand.Seed(time.Now().UnixNano())
	probabilidad := rand.Intn(100)

	if probabilidad < 81 {
		return true
	} else {
		return false
	}

}

/** generarRegistro
** Parámetros **
- idCamion: id del camión que va a registrar información
- fecha: fecha y hora de la entrega realizada
- paquete: información del paquete
** Retorno **
Ninguno

** Descripción **
Escribe en un archivo .csv el registro de las entregas realizadas por el camion asignado.
Contiene el ID, tipo, valor, origen, destino, intentos y la fecha de entrega

**/
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

	// log.Printf("Generando linea en archivo del camion xd")

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


/** entregaNormal
** Parámetros **
- conn: conexión gRPC con el servidor
- truck: información del camion
- tiempoEntrega: tiempo que demora el camión en realizar un envío
** Retorno **
Ninguno

** Descripción **
Realiza los envíos de los paquetes asignados al camión del tipo normal. 
A su vez, genera registros y envía la información al servidor, con el uso de otras funciones.
**/
func entregaNormal(conn *grpc.ClientConn, truck *Camion, tiempoEntrega int) {
	var intentosTotales int
	intentosTotales = 0

	for intentosTotales < 2 {
		// log.Println("Cantidad de intentos entrega normal: &v", strconv.Itoa(intentosTotales))
		if truck.infoPaquete1.Valor > truck.infoPaquete2.Valor && truck.infoPaquete1.entregado == false {
			//
			if truck.infoPaquete1.Valor > truck.infoPaquete1.penalizacion {
				truck.infoPaquete1.Intentos++
				time.Sleep(time.Duration(tiempoEntrega) * time.Second)
				if ochentaPorcientoXD() == true {
					timestamp := time.Now()
					truck.infoPaquete1.entregado = true
					generarRegistro(strconv.Itoa(truck.Id), timestamp.String(), truck.infoPaquete1)
					actualizarSeguimiento(conn, truck.infoPaquete1)
					truck.cantPaquetes--

				} else {
					truck.infoPaquete1.penalizacion += 10
				}
			}

			if truck.infoPaquete2.Valor > truck.infoPaquete2.penalizacion && truck.infoPaquete2.entregado == false {
				truck.infoPaquete2.Intentos++
				time.Sleep(time.Duration(tiempoEntrega) * time.Second)
				if ochentaPorcientoXD() == true {
					timestamp := time.Now()
					truck.infoPaquete2.entregado = true
					generarRegistro(strconv.Itoa(truck.Id), timestamp.String(), truck.infoPaquete2)
					actualizarSeguimiento(conn, truck.infoPaquete2)
					truck.cantPaquetes--

				} else {
					truck.infoPaquete2.penalizacion += 10
				}
			}
		} else if truck.infoPaquete2.Valor > truck.infoPaquete1.Valor && truck.infoPaquete2.entregado == false {
			//
			if truck.infoPaquete2.Valor > truck.infoPaquete2.penalizacion {
				truck.infoPaquete2.Intentos++
				time.Sleep(time.Duration(tiempoEntrega) * time.Second)
				if ochentaPorcientoXD() == true {
					timestamp := time.Now()
					truck.infoPaquete2.entregado = true
					generarRegistro(strconv.Itoa(truck.Id), timestamp.String(), truck.infoPaquete2)
					actualizarSeguimiento(conn, truck.infoPaquete2)
					truck.cantPaquetes--

				} else {
					truck.infoPaquete2.penalizacion += 10
				}
			}

			if truck.infoPaquete1.Valor > truck.infoPaquete1.penalizacion && truck.infoPaquete1.entregado == false {
				truck.infoPaquete1.Intentos++
				time.Sleep(time.Duration(tiempoEntrega) * time.Second)
				if ochentaPorcientoXD() == true {
					timestamp := time.Now()
					truck.infoPaquete1.entregado = true
					generarRegistro(strconv.Itoa(truck.Id), timestamp.String(), truck.infoPaquete1)
					actualizarSeguimiento(conn, truck.infoPaquete1)
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
		actualizarSeguimiento(conn, truck.infoPaquete1)

	}

	if truck.infoPaquete2.entregado == false {
		generarRegistro(strconv.Itoa(truck.Id), "0", truck.infoPaquete2)
		actualizarSeguimiento(conn, truck.infoPaquete2)
	}

	fmt.Printf("Entrega normal/prioritaria \n")
	fmt.Printf("Estado paquete 1: %t\n", truck.infoPaquete1.entregado)
	fmt.Printf("Estado paquete 2: %t\n", truck.infoPaquete2.entregado)
	truck.cantPaquetes = 0
	truck.infoPaquete1 = infoPaquete{}
	truck.infoPaquete2 = infoPaquete{}
}


/** entregaRetail
** Parámetros **
- conn: conexión gRPC con el servidor
- truck: información del camion
- tiempoEntrega: tiempo que demora el camión en realizar un envío
** Retorno **
Ninguno

** Descripción **
Realiza los envíos de los paquetes asignados al camión del tipo retail. 
A su vez, genera registros y envía la información al servidor, con el uso de otras funciones.
**/
func entregaRetail(conn *grpc.ClientConn, truck *Camion, tiempoEntrega int) {
	var intentosTotales int
	intentosTotales = 0

	for intentosTotales < 3 {
		// log.Println("Cantidad de intentos entrega retail: &v", strconv.Itoa(intentosTotales))
		if truck.infoPaquete1.Valor > truck.infoPaquete2.Valor && truck.infoPaquete1.entregado == false {
			//
			if truck.infoPaquete1.Valor > truck.infoPaquete1.penalizacion {
				truck.infoPaquete1.Intentos++
				time.Sleep(time.Duration(tiempoEntrega) * time.Second)
				if ochentaPorcientoXD() == true {
					timestamp := time.Now()
					truck.infoPaquete1.entregado = true
					generarRegistro(strconv.Itoa(truck.Id), timestamp.String(), truck.infoPaquete1)
					actualizarSeguimiento(conn, truck.infoPaquete1)
					truck.cantPaquetes--

				} else {
					truck.infoPaquete1.penalizacion += 10
				}
			}

			if truck.infoPaquete2.Valor > truck.infoPaquete2.penalizacion && truck.infoPaquete2.entregado == false {
				truck.infoPaquete2.Intentos++
				time.Sleep(time.Duration(tiempoEntrega) * time.Second)
				if ochentaPorcientoXD() == true {
					timestamp := time.Now()
					truck.infoPaquete2.entregado = true
					generarRegistro(strconv.Itoa(truck.Id), timestamp.String(), truck.infoPaquete2)
					actualizarSeguimiento(conn, truck.infoPaquete2)
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
				time.Sleep(time.Duration(tiempoEntrega) * time.Second)
				if ochentaPorcientoXD() == true {
					timestamp := time.Now()
					truck.infoPaquete2.entregado = true
					generarRegistro(strconv.Itoa(truck.Id), timestamp.String(), truck.infoPaquete2)
					actualizarSeguimiento(conn, truck.infoPaquete2)
					truck.cantPaquetes--

				} else {
					truck.infoPaquete2.penalizacion += 10
				}
			}

			if truck.infoPaquete1.Valor > truck.infoPaquete1.penalizacion && truck.infoPaquete1.entregado == false {
				truck.infoPaquete1.Intentos++
				time.Sleep(time.Duration(tiempoEntrega) * time.Second)
				if ochentaPorcientoXD() == true {
					timestamp := time.Now()
					truck.infoPaquete1.entregado = true
					generarRegistro(strconv.Itoa(truck.Id), timestamp.String(), truck.infoPaquete1)
					actualizarSeguimiento(conn, truck.infoPaquete1)
					truck.cantPaquetes--

				} else {
					truck.infoPaquete1.penalizacion += 10
				}
			}

		}
		intentosTotales++
	}

	if truck.infoPaquete1.Tipo == "Retail" && truck.infoPaquete2.Tipo == "Retail" {
		truck.entregaRetail = true
	}

	if truck.infoPaquete1.entregado == false {
		generarRegistro(strconv.Itoa(truck.Id), "0", truck.infoPaquete1)
		actualizarSeguimiento(conn, truck.infoPaquete1)
	}

	if truck.infoPaquete2.entregado == false {
		generarRegistro(strconv.Itoa(truck.Id), "0", truck.infoPaquete2)
		actualizarSeguimiento(conn, truck.infoPaquete2)
	}

	fmt.Printf("Entrega retail \n")
	fmt.Printf("Estado paquete 1: %t\n", truck.infoPaquete1.entregado)
	fmt.Printf("Estado paquete 2: %t\n", truck.infoPaquete2.entregado)
	truck.cantPaquetes = 0
	truck.infoPaquete1 = infoPaquete{}
	truck.infoPaquete2 = infoPaquete{}
}

/** pedirPaquete
** Parámetros **
- conn: conexión gRPC con el servidor
- truck: información del camion
** Retorno **
infoPaquete: información del paquete recibido

** Descripción **
Se comunica con el servidor para pedir un paquete a entregar.
**/
func pedirPaquete(conn *grpc.ClientConn, truck *Camion) infoPaquete {
	c := pb.NewLogisticaServiceClient(conn)

	paquete, errorPaquete := c.SolicitarPaquete(context.Background(), &pb.Camion{
		Id:            int32(truck.Id),
		Tipo:          truck.Tipo,
		EntregaRetail: truck.entregaRetail})
	if errorPaquete != nil {
		// log.Println("")
		return infoPaquete{}
	}
	fmt.Println("Se recibio exitosamente el paquete. Su ID es: ", paquete.GetId())
	fmt.Println("ID Seguimiento: ", paquete.GetSeguimiento())

	infoPaquete := infoPaquete{
		Id:           paquete.GetId(),
		Tipo:         paquete.GetTipo(),
		Valor:        paquete.GetValor(),
		Origen:       paquete.GetOrigen(),
		Destino:      paquete.GetDestino(),
		Seguimiento:  int(paquete.GetSeguimiento()),
		Intentos:     0,
		penalizacion: 0}
	truck.cantPaquetes++
	return infoPaquete
}

/** actualizarSeguimiento
** Parámetros **
- conn: conexión gRPC con el servidor
- paquete; información del paquete
** Retorno **
Ninguno

** Descripción **
Se comunica con el servidor para informarle del estado del paquete (si fue recibido o no), junto a los intentos.
**/
func actualizarSeguimiento(conn *grpc.ClientConn, paquete infoPaquete) {
	c := pb.NewLogisticaServiceClient(conn)

	// fmt.Println("Actualizando estado del paquete.")
	// fmt.Println("Entregado: ", paquete.entregado)
	// fmt.Println("Seguimiento: ", paquete.Seguimiento)
	// fmt.Println("Intentos: ", paquete.Intentos)

	_, errorStatus := c.ActualizarSeguimiento(context.Background(), &pb.UpdateSeguimiento{
		Entregado:   paquete.entregado,
		Seguimiento: int32(paquete.Seguimiento),
		Intentos:    paquete.Intentos})

	if errorStatus != nil {
		log.Println("Error al actualizar estado del paquete")
	}

}


/** cargarCamion
** Parámetros **
- conn: conexión gRPC con el servidor
- truck: información del camion
- waitTime; tiempo de espera del camión para recibir un segundo paquete
** Retorno **
Ninguno

** Descripción **
Se comunica con el servidor para pedir los paquetes. Pide máximo dos.
**/
func cargarCamion(conn *grpc.ClientConn, truck *Camion, waitTime int) {
	truck.infoPaquete1 = pedirPaquete(conn, truck)
	if truck.cantPaquetes == 1 {
		time.Sleep(time.Duration(waitTime) * time.Second)
		truck.infoPaquete2 = pedirPaquete(conn, truck)
	}
}


/** main
** Parámetros **
Ninguno
** Retorno **
Ninguno

** Descripción **
Crea los tres camiones, pide tiempos de espera para un segundo paquete y de demora del envío
y realiza un loop infinito de carga de camiones y entrega de paquetes
**/
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

	var tiempoEntrega int
	log.Printf("Ingrese el tiempo que demora en entregar un paquete: ")
	fmt.Scanln(&tiempoEntrega)

	for {
		// log.Println("Comienzo de carga")
		// log.Println("Cantidad paquetes camion 1: %v", strconv.Itoa(camion1.cantPaquetes))
		// log.Println("Cantidad paquetes camion 2: %v", strconv.Itoa(camion2.cantPaquetes))
		// log.Println("Cantidad paquetes camion 3: %v", strconv.Itoa(camion3.cantPaquetes))
		// carga de paquetes
		if camion1.cantPaquetes == 0 {
			cargarCamion(conn, camion1, waitTime)
		}
		if camion2.cantPaquetes == 0 {
			cargarCamion(conn, camion2, waitTime)
		}
		if camion3.cantPaquetes == 0 {
			cargarCamion(conn, camion3, waitTime)
		}

		// entrega de paquetes
		if camion1.cantPaquetes != 0 {
			fmt.Println("Camion 1 cargado, procediendo a entrega. Cantidad de paquetes: ", camion1.cantPaquetes)
			entregaRetail(conn, camion1, tiempoEntrega)
		}
		if camion2.cantPaquetes != 0 {
			fmt.Println("Camion 2 cargado, procediendo a entrega. Cantidad de paquetes: ", camion2.cantPaquetes)
			entregaRetail(conn, camion2, tiempoEntrega)
		}
		if camion3.cantPaquetes != 0 {
			fmt.Println("Camion 3 cargado, procediendo a entrega. Cantidad de paquetes: ", camion3.cantPaquetes)
			entregaNormal(conn, camion3, tiempoEntrega)
		}
		// log.Println("Fin de entrega")
	}
}
