package maximumsumcircularsubarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	nums     []int
	expected int
}

func TestMaxSubarraySumCircular(t *testing.T) {
	testcaes := []testcase{
		{
			nums:     []int{5, -3, 5},
			expected: 10,
		},
		{
			nums:     []int{-3, -2, -3},
			expected: -2,
		},
		{
			nums:     []int{1, -2, 3, -2},
			expected: 3,
		},
	}

	for i, tc := range testcaes {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, maxSubarraySumCircular(tc.nums))
		})

	}
}
