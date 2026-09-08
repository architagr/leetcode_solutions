package sumofnodeswithevenvaluedgrandparent

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	root     *TreeNode
	expected int
}

func TestSumEvenGrandparent(t *testing.T) {
	testcases := []testcase{
		{
			root: &TreeNode{
				Val: 1,
			},
			expected: 0,
		},
		{
			root: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val: 7,
					Left: &TreeNode{
						Val:  2,
						Left: &TreeNode{Val: 9},
					},
					Right: &TreeNode{
						Val: 7,
						Left: &TreeNode{
							Val: 1,
						},
						Right: &TreeNode{
							Val: 4,
						},
					},
				},
				Right: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val:   3,
						Right: &TreeNode{Val: 5},
					},
				},
			},
			expected: 18,
		},
	}
	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := sumEvenGrandparent(tc.root)
			assert.Equal(t, tc.expected, got)
		})
	}
}
