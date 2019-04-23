package main

import "fmt"

func DeferTest() int {

	defer func() {
		fmt.Println("defer called")
	}()
	return 1
	//return 2
}

func main() {
	fmt.Println("Value: ", DeferTest())
}
