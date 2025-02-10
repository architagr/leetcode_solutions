package alloonedatastructure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllOne(t *testing.T) {
	obj := Constructor()
	obj.Inc("hello")
	obj.Inc("hello")
	assert.Equal(t, "hello", obj.GetMaxKey())
	assert.Equal(t, "hello", obj.GetMinKey())
}
