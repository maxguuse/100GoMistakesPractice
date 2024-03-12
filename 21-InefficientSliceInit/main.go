package main

import "fmt"

func main() {
	foos := make([]Foo, 1000)
	for i := 0; i < 1000; i++ {
		foos[i] = Foo{}
	}

	bars := convert(foos)         // Inefficient
	bars2 := convertWithCap(foos) // More efficient
	bars3 := convertWithLen(foos) // Sligtly more efficient

	fmt.Println(len(bars), len(bars2), len(bars3))
}

type Foo struct{}

type Bar struct{}

func fooToBar(_ Foo) Bar {
	return Bar{}
}

func convert(foos []Foo) []Bar {
	bars := make([]Bar, 0)
	for _, foo := range foos {
		bars = append(bars, fooToBar(foo))
	}

	return bars
}

func convertWithCap(foos []Foo) []Bar {
	bars := make([]Bar, 0, len(foos))
	for _, foo := range foos {
		bars = append(bars, fooToBar(foo))
	}

	return bars
}

func convertWithLen(foos []Foo) []Bar {
	bars := make([]Bar, len(foos))
	for i, foo := range foos {
		bars[i] = fooToBar(foo)
	}

	return bars
}
