package median_of_two_sorted_arrays

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input1, input2 []int
	expected       float64
}

var testCases []TestCase = []TestCase{
	{
		input1:   []int{1, 3},
		input2:   []int{2},
		expected: 2.0,
	},
	{
		input1:   []int{1, 2},
		input2:   []int{3, 4},
		expected: 2.5,
	},
	{
		input1:   []int{1, 2, 5},
		input2:   []int{3, 4},
		expected: 3.0,
	},
}

func TestFindMedianSortedArrays(t *testing.T) {

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := FindMedianSortedArrays(testCase.input1, testCase.input2)
			if got != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
