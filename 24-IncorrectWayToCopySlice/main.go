package main

import "fmt"

func main() {
	src := []int{0, 1, 2}
	var dst []int
	copy(dst, src)
	fmt.Println(dst)

	dst2 := make([]int, len(src))
	copy(dst2, src)
	fmt.Println(dst2)
}
