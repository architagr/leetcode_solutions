package getequalsubstringswithinbudget

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputA, inputB string
	max            int
	expected       int
}

func TestEqualSubstring(t *testing.T) {
	testCases := []TestCase{
		{
			inputA:   "abcd",
			inputB:   "bcdf",
			max:      3,
			expected: 3,
		},
		{
			inputA:   "abcd",
			inputB:   "cdef",
			max:      3,
			expected: 1,
		},
		{
			inputA:   "abcd",
			inputB:   "acde",
			max:      0,
			expected: 1,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := equalSubstring(testCase.inputA, testCase.inputB, testCase.max)

			if got != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
