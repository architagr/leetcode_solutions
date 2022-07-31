package first_missing_positive

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    []int
	expected int
}

var testCases []TestCase = []TestCase{
	{[]int{2, 1}, 3},
}

func TestFirstMissingPositiveNumber(t *testing.T) {

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := FirstMissingPositiveNumber(testCase.input)
			if got != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
