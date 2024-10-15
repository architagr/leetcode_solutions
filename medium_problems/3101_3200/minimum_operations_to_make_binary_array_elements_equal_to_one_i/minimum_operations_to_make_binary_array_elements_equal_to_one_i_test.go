package minimumoperationstomakebinaryarrayelementsequaltoonei

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputNum []int
	expected int
}

func TestMinOperations(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		inputNum: []int{0, 1, 1, 1, 0, 0},
		expected: 3,
	})
	testCases = append(testCases, TestCase{
		inputNum: []int{0, 1, 1, 1},
		expected: -1,
	})
	testCases = append(testCases, TestCase{
		inputNum: []int{1, 0, 0, 1, 1, 0, 1, 1, 1},
		expected: -1,
	})
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := minOperations(testCase.inputNum)
			if got != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
