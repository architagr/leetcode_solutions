package cousionsinbinarytree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	root     *TreeNode
	x, y     int
	expected bool
}

func TestIsCousins(t *testing.T) {
	testcases := []testcase{
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			x:        4,
			y:        3,
			expected: false,
		},
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 5,
					},
				},
			},
			x:        5,
			y:        4,
			expected: true,
		},
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			x:        2,
			y:        3,
			expected: false,
		},
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,

					Right: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 5,
						},
					},
				},
			},
			x:        1,
			y:        2,
			expected: false,
		},
	}
	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := isCousins(tc.root, tc.x, tc.y)
			assert.Equal(t, tc.expected, got)
		})
	}
}
