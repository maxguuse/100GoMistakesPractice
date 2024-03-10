package main

import "fmt"

func main() {
	fmt.Println("s1", "\t\t", "s2")
	fmt.Println("---", "\t\t", "---")

	s1 := make([]int, 3, 6)
	s2 := s1[1:3]
	fmt.Println("Initial values")
	fmt.Println(s1, "\t", s2)

	s1[1] = 1
	fmt.Println("Index 1 set to 1")
	fmt.Println(s1, "\t", s2)

	s2 = append(s2, 2, 3, 4)
	fmt.Println("Appended values to s2")
	fmt.Println(s1, "\t", s2)

	s1 = append(s1, 100)
	fmt.Println("Appended value to s1")
	fmt.Println(s1, "\t", s2)

	s2 = append(s2, 5)
	fmt.Println("Appended value to s2, cap exceeded, new underlying array created")
	fmt.Println(s1, "\t", s2)

	s1 = append(s1, 200)
	fmt.Println("Appended value to s1")
	fmt.Println(s1, "", s2)
}
