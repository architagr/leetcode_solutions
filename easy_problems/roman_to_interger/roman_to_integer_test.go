package roman_to_integer

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    string
	expected int
}

var testCases []TestCase = []TestCase{
	{"I", 1},
	{"III", 3},
	{"IV", 4},
	{"IX", 9},
	{"LVIII", 58},
	{"MCMXCIV", 1994},
}

func TestRomanToInt(t *testing.T) {

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %s", testCase.input), func(test *testing.T) {
			if got := romanToInt(testCase.input); got != testCase.expected {
				test.Errorf("%s = %d but got %d", testCase.input, testCase.expected, got)
			}
		})
	}
}
