package valid_palindrome

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    string
	expected bool
}

var testCases []TestCase = []TestCase{
	{input: "A man, a plan, a canal: Panama", expected: true},
	{input: "race a car", expected: false},
	{input: " ", expected: true},
	{input: "a ", expected: true},
	{input: ".,", expected: true},
	{input: "0A", expected: false},
}

func TestCheckPalindrome(t *testing.T) {

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			if got := CheckPalindrome(testCase.input); got != testCase.expected {
				test.Fatalf("testing %d expected %t but got %t", (i + 1), testCase.expected, got)
			}
		})
	}
}
