package search_a_2d_matrix

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	input    [][]int
	target   int
	expected bool
}

var testCases []TestCase = []TestCase{
	{[][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 0, false},
	{[][]int{{5, 6, 9}, {9, 10, 11}, {11, 14, 18}}, 9, true},
	{[][]int{{2, 5}, {2, 8}, {7, 9}, {7, 11}, {9, 11}}, 7, true},
	{[][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 5, true},
	{[][]int{{1, 4, 7, 11, 15}}, 5, false},
	{[][]int{{1}, {2}, {3}, {10}, {18}}, 5, false},
	{[][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 20, false},
}

func TestSearchMatrix(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := SearchMatrix(testCase.input, testCase.target)
			if !reflect.DeepEqual(got, testCase.expected) {
				test.Fatalf("tested %d expected %+v bti got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
