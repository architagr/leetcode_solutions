package narytreelevelordertraversal

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	root     *Node
	expected [][]int
}

func TestLevelOrder(t *testing.T) {
	testcases := []testcase{
		{
			root: &Node{
				Val: 1,
				Children: []*Node{
					{
						Val: 3,
						Children: []*Node{
							{
								Val: 5,
							},
							{
								Val: 6,
							},
						},
					},
					{
						Val: 2,
					},
					{
						Val: 4,
					},
				},
			},
			expected: [][]int{
				{1},
				{3, 2, 4},
				{5, 6},
			},
		},
		{
			root: &Node{
				Val: 1,
				Children: []*Node{
					{
						Val: 2,
					},
					{
						Val: 3,
						Children: []*Node{
							{
								Val: 6,
							},
							{
								Val: 7,
								Children: []*Node{
									{
										Val: 11,
										Children: []*Node{
											{Val: 14},
										},
									},
								},
							},
						},
					},
					{
						Val: 4,
						Children: []*Node{
							{
								Val: 8,
								Children: []*Node{
									{Val: 12},
								},
							},
						},
					},
					{
						Val: 5,
						Children: []*Node{
							{
								Val: 9,
								Children: []*Node{
									{Val: 13},
								},
							},
							{
								Val: 10,
							},
						},
					},
				},
			},
			expected: [][]int{
				{1}, {2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13}, {14},
			},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, levelOrder(tc.root))
		})
	}
}
