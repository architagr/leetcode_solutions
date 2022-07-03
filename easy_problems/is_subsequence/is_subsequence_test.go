package is_subsequence

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputS, inputT string
	expected       bool
}

var testCases []TestCase = []TestCase{
	{inputS: "abc", inputT: "ahbgdc", expected: true},
	{inputS: "abc", inputT: "babgdc", expected: true},
	{inputS: "abc", inputT: "bbgadc", expected: false},
	{inputS: "axc", inputT: "ahbgdc", expected: false},
}

func TestIsSubsequence(t *testing.T) {

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			if got := IsSubsequence(testCase.inputS, testCase.inputT); got != testCase.expected {
				test.Fatalf("testing %d expected %t but got %t", (i + 1), testCase.expected, got)
			}
		})
	}
}
