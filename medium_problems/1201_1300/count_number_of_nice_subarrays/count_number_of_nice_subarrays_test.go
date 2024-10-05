package countnumberofnicesubarrays

import (
	"fmt"
	"testing"
)

type TestCase struct {
	nums     []int
	k        int
	expected int
}

func TestNumberOfSubarrays(t *testing.T) {
	testCases := []TestCase{
		{
			nums:     []int{1, 1, 2, 1, 1},
			k:        3,
			expected: 2,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d", i), func(t1 *testing.T) {
			if got := numberOfSubarrays(tc.nums, tc.k); got != tc.expected {
				t1.Errorf("expected %d but got %d", tc.expected, got)
			}
		})
	}
}
