package constructbinarytreefrompreorderandpostordertraversal

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	preOrder, postOrder []int
	expected            *TreeNode
}

func TestConstructFromPrePost(t *testing.T) {
	testcases := []testcase{
		{
			preOrder:  []int{2, 1},
			postOrder: []int{1, 2},
			expected: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 1},
			},
		},
		{
			preOrder:  []int{1, 2, 4, 5, 3, 6, 7},
			postOrder: []int{4, 5, 2, 6, 7, 3, 1},
			expected: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 6,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
		},
		{
			preOrder:  []int{1},
			postOrder: []int{1},
			expected: &TreeNode{
				Val: 1,
			},
		},
	}
	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := constructFromPrePost(tc.preOrder, tc.postOrder)
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
