package array

import (
	"fmt"
	"math"
)

type Array[T any] struct {
	values []T
}

func New[T any](size int) *Array[T] {
	return &Array[T]{values: make([]T, size)}
}

func CreateCopy[T any](toCopy *Array[T]) *Array[T] {
	arr := New[T](len(toCopy.values))
	copy(arr.values, toCopy.values)
	return arr
}

func evalPos(pos, sz int) int {
	if pos < 0 { // If index is negative, make it positive
		pos = sz - (int(math.Abs(float64(pos))) % sz) // And count backwards. F.E: -1 points to the last element
	}

	return pos % sz // Keeps pos in the bounds
}

func (arr *Array[T]) At(pos int) T {
	return arr.values[evalPos(pos, len(arr.values))]
}

func (arr *Array[T]) Set(pos int, val T) {
	arr.values[evalPos(pos, len(arr.values))] = val
}

func (arr *Array[T]) Size() int {
	return len(arr.values)
}

func (arr *Array[T]) Fill(val T) {
	for K := range arr.values {
		arr.values[K] = val
	}
}

func (arr *Array[T]) Empty() bool {
	return len(arr.values) == 0
}

func (arr *Array[T]) Front() T {
	return arr.values[0]
}

func (arr *Array[T]) Back() T {
	return arr.values[len(arr.values)-1]
}

func (arr *Array[T]) SwapArrays(toSwap *Array[T]) {
	if len(arr.values) != len(toSwap.values) {
		panic(fmt.Errorf("sizes of arrays must match"))
	}

	arr.values, toSwap.values = toSwap.values, arr.values
}

func (arr *Array[T]) String() string {
	return fmt.Sprintf("%v", arr.values)
}
