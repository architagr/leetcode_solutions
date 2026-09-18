package maximumdifferencebetweennodesandancestor

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	root     *TreeNode
	expected int
}

func TestMaxAncestorDiff(t *testing.T) {
	testcases := []testcase{
		{
			root: &TreeNode{
				Val: 8,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 6,
						Left: &TreeNode{
							Val: 4,
						},
						Right: &TreeNode{
							Val: 7,
						},
					},
				},
				Right: &TreeNode{
					Val: 10,
					Right: &TreeNode{
						Val: 14,
						Left: &TreeNode{
							Val: 13,
						},
					},
				},
			},
			expected: 7,
		},
		{
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 0,
						Left: &TreeNode{
							Val: 3,
						},
					},
				},
			},
			expected: 3,
		},
		{
			root: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val:  3,
					Left: &TreeNode{Val: 0},
					Right: &TreeNode{
						Val: 5,
						Left: &TreeNode{
							Val: 7,
						},
					},
				},
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val:   1,
						Right: &TreeNode{Val: 6},
					},
				},
			},
			expected: 5,
		},
	}
	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := maxAncestorDiff(tc.root)
			assert.Equal(t, tc.expected, got)
		})
	}

}
