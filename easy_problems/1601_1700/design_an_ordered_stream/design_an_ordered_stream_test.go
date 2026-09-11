package design_an_ordered_stream

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	size     int
	inputA   []int
	inputS   []string
	expected [][]string
}

func TestOrderedStreamInsert(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, case1())
	testCases = append(testCases, case2())

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", i+1), func(tb *testing.T) {
			obj := Constructor(testCase.size)
			for j := 0; j < testCase.size; j++ {
				x := obj.Insert(testCase.inputA[j], testCase.inputS[j])
				if !reflect.DeepEqual(x, testCase.expected[j]) {
					tb.Errorf("Testing %d error in input %d expected %+v got %+v", (i + 1), j, testCase.expected[j], x)
				}
			}
		})
	}
}

func case1() TestCase {
	return TestCase{
		size:     5,
		inputA:   []int{3, 1, 2, 5, 4},
		inputS:   []string{"ccccc", "aaaaa", "bbbbb", "eeeee", "ddddd"},
		expected: [][]string{{}, {"aaaaa"}, {"bbbbb", "ccccc"}, {}, {"ddddd", "eeeee"}},
	}
}
func case2() TestCase {
	return TestCase{
		size:     4,
		inputA:   []int{2, 1, 3, 4},
		inputS:   []string{"iknuk", "bykll", "gxfcf", "rgwzt"},
		expected: [][]string{{}, {"bykll", "iknuk"}, {"gxfcf"}, {"rgwzt"}},
	}
}
