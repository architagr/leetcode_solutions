package maximumaveragesubarrayi

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputNum []int
	inputk   int
	expected float64
}

func TestFindMaxAverage(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		inputNum: []int{1, 12, -5, -6, 50, 3},
		inputk:   4,
		expected: 12.7500,
	})
	testCases = append(testCases, TestCase{
		inputNum: []int{5},
		inputk:   1,
		expected: 5.00,
	})
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := findMaxAverage(testCase.inputNum, testCase.inputk)
			if got != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
