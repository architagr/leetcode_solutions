package count_num_pairs_with_absolute_diff

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	nums   []int
	target int
	output int
}

func TestCountDiff(t *testing.T) {
	testCases := []TestCase{
		{nums: []int{1, 2, 2, 1}, target: 1, output: 4},
		{nums: []int{1, 3}, target: 3, output: 0},
		{nums: []int{3, 2, 1, 5, 4}, target: 2, output: 3},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("run itration number %d", (i+1)), func(test *testing.T) {
			output := CountDiff(testCase.nums, testCase.target)
			if !reflect.DeepEqual(output, testCase.output) {
				test.Errorf("expected %+v, got %+v", testCase.output, output)
			}
		})

	}
}
