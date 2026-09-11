package running_sum

import (
	"fmt"
	"reflect"
	tpkg "testing"
)

type TestCase struct {
	input, expected []int
}

var testCases []TestCase = []TestCase{
	{[]int{1, 2, 3, 4}, []int{1, 3, 6, 10}},
}

func TestRunningSum(t *tpkg.T) {

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *tpkg.T) {
			test.Logf("Start testing %d", testCase.input)

			if got := RunningSum(testCase.input); !reflect.DeepEqual(got, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
