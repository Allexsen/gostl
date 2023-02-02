package bilist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Add Tests
func TestList(t *testing.T) {
	list := New[int]()
	assert.True(t, list.Empty())
	list.PushBack(1)

	assert.Equal(t, 1, list.Head().Value)
	assert.Equal(t, 1, list.Tail().Value)

	list.PushFront(2)
	assert.Equal(t, 2, list.Len())

	list.PushBack(3)
	list.PushFront(4)

	assert.Equal(t, 4, list.Head().Value)
	assert.Equal(t, 3, list.Tail().Value)

	t.Logf("list: %v", list)
	list.Remove(list.Head())

	t.Logf("list: %v", list)
	assert.Equal(t, "[nil<-2<->1<->3->nil]", list.String())

	list.Remove(list.Tail())
	assert.Equal(t, "[nil<-2<->1->nil]", list.String())

	list.PushBack(5)
	list.PushBack(6)
	list.InsertAfter(7, list.Head())
	t.Logf("list: %v", list)
	assert.Equal(t, "[nil<-2<->7<->1<->5<->6->nil]", list.String())

	list.InsertBefore(8, list.Tail().Prev())

	assert.Equal(t, "[nil<-2<->7<->1<->8<->5<->6->nil]", list.String())

	list.Remove(list.Head().Next().Next())
	assert.Equal(t, "[nil<-2<->7<->8<->5<->6->nil]", list.String())

	list.PopFront()
	assert.Equal(t, "[nil<-7<->8<->5<->6->nil]", list.String())

	list.PopBack()
	assert.Equal(t, "[nil<-7<->8<->5->nil]", list.String())

	list.PushBack(2)
	list.PushBack(1)
	assert.Equal(t, "[nil<-7<->8<->5<->2<->1->nil]", list.String())

	list.SetTail(list.Head().Next())
	assert.Equal(t, "[nil<-7<->5<->2<->1<->8->nil]", list.String())

	list.SetHead(list.Head().Next())
	assert.Equal(t, "[nil<-5<->7<->2<->1<->8->nil]", list.String())

	list.SetHead(list.Tail())
	assert.Equal(t, "[nil<-8<->5<->7<->2<->1->nil]", list.String())

	list.SetTail(list.Head())
	assert.Equal(t, "[nil<-5<->7<->2<->1<->8->nil]", list.String())

	list.SetHead(list.Head())
	assert.Equal(t, "[nil<-5<->7<->2<->1<->8->nil]", list.String())

	list.SetTail(list.Tail())
	assert.Equal(t, "[nil<-5<->7<->2<->1<->8->nil]", list.String())
}
