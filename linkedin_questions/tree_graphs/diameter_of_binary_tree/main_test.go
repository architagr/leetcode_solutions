package diameterofbinarytree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	root     *TreeNode
	expected int
}

func TestDiameterOfBinaryTree(t *testing.T) {
	testcases := []testcase{
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
			expected: 3,
		},
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
			},
			expected: 1,
		},
		{
			root: &TreeNode{
				Val: 1,
			},
			expected: 0,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, diameterOfBinaryTree(tc.root))
		})
	}
}
