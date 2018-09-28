package main

import (
  "fmt"
  "net"
  "bufio"
)

func main() {
  conn, err := net.Dial("tcp", "google.com:80")
  if err != nil {
    panic(err)
  }
  fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
  status, err := bufio.NewReader(conn).ReadString('\n')
  if err != nil {
    panic(err)
  }
  fmt.Println(status)
}