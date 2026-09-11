package numberofconnectedcomponentsinanundirectedgraph

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	n        int
	edges    [][]int
	expected int
}

func TestCountComponents(t *testing.T) {
	testcases := []testcase{
		{
			n: 5,
			edges: [][]int{
				{0, 1}, {1, 2}, {3, 4},
			},
			expected: 2,
		},
		{
			n: 5,
			edges: [][]int{
				{0, 1}, {1, 2}, {2, 3}, {3, 4},
			},
			expected: 1,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, countComponents(tc.n, tc.edges))
		})
	}
}
