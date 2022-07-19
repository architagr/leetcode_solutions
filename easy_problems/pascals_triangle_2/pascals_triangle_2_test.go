package pascalstriangle2

import (
	"fmt"
	"reflect"
	tpkg "testing"
)

type TestCase struct {
	input    int
	expected []int
}

var testCases []TestCase = []TestCase{
	{1, []int{1, 1}},
	{2, []int{1, 2, 1}},
	{3, []int{1, 3, 3, 1}},
	{4, []int{1, 4, 6, 4, 1}},
	{5, []int{1, 5, 10, 10, 5, 1}},
	{6, []int{1, 6, 15, 20, 15, 6, 1}},
}

func TestGeneratePascaleTriangle(t *tpkg.T) {

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", testCase.input), func(test *tpkg.T) {
			got := GeneratePascaleTriangle(testCase.input)
			if !reflect.DeepEqual(got, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", testCase.input, testCase.expected, got)
			}
		})
	}
}
