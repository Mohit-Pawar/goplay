package main

import (
  "fmt"
  "net"
  "time"
)

func main() {
  ln, err := net.Listen("tcp", ":8080")
  if err != nil {
    panic(err)
  }
  for {
    conn, err := ln.Accept()
    err = conn.SetDeadline(time.Now().Add(1 * time.Second))
    if err != nil {
      panic(err)
    }
    if err != nil {
      panic(err)
    }
    go handleConnection(conn)
  }
}

func handleConnection(conn net.Conn) {
  conn.Write([]byte{'h', 'e', 'l', 'l', 'o', ',', ' ', 'w', 'o', 'r', 'l', 'd'})
  fmt.Println("LocalAddr :- ", conn.LocalAddr())
  fmt.Println("RemoveAddr :- ", conn.RemoteAddr())
  conn.Close()
}