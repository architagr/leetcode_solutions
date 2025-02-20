package minstack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinStack(t *testing.T) {
	minStackObj := Constructor()
	minStackObj.Push(-2)
	minStackObj.Push(0)
	minStackObj.Push(-3)
	assert.Equal(t, -3, minStackObj.GetMin())
	minStackObj.Pop()
	assert.Equal(t, 0, minStackObj.Top())
	assert.Equal(t, -2, minStackObj.GetMin())
}
