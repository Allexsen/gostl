package container

type Container[T any] interface {
	PushBack(T)
	PushFront(T)
	Back() T
	Front() T
	PopBack() T
	PopFront() T
	Clear()
	Size() int
	Empty() bool
	String() string
}
