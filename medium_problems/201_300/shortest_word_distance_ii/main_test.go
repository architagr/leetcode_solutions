package shortestworddistanceii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortestWordDistanceII_1(t *testing.T) {
	obj := Constructor([]string{"practice", "makes", "perfect", "coding", "makes"})
	assert.Equal(t, 3, obj.Shortest("coding", "practice"))
	assert.Equal(t, 1, obj.Shortest("makes", "coding"))
}

func TestShortestWordDistanceII_2(t *testing.T) {
	obj := Constructor([]string{"a", "c", "b", "b", "a"})
	assert.Equal(t, 1, obj.Shortest("a", "b"))
	assert.Equal(t, 1, obj.Shortest("b", "a"))
}
func TestShortestWordDistanceII_3(t *testing.T) {
	obj := Constructor([]string{"a", "a", "b", "b"})
	assert.Equal(t, 1, obj.Shortest("a", "b"))
	assert.Equal(t, 1, obj.Shortest("b", "a"))
}
