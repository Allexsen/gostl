package stack

import (
	"errors"
	"fmt"
)

var ErrEmpty = errors.New("stack is empty")

type Stack[T any] struct {
	values []T
}

func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (stk *Stack[T]) Empty() bool {
	return len(stk.values) == 0
}

func (stk *Stack[T]) Size() int {
	return len(stk.values)
}

func (stk *Stack[T]) Top() T {
	if stk.Empty() {
		panic(ErrEmpty)
	}

	return stk.values[len(stk.values)-1]
}

func (stk *Stack[T]) Pop() T {
	if stk.Empty() {
		panic(ErrEmpty)
	}

	val := stk.values[len(stk.values)-1]
	stk.values = stk.values[:len(stk.values)-1]
	return val
}

func (stk *Stack[T]) Push(val T) {
	stk.values = append(stk.values, val)
}

func (stk *Stack[T]) Clear() {
	stk.values = nil
}

func (stk *Stack[T]) String() string {
	return fmt.Sprintf("%v", stk.values)
}
