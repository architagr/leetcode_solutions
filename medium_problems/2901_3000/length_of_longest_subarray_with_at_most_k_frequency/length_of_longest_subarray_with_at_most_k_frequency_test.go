package lengthoflongestsubarraywithatmostkfrequency

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputNums []int
	first     int
	expected  int
}

func TestMaxSubarrayLength(t *testing.T) {
	testCases := []TestCase{
		{
			inputNums: []int{1, 2, 3, 1, 2, 3, 1, 2},
			first:     2,
			expected:  6,
		},
		{
			inputNums: []int{1, 2, 1, 2, 1, 2, 1, 2},
			first:     1,
			expected:  2,
		},
		{
			inputNums: []int{5, 5, 5, 5, 5, 5, 5},
			first:     4,
			expected:  4,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := maxSubarrayLength(testCase.inputNums, testCase.first)

			if got != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
