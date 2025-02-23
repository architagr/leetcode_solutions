package partitiontokequalsumsubsets

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	nums     []int
	k        int
	expected bool
}

func TestMinCost(t *testing.T) {
	testcases := []testcase{
		{
			nums:     []int{3, 3, 10, 2, 6, 5, 10, 6, 8, 3, 2, 1, 6, 10, 7, 2},
			k:        6,
			expected: false,
		},
		{
			nums:     []int{1, 1, 1, 1, 2, 2, 2, 2},
			k:        2,
			expected: true,
		},
		{
			nums:     []int{4, 3, 2, 3, 5, 2, 1},
			k:        4,
			expected: true,
		},
		{
			nums:     []int{1, 2, 3, 4},
			k:        4,
			expected: false,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			defer func(i int, startTime time.Time) {
				t.Logf("time taken by %d is %+v", i, time.Since(startTime))
			}(i, time.Now())
			assert.Equal(t, tc.expected, canPartitionKSubsets(tc.nums, tc.k))
		})
	}
}
