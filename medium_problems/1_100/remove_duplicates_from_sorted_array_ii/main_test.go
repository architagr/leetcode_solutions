package removeduplicatesfromsortedarrayii

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	nums         []int
	expectedNums []int
}

func TestRemoveDuplicates(t *testing.T) {
	testcases := []testcase{
		{
			nums:         []int{1, 1, 1, 2, 2, 3},
			expectedNums: []int{1, 1, 2, 2, 3},
		},
		{
			nums:         []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			expectedNums: []int{0, 0, 1, 1, 2, 3, 3},
		},
		{
			nums:         []int{0, 0, 1, 1, 1, 1, 2, 3, 3, 3},
			expectedNums: []int{0, 0, 1, 1, 2, 3, 3},
		},
	}
	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := removeDuplicates(tc.nums)
			assert.Equal(t, len(tc.expectedNums), got)
			for j, n := range tc.expectedNums {
				assert.Equal(t, n, tc.nums[j])
			}
		})
	}
}
