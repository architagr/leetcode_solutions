package maximumlengthsubstringwithtwooccurrences

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    string
	expected int
}

func TestMaximumLengthSubstring(t *testing.T) {
	testCases := []TestCase{
		{
			input:    "bcbbbcba",
			expected: 4,
		},
		{
			input:    "aaaa",
			expected: 2,
		},
		{
			input:    "adaddccdb",
			expected: 5,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := maximumLengthSubstring(testCase.input)

			if got != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
