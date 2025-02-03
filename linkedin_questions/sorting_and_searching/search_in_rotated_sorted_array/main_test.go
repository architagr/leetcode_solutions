package searchinrotatedsortedarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcaseFindPivot struct {
	nums     []int
	expected int
}

func TestFindPivotIndex(t *testing.T) {
	testcases := []testcaseFindPivot{
		{
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			expected: 4,
		},
		{
			nums:     []int{0, 1, 2, 4, 5, 6, 7},
			expected: 0,
		},
		{
			nums:     []int{1, 2, 4, 5, 6, 7, 0},
			expected: 6,
		},
		{
			nums:     []int{7, 0, 1, 2, 4, 5, 6},
			expected: 1,
		},
	}
	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, findPivotIndex(tc.nums))
		})
	}
}

type testcaseSearchSubArray struct {
	nums       []int
	start, end int
	target     int
	expected   int
}

func TestSearchSubArray(t *testing.T) {
	testcases := []testcaseSearchSubArray{
		{
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			start:    0,
			end:      3,
			target:   4,
			expected: 0,
		},
		{
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			start:    0,
			end:      3,
			target:   5,
			expected: 1,
		},
		{
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			start:    0,
			end:      3,
			target:   8,
			expected: -1,
		},
		{
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			start:    0,
			end:      3,
			target:   7,
			expected: 3,
		},
		{
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			start:    4,
			end:      6,
			target:   0,
			expected: 4,
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
	expected int
}

func TestSearch(t *testing.T) {
	testcases := []testcase{
		{
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   0,
			expected: 4,
		},
		{
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   3,
			expected: -1,
		},
		{
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   5,
			expected: 1,
		},
		{
			nums:     []int{4},
			target:   8,
			expected: -1,
		},
		{
			nums:     []int{4},
			target:   4,
			expected: 0,
		},
		{
			nums:     []int{1, 3},
			target:   0,
			expected: -1,
		},
		{
			nums:     []int{1, 2, 3},
			target:   0,
			expected: -1,
		},
	}
	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, search(tc.nums, tc.target))
		})
	}
}
