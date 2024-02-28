package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	shadowing()
	correctUsage()
	correctUsageWithoutTempVar()
}

func shadowing() {
	var client *http.Client
	timeout := true

	if timeout {
		client, err := createClientWithTimeout()
		if err != nil {
			panic(err)
		}
		log.Println(client)
	} else {
		client, err := createDefaultclient()
		if err != nil {
			panic(err)
		}
		log.Println(client)
	}

	log.Println(client)
}

func correctUsage() {
	var client *http.Client
	timeout := true

	if timeout {
		c, err := createClientWithTimeout()
		if err != nil {
			panic(err)
		}
		log.Println(c)
		client = c
	} else {
		c, err := createDefaultclient()
		if err != nil {
			panic(err)
		}
		log.Println(c)
		client = c
	}

	log.Println(client)
}

func correctUsageWithoutTempVar() {
	var client *http.Client
	var err error
	timeout := true

	if timeout {
		client, err = createClientWithTimeout()
		log.Println(client)
	} else {
		client, err = createDefaultclient()
		log.Println(client)
	}
	if err != nil {
		panic(err)
	}

	log.Println(client)
}

func createClientWithTimeout() (*http.Client, error) {
	return &http.Client{
		Timeout: 5 * time.Second,
	}, nil
}

func createDefaultclient() (*http.Client, error) {
	return http.DefaultClient, nil
}
