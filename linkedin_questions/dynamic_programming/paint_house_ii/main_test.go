package painthouseii

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	input    [][]int
	expected int
}

func TestMinCost(t *testing.T) {
	testcases := []testcase{
		{
			input: [][]int{
				{17, 2, 17},
				{16, 16, 5},
				{14, 3, 19},
			},
			expected: 10,
		},
		{
			input: [][]int{
				{7, 6, 2},
			},
			expected: 2,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, minCostII(tc.input))
		})
	}
}
