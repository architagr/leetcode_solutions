package sumroottoleafnumbers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	root     *TreeNode
	expected int
}

func TestSumNumbers(t *testing.T) {
	testcases := []testcase{
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			expected: 25,
		},
		{
			root: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 9,
					Left: &TreeNode{
						Val: 5,
					},
					Right: &TreeNode{Val: 1},
				},
				Right: &TreeNode{
					Val: 0,
				},
			},
			expected: 1026,
		},
		{
			root: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 9,
					Left: &TreeNode{
						Val: 5,
						Left: &TreeNode{
							Val: 1,
						},
					},
					Right: &TreeNode{Val: 1},
				},
				Right: &TreeNode{
					Val: 0,
				},
			},
			expected: 4951 + 491 + 40,
		},
	}
	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := sumNumbers(tc.root)
			assert.Equal(t, tc.expected, got)
		})
	}
}
