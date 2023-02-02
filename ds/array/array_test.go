package array

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestArray(t *testing.T) {
	a := New[int](10)
	assert.Equal(t, 10, a.Size())
	assert.Equal(t, false, a.Empty())

	va := 10
	a.Fill(va)
	for i := 0; i < a.Size(); i++ {
		val := a.At(i)
		assert.Equal(t, va, val)
	}

	b := New[int](10)
	vb := 66
	b.Fill(vb)
	a.SwapArrays(b)

	for i := 0; i < a.Size(); i++ {
		assert.Equal(t, vb, a.At(i))
		assert.Equal(t, va, b.At(i))
	}

	for i := 0; i < a.Size(); i++ {
		a.Set(i, i)
	}

	neg := a.At(-1)
	assert.Equal(t, neg, a.Back())
}

func TestNewFromArray(t *testing.T) {
	a := New[int](10)
	for i := 0; i < 10; i++ {
		a.Set(i, i*10)
	}
	b := CreateCopy(a)

	assert.Equal(t, a.Size(), b.Size())
	for i := 0; i < 10; i++ {
		assert.Equal(t, a.At(i), b.At(i))
	}
}
