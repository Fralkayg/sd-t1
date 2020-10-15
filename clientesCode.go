package main
import (
  "os"
  "log"
  "encoding/csv"
  "bufio"
  "io"
  "github.com/" //ingresar tu repo chris 
  "golang.org/x/net/context"
  "google.golang.org/grpc"
  "strconv"
  "math/rand"
  "time"
  "fmt"
)


func pedir_retail(){
  
}

func pedir_pyme(conn *grpc.ClientConn){
  //c := comms.NewCommsClient(conn)
  
}

func pedir_seguimiento(conn *grpc.ClientConn, code_seguimiento int){
  //c := comms.NewCommsClient(conn)

}

func main() {
  var conn *grpc.ClientConn
  conn, err := grpc.Dial("dist54:9000", grpc.WithInsecure())
  if err != nil {
    log.Fatalf("No hubo conexion: %s", err)
  }

  var periodo int
  log.Printf("Ingrese el tiempo entre ordenes del cliente:")
  fmt.Scanln(&periodo)

  var code_seguimiento[50] int
  var cant_pedidos int
  var cant_pedidos_pyme int

  for cant_pedidos < 51{ //while algo pase xd 50 pedidos maybe?
      opcion=rand.Intn(3)

      if opcion == 0{
        // orden pyme
        pyme_realizada = pedir_pyme() //entrega el codigo de seguimiento
        if pyme_realizada != 0{
          code_seguimiento[cant_pedidos_pyme] = pyme_realizada
          cant_pedidos_pyme++
          cant_pedidos++
        }
        

      }
      else if opcion == 1{
        // orden retail
        retail_realizado = pedir_retail() //algo entregara xd
        if retail_realizado != 0{
          cant_pedidos++
        }

      }
      else{
        // pedir seguimiento
        if cant_pedidos_pyme > 0{
          rand_seguimiento = rand.Intn(cant_pedidos_pyme)
          pedir_seguimiento(conn, code_seguimiento[rand_seguimiento])
        }
      }
      time.Sleep(time.Duration(periodo) * time.Second)
  }

}