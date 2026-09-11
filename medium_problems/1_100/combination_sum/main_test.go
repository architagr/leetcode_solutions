package combimationsum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	candidate []int
	target    int
	expected  [][]int
}

func TestCombinationSum(t *testing.T) {
	testcases := []testcase{
		{
			candidate: []int{2, 3, 6, 7},
			target:    7,
			expected: [][]int{
				{2, 2, 3},
				{7},
			},
		},
		{
			candidate: []int{2, 3, 5},
			target:    8,
			expected: [][]int{
				{2, 2, 2, 2},
				{2, 3, 3},
				{3, 5},
			},
		},
		{
			candidate: []int{2},
			target:    1,
			expected:  [][]int{},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := combinationSum(tc.candidate, tc.target)
			assert.Equal(t, len(tc.expected), len(got))
			assert.Equal(t, tc.expected, got)
		})
	}
}

// func testSort(arr [][]int) {
// 	sort.Slice(arr, func(i, j int) bool {
// 		sort.Ints()
// 	})
// }
