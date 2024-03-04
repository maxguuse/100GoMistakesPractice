package main

import (
	"fmt"
	"io"

	"github.com/maxguuse/100GoMistakePractice/9-GenericsMess/list"
)

func main() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	keys, err := getKeys(m)
	if err != nil {
		panic(err)
	}
	fmt.Println(keys)

	m2 := map[int]string{
		1: "a",
		2: "b",
		3: "c",
	}
	keys2 := getKeysGeneric(m2)
	fmt.Println(keys2)

	m3 := map[key]string{
		1: "a",
		2: "b",
		3: "c",
	}
	keys3 := getKeysCustom(m3)
	fmt.Println(keys3)

	l := list.New[int]()
	l.Add(1)
	l.Add(100)
	l.Add(3)
	fmt.Println(l)
}

type key int

func getKeys(m any) ([]any, error) {
	switch t := m.(type) {
	default:
		return nil, fmt.Errorf("unknown type: %T", t)
	case map[string]int:
		keys := make([]any, 0, len(t))
		for k := range t {
			keys = append(keys, k)
		}
		return keys, nil
	case map[int]string:
		keys := make([]any, 0, len(t))
		for k := range t {
			keys = append(keys, k)
		}
		return keys, nil
	}
}

func getKeysGeneric[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

type customConstraint interface {
	~int | ~string
}

func getKeysCustom[K customConstraint, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// Why the fuck use generic here???
func foo[T io.Writer](w T) {
	b := []byte{1, 2, 3}
	_, _ = w.Write(b)
}
