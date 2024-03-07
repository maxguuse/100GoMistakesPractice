package main

import "fmt"

func main() {
	sum := 100 + 010 // Expect 110, got 108
	fmt.Println(sum)

	sum = 100 + 0o10 // More explicit syntax
	fmt.Println(sum)

	//Also be aware of
	bin := 0b1000
	hex := 0x8
	img := 0 + 8i
	fmt.Println(bin, hex, img)
}
