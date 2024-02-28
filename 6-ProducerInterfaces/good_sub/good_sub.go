package goodsub

import "fmt"

type Sub struct{}

func (s *Sub) DoWork() {
	fmt.Println("Do Work good sub")
}
