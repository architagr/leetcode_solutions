package allelementsintwobinarysearchtrees

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	root1, root2 *TreeNode
	expected     []int
}

func TestGetAllElements(t *testing.T) {
	testcases := []testcase{
		{
			root1: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
			root2: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 0,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			expected: []int{0, 1, 1, 2, 3, 4},
		},
	}
	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := getAllElements(tc.root1, tc.root2)
			assert.EqualValues(t, tc.expected, got)
		})
	}
}
