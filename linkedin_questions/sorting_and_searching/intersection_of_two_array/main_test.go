package intersectionoftwoarray

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	nums1, nums2 []int
	expected     []int
}

func TestIntersection(t *testing.T) {
	testcases := []testcase{
		{
			nums1:    []int{1, 2, 2, 1},
			nums2:    []int{2, 2},
			expected: []int{2},
		},
		{
			nums1:    []int{4, 9, 5},
			nums2:    []int{9, 4, 9, 8, 4},
			expected: []int{9, 4},
		},
	}
	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := intersection(tc.nums1, tc.nums2)
			sort.Ints(got)
			sort.Ints(tc.expected)
			assert.Equal(t, tc.expected, got)
		})
	}
}
