package task_scheduler_II

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    []int
	space    int
	expected int64
}

func TestTaskSchedulerII(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    []int{1, 2, 1, 2, 3, 1},
		space:    3,
		expected: 9,
	})
	testCases = append(testCases, TestCase{
		input:    []int{5, 8, 8, 5},
		space:    2,
		expected: 6,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := TaskSchedulerII(testCase.input, testCase.space)
			if got != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
