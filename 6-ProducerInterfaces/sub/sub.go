package sub

import "fmt"

type Worker interface {
	DoWork()
}

type Sub struct{}

func (s *Sub) DoWork() {
	fmt.Println("Do Work sub")
}
