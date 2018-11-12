package main

import (
	"fmt"
)

func main() {
	PrintBoolValues()
	PrintIota()
	NilAsZeroValue()
	AppendExample()
	CapAndLenExample()
	CopyExample()
	DeleteExample()
}

func PrintBoolValues() {
	fmt.Println(true)
	fmt.Println(false)
	// Output
	// true
	// false
}

func PrintIota() {
	// fmt.Println(iota)
	// 0
}

// nil is Zero Value for pointer, channel, func, interface, map, or slice type
func NilAsZeroValue() {
	fmt.Println("NilAsZeroValue")
	var s []string = nil
	var p *int
	var ch chan int
	var fn func()
	var i interface{}
	var m map[string]string

	if s == nil {
		fmt.Println("Zero Value for Slice is nil")
	}
	if p == nil {
		fmt.Println("Zero Value for pointer is nil")
	}
	if ch == nil {
		fmt.Println("Zero Value for channel is nil")
	}
	if fn == nil {
		fmt.Println("Zero Value for func is nil")
	}
	if i == nil {
		fmt.Println("Zero Value for interface is nil")
	}
	if m == nil {
		fmt.Println("Zero Value for map is nil")
	}
}

// appends elements to the end of a slice
func AppendExample() {
	fmt.Println("Append Example")

	var s []string = []string{"a", "b"}
	fmt.Printf("len=%d, cap=%d, %v\n", len(s), cap(s), s)
	// Output
	// len=2, cap=2, [a b]

	// One Element
	s = append(s, "c")
	fmt.Printf("len=%d, cap=%d, %v\n", len(s), cap(s), s)
	// Output
	// len=3, cap=4, [a b c]

	// More Than One Element
	s = append(s, "d", "e")
	fmt.Printf("len=%d, cap=%d, %v\n", len(s), cap(s), s)
	// Output
	// len=5, cap=8, [a b c d e]

	// Another Slice with ...
	s = append(s, []string{"f", "g", "h", "i", "j"}...)
	fmt.Printf("len=%d, cap=%d, %v\n", len(s), cap(s), s)
	// Output
	// len=10, cap=16, [a b c d e f g h i j]

	// append to nil slice
	var nilS []string
	nilS = append(nilS, "a")
	fmt.Printf("len=%d, cap=%d, %v\n", len(nilS), cap(nilS), nilS)
	// Output
	// len=1, cap=1, [a]
}

// cap and len

func CapAndLenExample() {
	fmt.Println("CapExample")

	// Array
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("len=%d, cap=%d, %v\n", len(a), cap(a), a)
	// Output
	// len=5, cap=5, [1 2 3 4 5]

	// Pointer to Array
	var pa = &[5]int{1, 2, 3, 4, 5}
	fmt.Printf("len=%d, cap=%d, %v\n", len(pa), cap(pa), pa)
	// Output
	// len=5, cap=5, &[1 2 3 4 5]

	// Slice
	var s []string = []string{"a", "b", "c", "d"}
	fmt.Printf("len=%d, cap=%d, %v\n", len(s), cap(s), s)
	// Output
	//len=4, cap=4, [a b c d]

	// Capacity of Nil Slice
	var nilS []string
	fmt.Printf("len=%d, cap=%d, %v\n", len(nilS), cap(nilS), nilS)
	// Output
	// len=0, cap=0, []

	// UnBuffered Channel
	var ch chan int = make(chan int)
	fmt.Printf("len=%d, cap=%d, %v\n", len(ch), cap(ch), ch)
	// Output
	// len=0, cap=0, 0xc0000260c0

	// Buffered Channel
	var bCh chan int = make(chan int, 10)
	fmt.Printf("len=%d, cap=%d, %v\n", len(bCh), cap(bCh), bCh)
	// Output
	// len=0, cap=10, 0xc000024160

	// Nil Channel
	var nilCh chan int
	fmt.Printf("len=%d, cap=%d, %v\n", len(nilCh), cap(nilCh), nilCh)
	// Output
	// len=0, cap=0, <nil>
}

func CopyExample() {

	// dest > src
	var s []byte = []byte{'a', 'b', 'c', 'd'}
	var d []byte = []byte{'p', 'q', 'r', 's', 'x', 'y', 'z'}
	n := copy(d, s)
	fmt.Println("N: ", n, "S: ", s, "D: ", d)
	// Output
	// N:  4 S:  [97 98 99 100] D:  [97 98 99 100 120 121 122]

	// dest < src
	s = []byte{'a', 'b', 'c', 'd'}
	d = []byte{'x', 'y', 'z'}
	n = copy(d, s)
	fmt.Println("N: ", n, "S: ", s, "D: ", d)
	// Output
	// N:  3 S:  [97 98 99 100] D:  [97 98 99]

	// empty dest
	s = []byte{'a', 'b'}
	d = []byte{}
	n = copy(d, s)
	fmt.Println("N: ", n, "S: ", s, "D: ", d)
	// Output
	// N:  0 S:  [97 98] D:  []

	// empty src
	s = []byte{}
	d = []byte{'a', 'b', 'c'}
	n = copy(d, s)
	fmt.Println("N: ", n, "S: ", s, "D: ", d)
	// Output
	// N:  0 S:  [] D:  [97 98 99]

	// nil dest
	s = []byte{'a', 'b'}
	d = nil
	n = copy(d, s)
	fmt.Println("N: ", n, "S: ", s, "D: ", d)
	// Output
	// N:  0 S:  [97 98] D:  []

	// nil src
	s = nil
	d = []byte{'a', 'b'}
	n = copy(d, s)
	fmt.Println("N: ", n, "S: ", s, "D: ", d)
	// Output
	// N:  0 S:  [] D:  [97 98]
}

func DeleteExample() {
	var emptyMap map[string]int = map[string]int{}
	var nilMap map[string]int
	var m map[string]int = map[string]int{"a": 1, "b": 2}
	delete(emptyMap, "a")
	delete(nilMap, "a")
	delete(m, "a")
	fmt.Println("M: ", m)
	// Output
	// M:  map[b:2]
	delete(m, "c")
	fmt.Println("M: ", m)
	// Output
	// M:  map[b:2]
}

// ONLY for Slice, Map and Channel
func make() {
	var s1 []string = make([]string)
	var s2 []string = make([]string, 10)
}
