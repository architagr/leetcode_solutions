package search_in_rotated_sorted_array

import (
	"fmt"
	"testing"
)

type TestCase struct {
	a []int

	expected, b int
}

func TestRotatedSortedArray(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		a:        []int{4, 5, 6, 7, 0, 1, 2, 3},
		b:        4,
		expected: 0,
	})

	testCases = append(testCases, TestCase{
		a:        []int{1},
		b:        4,
		expected: -1,
	})

	testCases = append(testCases, TestCase{
		a:        []int{1},
		b:        1,
		expected: 0,
	})

	testCases = append(testCases, TestCase{
		a:        []int{1, 3},
		b:        3,
		expected: 1,
	})

	testCases = append(testCases, TestCase{
		a:        []int{3, 1},
		b:        1,
		expected: 1,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := Search(testCase.a, testCase.b)

			if got != testCase.expected {
				tb.Errorf("testing %d, expected %+v, got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
