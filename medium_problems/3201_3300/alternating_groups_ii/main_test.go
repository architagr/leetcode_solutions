package alternatinggroupii

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	colors   []int
	k        int
	expected int
}

func TestNumberOfAlternatingGroups(t *testing.T) {
	testcases := []testcase{
		{
			colors:   []int{0, 1, 0, 1, 0},
			k:        3,
			expected: 3,
		},
		{
			colors:   []int{0, 0, 0, 1},
			k:        3,
			expected: 1,
		},
		{
			colors:   []int{0, 0, 0, 1},
			k:        4,
			expected: 0,
		},
		{
			colors:   []int{0, 1, 0, 0, 1, 0, 1},
			k:        6,
			expected: 2,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, numberOfAlternatingGroups(tc.colors, tc.k))
		})
	}
}
