package sum_of_absolute_difference_in_a_sorted_array

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	input    []int
	expected []int
}

func TestGetSumAbsoluteDifferences(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    []int{2, 3, 5},
		expected: []int{4, 3, 5},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := GetSumAbsoluteDifferences(testCase.input)

			if !reflect.DeepEqual(got, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
