package main

import (
  "fmt"
)

func main() {
  var s string = "hello"
  s := "hello"
  var a []byte = []byte{'A', 'B'}
  printSlice(a)
}

func len

func printSlice(a []byte) {
  fmt.Printf("%d %d %c", len(a), len(a), a)
}

