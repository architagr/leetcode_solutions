package min_cost_climbing_stairs

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    []int
	expected int
}

var testCases []TestCase = []TestCase{
	{
		input:    []int{10, 15, 20},
		expected: 15,
	},
	{
		input:    []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
		expected: 6,
	},
}

func TestMinCostClimbingStairs(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			if got := MinCostClimbingStairs(testCase.input); got != testCase.expected {
				test.Fatalf("tested %d expected %d but got %d", (i + 1), testCase.expected, got)
			}
		})
	}
}
