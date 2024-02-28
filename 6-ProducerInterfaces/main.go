package main

import (
	goodsub "github.com/maxguuse/100GoMistakePractice/6-ProducerInterfaces/good_sub"
	"github.com/maxguuse/100GoMistakePractice/6-ProducerInterfaces/sub"
)

type Worker interface {
	DoWork()
}

func main() {
	// Bad interface - using interface from producer package
	var badSub sub.Worker = &sub.Sub{}
	badSub.DoWork()

	// Good interface - using interface from consumer package
	var goodSub Worker = &goodsub.Sub{}
	goodSub.DoWork()
}
