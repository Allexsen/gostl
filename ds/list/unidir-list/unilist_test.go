package unilist

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestList(t *testing.T) {
	list := New[int]()
	assert.Equal(t, 0, list.Len())
	list.PushBack(1)
	assert.Equal(t, 1, list.Len())
	assert.Equal(t, 1, list.Head().Value)
	assert.Equal(t, 1, list.Tail().Value)
	list.PushFront(2)

	assert.Equal(t, 2, list.Len())
	assert.Equal(t, list.String(), "[2->1->nil]")

	list.PushBack(3)
	list.PushBack(4)
	assert.Equal(t, list.String(), "[2->1->3->4->nil]")

	list.SetHead(list.Head(), list.Head().Next())

	assert.Equal(t, list.String(), "[1->2->3->4->nil]")

	list.SetTail(list.Head(), list.Head().Next())

	assert.Equal(t, list.String(), "[1->3->4->2->nil]")
}

func TestList_InsertAfter(t *testing.T) {
	list := New[int]()
	for i := 1; i <= 5; i++ {
		list.PushBack(i)
	}
	list.InsertAfter(6, list.Head())
	assert.Equal(t, list.String(), "[1->6->2->3->4->5->nil]")

	list.InsertAfter(7, list.Head().Next())
	assert.Equal(t, list.String(), "[1->6->7->2->3->4->5->nil]")

	list.InsertAfter(8, list.Tail())
	assert.Equal(t, list.String(), "[1->6->7->2->3->4->5->8->nil]")
}

func TestList_Remove(t *testing.T) {
	list := New[int]()
	for i := 1; i <= 5; i++ {
		list.PushBack(i)
	}

	list.Remove(nil, list.Head())
	assert.Equal(t, list.String(), "[2->3->4->5->nil]")

	list.Remove(list.Head(), list.Head().Next())
	assert.Equal(t, list.String(), "[2->4->5->nil]")

	list.PushFront(6)
}

func TestList_PopFront(t *testing.T) {
	list := New[int]()
	list.PushBack(1)
	list.PushBack(2)

	list.PopFront()
	assert.Equal(t, list.String(), "[2->nil]")

	list.PushFront(3)
	assert.Equal(t, list.String(), "[3->2->nil]")

	list.PopFront()
	list.PopFront()
	assert.Equal(t, list.String(), "[nil]")
}
