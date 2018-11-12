package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	msg := "boring! "
	// BoringExample(msg)
	// BoringWithGoRoutine(msg)
	// BoringWithChannel(msg)
	GeneratorPattern(msg)
}

func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func BoringExample(msg string) {
	boring(msg)
}

func BoringWithGoRoutine(msg string) {
	go boring(msg)
}

func boringWithChannel(msg string, ch chan string) {
	for i := 0; ; i++ {
		ch <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func BoringWithChannel(msg string) {
	ch := make(chan string)
	go boringWithChannel(msg, ch)
	for i := 0; i < 5; i += 1 {
		fmt.Printf("You say: %q\n", <-ch)
	}
	fmt.Println("You're boring i am leaving.")
}

func boringGenerator(msg string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return ch
}

func GeneratorPattern(msg string) {
	ch := boringGenerator(msg)
	for i := 0; i < 5; i += 1 {
		fmt.Printf("You say: %q\n", <-ch)
	}
	fmt.Println("You're boring i am leaving.")
}
