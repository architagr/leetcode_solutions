package allintegersrangearecovered

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input       [][]int
	left, right int
	expected    bool
}

func TestIsCovered(t *testing.T) {

	testcases := make([]TestCase, 0)
	testcases = append(testcases, TestCase{
		input:    [][]int{{1, 2}, {3, 4}, {5, 6}},
		left:     2,
		right:    5,
		expected: true,
	})
	testcases = append(testcases, TestCase{
		input:    [][]int{{1, 10}, {10, 20}},
		left:     21,
		right:    21,
		expected: false,
	})
	testcases = append(testcases, TestCase{
		input:    [][]int{{36, 50}, {14, 28}, {4, 31}, {24, 37}, {13, 36}, {27, 33}, {23, 32}, {23, 27}, {1, 35}},
		left:     35,
		right:    40,
		expected: true,
	})
	for i, caseObj := range testcases {
		t.Run(fmt.Sprintf("running %d", i), func(tb *testing.T) {
			got := isCovered(caseObj.input, caseObj.left, caseObj.right)
			if got != caseObj.expected {
				tb.Fatalf("tested %d expected %+v but got %+v", (i + 1), caseObj.expected, got)
			}
		})
	}

}
