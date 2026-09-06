package deepestleavessum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	root     *TreeNode
	expected int
}

func TestDeepestLeavesSum(t *testing.T) {
	testcases := []testcase{
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:  4,
						Left: &TreeNode{Val: 7},
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 6,
						Right: &TreeNode{
							Val: 8,
						},
					},
				},
			},
			expected: 15,
		},
	}
	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := deepestLeavesSum(tc.root)
			assert.Equal(t, tc.expected, got)
		})
	}
}
