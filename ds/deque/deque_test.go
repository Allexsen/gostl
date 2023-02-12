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

	assert.Equal(t, q.PopFront(), 2)
	assert.Equal(t, q.String(), "[1]")

	t.Logf("q: %v", q)
}
