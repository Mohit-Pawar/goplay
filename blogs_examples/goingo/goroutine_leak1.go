package main

import (
	"fmt"
)

func leak() {
	ch := make(chan int)
	go func() {
		val := <-ch
		fmt.Println("We receive Value :-", val)
	}()
}

func main() {
	leak()
}
