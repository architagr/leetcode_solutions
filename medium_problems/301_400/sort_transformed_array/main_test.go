package sorttransformedarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	nums     []int
	a        int
	b        int
	c        int
	expected []int
}

func TestSortTransformedArray(t *testing.T) {
	testcases := []testcase{
		{
			nums:     []int{-99, -94, -90, -88, -84, -83, -79, -68, -58, -52, -52, -50, -47, -45, -35, -29, -5, -1, 9, 12, 13, 25, 27, 33, 36, 38, 40, 46, 47, 49, 57, 57, 61, 63, 73, 75, 79, 97, 98},
			a:        -2,
			b:        44,
			c:        -56,
			expected: []int{-24014, -21864, -20216, -19416, -17864, -17486, -16014, -14952, -14606, -12296, -9336, -9062, -8006, -7752, -7752, -7502, -7256, -6542, -6086, -5222, -4814, -4046, -4046, -4046, -3014, -2702, -2406, -2264, -1496, -1272, -1064, -782, -326, -326, -206, -102, 178, 178, 184},
		},
		{
			nums:     []int{-4},
			a:        0,
			b:        -1,
			c:        5,
			expected: []int{9},
		},
		{
			nums:     []int{-4, -2, -2, -2},
			a:        0,
			b:        -1,
			c:        5,
			expected: []int{7, 7, 7, 9},
		},
		{
			nums:     []int{-4, -2, 2, 4},
			a:        1,
			b:        3,
			c:        5,
			expected: []int{3, 9, 15, 33},
		},
		{
			nums:     []int{-4, -2, 2, 4},
			a:        -1,
			b:        3,
			c:        5,
			expected: []int{-23, -5, 1, 7},
		},
		{
			nums:     []int{-4, -2, 2, 4},
			a:        -1,
			b:        -3,
			c:        5,
			expected: []int{-23, -5, 1, 7},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := sortTransformedArray(tc.nums, tc.a, tc.b, tc.c)
			assert.Equal(t, tc.expected, got)
		})
	}
}
