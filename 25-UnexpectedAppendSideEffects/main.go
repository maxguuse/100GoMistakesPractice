package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	f(a[:2]) // Changes a
	fmt.Println(a)

	b := []int{1, 2, 3}
	f(b[:2:2]) // Doesn't change b
	fmt.Println(b)
}

func f(i []int) {
	_ = append(i, 4)
}
