package lowestcommonancestorofdeepestleaves

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	root     *TreeNode
	expected *TreeNode
}

func TestLcaDeepestLeaves1(t *testing.T) {
	expected := &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 7},
		Right: &TreeNode{Val: 4},
	}
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   5,
			Left:  &TreeNode{Val: 6},
			Right: expected,
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 8,
			},
		},
	}

	got := lcaDeepestLeaves(root)
	assert.Equal(t, expected, got)
	assert.True(t, equalBinaryTree(expected, got))
}

func TestLcaDeepestLeaves2(t *testing.T) {
	expected := &TreeNode{
		Val: 1,
	}
	root := expected

	got := lcaDeepestLeaves(root)
	assert.Equal(t, expected, got)
	assert.True(t, equalBinaryTree(expected, got))
}

func TestLcaDeepestLeaves3(t *testing.T) {
	expected := &TreeNode{
		Val: 2,
	}
	root := &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val:  1,
			Left: expected,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	got := lcaDeepestLeaves(root)
	assert.Equal(t, expected, got)
	assert.True(t, equalBinaryTree(expected, got))
}

func TestLcaDeepestLeaves4(t *testing.T) {
	expected := &TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val:  8,
			Left: &TreeNode{Val: 10},
		},
		Right: &TreeNode{
			Val:  7,
			Left: &TreeNode{Val: 9},
		},
	}
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  expected,
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{
			Val:  4,
			Left: &TreeNode{Val: 5},
		},
	}

	got := lcaDeepestLeaves(root)
	assert.Equal(t, expected, got)
	assert.True(t, equalBinaryTree(expected, got))
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
