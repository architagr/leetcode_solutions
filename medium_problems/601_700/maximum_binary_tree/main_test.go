package maximumbinarytree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	expected *TreeNode
	nums     []int
}

func TestConstructMaximumBinaryTree(t *testing.T) {
	testcases := []testcase{
		{
			nums: []int{3, 2, 1, 6, 0, 5},
			expected: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 2,
						Right: &TreeNode{
							Val: 1,
						},
					},
				},
				Right: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 0,
					},
				},
			},
		},
		{
			nums: []int{3, 2, 1},
			expected: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 1,
					},
				},
			},
		},
	}
	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := constructMaximumBinaryTree(tc.nums)
			assert.True(t, equalBinaryTree(tc.expected, got))
		})
	}
}

// equalBinaryTree reports whether root and subRoot have the exact same
// shape and node values, walked in lockstep.
func equalBinaryTree(root *TreeNode, subRoot *TreeNode) bool {
	// Both sides empty at the same position -> equal so far.
	if root == nil {
		return subRoot == nil
	}
	// root is empty but subRoot isn't (root == nil is already false here).
	if subRoot == nil {
		return root == nil
	}
	// Values differ at this position: not equal, stop immediately.
	if root.Val != subRoot.Val {
		return false
	}
	// Both children must match, left-to-left and right-to-right.
	return equalBinaryTree(root.Left, subRoot.Left) && equalBinaryTree(root.Right, subRoot.Right)
}
