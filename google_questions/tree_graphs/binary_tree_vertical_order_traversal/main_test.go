package binarytreeverticalordertraversal

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	root     *TreeNode
	expected [][]int
}

func TestMinDepth(t *testing.T) {
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
			expected: [][]int{
				{9},
				{3, 15},
				{20},
				{7},
			},
		},
		{
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 0,
					},
				},
				Right: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			expected: [][]int{
				{4},
				{9},
				{3, 0, 1},
				{8},
				{7},
			},
		},
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
						Right: &TreeNode{
							Val: 5,
							Right: &TreeNode{
								Val: 6,
							},
						},
					},
					Right: &TreeNode{
						Val: 10,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 9,
					},
					Right: &TreeNode{
						Val: 11,
					},
				},
			},
			expected: [][]int{
				{4},
				{2, 5},
				{1, 10, 9, 6},
				{3},
				{11},
			},
		},
	}
	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, verticalOrder(tc.root))
		})
	}
}
