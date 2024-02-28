package main

// Bad interface, mixed all kinds of methods
type GenericInteface interface {
	A() int
	B() string
	C() ([]byte, error)
}

// Better usage, smaller the interface better the readability
type DoerA interface {
	A() int
}

type DoerB interface {
	B() string
}

type DoerC interface {
	C() ([]byte, error)
}

func main() {

}
