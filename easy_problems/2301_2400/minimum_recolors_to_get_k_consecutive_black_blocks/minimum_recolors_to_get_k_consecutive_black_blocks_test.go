package minimumrecolorstogetkconsecutiveblackblocks

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    string
	k        int
	expected int
}

func TestMinimumRecolors(t *testing.T) {
	testCases := []TestCase{
		{
			input:    "WBBWWBBWBW",
			k:        7,
			expected: 3,
		},
		{
			input:    "WBWBBBW",
			k:        2,
			expected: 0,
		},
		{
			input:    "WWWWWWW",
			k:        2,
			expected: 2,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := minimumRecolors(testCase.input, testCase.k)

			if got != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
