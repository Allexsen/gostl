package vector

import "fmt"

type Vector[T any] struct {
	values []T
}

func New[T any]() *Vector[T] {
	return &Vector[T]{}
}

func (v *Vector[T]) Size() int {
	return len(v.values)
}

func (v *Vector[T]) Empty() bool {
	return len(v.values) == 0
}

func (v *Vector[T]) PushBack(val T) {
	v.values = append(v.values, val)
}

func (v *Vector[T]) Set(pos int, val T) {
	if pos >= 0 && pos < len(v.values) {
		v.values[pos] = val
	}
}

func (v *Vector[T]) Insert(pos int, val T) {
	if pos < 0 || pos > len(v.values) {
		return
	}

	v.values = append(v.values, val)
	for K := len(v.values) - 1; K > pos; K-- {
		v.values[K] = v.values[K-1]
	}

	v.values[pos] = val
}

func (v *Vector[T]) Erase(pos int) {
	v.EraseRange(pos, pos)
}

func (v *Vector[T]) EraseRange(from, to int) {
	if from > to || from < 0 || to > len(v.values) {
		return
	}

	left := v.values[:from]
	right := v.values[to:]
	v.values = append(left, right...)
}

func (v *Vector[T]) At(pos int) T {
	if pos < 0 || pos >= len(v.values) {
		panic("out of range")
	}

	return v.values[pos]
}

func (v *Vector[T]) Front() T {
	return v.At(0)
}

func (v *Vector[T]) Back() T {
	return v.At(len(v.values) - 1)
}

func (v *Vector[T]) PopBack() T {
	if len(v.values) == 0 {
		panic("out of range")
	}

	back := v.values[len(v.values)-1]
	v.values = v.values[:len(v.values)-1]
	return back
}

func (v *Vector[T]) Reverse() {
	sz := len(v.values) - 1
	rng := sz / 2
	for K := 0; K <= rng; K++ {
		v.values[K], v.values[sz-K] = v.values[sz-K], v.values[K]
	}
}

func (v *Vector[T]) Clear() {
	v.values = v.values[:0]
}

func (v *Vector[T]) Values() []T {
	return v.values
}

func (v *Vector[T]) String() string {
	return fmt.Sprintf("%v", v.values)
}
