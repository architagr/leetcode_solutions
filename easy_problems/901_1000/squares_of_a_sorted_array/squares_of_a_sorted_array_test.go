package squares_of_a_sorted_array

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	inputA   []int
	expected []int
}

func TestSortedSquares(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		inputA:   []int{2, 4, 5, 7},
		expected: []int{4, 16, 25, 49},
	})
	testCases = append(testCases, TestCase{
		inputA:   []int{-2, -1},
		expected: []int{1, 4},
	})
	testCases = append(testCases, TestCase{
		inputA:   []int{-2, -1, 1, 3},
		expected: []int{1, 1, 4, 9},
	})

	testCases = append(testCases, TestCase{
		inputA:   []int{-4, -1, 0, 3, 10},
		expected: []int{0, 1, 9, 16, 100},
	})
	testCases = append(testCases, TestCase{
		inputA:   []int{-7, -3, 2, 3, 11},
		expected: []int{4, 9, 9, 49, 121},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := SortedSquares(testCase.inputA)
			if !reflect.DeepEqual(got, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
