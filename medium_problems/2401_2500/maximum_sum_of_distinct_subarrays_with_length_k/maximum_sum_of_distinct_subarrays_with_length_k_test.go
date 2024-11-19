package maximumsumofdistinctsubarrayswithlengthk

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputNums []int
	first     int
	expected  int64
}

func TestMaximumSubarraySum(t *testing.T) {
	testCases := []TestCase{
		{
			inputNums: []int{1, 5, 4, 2, 9, 9, 9},
			first:     3,
			expected:  15,
		},
		{
			inputNums: []int{4, 4, 4},
			first:     3,
			expected:  0,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := maximumSubarraySum(testCase.inputNums, testCase.first)

			if got != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
