package sumofsquarenumbers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	input    int
	expected bool
}

func TestJudgeSquareSum(t *testing.T) {
	testcases := []testcase{
		{
			input:    4,
			expected: true,
		},
		{
			input:    5,
			expected: true,
		},
		{
			input:    3,
			expected: false,
		},
		{
			input:    45,
			expected: true,
		},
		{
			input:    46,
			expected: false,
		},
	}
	for i, tc := range testcases {
		t.Run(fmt.Sprintf("%d-%d", i, tc.input), func(t *testing.T) {
			assert.Equal(t, tc.expected, judgeSquareSum(tc.input))
		})
	}
}
