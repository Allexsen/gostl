package vector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVectorBase(t *testing.T) {
	v := New[int]()
	assert.True(t, v.Empty())
	v.PushBack(1)
	v.PushBack(2)

	assert.False(t, v.Empty())
	assert.Equal(t, 2, v.Size())

	assert.Equal(t, 1, v.Front())
	assert.Equal(t, 2, v.Back())
}

func TestModifyVector(t *testing.T) {
	v := New[int]()
	v.PushBack(1)
	v.PushBack(2)
	v.PushBack(3)
	//[1,2,3]
	assert.Equal(t, 3, v.PopBack())
	//[1 2]
	v.PushBack(4)
	//[1 2 4]

	v.Set(1, 9)
	assert.Equal(t, 9, v.At(1))
	//[1 9 4]

	v.Insert(0, 8)
	//[8 1 9 4]

	assert.Equal(t, "[8 1 9 4]", v.String())
	v.Clear()
}
