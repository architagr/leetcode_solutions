package rearrange_array_elements_by_sign

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	inputA   []int
	expected []int
}

func TestRearrangeArray(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		inputA:   []int{3, 1, -2, -5, 2, -4},
		expected: []int{3, -2, 1, -5, 2, -4},
	})
	testCases = append(testCases, TestCase{
		inputA:   []int{-1, 1},
		expected: []int{1, -1},
	})
	testCases = append(testCases, TestCase{
		inputA:   []int{28, -41, 22, -8, -37, 46, 35, -9, 18, -6, 19, -26, -37, -10, -9, 15, 14, 31},
		expected: []int{28, -41, 22, -8, 46, -37, 35, -9, 18, -6, 19, -26, 15, -37, 14, -10, 31, -9},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := RearrangeArray(testCase.inputA)
			if !reflect.DeepEqual(got, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
