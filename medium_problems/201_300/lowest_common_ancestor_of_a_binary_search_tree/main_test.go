package lowestcommonancestorofabinarysearchtree

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
				Val: 6,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 0,
					},
					Right: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 3,
						},
						Right: &TreeNode{
							Val: 5,
						},
					},
				},
				Right: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: 7,
					},
					Right: &TreeNode{
						Val: 9,
					},
				},
			},
			p: &TreeNode{
				Val: 2,
			},
			q: &TreeNode{
				Val: 8,
			},
			expected: &TreeNode{
				Val: 6,
			},
		},
		{
			root: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 0,
					},
					Right: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 3,
						},
						Right: &TreeNode{
							Val: 5,
						},
					},
				},
				Right: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: 7,
					},
					Right: &TreeNode{
						Val: 9,
					},
				},
			},
			p: &TreeNode{
				Val: 2,
			},
			q: &TreeNode{
				Val: 4,
			},
			expected: &TreeNode{
				Val: 2,
			},
		},
	}
	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := lowestCommonAncestor(tc.root, tc.p, tc.q)
			if assert.NotNil(t, got) {
				assert.Equal(t, tc.expected.Val, got.Val)
			}
		})
	}
}
