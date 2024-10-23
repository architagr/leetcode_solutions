package minimumdifferencebetweenhighestandlowestofkscores

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    []int
	k        int
	expected int
}

func TestMinimumRecolors(t *testing.T) {
	testCases := []TestCase{
		{
			input:    []int{90},
			k:        1,
			expected: 0,
		},
		{
			input:    []int{9, 4, 1, 7},
			k:        2,
			expected: 2,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := minimumDifference(testCase.input, testCase.k)

			if got != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
