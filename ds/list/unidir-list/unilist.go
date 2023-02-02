package unilist

import (
	"fmt"
	"strings"
)

// List Node
type Node[T any] struct {
	Value T
	next  *Node[T]
}

func (node *Node[T]) Next() *Node[T] {
	return node.next
}

// Unidirectional(Singly) Linked List
type List[T any] struct {
	head *Node[T]
	tail *Node[T]
	len  int
}

func New[T any]() *List[T] {
	return &List[T]{} // Declaration with zero value
}

func (list *List[T]) Head() *Node[T] {
	return list.head
}

func (list *List[T]) Tail() *Node[T] {
	return list.tail
}

func (list *List[T]) Len() int {
	return list.len
}

func (list *List[T]) PushFront(val T) {
	node := &Node[T]{Value: val}
	if list.len == 0 {
		list.head = node
		list.tail = node
	} else {
		node.next = list.head
		list.head = node
	}

	list.len++
}

func (list *List[T]) PushBack(val T) {
	node := &Node[T]{Value: val}
	if list.len == 0 {
		list.head = node
		list.tail = node
	} else {
		list.tail.next = node
		list.tail = node
	}

	list.len++
}

func (list *List[T]) InsertAfter(val T, node *Node[T]) *Node[T] {
	newNode := &Node[T]{Value: val}
	if node == nil {
		newNode.next = list.head
		list.head = newNode
	} else {
		newNode.next = node.next
		node.next = newNode
	}

	if newNode.next == nil {
		list.tail = newNode
	}

	list.len++
	return newNode
}

// prev MUST be right before the node, otherwise list.len will miscount
func (list *List[T]) Remove(prev, node *Node[T]) *Node[T] {
	if node == nil {
		return nil
	}

	if prev == nil {
		list.head = node.next
		if list.head == nil {
			list.tail = nil
		}
	} else {
		prev.next = node.next
		if prev.next == nil {
			list.tail = prev
		}
	}

	list.len--
	return node
}

func (list *List[T]) MakeHead(prev, node *Node[T]) {
	if node == nil || prev == nil || prev.next != node {
		return
	}

	prev.next = node.next
	if prev.next == nil {
		list.tail = prev
	}

	node.next = list.head
	list.head = node
}

func (list *List[T]) MakeTail(prev, node *Node[T]) {
	if node == nil || list.len < 2 {
		return
	}

	if prev == nil {
		list.head = node.next
	} else {
		prev.next = node.next
	}

	list.tail.next = node
	list.tail = node
	node.next = nil
}

func (list *List[T]) String() string {
	var sb strings.Builder
	for node := list.head; node != nil; node = node.next {
		sb.WriteString(fmt.Sprintf("%v", node.Value))
		sb.WriteString("->")
	}

	sb.WriteString("nil")
	return sb.String()
}
