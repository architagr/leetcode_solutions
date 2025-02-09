package binarytreezigzaglevelordertraversal

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
			expected: [][]int{{3}, {20, 9}, {15, 7}},
		},
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
			},
			expected: [][]int{{1}, {3, 2}, {4, 5}},
		},
	}
	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := zigzagLevelOrder(tc.root)
			assert.Equal(t, tc.expected, got)
		})
	}
}
