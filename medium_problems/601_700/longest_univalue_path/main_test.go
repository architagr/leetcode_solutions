package longestunivaluepath

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	root     *TreeNode
	expected int
}

func TestLongestUnivaluePath(t *testing.T) {
	testcases := []testcase{
		{
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 1,
					},
				},
				Right: &TreeNode{
					Val: 5,
					Right: &TreeNode{
						Val: 5,
					},
				},
			},
			expected: 2,
		},
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 5,
					Right: &TreeNode{
						Val: 5,
					},
				},
			},
			expected: 2,
		},
	}
	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := longestUnivaluePath(tc.root)
			assert.Equal(t, tc.expected, got)
		})
	}
}
