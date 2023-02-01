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
	assert.Equal(t, "[2 1]", list.String())

	list.PushBack(3)
	list.PushBack(4)
	assert.Equal(t, "[2 1 3 4]", list.String())

	list.MakeHead(list.Head(), list.Head().Next())

	assert.Equal(t, "[1 2 3 4]", list.String())

	list.MakeTail(list.Head(), list.Head().Next())

	assert.Equal(t, "[1 3 4 2]", list.String())
}

func TestList_InsertAfter(t *testing.T) {
	list := New[int]()
	for i := 1; i <= 5; i++ {
		list.PushBack(i)
	}
	list.InsertAfter(6, list.Head())
	assert.Equal(t, "[1 6 2 3 4 5]", list.String())

	list.InsertAfter(7, list.Head().Next())
	assert.Equal(t, "[1 6 7 2 3 4 5]", list.String())

	list.InsertAfter(8, list.Tail())
	assert.Equal(t, "[1 6 7 2 3 4 5 8]", list.String())
}

func TestList_Remove(t *testing.T) {
	list := New[int]()
	for i := 1; i <= 5; i++ {
		list.PushBack(i)
	}
	list.Remove(nil, list.Head())
	assert.Equal(t, "[2 3 4 5]", list.String())

	list.Remove(list.Head(), list.Head().Next())
	assert.Equal(t, "[2 4 5]", list.String())

	list.PushFront(6)
}
