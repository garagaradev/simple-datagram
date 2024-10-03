package main
import (
  "fmt"
  "net"
)

func main(){
  //UDP server address
  serverAddr := "localhost:8080"
  udpAddr, err := net.ResolveUDPAddr("udp", serverAddr)
  if err != nil {
    fmt.Println("Error resolving address:", err)
    return
  }

  //make a UDP connection to the server
  conn, err := net.DialUDP("udp", nil, udpAddr)
  if err != nil {
    fmt.Println("Error connecting to UDP server:", err)
    return
  }

  defer conn.Close()

  // the datagram
  message := []byte("Hello, this is a datagram!")
  //sending the datagram
  _, err := conn.Write(message)
  if err != nil {
    fmt.Println("Error sending datagram:", err)
    return
  }

  fmt.Printf("Sent message: %s\n", message)

}
