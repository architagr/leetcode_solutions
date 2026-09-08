package first_n_last_position_element_sorted_array

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	input, expected []int
	target          int
}

func TestSearchRange(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    []int{5, 7, 7, 8, 8, 10},
		target:   8,
		expected: []int{3, 4},
	})
	testCases = append(testCases, TestCase{
		input:    []int{5, 7, 7, 8, 8, 10},
		target:   6,
		expected: []int{-1, -1},
	})

	testCases = append(testCases, TestCase{
		input:    []int{},
		target:   6,
		expected: []int{-1, -1},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := SearchRange(testCase.input, testCase.target)
			if !reflect.DeepEqual(got, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
