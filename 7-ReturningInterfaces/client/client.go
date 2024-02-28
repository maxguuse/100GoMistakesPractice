package client

// Can't do that because of circular import
// import "github.com/maxguuse/100GoMistakePractice/7-ReturningInterfaces/sub"
import goodsub "github.com/maxguuse/100GoMistakePractice/7-ReturningInterfaces/good_sub"

type Storer interface {
	Store(string)
}

func Example() {
	/*
		Can't use NewInMemoryStore from sub package because
		of circular import because sub package returns interface
		declared in client package
	*/
	// var store Storer := sub.NewInMemoryStore()
	// store.Store("hello")

	// Can do that because goodsub package returns implementation instead of interface
	var store Storer = goodsub.NewInMemoryStore()
	store.Store("hello")
}
