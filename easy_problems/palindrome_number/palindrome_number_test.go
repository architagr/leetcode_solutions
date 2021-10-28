package palindrome_number

import (
	"fmt"
	tpkg "testing"
)

type TestCase struct {
	input    int
	expected bool
}

var testCases []TestCase = []TestCase{
	{121, true},
	{1234554321, true},
	{1234564321, false},
	{12345654321, true},
	{12345664321, false},
	{-121, false},
	{10, false},
	{-10014, false},
}

func TestIsPalindrome(t *tpkg.T) {

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", testCase.input), func(test *tpkg.T) {
			test.Logf("Start testing %d", testCase.input)

			if got := isPalindrome(testCase.input); got != testCase.expected {
				test.Fatalf("%d's palindrome status should have been %t but got %t", testCase.input, testCase.expected, got)
			}
		})
	}
}
