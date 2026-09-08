package containerwithmostwater

import (
	"fmt"
	"testing"
)

type TestCase struct {
	a []int

	expected int
}

func TestRotatedSortedArray(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		a:        []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
		expected: 49,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := maxArea(testCase.a)
			if got != testCase.expected {
				tb.Errorf("testing %d, expected %+v, got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
