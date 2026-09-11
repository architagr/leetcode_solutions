package wiggle_subsequence

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    []int
	expected int
}

var testCases []TestCase = []TestCase{
	{[]int{1, 7, 4, 9, 2, 5}, 6},
	{[]int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}, 7},
	{[]int{2, 2, 2, 2}, 1},
	{[]int{2, 2}, 1},
	{[]int{2, 3}, 2},
	{[]int{3, 3, 3, 2, 5}, 3},
}

func TestWiggleMaxLength(t *testing.T) {

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			if got := WiggleMaxLength(testCase.input); got != testCase.expected {
				test.Fatalf("tested %d expected %d but got %d", (i + 1), testCase.expected, got)
			}
		})
	}
}
