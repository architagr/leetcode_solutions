package arithmeticslices

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    []int
	expected int
}

func TestNumberOfSubstrings(t *testing.T) {
	testCases := []TestCase{
		{
			input:    []int{1, 2, 3, 4},
			expected: 3,
		},
		{
			input:    []int{1, 2, 3},
			expected: 1,
		},
		{
			input:    []int{1},
			expected: 0,
		},
		{
			input:    []int{1, 2, 3, 5, 7, 9},
			expected: 4,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := numberOfArithmeticSlices(testCase.input)

			if got != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
