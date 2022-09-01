package count_triplets_that_can_form_two_arrays_of_equal_XOR

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	input    []int
	expected int
}

func TestCountTriplets(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    []int{2, 3, 1, 6, 7},
		expected: 4,
	})

	testCases = append(testCases, TestCase{
		input:    []int{1, 1, 1, 1, 1},
		expected: 10,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := CountTriplets(testCase.input)

			if !reflect.DeepEqual(got, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
