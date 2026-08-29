package findacorrespondingnodeofabinarytreeinclone

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	original, cloned, target *TreeNode
}

func TestGetTargetCopy1(t *testing.T) {
	testcase := &testcase{
		original: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 6,
				},
				Right: &TreeNode{
					Val: 19,
				},
			},
		},
		target: &TreeNode{Val: 3},
	}

	expected := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 6,
		},
		Right: &TreeNode{
			Val: 19,
		},
	}
	testcase.cloned = &TreeNode{
		Val: 7,
		Left: &TreeNode{
			Val: 4,
		},
		Right: expected,
	}

	got := getTargetCopy(testcase.original, testcase.cloned, testcase.target)
	assert.Equal(t, expected, got)

}

func TestGetTargetCopy2(t *testing.T) {
	testcase := &testcase{
		original: &TreeNode{
			Val: 7,
		},
		target: &TreeNode{Val: 7},
	}

	expected := &TreeNode{
		Val: 7,
	}
	testcase.cloned = expected

	got := getTargetCopy(testcase.original, testcase.cloned, testcase.target)
	assert.Equal(t, expected, got)

}

func TestGetTargetCopy3(t *testing.T) {
	testcase := &testcase{
		original: &TreeNode{
			Val: 8,
			Right: &TreeNode{
				Val: 6,
				Right: &TreeNode{
					Val: 5,
					Right: &TreeNode{
						Val: 4,
						Right: &TreeNode{
							Val: 3,
							Right: &TreeNode{
								Val: 2,
								Right: &TreeNode{
									Val: 1,
								},
							},
						},
					},
				},
			},
		},
		target: &TreeNode{Val: 4},
	}

	expected := &TreeNode{
		Val: 4,
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
	}

	testcase.cloned = &TreeNode{
		Val: 8,
		Right: &TreeNode{
			Val: 6,
			Right: &TreeNode{
				Val:   5,
				Right: expected,
			},
		},
	}

	got := getTargetCopy(testcase.original, testcase.cloned, testcase.target)
	assert.Equal(t, expected, got)

}
