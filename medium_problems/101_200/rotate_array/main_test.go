package rotatearray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	nums     []int
	k        int
	expected []int
}

func TestRotate(t *testing.T) {
	testcases := []testcase{
		{
			nums:     []int{1, 2, 3, 4, 5, 6, 7},
			k:        3,
			expected: []int{5, 6, 7, 1, 2, 3, 4},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			rotate(tc.nums, tc.k)
			assert.Equal(t, tc.expected, tc.nums)
		})
	}
}
