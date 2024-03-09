package main

import (
	"fmt"
	"math"
)

func main() {
	var counter int32 = math.MaxInt32
	counter++
	fmt.Printf("counter: %d\n", counter) // Overflow

	counter = math.MaxInt32
	counter = Inc32(counter) // Panics on overflow
	fmt.Printf("counter: %d\n", counter)

	r := AddInt(int(counter), 1) // Panics on overflow
	fmt.Printf("r: %d\n", r)

	r = MultiplyInt(math.MinInt, 2) // Panics on overflow
	fmt.Printf("r: %d\n", r)
}

func Inc32(counter int32) int32 {
	if counter == math.MaxInt32 {
		panic("counter int32 overflow")
	}

	return counter + 1
}

func AddInt(a, b int) int {
	if a > math.MaxInt-b {
		panic("integer overflow")
	}

	return a + b
}

func MultiplyInt(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}

	result := a * b

	if a == 1 || b == 1 {
		return result
	}

	if result/b != a {
		panic("integer overflow")
	}

	return result
}
