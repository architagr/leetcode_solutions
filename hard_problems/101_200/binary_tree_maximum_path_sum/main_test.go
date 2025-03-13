package binarytreemaximumpathsum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	root     *TreeNode
	expected int
}

func TestMaxPathSum(t *testing.T) {
	testcases := []testcase{
		{
			root: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: -1,
				},
			},
			expected: 2,
		},
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			expected: 6,
		},
		{
			root: &TreeNode{
				Val: -10,
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
			expected: 42,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, maxPathSum(tc.root))
		})
	}
}
