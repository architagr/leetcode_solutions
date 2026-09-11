package sort_array_by_parity_II

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	inputA   []int
	expected []int
}

func TestSortArrayByParityII(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		inputA:   []int{4, 2, 5, 7},
		expected: []int{4, 5, 2, 7},
	})
	testCases = append(testCases, TestCase{
		inputA:   []int{2, 3},
		expected: []int{2, 3},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := SortArrayByParityII(testCase.inputA)
			if !reflect.DeepEqual(got, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
