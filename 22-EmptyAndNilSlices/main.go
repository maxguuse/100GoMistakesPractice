package main

import (
	"encoding/json"
	"fmt"
)

type customer struct {
	ID         string
	Operations []float32
}

func main() {
	var s []string
	log(1, s)

	s = []string(nil)
	log(2, s)

	s = []string{}
	log(3, s)

	s = make([]string, 0)
	log(4, s)

	var s1 []float32
	customer1 := customer{
		ID:         "foo",
		Operations: s1,
	}
	b, _ := json.Marshal(customer1)
	fmt.Println(string(b)) // {"ID":"foo","Operations":null}

	s2 := make([]float32, 0)
	customer2 := customer{
		ID:         "bar",
		Operations: s2,
	}
	b, _ = json.Marshal(customer2)
	fmt.Println(string(b)) // {"ID":"bar","Operations":[]}
}

func log(i int, s []string) {
	fmt.Printf("%d: empty=%t, nil=%t\n", i, len(s) == 0, s == nil)
}
