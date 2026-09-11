package string_to_integer_atoi

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    string
	expected int
}

var testCases []TestCase = []TestCase{
	{
		input:    "+1",
		expected: 1,
	},
	{
		input:    "+-1",
		expected: 0,
	},
	{
		input:    "   -1",
		expected: -1,
	},
	{
		input:    "szdds   -1",
		expected: 0,
	},
	{
		input:    "-2szdds   -1",
		expected: -2,
	},
	{
		input:    "   +0 123",
		expected: 0,
	},
	{
		input:    "   +1 123",
		expected: 1,
	},
	{
		input:    "9223372036854775808",
		expected: 2147483647,
	},
	{
		input:    "-91283472332",
		expected: -2147483648,
	},
	{
		input:    "-2147483648",
		expected: -2147483648,
	},
}

func TestMyAtoi(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := MyAtoi(testCase.input)
			if got != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
