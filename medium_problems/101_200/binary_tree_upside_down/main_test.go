package binarytreeupsidedown

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	root     *TreeNode
	expected *TreeNode
}

func TestUpsideDownBinaryTree(t *testing.T) {
	testcases := []testcase{
		{
			root:     nil,
			expected: nil,
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
				},
			},
			expected: &TreeNode{
				Val: 4,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 1,
					},
				},
				Left: &TreeNode{
					Val: 5,
				},
			},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := upsideDownBinaryTree(tc.root)
			assert.Equal(t, inOrdertraversal(tc.expected, []int{}), inOrdertraversal(got, []int{}))
			assert.Equal(t, preOrderTraversal(tc.expected, []int{}), preOrderTraversal(got, []int{}))
		})
	}
}

func inOrdertraversal(head *TreeNode, arr []int) []int {
	if head == nil {
		return arr
	}

	arr = inOrdertraversal(head.Left, arr)
	arr = append(arr, head.Val)
	arr = inOrdertraversal(head.Right, arr)

	return arr
}

func preOrderTraversal(A *TreeNode, arr []int) []int {
	if A == nil {
		return arr
	}
	arr = append(arr, A.Val)
	arr = preOrderTraversal(A.Left, arr)
	arr = preOrderTraversal(A.Right, arr)
	return arr
}
