package allpathsfromsourceleadtodestination

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	n           int
	edges       [][]int
	source      int
	destination int
	expected    bool
}

func TestLeadsToDestination(t *testing.T) {
	testcases := []testcase{
		{
			n: 3,
			edges: [][]int{
				{0, 1}, {0, 2},
			},
			source:      0,
			destination: 2,
			expected:    false,
		},
		{
			n: 4,
			edges: [][]int{
				{0, 1}, {0, 3}, {1, 2}, {2, 1},
			},
			source:      0,
			destination: 3,
			expected:    false,
		},
		{
			n: 4,
			edges: [][]int{
				{0, 1}, {0, 2}, {1, 3}, {2, 3},
			},
			source:      0,
			destination: 3,
			expected:    true,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, leadsToDestination(tc.n, tc.edges, tc.source, tc.destination))
		})
	}
}
