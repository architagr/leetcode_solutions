package minimum_XOR_sum_of_two_arrays

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input1, input2 []int
	expected       int
}

var testCases []TestCase = []TestCase{
	// {
	// 	input1:   []int{1, 2},
	// 	input2:   []int{2, 3},
	// 	expected: 2,
	// },
	// {
	// 	input1:   []int{1, 0, 3},
	// 	input2:   []int{5, 3, 4},
	// 	expected: 8,
	// },
	{
		input1:   []int{72, 97, 8, 32, 15},
		input2:   []int{63, 97, 57, 60, 83},
		expected: 152,
	},
}

func TestMinimumXORSum(t *testing.T) {

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := MinimumXORSum(testCase.input1, testCase.input2)
			if got != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
