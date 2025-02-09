package binarytreelevelordertraversal

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	root     *TreeNode
	expected [][]int
}

func TestLevelOrder(t *testing.T) {
	testcases := []testcase{
		{
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			expected: [][]int{{3}, {9, 20}, {15, 7}},
		},
	}
	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := levelOrder(tc.root)
			assert.Equal(t, tc.expected, got)
		})
	}
}
