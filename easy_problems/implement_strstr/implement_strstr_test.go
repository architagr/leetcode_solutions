package implement_strstr

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputS, inputT string
	expected       int
}

var testCases []TestCase = []TestCase{
	{inputS: "hello", inputT: "ll", expected: 2},
	{inputS: "helallo", inputT: "ll", expected: 4},
	{inputS: "aaa", inputT: "aaaa", expected: -1},
	{inputS: "aaaaa", inputT: "bba", expected: -1},
	{inputS: "mississippi", inputT: "issip", expected: 4},
}

func TestStrStr(t *testing.T) {

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			if got := StrStrApproch2(testCase.inputS, testCase.inputT); got != testCase.expected {
				test.Fatalf("testing %d expected %d but got %d", (i + 1), testCase.expected, got)
			}
		})
	}
}
