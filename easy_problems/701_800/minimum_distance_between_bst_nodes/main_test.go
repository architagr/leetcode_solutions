package minimumdistancebetweenbstnodes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	root     *TreeNode
	expected int
}

func TestMinDiffInBST(t *testing.T) {
	testcases := []testcase{
		{
			root: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 6,
				},
			},

			expected: 1,
		},
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 0,
				},
				Right: &TreeNode{
					Val: 48,
					Left: &TreeNode{
						Val: 12,
					},
					Right: &TreeNode{
						Val:  49,
						Left: &TreeNode{Val: 0},
					},
				},
			},
			expected: 1,
		},
		{
			root: &TreeNode{
				Val: 90,
				Left: &TreeNode{
					Val: 69,
					Left: &TreeNode{
						Val: 49,
						Right: &TreeNode{
							Val: 52,
						},
					},
					Right: &TreeNode{Val: 89},
				},
			},
			expected: 1,
		},
	}
	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := minDiffInBST(tc.root)
			assert.Equal(t, tc.expected, got)
		})
	}
}
