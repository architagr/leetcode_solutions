package number_of_substrings_with_only_1s

import (
	"fmt"
	"testing"
)

type TestCase struct {
	s      string
	output int
}

func TestNumSub(t *testing.T) {
	testCases := []TestCase{
		{
			s:      "0110111",
			output: 9,
		},
		{
			s:      "101",
			output: 2,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("run itration number %d", (i+1)), func(test *testing.T) {
			output := NumSub(testCase.s)
			if output != testCase.output {
				test.Errorf("expected %+v, got %+v", testCase.output, output)
			}
		})

	}
}
