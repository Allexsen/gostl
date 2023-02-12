package deque

import (
	"fmt"
	"strings"

	bilist "github.com/Allexsen/gostl/ds/list/bidir-list"
)

type Deque[T any] struct {
	values *bilist.List[T]
}

func New[T any]() *Deque[T] {
	return &Deque[T]{bilist.New[T]()} // Declaration with zero value
}

func (dq *Deque[T]) Size() int {
	return dq.values.Len()
}

func (dq *Deque[T]) Empty() bool {
	return dq.values.Empty()
}

func (dq *Deque[T]) PushFront(val T) {
	dq.values.PushFront(val)
}

func (dq *Deque[T]) PushBack(val T) {
	dq.values.PushBack(val)
}

func (dq *Deque[T]) Front() T {
	return dq.values.Head().Value
}

func (dq *Deque[T]) Back() T {
	return dq.values.Tail().Value
}

func (dq *Deque[T]) PopFront() T {
	return dq.values.PopFront().Value
}

func (dq *Deque[T]) PopBack() T {
	return dq.values.PopBack().Value
}

func (dq *Deque[T]) Clear() {
	dq.values.Clear()
}

func (dq *Deque[T]) String() string {
	var sb strings.Builder
	sb.WriteString("[")
	for node := dq.values.Head(); node != nil; node = node.Next() {
		sb.WriteString(fmt.Sprintf("%v", node.Value))
		if node.Next() != nil {
			sb.WriteString(" ")
		}
	}

	sb.WriteString("]")
	return sb.String()
}
