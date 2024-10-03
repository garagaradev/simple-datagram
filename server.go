package main

import (
  "fmt"
  "net"
)

func main(){
  //create a UDP socket on PORT 8080
  address := ":8080"
  udpAddr, err := net.ResolveUDPAddr("udp", address)
  if err != nil {
    fmt.Println("Error resolving address:", err)
    return
  }
  //listening to UDP connection 
  conn, err := net.ListenUDP("udp", udpAddr)
  if err != nil {
    fmt.Println("Error listening:", err)
    return
  }
  defer conn.Close()

  fmt.Printff("Listening on %s\n", address)

  //buffer for accepting data
  buffer := make([]byte, 1024)

  for {
    //accepting datagram
    n, addr, err := conn.ReadFromUDP(buffer)
    if err != nil {
      fmt.Println("Error reading from UDP:", err)
      return
    }
    //displaying the data accepted
    fmt.Printf("Received %s from %s\n",string(buffer[:n]), addr)
  }

}
