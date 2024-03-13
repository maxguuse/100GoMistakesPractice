package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GC()
	printAlloc()
	a := make([]int, 100000)
	printAlloc()
	b := f(a)
	printAlloc()
	runtime.GC() // Doesn't collect extra elements in b underlying array
	printAlloc()
	runtime.KeepAlive(b)

	foos := make([]Foo, 500)
	printAlloc()
	for i := 0; i < len(foos); i++ {
		foos[i] = Foo{
			v: make([]byte, 1024*1024),
		}
	}
	printAlloc()
	two := k(foos)
	runtime.GC() // Doesn't collect extra elements in two underlying array
	printAlloc()
	runtime.KeepAlive(two)
}

func f(i []int) []int {
	return i[:5:5]
}

type Foo struct {
	v []byte
}

func k(i []Foo) []Foo {
	return i[:2]
}

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d KB\n", m.Alloc/1024)
}
