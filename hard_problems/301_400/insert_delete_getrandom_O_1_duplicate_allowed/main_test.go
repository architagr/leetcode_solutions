package insertdeletegetrandomo1duplicateallowed

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomizedCollection_1(t *testing.T) {
	obj := Constructor()
	assert.Equal(t, true, obj.Insert(1))
	assert.Equal(t, false, obj.Insert(1))
	assert.Equal(t, true, obj.Insert(2))
	t.Log(obj.GetRandom())
	assert.Equal(t, true, obj.Remove(1))
	t.Log(obj.GetRandom())
}

func TestRandomizedCollection_2(t *testing.T) {
	obj := Constructor()
	assert.Equal(t, true, obj.Insert(1))
	assert.Equal(t, false, obj.Insert(1))
	assert.Equal(t, true, obj.Remove(1))
	t.Log(obj.GetRandom())
}

func TestRandomizedCollection_3(t *testing.T) {
	obj := Constructor()
	assert.Equal(t, true, obj.Insert(1))
	assert.Equal(t, true, obj.Remove(1))
	assert.Equal(t, true, obj.Insert(1))
}
