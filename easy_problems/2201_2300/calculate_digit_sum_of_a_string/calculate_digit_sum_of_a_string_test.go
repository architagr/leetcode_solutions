package calculate_digit_sum_of_a_string

import (
	"fmt"
	"testing"
)

type TestCase struct {
	nums   string
	target int
	output string
}

func TestDigitSum(t *testing.T) {
	testCases := []TestCase{
		{nums: "11111222223", target: 3, output: "135"},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("run itration number %d", (i+1)), func(test *testing.T) {
			output := DigitSum(testCase.nums, testCase.target)
			if output != testCase.output {
				test.Errorf("expected %+v, got %+v", testCase.output, output)
			}
		})

	}
}
