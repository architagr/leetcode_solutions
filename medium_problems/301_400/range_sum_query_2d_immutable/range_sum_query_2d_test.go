package range_sum_query_2d

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	matrix   [][]int
	queries  [][]int
	expected []int
}

func TestRangeSumQuery2d(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		matrix:   [][]int{{3, 0, 1, 4, 2}, {5, 6, 3, 2, 1}, {1, 2, 0, 1, 5}, {4, 1, 0, 1, 7}, {1, 0, 3, 0, 5}},
		queries:  [][]int{{2, 1, 4, 3}, {1, 1, 2, 2}, {1, 2, 2, 4}, {0, 2, 2, 4}},
		expected: []int{8, 11, 12, 19},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			obj := Constructor(testCase.matrix)
			result := make([]int, len(testCase.expected))
			for j, value := range testCase.queries {
				result[j] = obj.SumRegion(value[0], value[1], value[2], value[3])
			}

			if !reflect.DeepEqual(result, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, result)
			}
		})
	}
}
