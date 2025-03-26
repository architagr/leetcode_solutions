package deleteandearn

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	nums     []int
	expected int
}

func TestDeleteAndEarn(t *testing.T) {
	testcases := []testcase{
		{
			nums:     []int{3, 4, 2},
			expected: 6,
		},
		{
			nums:     []int{2, 2, 3, 3, 3, 4},
			expected: 9,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, deleteAndEarn(tc.nums))
		})
	}
}
