package main

import (
	"bufio"
	"errors"
	"fmt"
)

func main() {
	CustomWriterExample()
	CustomWriterWithError()
	BufferedExample()
	LargeWriteExample()
}

type Writer int

func (w *Writer) Write(p []byte) (n int, err error) {
	fmt.Println(len(p))
	return len(p), nil
}

// CustomWriterExample With Bufio and Without Bufio
func CustomWriterExample() {
	fmt.Println("CustomWriterExample With Bufio and Without Bufio")
	w := new(Writer)
	w.Write([]byte{'a'})
	w.Write([]byte{'b'})
	w.Write([]byte{'c'})
	w.Write([]byte{'d'})
	bw := bufio.NewWriterSize(w, 3)
	bw.Write([]byte{'a'})
	bw.Write([]byte{'b'})
	bw.Write([]byte{'c'})
	bw.Write([]byte{'d'})
	err := bw.Flush()
	if err != nil {
		panic(err)
	}
	// Output
	// 1
	// 1
	// 1
	// 1
	// 3
	// 1
}

type WriterWithError int

func (w *WriterWithError) Write(p []byte) (n int, err error) {
	fmt.Printf("Write: %q\n", p)
	return 0, errors.New("boom!")
}

func CustomWriterWithError() {
	fmt.Println("CustomWriterWithError Example")
	w := new(WriterWithError)
	bw := bufio.NewWriterSize(w, 3)
	bw.Write([]byte{'a'})
	bw.Write([]byte{'b'})
	bw.Write([]byte{'c'})
	bw.Write([]byte{'d'})
	err := bw.Flush()
	if err != nil {
		fmt.Println(err)
	}

	// Output
	// Write: "abc"
	// boom!
}

func BufferedExample() {
	fmt.Println("Buffered Example")
	w := new(Writer)
	bw := bufio.NewWriterSize(w, 3)
	bw.Write([]byte{'a'})
	fmt.Println(bw.Buffered())
	bw.Write([]byte{'b'})
	fmt.Println(bw.Buffered())
	bw.Write([]byte{'c'})
	fmt.Println(bw.Buffered())
	bw.Write([]byte{'d'})
	fmt.Println(bw.Buffered())

	// Output
	// 0
	// 1
	// 2
	// 3
	// 1
}

type LargeWrite int

func (w *LargeWrite) Write(p []byte) (n int, err error) {
	fmt.Printf("%q\n", p)
	return len(p), nil
}
func LargeWriteExample() {
	fmt.Println("LargeWriteExample")
	w := new(LargeWrite)
	bw := bufio.NewWriterSize(w, 3)
	bw.Write([]byte("abcd"))
	// Output
	// abcd
}
