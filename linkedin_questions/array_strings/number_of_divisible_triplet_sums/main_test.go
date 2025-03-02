package numberofdivisibletripletsums

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	nums     []int
	d        int
	expected int
}

func TestDivisibleTripletCount(t *testing.T) {
	testcases := []testcase{
		{
			nums:     []int{3, 3, 4, 7, 8},
			d:        5,
			expected: 3,
		},
		{
			nums:     []int{3, 3, 3, 3},
			d:        3,
			expected: 4,
		},
		{
			nums:     []int{3, 3, 3, 3},
			d:        6,
			expected: 0,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, divisibleTripletCount(tc.nums, tc.d))
		})
	}
}
