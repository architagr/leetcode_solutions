package findmissingandrepeatedvalues

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMissingAndRepeatedValues(t *testing.T) {
	testcase := []struct {
		input    [][]int
		expected []int
	}{
		{
			input:    [][]int{{1, 3}, {2, 2}},
			expected: []int{2, 4},
		},
		{
			input:    [][]int{{9, 1, 7}, {8, 9, 2}, {3, 4, 6}},
			expected: []int{9, 5},
		},
	}
	for i, tc := range testcase {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, findMissingAndRepeatedValues(tc.input))
		})
	}
}
