package countsubarrayswheremaxelementappearsatleastktimes

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    []int
	k        int
	expected int64
}

func TestCountSubarrays(t *testing.T) {
	testCases := []TestCase{
		{
			input:    []int{1, 3, 2, 3, 3},
			k:        2,
			expected: 6,
		},
		{
			input:    []int{1, 4, 2, 1},
			k:        3,
			expected: 0,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := countSubarrays(testCase.input, testCase.k)

			if got != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
