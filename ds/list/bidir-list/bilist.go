package bilist

import (
	"fmt"
	"strings"
)

// List Node
type Node[T any] struct {
	next  *Node[T] // Next node
	prev  *Node[T] // Previous node
	Value T
}

func (node *Node[T]) Next() *Node[T] {
	return node.next
}

func (node *Node[T]) Prev() *Node[T] {
	return node.prev
}

// Bidirectional(Doubly) Linked List
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

func (list *List[T]) Empty() bool {
	return list.len == 0
}

func (list *List[T]) Len() int {
	return list.len
}

// Inserts a new node as a new head
func (list *List[T]) PushFront(val T) {
	newNode := &Node[T]{Value: val}
	if list.len == 0 { // List is empty
		list.head = newNode
		list.tail = newNode
	} else { // node becomes new head
		newNode.next = list.head
		newNode.next.prev = newNode
		list.head = newNode
	}

	list.len++
}

// Appends a new node as a new tail
func (list *List[T]) PushBack(val T) {
	newNode := &Node[T]{Value: val}
	if list.len == 0 { // List is empty
		list.head = newNode
		list.tail = newNode
	} else { // node becomes new tail
		newNode.prev = list.tail
		newNode.prev.next = newNode
		list.tail = newNode
	}

	list.len++
}

// Deletes head
func (list *List[T]) PopFront() *Node[T] {
	return list.Remove(list.head)
}

// Deletes tail
func (list *List[T]) PopBack() *Node[T] {
	return list.Remove(list.tail)
}

// Inserts a new node after the given node
func (list *List[T]) InsertAfter(val T, node *Node[T]) *Node[T] {
	newNode := &Node[T]{Value: val}
	if list.len == 0 { // List is empty
		list.head = newNode
		list.tail = newNode
	} else if node == nil { // node becomes new head
		newNode.next = list.head
		newNode.next.prev = newNode
		list.head = newNode
	} else if node.next == nil { // node becomes new tail
		newNode.prev = node
		newNode.prev.next = newNode
		list.tail = newNode
	} else { // node is inbetween two other nodes
		newNode.next = node.next
		newNode.prev = node
		newNode.prev.next = newNode
		newNode.next.prev = newNode
	}

	list.len++
	return newNode
}

// Insert before the given node
func (list *List[T]) InsertBefore(val T, node *Node[T]) *Node[T] {
	var newNode *Node[T]
	if node == nil { // node becomes new tail
		newNode = list.InsertAfter(val, list.tail)
	} else { // node might be a new head or inbetween two other nodes
		newNode = list.InsertAfter(val, node.prev)
	}

	// list.len is increased in 'list.InsertAfter()' method
	return newNode
}

// Delete the given node
func (list *List[T]) Remove(node *Node[T]) *Node[T] { // Returns the removed node
	if node == nil { // Node doesn't exist
		return nil
	}

	list.len--
	if list.len == 1 { // Node to remove is the only node
		list.head = nil
		list.tail = nil
	} else if list.head == node { // Node to remove is a head
		node = node.next
		list.head = node
		node.prev = nil
	} else if list.tail == node { // Node to remove isa  tail
		node = node.prev
		list.tail = node
		node.next = nil
	} else { // Node to remove is inbetween two other nodes
		node.next.prev = node.prev
		node.prev.next = node.next
	}

	return node
}

// Set the given node as a new head
func (list *List[T]) SetHead(node *Node[T]) {
	if node == nil || node.prev == nil { // Node doesn't exist or it's a head already
		return
	}

	node.prev.next = node.next
	if node.next == nil { // Node is a tail
		list.tail = node.prev
	} else { // Node is inbetween two other nodes
		node.next.prev = node.prev
	}

	list.head.prev = node
	node.next = list.head
	node.prev = nil
	list.head = node
}

// Set the given node as a new tail
func (list *List[T]) SetTail(node *Node[T]) {
	if node == nil || node.next == nil { // Node doesn't exist or it's a tail already
		return
	}

	node.next.prev = node.prev
	if node.prev == nil { // Node is a head
		list.head = node.next
	} else { // Node is inbetween two other nodes
		node.prev.next = node.next
	}

	list.tail.next = node
	node.next = nil
	node.prev = list.tail
	list.tail = node
}

// Stringify the list
func (list *List[T]) String() string {
	var sb strings.Builder
	sb.WriteString("nil->")
	for node := list.head; node != nil; node = node.next {
		sb.WriteString(fmt.Sprintf("%v", node.Value))
		if node.next != nil {
			sb.WriteString("<->")
		}
	}

	sb.WriteString("->nil")
	return sb.String()
}
