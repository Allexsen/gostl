package deque

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestPushPop(t *testing.T) {
	q := New[int]()

	q.PushBack(1)  //[1]
	q.PushFront(2) //[2 1]
	q.PushBack(3)  //[2 1 3]

	t.Logf("q: %v", q)

	assert.Equal(t, q.Size(), 3)
	assert.Equal(t, q.Front(), 2)
	assert.Equal(t, q.Back(), 3)

	assert.Equal(t, q.String(), "[2 1 3]")

	assert.Equal(t, q.PopBack(), 3)
	assert.Equal(t, q.String(), "[2 1]")

	assert.Equal(t, q.PopFront(), 5)
	assert.Equal(t, q.String(), "[1]")

	t.Logf("q: %v", q)
}

func TestInsert1(t *testing.T) {
	q := New[int]()
	for i := 0; i < 5; i++ {
		q.PushBack(i)
	}
	assert.Equal(t, "[0 1 2 3 4]", q.String())

	q = New[int]()
	for i := 0; i < 6; i++ {
		q.PushBack(i)
	}
	assert.Equal(t, "[0 1 2 3 4 5]", q.String())
}
