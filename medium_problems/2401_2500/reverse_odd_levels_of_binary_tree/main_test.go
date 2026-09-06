package reverseoddlevelsofbinarytree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	root     *TreeNode
	expected *TreeNode
}

func TestReverseOddLevels(t *testing.T) {
	testcases := []testcase{
		{
			root: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 8,
					},
					Right: &TreeNode{
						Val: 13,
					},
				},
				Right: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 21,
					},
					Right: &TreeNode{
						Val: 34,
					},
				},
			},
			expected: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 8,
					},
					Right: &TreeNode{
						Val: 13,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 21,
					},
					Right: &TreeNode{
						Val: 34,
					},
				},
			},
		},
		{
			root: &TreeNode{
				Val: 7,
				Left: &TreeNode{
					Val: 13,
				},
				Right: &TreeNode{
					Val: 11,
				},
			},
			expected: &TreeNode{
				Val: 7,
				Left: &TreeNode{
					Val: 11,
				},
				Right: &TreeNode{
					Val: 13,
				},
			},
		},
	}
	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := reverseOddLevels(tc.root)
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
