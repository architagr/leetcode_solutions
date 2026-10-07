package maximumwidthofbinarytree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	root     *TreeNode
	expected int
}

func TestWidthOfBinaryTree(t *testing.T) {
	testcases := []testcase{
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 5,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 9,
					},
				},
			},
			expected: 4,
		},
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 5,
						Left: &TreeNode{
							Val: 3,
						},
					},
				},
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val:  9,
						Left: &TreeNode{Val: 7},
					},
				},
			},
			expected: 7,
		},
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
			expected: 2,
		},
		{
			root: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: 0,
					Right: &TreeNode{
						Val: 0,
						Right: &TreeNode{
							Val: 0,
							Right: &TreeNode{
								Val: 0,
								Right: &TreeNode{
									Val: 0,
									Right: &TreeNode{
										Val: 0,
										Right: &TreeNode{
											Val: 0,
											Right: &TreeNode{
												Val: 0,
												Right: &TreeNode{
													Val: 0,
													Right: &TreeNode{
														Val: 0,
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
				Right: &TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val: 0,
						Left: &TreeNode{
							Val: 0,
							Left: &TreeNode{
								Val: 0,
								Left: &TreeNode{
									Val: 0,
									Left: &TreeNode{
										Val: 0,
										Left: &TreeNode{
											Val: 0,
											Left: &TreeNode{
												Val: 0,
												Left: &TreeNode{
													Val: 0,
													Left: &TreeNode{
														Val: 0,
														Left: &TreeNode{
															Val: 0,
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expected: 2,
		},
	}
	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := widthOfBinaryTree(tc.root)
			assert.Equal(t, tc.expected, got)
		})
	}
}
