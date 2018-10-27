package main

import (
	"fmt"
	"strconv"
)

func main() {
	ExampleConversion()
}

func ExampleConversion() {
	i, err := strconv.Atoi("65")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Integer %d", i)
	s := strconv.Itoa(i)
	fmt.Println("String %s", s)

	b, err := strconv.ParseBool("true")
	fmt.Println("Boolean ", b)
	f, err := strconv.ParseFloat("1.3", 64)
	fmt.Println("Fload64", f)
	i64, err := strconv.ParseInt("-42", 10, 64)
	fmt.Println("Int64 ", i64)
}
