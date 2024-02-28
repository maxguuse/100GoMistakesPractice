package main

import (
	"errors"
	"log"
)

func main() {
	r, err := join("Hello, ", "World!", 10)
	if err != nil {
		panic(err)
	}
	log.Println(r)

	r, err = correctUsage("Hello, ", "World!", 10)
	if err != nil {
		panic(err)
	}
	log.Println(r)
}

func join(s1, s2 string, max int) (string, error) {
	if s1 == "" {
		return "", errors.New("s1 is empty")
	} else {
		if s2 == "" {
			return "", errors.New("s2 is empty")
		} else {
			concat, err := concatenate(s1, s2)
			if err != nil {
				return "", err
			} else {
				if len(concat) > max {
					return concat[:max], nil
				} else {
					return concat, nil
				}
			}
		}
	}
}

func correctUsage(s1, s2 string, max int) (string, error) {
	if s1 == "" {
		return "", errors.New("s1 is empty")
	}

	if s2 == "" {
		return "", errors.New("s2 is empty")
	}

	concat, err := concatenate(s1, s2)
	if err != nil {
		return "", err
	}

	if len(concat) > max {
		return concat[:max], nil
	}

	return concat, nil
}

func concatenate(s1, s2 string) (string, error) {
	return s1 + s2, nil
}
