package max_area_cake_piece

import (
	"fmt"
	"testing"
)

type TestCase struct {
	h, w                 int
	horizontal, vertical []int
	expected             int
}

var testCases []TestCase = []TestCase{
	{5, 4, []int{3, 1}, []int{1}, 6},
	{5, 4, []int{1, 2, 4}, []int{1, 3}, 4},
	{5, 4, []int{3}, []int{3}, 9},
}

func TestMaxArea(t *testing.T) {

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			if got := MaxArea(testCase.h, testCase.w, testCase.horizontal, testCase.vertical); got != testCase.expected {
				test.Fatalf("%d's palindrome status should have been %d but got %d", (i + 1), testCase.expected, got)
			}
		})
	}
}
