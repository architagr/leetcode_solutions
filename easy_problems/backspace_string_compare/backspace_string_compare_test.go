package backspace_string_compare

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputS, inputG string
	expected       bool
}

func TestBackspaceCompare(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		inputS:   "ab#c",
		inputG:   "ad#c",
		expected: true,
	})

	testCases = append(testCases, TestCase{
		inputS:   "ab##",
		inputG:   "c#d#",
		expected: true,
	})

	testCases = append(testCases, TestCase{
		inputS:   "#ab##",
		inputG:   "c#d#",
		expected: true,
	})

	testCases = append(testCases, TestCase{
		inputS:   "a#c",
		inputG:   "b",
		expected: false,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := BackspaceCompare(testCase.inputS, testCase.inputG)

			if got != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
