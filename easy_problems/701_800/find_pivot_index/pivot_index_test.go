package pivot_index

import (
	"fmt"
	tpkg "testing"
)

type TestCase struct {
	input    []int
	expected int
}

var testCases []TestCase = []TestCase{
	// {[]int{1, 7, 3, 6, 5, 6}, 3},
	// {[]int{1, 2, 3}, -1},
	{[]int{2, 1, -1}, 0},
}

func TestPivotIndex(t *tpkg.T) {

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *tpkg.T) {
			test.Logf("Start testing %d", testCase.input)

			if got := PivotIndex(testCase.input); got != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
