package sum_of_distances_in_tree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	n        int
	expected []*TreeNode
}

// LeetCode accepts the trees in any order, so the cases compare the two sets of
// trees rather than the two slices.
func TestAllPossibleFBT(t *testing.T) {
	testcases := []testcase{
		{
			n: 7,
			expected: []*TreeNode{
				{
					Val: 0,
					Left: &TreeNode{
						Val: 0,
					},
					Right: &TreeNode{
						Val: 0,
						Left: &TreeNode{
							Val: 0,
						},
						Right: &TreeNode{
							Val: 0,
							Left: &TreeNode{
								Val: 0,
							},
							Right: &TreeNode{
								Val: 0,
							},
						},
					},
				},
				{
					Val: 0,
					Left: &TreeNode{
						Val: 0,
					},
					Right: &TreeNode{
						Val: 0,
						Left: &TreeNode{
							Val: 0,
							Left: &TreeNode{
								Val: 0,
							},
							Right: &TreeNode{
								Val: 0,
							},
						},
						Right: &TreeNode{
							Val: 0,
						},
					},
				},
				{
					Val: 0,
					Left: &TreeNode{
						Val: 0,
						Left: &TreeNode{
							Val: 0,
						},
						Right: &TreeNode{
							Val: 0,
						},
					},
					Right: &TreeNode{
						Val: 0,
						Left: &TreeNode{
							Val: 0,
						},
						Right: &TreeNode{
							Val: 0,
						},
					},
				},
				{
					Val: 0,
					Left: &TreeNode{
						Val: 0,
						Left: &TreeNode{
							Val: 0,
						},
						Right: &TreeNode{
							Val: 0,
							Left: &TreeNode{
								Val: 0,
							},
							Right: &TreeNode{
								Val: 0,
							},
						},
					},
					Right: &TreeNode{
						Val: 0,
					},
				},
				{
					Val: 0,
					Left: &TreeNode{
						Val: 0,
						Left: &TreeNode{
							Val: 0,
							Left: &TreeNode{
								Val: 0,
							},
							Right: &TreeNode{
								Val: 0,
							},
						},
						Right: &TreeNode{
							Val: 0,
						},
					},
					Right: &TreeNode{
						Val: 0,
					},
				},
			},
		},
		{
			n: 3,
			expected: []*TreeNode{
				{
					Val: 0,
					Left: &TreeNode{
						Val: 0,
					},
					Right: &TreeNode{
						Val: 0,
					},
				},
			},
		},
	}
	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := allPossibleFBT(tc.n)
			assert.ElementsMatch(t, tc.expected, got)
		})
	}
}

// [
// 	[0,0,0,0,0,null,null,0,0],
// 	[0,0,0,0,0,null,null,0,0,0,0],
// 	[0,0,0,null,null,0,0,0,0],
// 	[0,0,0,null,null,0,0,0,0,0,0]
// ]

// [
// 	[0,0,0,null,null,0,0,null,null,0,0],
// 	[0,0,0,null,null,0,0,0,0],
// 	[0,0,0,0,0,0,0],
// 	[0,0,0,0,0,null,null,null,null,0,0],
// 	[0,0,0,0,0,null,null,0,0]
// ]
