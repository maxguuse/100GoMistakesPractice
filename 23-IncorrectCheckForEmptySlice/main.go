package main

import "fmt"

func main() {
	handleOperations("")
	handleOperationsCorrect("")
}
func handleOperations(id string) {
	operations := getOperations(id)
	if operations != nil {
		handle(operations)
	}
}

func handleOperationsCorrect(id string) {
	operations := getOperations(id)
	if len(operations) == 0 { // It's better to check len instead of nil
		fmt.Println("No operations to handle")

		return
	}

	handle(operations)
}

func handle(operations []float32) {
	fmt.Println("Do some work with", operations)
}

func getOperations(id string) []float32 {
	operations := make([]float32, 0)
	if id == "" {
		return operations
	}

	operations = append(operations, 1.0)

	return operations
}
