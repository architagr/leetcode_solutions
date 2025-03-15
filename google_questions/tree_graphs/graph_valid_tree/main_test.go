package graphvalidtree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	n        int
	edges    [][]int
	expected bool
}

func TestValidTree(t *testing.T) {
	testcases := []testcase{
		{
			n: 5,
			edges: [][]int{
				{0, 1}, {0, 2}, {0, 3}, {1, 4},
			},
			expected: true,
		},
		{
			n: 5,
			edges: [][]int{
				{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4},
			},
			expected: false,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, validTree(tc.n, tc.edges))
		})
	}
}
