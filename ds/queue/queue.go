package queue

import (
	"fmt"
	"strings"

	bilist "github.com/Allexsen/gostl/ds/list/bidir-list"
)

type Queue[T any] struct {
	values *bilist.List[T]
}

func New[T any]() *Queue[T] {
	return &Queue[T]{bilist.New[T]()}
}

func (q *Queue[T]) Size() int {
	return q.values.Len()
}

func (q *Queue[T]) Empty() bool {
	return q.values.Empty()
}

func (q *Queue[T]) Push(val T) {
	q.values.PushBack(val)
}

func (q *Queue[T]) Front() T {
	return q.values.Head().Value
}

func (q *Queue[T]) Back() T {
	return q.values.Tail().Value
}

func (q *Queue[T]) Pop() T {
	return q.values.PopFront().Value
}

func (q *Queue[T]) Clear() {
	q.values.Clear()
}

func (q *Queue[T]) String() string {
	var sb strings.Builder
	sb.WriteString("[")
	for node := q.values.Head(); node != nil; node = node.Next() {
		sb.WriteString(fmt.Sprintf("%v", node.Value))
		if node.Next() != nil {
			sb.WriteString(" ")
		}
	}

	sb.WriteString("]")
	return sb.String()
}
