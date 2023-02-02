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

func (list *List[T]) Len() int {
	return list.len
}

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

func (list *List[T]) Remove(node *Node[T]) *Node[T] { // Returns a previous node. If head is removed, returns new head
	if node == nil {
		return nil
	}

	list.len--
	if list.len == 1 { // List becomes empty
		list.head = nil
		list.tail = nil
		node = nil
		return nil
	}

	if list.head == node { // Node to remove is head
		node = node.next
		list.head = node
		node.prev = nil
		return list.head
	}

	if list.tail == node { // Node to remove is tail
		node = node.prev
		list.tail = node
		node.next = nil
		return list.tail
	}

	// Node to remove inbetween two other nodes
	node.next.prev = node.prev
	node.prev.next = node.next
	prevNode := node.prev
	node = nil
	return prevNode
}

func (list *List[T]) SetHead(node *Node[T]) {
	if node == nil || node.prev == nil { // Node doesn't exist or it's a head already
		return
	}

	node.prev.next = node.next
	if node.next == nil { // Node is tail
		list.tail = node.prev
	} else { // Node is inbetween two other nodes
		node.next.prev = node.prev
	}

	list.head.prev = node
	node.next = list.head
	node.prev = nil
	list.head = node
}

func (list *List[T]) SetTail(node *Node[T]) {
	if node == nil || node.next == nil { // Node doesn't exist or it's a tail already
		return
	}

	node.next.prev = node.prev
	if node.prev == nil { // Node is head
		list.head = node.next
	} else { // Node is inbetween two other nodes
		node.prev.next = node.next
	}

	list.tail.next = node
	node.next = nil
	node.prev = list.tail
	list.tail = node
}

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
