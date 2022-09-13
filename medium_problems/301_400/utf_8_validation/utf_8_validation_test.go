package utf_8_validation

import (
	"fmt"
	"testing"
)

type TestCase struct {
	data     []int
	expected bool
}

func TestRangeSumQuery2d(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		data:     []int{197, 130, 1},
		expected: true,
	})

	testCases = append(testCases, TestCase{
		data:     []int{235, 140, 4},
		expected: false,
	})

	testCases = append(testCases, TestCase{
		data:     []int{255},
		expected: false,
	})

	testCases = append(testCases, TestCase{
		data:     []int{145},
		expected: false,
	})

	testCases = append(testCases, TestCase{
		data:     []int{248, 130, 130, 130},
		expected: false,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := ValidateUTF8(testCase.data)
			if got != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
