package lrucache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLru_TC1(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	assert.Equal(t, 1, cache.Get(1))
	cache.Put(3, 3)
	assert.Equal(t, -1, cache.Get(2))
	cache.Put(4, 4)
	assert.Equal(t, -1, cache.Get(1))
	assert.Equal(t, 3, cache.Get(3))
	assert.Equal(t, 4, cache.Get(4))
}

func TestLru_TC2(t *testing.T) {
	cache := Constructor(2)
	cache.Put(2, 1)
	cache.Put(2, 2)
	assert.Equal(t, 2, cache.Get(2))
	cache.Put(1, 1)
	cache.Put(4, 1)
	assert.Equal(t, -1, cache.Get(2))
}
