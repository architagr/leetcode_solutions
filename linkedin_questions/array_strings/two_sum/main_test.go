package twosum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	nums     []int
	target   int
	expected []int
}

func TestTwoSum(t *testing.T) {
	testCases := []testCase{
		{
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			nums:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := twoSum(tc.nums, tc.target)
			if !assert.ObjectsAreEqual(tc.expected, got) {
				t.Error()
			}
		})
	}
}
