package maxstack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxStack1(t *testing.T) {
	obj := Constructor()
	obj.Push(5)
	obj.Push(1)
	obj.Push(5)
	assert.Equal(t, 5, obj.Top())
	assert.Equal(t, 5, obj.PopMax())
	assert.Equal(t, 1, obj.Top())
	assert.Equal(t, 5, obj.PeekMax())
	assert.Equal(t, 1, obj.Pop())
	assert.Equal(t, 5, obj.Top())
}
