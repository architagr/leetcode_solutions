package nestedlistweightsum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	input    []*NestedInteger
	expected int
}

func TestDepthSum(t *testing.T) {
	testcases := []testcase{
		{
			input: []*NestedInteger{
				{
					list: []*NestedInteger{
						{n: 1},
						{n: 1},
					},
				},
				{
					n: 2,
				},
				{
					list: []*NestedInteger{
						{n: 1},
						{n: 1},
					},
				},
			},
			expected: 10,
		},
		{
			input: []*NestedInteger{
				{n: 1},
				{list: []*NestedInteger{
					{n: 4},
					{list: []*NestedInteger{
						{n: 6},
					}},
				}},
			},
			expected: 27,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, depthSum(tc.input))
		})
	}
}
