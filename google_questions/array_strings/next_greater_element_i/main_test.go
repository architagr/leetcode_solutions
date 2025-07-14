package nextgreaterelementi

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	nums1, nums2, expected []int
}

func TestNextGreaterElement(t *testing.T) {
	testcases := []testcase{
		{
			nums1:    []int{4, 1, 2, 0},
			nums2:    []int{3, 4, 2, 0, 1},
			expected: []int{-1, -1, -1, 1},
		},
		{
			nums1:    []int{4, 1, 2},
			nums2:    []int{1, 3, 4, 2},
			expected: []int{-1, 3, -1},
		},
		{
			nums1:    []int{4, 1, 2},
			nums2:    []int{1, 3, 4, 2, 5},
			expected: []int{5, 3, 5},
		},
		{
			nums1:    []int{2, 4},
			nums2:    []int{1, 2, 3, 4},
			expected: []int{3, -1},
		},
	}
	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := nextGreaterElement(tc.nums1, tc.nums2)
			assert.EqualValues(t, tc.expected, got)
		})
	}
}
