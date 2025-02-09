package factorcombinations

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	n        int
	expected [][]int
}

func TestPermute(t *testing.T) {
	testcases := []testcase{
		{
			n: 8,
			expected: [][]int{
				{2, 4},
				{2, 2, 2},
			},
		},
		{
			n: 12,
			expected: [][]int{
				{2, 6},
				{2, 2, 3},
				{3, 4},
			},
		},
		{
			n:        2,
			expected: [][]int{},
		},
		{
			n:        3,
			expected: [][]int{},
		},
		{
			n: 4,
			expected: [][]int{
				{2, 2},
			},
		},

		{
			n:        37,
			expected: [][]int{},
		},
		{
			n:        1,
			expected: [][]int{},
		},
	}
	for i, tc := range testcases {
		t.Run(fmt.Sprint(i, tc.n), func(t *testing.T) {
			got := getFactors(tc.n)
			assert.Equal(t, tc.expected, got)
		})
	}
}
