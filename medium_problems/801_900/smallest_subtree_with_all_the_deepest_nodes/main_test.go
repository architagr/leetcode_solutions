package smallestsubtreewithallthedeepestnodes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	root     *TreeNode
	expected *TreeNode
}

func TestSubtreeWithAllDeepest(t *testing.T) {
	testcases := []testcase{
		{
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:  5,
					Left: &TreeNode{Val: 6},
					Right: &TreeNode{
						Val:   2,
						Left:  &TreeNode{Val: 7},
						Right: &TreeNode{Val: 4},
					},
				},
				Right: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 0},
					Right: &TreeNode{Val: 8},
				},
			},
			expected: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: 4},
			},
		},
		{
			root: &TreeNode{
				Val: 1,
			},
			expected: &TreeNode{
				Val: 1,
			},
		},
		{
			root: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 2,
					},
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			expected: &TreeNode{
				Val: 2,
			},
		},
	}
	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := subtreeWithAllDeepest(tc.root)
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
