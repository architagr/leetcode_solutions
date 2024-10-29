package grumpybookstoreowner

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputC   []int
	inputG   []int
	inputM   int
	expected int
}

func TestNumberOfSubstrings(t *testing.T) {
	testCases := []TestCase{
		{
			inputC:   []int{1, 0, 1, 2, 1, 1, 7, 5},
			inputG:   []int{0, 1, 0, 1, 0, 1, 0, 1},
			inputM:   3,
			expected: 16,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := maxSatisfied(testCase.inputC, testCase.inputG, testCase.inputM)

			if got != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
