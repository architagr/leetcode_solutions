package deleteleaveswithagivenvalue

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	root     *TreeNode
	target   int
	expected *TreeNode
}

func TestRemoveLeafNodes(t *testing.T) {
	testcases := []testcase{
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 2},
				},
				Right: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 2},
					Right: &TreeNode{Val: 4},
				},
			},
			target: 2,
			expected: &TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}},
			},
		},
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 2},
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			target: 3,
			expected: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   3,
					Right: &TreeNode{Val: 2},
				},
			},
		},
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 2,
						},
					},
				},
			},
			target: 2,
			expected: &TreeNode{
				Val: 1,
			},
		},
		{
			root: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 2,
						},
					},
				},
			},
			target:   2,
			expected: nil,
		},
	}
	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := removeLeafNodes(tc.root, tc.target)
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
