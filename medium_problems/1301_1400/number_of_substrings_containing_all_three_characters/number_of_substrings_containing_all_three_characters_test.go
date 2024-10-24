package numberofsubstringscontainingallthreecharacters

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    string
	expected int
}

func TestNumberOfSubstrings(t *testing.T) {
	testCases := []TestCase{
		{
			input:    "abcabc",
			expected: 10,
		},
		{
			input:    "aaacb",
			expected: 3,
		},
		{
			input:    "abc",
			expected: 1,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := numberOfSubstrings(testCase.input)

			if got != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
