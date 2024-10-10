package defusethebomb

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	inputCode []int
	inputK    int
	expected  []int
}

func TestDecrypt(t *testing.T) {
	testCases := make([]TestCase, 0)
	// testCases = append(testCases, TestCase{
	// 	inputCode: []int{5, 7, 1, 4},
	// 	inputK:    3,
	// 	expected:  []int{12, 10, 16, 13},
	// })
	// testCases = append(testCases, TestCase{
	// 	inputCode: []int{1, 2, 3, 4},
	// 	inputK:    0,
	// 	expected:  []int{0, 0, 0, 0},
	// })

	testCases = append(testCases, TestCase{
		inputCode: []int{2, 4, 9, 3},
		inputK:    -2,
		expected:  []int{12, 5, 6, 13},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := decrypt(testCase.inputCode, testCase.inputK)
			if !reflect.DeepEqual(got, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
