package sqrtx

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	nums     int
	expected int
}

func TestMySqrt(t *testing.T) {
	testcases := []testcase{
		{
			nums:     4,
			expected: 2,
		},
		{
			nums:     8,
			expected: 2,
		},
		{
			nums:     9,
			expected: 3,
		},
		{
			nums:     25,
			expected: 5,
		},
		{
			nums:     26,
			expected: 5,
		},
		{
			nums:     27,
			expected: 5,
		},
		{
			nums:     28,
			expected: 5,
		},
		{
			nums:     29,
			expected: 5,
		},
		{
			nums:     30,
			expected: 5,
		},
		{
			nums:     31,
			expected: 5,
		},
		{
			nums:     32,
			expected: 5,
		},
		{
			nums:     33,
			expected: 5,
		},
		{
			nums:     34,
			expected: 5,
		},
		{
			nums:     35,
			expected: 5,
		},
		{
			nums:     36,
			expected: 6,
		},
	}
	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, mySqrt(tc.nums))
		})
	}
}
