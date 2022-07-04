package candy

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    []int
	expected int
}

var testCases []TestCase = []TestCase{
	// {[]int{1, 6, 10, 8, 7, 3, 2}, 18},
	// {[]int{1, 0, 3}, 5},
	// {[]int{1, 2, 3}, 6},
	// {[]int{1, 2, 2}, 4},
	// {[]int{1, 3, 2}, 4},
	{[]int{1, 3, 6, 8, 9, 5, 3, 6, 8, 5, 4, 2, 2, 3, 7, 7, 9, 8, 6, 6, 6, 4, 2}, 50},
}

func TestCandy(t *testing.T) {

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			if got := Candy(testCase.input); got != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
