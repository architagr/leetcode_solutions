package countnodesequaltoaverageofsubtree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	root     *TreeNode
	expected int
}

func TestAverageOfSubtree(t *testing.T) {
	testcases := []testcase{
		{
			root: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: 0,
					},
					Right: &TreeNode{
						Val: 1,
					},
				},
				Right: &TreeNode{
					Val: 5,
					Right: &TreeNode{
						Val: 6,
					},
				},
			},
			expected: 5,
		},
		{
			root: &TreeNode{
				Val: 1,
			},
			expected: 1,
		},
	}
	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := averageOfSubtree(tc.root)
			assert.Equal(t, tc.expected, got)
		})
	}
}
