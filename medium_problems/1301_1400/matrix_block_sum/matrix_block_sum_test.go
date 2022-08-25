package matrix_block_sum

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	input    [][]int
	k        int
	expected [][]int
}

func TestMatrixBlockSum(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		k:        1,
		expected: [][]int{{12, 21, 16}, {27, 45, 33}, {24, 39, 28}},
	})
	testCases = append(testCases, TestCase{
		input:    [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		k:        2,
		expected: [][]int{{45, 45, 45}, {45, 45, 45}, {45, 45, 45}},
	})
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := MatrixBlockSum(testCase.input, testCase.k)

			if !reflect.DeepEqual(got, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
