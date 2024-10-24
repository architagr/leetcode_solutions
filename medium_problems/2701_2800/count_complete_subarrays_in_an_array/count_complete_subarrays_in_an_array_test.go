package countcompletesubarraysinanarray

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    []int
	expected int
}

func TestCountCompleteSubarrays(t *testing.T) {
	testCases := []TestCase{
		{
			input:    []int{1, 3, 1, 2, 2},
			expected: 4,
		},
		{
			input:    []int{5, 5, 5, 5},
			expected: 10,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := countCompleteSubarrays(testCase.input)

			if got != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
