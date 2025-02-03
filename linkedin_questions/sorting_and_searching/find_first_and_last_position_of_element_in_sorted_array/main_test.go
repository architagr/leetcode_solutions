package findfirstandlastpositionofelementinsortedarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcaseSearchSubArray struct {
	nums       []int
	start, end int
	target     int
	expected   int
}

func TestSearchSubArray(t *testing.T) {
	testcases := []testcaseSearchSubArray{
		{
			nums:     []int{5, 7, 7, 8, 8, 10},
			start:    0,
			end:      5,
			target:   8,
			expected: 3,
		},
		{
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			start:    0,
			end:      5,
			target:   6,
			expected: 2,
		},
		{
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			start:    0,
			end:      5,
			target:   16,
			expected: -1,
		},
		{
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			start:    0,
			end:      5,
			target:   4,
			expected: 0,
		},
		{
			nums:     []int{},
			start:    0,
			end:      0,
			target:   8,
			expected: -1,
		},
	}
	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, searchSubArray(tc.nums, tc.start, tc.end, tc.target))
		})
	}
}

type testcase struct {
	nums     []int
	target   int
	expected []int
}

func TestSearchRange(t *testing.T) {
	testcases := []testcase{
		{
			nums:     []int{5, 7, 7, 8, 8, 10},
			target:   8,
			expected: []int{3, 4},
		},
		{
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   6,
			expected: []int{2, 2},
		},
		{
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   16,
			expected: []int{-1, -1},
		},
		{
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   4,
			expected: []int{0, 0},
		},
		{
			nums:     []int{},
			target:   8,
			expected: []int{-1, -1},
		},
	}
	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, searchRange(tc.nums, tc.target))
		})
	}
}
