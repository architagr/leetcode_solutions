package maximumerasurevalue

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    []int
	expected int
}

func TestMaximumUniqueSubarray(t *testing.T) {
	testCases := []TestCase{
		{
			input:    []int{4, 2, 4, 5, 6},
			expected: 17,
		},
		{
			input:    []int{5, 2, 1, 2, 5, 2, 1, 2, 5},
			expected: 8,
		},
		{
			input:    []int{187, 470, 25, 436, 538, 809, 441, 167, 477, 110, 275, 133, 666, 345, 411, 459, 490, 266, 987, 965, 429, 166, 809, 340, 467, 318, 125, 165, 809, 610, 31, 585, 970, 306, 42, 189, 169, 743, 78, 810, 70, 382, 367, 490, 787, 670, 476, 278, 775, 673, 299, 19, 893, 817, 971, 458, 409, 886, 434},
			expected: 16911,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := maximumUniqueSubarray(testCase.input)

			if got != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
