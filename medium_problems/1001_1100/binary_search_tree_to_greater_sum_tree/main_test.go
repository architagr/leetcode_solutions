package binarysearchtreetogreatersumtree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	root     *TreeNode
	expected *TreeNode
}

func TestLevelOrder(t *testing.T) {
	testcases := []testcase{
		{
			root: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 0,
					},
					Right: &TreeNode{
						Val: 2,
						Right: &TreeNode{
							Val: 3,
						},
					},
				},
				Right: &TreeNode{
					Val: 6,
					Left: &TreeNode{
						Val: 5,
					},
					Right: &TreeNode{
						Val: 7,
						Right: &TreeNode{
							Val: 8,
						},
					},
				},
			},
			expected: &TreeNode{
				Val: 30,
				Left: &TreeNode{
					Val: 36,
					Left: &TreeNode{
						Val: 36,
					},
					Right: &TreeNode{
						Val: 35,
						Right: &TreeNode{
							Val: 33,
						},
					},
				},
				Right: &TreeNode{
					Val: 21,
					Left: &TreeNode{
						Val: 26,
					},
					Right: &TreeNode{
						Val: 15,
						Right: &TreeNode{
							Val: 8,
						},
					},
				},
			},
		},
	}
	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := bstToGst(tc.root)
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
