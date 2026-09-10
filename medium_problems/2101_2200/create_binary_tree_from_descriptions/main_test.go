package createbinarytreefromdescriptions

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	input    [][]int
	expected *TreeNode
}

func TestCreateBinaryTree(t *testing.T) {
	testcases := []testcase{
		{
			input: [][]int{
				{39, 70, 1},
				{13, 39, 1},
				{85, 74, 1},
				{74, 13, 1},
				{38, 82, 1},
				{82, 85, 1},
			},
			expected: &TreeNode{
				Val: 38,
				Left: &TreeNode{
					Val: 82,
					Left: &TreeNode{
						Val: 85,
						Left: &TreeNode{
							Val: 74,
							Left: &TreeNode{
								Val: 13,
								Left: &TreeNode{
									Val: 39,
									Left: &TreeNode{
										Val: 70,
									},
								},
							},
						},
					},
				},
			},
		},
		{
			input: [][]int{
				{20, 15, 1},
				{20, 17, 0},
				{50, 20, 1},
				{50, 80, 0},
				{80, 19, 1},
			},
			expected: &TreeNode{
				Val: 50,
				Left: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 17,
					},
				},
				Right: &TreeNode{
					Val: 80,
					Left: &TreeNode{
						Val: 19,
					},
				},
			},
		},
		{
			input: [][]int{
				{1, 2, 1},
				{2, 3, 0},
				{3, 4, 1},
			},
			expected: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val:  3,
						Left: &TreeNode{Val: 4},
					},
				},
			},
		},
	}
	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := createBinaryTree(tc.input)
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
