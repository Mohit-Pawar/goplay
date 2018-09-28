package main

import (
  "fmt"
  "net"
)

func main() {

  // InterfaceAddrs()
  fmt.Println("---net.InterfaceAddrs()---")
  addrs, err := net.InterfaceAddrs()
  if err != nil {
    panic(err)
  }
  for _, addr := range addrs {
    fmt.Println("Network :- ", addr.Network())
    fmt.Println("String :- ", addr.String())
  }

  // Interfaces()
  fmt.Println("---net.Interfaces()---")
  interfaces, err := net.Interfaces()
  if err != nil {
    panic(err)
  }
  for _, i := range interfaces {
    fmt.Println("Index :- ", i.Index)
    fmt.Println("MTU :- ", i.MTU)
    fmt.Println("Name :- ", i.Name)
    fmt.Println("Hardware Addr :-", i.HardwareAddr)
    fmt.Println("Flags :-", i.Flags)
  }


}