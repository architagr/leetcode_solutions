package validatebinarytreenodes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	n           int
	left, right []int
	expected    bool
}

func TestValidateBinaryTreeNodes(t *testing.T) {
	cases := []testcase{
		{4, []int{1, -1, 3, -1}, []int{2, -1, -1, -1}, true},
		{4, []int{1, -1, 3, -1}, []int{2, 3, -1, -1}, false},
		{2, []int{1, 0}, []int{-1, -1}, false},
		{6, []int{1, -1, -1, 4, -1, -1}, []int{2, -1, -1, 5, -1, -1}, false},
		{4, []int{3, -1, 1, -1}, []int{-1, -1, 0, -1}, true},
		{4, []int{1, 0, 3, -1}, []int{-1, -1, -1, -1}, false},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprint(i), func(tb *testing.T) {
			assert.Equal(tb, tc.expected, validateBinaryTreeNodes(tc.n, tc.left, tc.right))
		})
	}
}
