package nestedlistweightsumii

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcaseMaxDepth struct {
	input    []*NestedInteger
	expected int
}

func TestDepthSum(t *testing.T) {
	testcases := []testcaseMaxDepth{
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
			expected: 8,
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
			expected: 17,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, depthSumInverse(tc.input))
		})
	}
}
