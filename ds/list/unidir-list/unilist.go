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

// Inserts a new node as a new head
func (list *List[T]) PushFront(val T) {
	node := &Node[T]{Value: val}
	if list.len == 0 { // List is empty
		list.head = node
		list.tail = node
	} else { // Inserts as a new head
		node.next = list.head
		list.head = node
	}

	list.len++
}

// Appends a new node as a new tail
func (list *List[T]) PushBack(val T) {
	node := &Node[T]{Value: val}
	if list.len == 0 { // List is empty
		list.head = node
		list.tail = node
	} else { // Appends at the end
		list.tail.next = node
		list.tail = node
	}

	list.len++
}

// Deletes head
func (list *List[T]) PopFront() *Node[T] {
	return list.Remove(nil, list.head)
}

// Inserts a new node after the given node
func (list *List[T]) InsertAfter(val T, node *Node[T]) *Node[T] { // Returns the inserted node
	newNode := &Node[T]{Value: val}
	if node == nil { // Inserts as a new head
		newNode.next = list.head
		list.head = newNode
	} else { // Inserts somewhere else in the list
		newNode.next = node.next
		node.next = newNode
	}

	if newNode.next == nil { // Newly inserted node is a new tail
		list.tail = newNode
	}

	list.len++
	return newNode
}

// Removes the given node
// prev MUST be right before the node, otherwise list.len will miscount
func (list *List[T]) Remove(prev, node *Node[T]) *Node[T] { // Returns the removed node
	if node == nil { // Node doesn't exist
		return nil
	}

	list.len--
	if list.len == 0 { // Node to remove is the only node
		list.head = nil
		list.tail = nil
	} else if node == list.head { // Node to remove is a head
		list.head = node.next
	} else if node == list.tail { // Node to remove is a tail
		list.tail = prev
		prev.next = nil
	} else { // Node to remove is inbetween two other nodes
		prev.next = node.next
	}

	return node
}

// Sets the given node as a new head
func (list *List[T]) SetHead(prev, node *Node[T]) {
	if node == nil || prev == nil || prev.next != node {
		return
	}

	prev.next = node.next
	if prev.next == nil { // Node is a tail
		list.tail = prev
	}

	node.next = list.head
	list.head = node
}

// Sets the given node as a new tail
func (list *List[T]) SetTail(prev, node *Node[T]) {
	if node == nil || list.len < 2 { // Node doesn't exist or already is a tail
		return
	}

	if prev == nil { // Node to set is a current head
		list.head = node.next
	} else { // Node to set is inbetween two other nodes
		prev.next = node.next
	}

	list.tail.next = node
	list.tail = node
	node.next = nil
}

// Clears the list
func (list *List[T]) Clear() {
	list.head = nil
	list.len = 0
}

// Stringifies the list
func (list *List[T]) String() string {
	var sb strings.Builder
	sb.WriteString("[")
	for node := list.head; node != nil; node = node.next {
		sb.WriteString(fmt.Sprintf("%v", node.Value))
		sb.WriteString("->")
	}

	sb.WriteString("nil]")
	return sb.String()
}
