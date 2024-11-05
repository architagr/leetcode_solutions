package maximumsumoftwononoverlappingsubarrays

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputNums     []int
	first, second int
	expected      int
}

func TestNumberOfSubstrings(t *testing.T) {
	testCases := []TestCase{
		{
			inputNums: []int{0, 6, 5, 2, 2, 5, 1, 9, 4},
			first:     1,
			second:    2,
			expected:  20,
		},
		{
			inputNums: []int{3, 8, 1, 3, 2, 1, 8, 9, 0},
			first:     3,
			second:    2,
			expected:  29,
		},
		{
			inputNums: []int{2, 1, 5, 6, 0, 9, 5, 0, 3, 8},
			first:     4,
			second:    3,
			expected:  31,
		},
		{
			inputNums: []int{8, 20, 6, 2, 20, 17, 6, 3, 20, 8, 12},
			first:     5,
			second:    4,
			expected:  108,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := maxSumTwoNoOverlap(testCase.inputNums, testCase.first, testCase.second)

			if got != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
