package main

import "container/list"

func main() {
	list := list.New() // From now you can't use list.New cuz of collisions
	_ = list.PushBack(1)
}
