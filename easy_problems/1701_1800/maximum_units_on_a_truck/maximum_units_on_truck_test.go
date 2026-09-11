package maximum_units_on_truck

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputA   [][]int
	inputB   int
	expected int
}

var testCases []TestCase = []TestCase{
	{[][]int{{1, 3}, {2, 2}, {3, 1}}, 4, 8},
	{[][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}}, 10, 91},
}

func TestMaxUnitsOnTruck(t *testing.T) {

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			if got := MaxUnitsOnTruck(testCase.inputA, testCase.inputB); got != testCase.expected {
				test.Fatalf("%d's palindrome status should have been %d but got %d", (i + 1), testCase.expected, got)
			}
		})
	}
}
