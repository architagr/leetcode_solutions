package lowestcommonancestorofabinarytree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	root, p, q *TreeNode
	expected   *TreeNode
}

func TestLowestCommonAncestor(t *testing.T) {
	testcases := []testcase{
		{
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 6,
					},
					Right: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 7,
						},
						Right: &TreeNode{
							Val: 4,
						},
					},
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
			},
			p: &TreeNode{
				Val: 5,
			},
			q: &TreeNode{
				Val: 1,
			},
			expected: &TreeNode{
				Val: 3,
			},
		},
		{
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 6,
					},
					Right: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 7,
						},
						Right: &TreeNode{
							Val: 4,
						},
					},
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
			},
			p: &TreeNode{
				Val: 5,
			},
			q: &TreeNode{
				Val: 4,
			},
			expected: &TreeNode{
				Val: 5,
			},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := lowestCommonAncestor(tc.root, tc.p, tc.q)
			assert.Equal(t, tc.expected.Val, got.Val)
		})
	}
}
