package main

import "fmt"

func main() {
	var a float64
	posInf := 1 / a
	negInf := -1 / a
	nan := a / a
	fmt.Println(posInf, negInf, nan)

	r1 := f1(10) // Less precise
	r2 := f2(10) // More precise
	fmt.Println(r1, r2)

	x := 100000.001
	y := 1.0001
	z := 1.0002
	fmt.Println(x * (y + z)) // Less precise
	fmt.Println(x*y + x*z)   // More precise
}

func f1(n int) float64 {
	var result float64 = 10_000
	for i := 0; i < n; i++ {
		result += 1.0001
	}
	return result
}

func f2(n int) float64 {
	var result float64 = 0
	for i := 0; i < n; i++ {
		result += 0.0001
	}
	return result + 10_000
}
