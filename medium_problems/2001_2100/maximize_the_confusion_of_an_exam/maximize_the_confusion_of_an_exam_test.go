package maximizetheconfusionofanexam

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    string
	k        int
	expected int
}

func TestMaxConsecutiveAnswers(t *testing.T) {
	testCases := []TestCase{
		{
			input:    "TTFF",
			k:        2,
			expected: 4,
		},
		{
			input:    "TFFT",
			k:        1,
			expected: 3,
		},
		{
			input:    "TTFTTFTT",
			k:        1,
			expected: 5,
		},
		{
			input:    "TF",
			k:        2,
			expected: 2,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := maxConsecutiveAnswers(testCase.input, testCase.k)

			if got != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
