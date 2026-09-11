package spiral_matrix

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	input    [][]int
	expected []int
}

var testCases []TestCase = []TestCase{
	{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
	{[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}, []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
	{[][]int{{1}}, []int{1}},
	{[][]int{{6, 9, 7}}, []int{6, 9, 7}},
	{[][]int{{6}, {9}, {7}}, []int{6, 9, 7}},
}

func TestSpiralOrder(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := SpiralOrder(testCase.input)
			if !reflect.DeepEqual(got, testCase.expected) {
				test.Fatalf("tested %d expected %+v bti got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
