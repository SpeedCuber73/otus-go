package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// not full coverage test
func TestList(t *testing.T) {
	assert := assert.New(t)

	l := NewList() // []
	assert.Equal(0, l.Len())
	assert.Nil(l.First())
	assert.Nil(l.Last())

	l.Remove(&Item{}) // []
	assert.Equal(0, l.Len())

	one := "one"
	l.PushBack(one) // ["one"]
	assert.Equal(1, l.Len())
	assert.Equal(one, l.First().Value())
	assert.Equal(one, l.Last().Value())

	two := "two"
	l.PushFront(two) // ["two", "one"]
	assert.Equal(2, l.Len())
	assert.Equal(two, l.First().Value())
	assert.Equal(one, l.Last().Value())

	item := l.Last()
	l.Remove(item) // ["two"]
	assert.Equal(1, l.Len())
	assert.Equal(two, l.First().Value())
	assert.Equal(two, l.Last().Value())
}
