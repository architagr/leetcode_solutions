package isomorphic_strings

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputS, inputT string
	expected       bool
}

var testCases []TestCase = []TestCase{
	{inputS: "egg", inputT: "add", expected: true},
	{inputS: "foo", inputT: "bar", expected: false},
	{inputS: "paper", inputT: "title", expected: true},
}

func TestIsIsomorphic(t *testing.T) {

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			if got := IsIsomorphic(testCase.inputS, testCase.inputT); got != testCase.expected {
				test.Fatalf("testing %d expected %t but got %t", (i + 1), testCase.expected, got)
			}
		})
	}
}
