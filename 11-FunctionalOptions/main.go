package main

import (
	"log"

	"github.com/maxguuse/100GoMistakePractice/11-FunctionalOptions/builder"
	"github.com/maxguuse/100GoMistakePractice/11-FunctionalOptions/options"
)

func main() {
	// One way to handle ctor with uncertain amount of args
	// Builder pattern
	cfg, err := builder.New().Port(9000).Build()
	if err != nil {
		panic(err)
	}
	log.Println(cfg)

	s, err := builder.NewServer("localhost", cfg)
	if err != nil {
		panic(err)
	}
	log.Println(s)

	// Better way to handle ctor with uncertain amount of args
	// Functional options pattern
	s2, err := options.NewServer(
		"localhost",
		options.WithPort(9000),
	)
	if err != nil {
		panic(err)
	}
	log.Println(s2)
}
