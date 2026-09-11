package insert_interval

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	nums   [][]int
	target []int
	output [][]int
}

func TestInsertInterval(t *testing.T) {
	testCases := []TestCase{
		{
			nums:   [][]int{{1, 5}},
			target: []int{0, 0},
			output: [][]int{{0, 0}, {1, 5}},
		},
		{
			nums:   [][]int{{1, 5}},
			target: []int{0, 3},
			output: [][]int{{0, 5}},
		},
		{
			nums:   [][]int{{1, 5}},
			target: []int{2, 7},
			output: [][]int{{1, 7}},
		},
		{
			nums:   [][]int{{1, 5}},
			target: []int{2, 3},
			output: [][]int{{1, 5}},
		},
		{
			nums:   [][]int{{1, 5}},
			target: []int{6, 7},
			output: [][]int{{1, 5}, {6, 7}},
		},
		{
			nums:   [][]int{{1, 3}, {6, 9}},
			target: []int{0, 1},
			output: [][]int{{0, 3}, {6, 9}},
		},
		{
			nums:   [][]int{{1, 3}, {6, 9}},
			target: []int{2, 5},
			output: [][]int{{1, 5}, {6, 9}},
		},
		{
			nums:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			target: []int{4, 8},
			output: [][]int{{1, 2}, {3, 10}, {12, 16}},
		},
		{
			nums:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			target: []int{17, 18},
			output: [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}, {17, 18}},
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("run itration number %d", (i+1)), func(test *testing.T) {
			output := InsertInterval(testCase.nums, testCase.target)
			if !reflect.DeepEqual(output, testCase.output) {
				test.Errorf("expected %+v, got %+v", testCase.output, output)
			}
		})

	}
}
