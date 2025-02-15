package closestbinarysearchtreevalueii

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	root     *TreeNode
	target   float64
	k        int
	expected []int
}

func TestClosestKValues(t *testing.T) {
	testcases := []testcase{

		{
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
					},
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
			target:   2.0,
			k:        3,
			expected: []int{2, 1, 3},
		},
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
					Val: 5,
				},
			},
			target:   3.714286,
			k:        2,
			expected: []int{4, 3},
		},
		{
			root: &TreeNode{
				Val: 1,
			},
			target:   0.0,
			k:        1,
			expected: []int{1},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := closestKValues(tc.root, tc.target, tc.k)
			// assert.Equal(t, tc.expected, got)
			assert.ElementsMatch(t, tc.expected, got)
		})
	}
}
