package list

import (
	"fmt"
	"strings"
)

type List[T any] struct {
	Head *Node[T]
	Len  int
}

func New[T any]() *List[T] {
	return &List[T]{
		Head: nil,
		Len:  0,
	}
}

func (l *List[T]) Add(value T) {
	newNode := &Node[T]{Value: value, Next: nil}
	if l.Head == nil {
		l.Head = newNode
		return
	}
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode

	l.Len++
}

func (l *List[T]) Get(index int) T {
	if index < 0 || index >= l.Len {
		var zero T
		return zero
	}

	if l.Head == nil {
		var zero T
		return zero
	}

	current := l.Head
	for i := 0; i < index; i++ {
		current = current.Next
	}
	return current.Value
}

func (l *List[T]) String() string {
	if l.Head == nil {
		return "[]"
	}
	values := make([]string, 0, l.Len)
	current := l.Head
	for current != nil {
		values = append(values, fmt.Sprintf("%v", current.Value))
		current = current.Next
	}
	return fmt.Sprintf("[%s]", strings.Join(values, " "))
}

type Node[T any] struct {
	Value T
	Next  *Node[T]
}
