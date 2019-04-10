package main

import (
	"fmt"
	"runtime"
	"strings"
)

func processRecords(records []string) {
	input := make(chan string, len(records))
	for _, record := range records {
		input <- record
	}
	//close(input)
	output := make(chan string, len(records))
	workers := runtime.NumCPU()
	for i := 0; i < workers; i++ {
		go worker(i, input, output)
	}

	for i := 0; i < len(records); i++ {
		record := <-output
		fmt.Printf("[result ]: output %s]\n", record)
	}
}

func worker(i int, input <-chan string, output chan<- string) {
	for v := range input {
		fmt.Printf("Worker[%d]: input %s \n", i, v)
		output <- strings.ToUpper(v)
	}
	fmt.Printf("Worker[%d] shutting down\n", i)
}

func main() {
	processRecords([]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l"})
}
