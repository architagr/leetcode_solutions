package plus_one

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {

	var testCases = []struct {
		input    []int
		expected []int
	}{
		{input: []int{1, 2, 3}, expected: []int{1, 2, 4}},
		{input: []int{1, 2, 9}, expected: []int{1, 3, 0}},
		{input: []int{9, 9, 9}, expected: []int{1, 0, 0, 0}},
		{input: []int{9}, expected: []int{1, 0}},
	}

	for _, testcase := range testCases {
		t.Run(fmt.Sprintf("%+v as input", testcase.input), func(test *testing.T) {
			if got := plusOne(testcase.input); !reflect.DeepEqual(got, testcase.expected) {
				test.Fatalf("for input %+v expected is %+v got %+v", testcase.input, testcase.expected, got)
			}
		})
	}
}
