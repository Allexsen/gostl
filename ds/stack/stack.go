package stack

import (
	"errors"
	"fmt"
	"strings"

	unilist "github.com/Allexsen/gostl/ds/list/unidir-list"
)

var ErrEmpty = errors.New("stack is empty")

type Stack[T any] struct {
	values *unilist.List[T]
}

func New[T any]() *Stack[T] {
	return &Stack[T]{values: unilist.New[T]()} // Zero value declaration
}

func (stk *Stack[T]) Empty() bool {
	return stk.values.Len() == 0
}

func (stk *Stack[T]) Size() int {
	return stk.values.Len()
}

// Returns the last element in the stack
func (stk *Stack[T]) Top() T {
	if stk.Empty() {
		panic(ErrEmpty)
	}

	return stk.values.Head().Value // The last element is always the head of the list
}

// Pops the last element in the stack
func (stk *Stack[T]) Pop() T {
	if stk.Empty() {
		panic(ErrEmpty)
	}

	return stk.values.PopFront().Value // The last element is always the head in the list
}

// Adds a new value on top of the stack
func (stk *Stack[T]) Push(val T) {
	stk.values.PushFront(val) // New value will always be a new head
}

// Removes all elements from the stack
func (stk *Stack[T]) Clear() {
	stk.values.Clear()
}

// Stringifies the stack
func (stk *Stack[T]) String() string {
	var sb strings.Builder
	sb.WriteString("[")
	for node := stk.values.Head(); node != nil; node = node.Next() {
		sb.WriteString(fmt.Sprintf("%v", node.Value))
		if node.Next() != nil {
			sb.WriteString(" ")
		}
	}

	sb.WriteString("]")

	// Reverse the values, since stack is LIFO
	// It just makes more sense to me personally
	bytes := []byte(sb.String())
	for l, r := 1, len(bytes)-2; l < r; l++ {
		bytes[l], bytes[r] = bytes[r], bytes[l]
		r--
	}

	return string(bytes)
}
