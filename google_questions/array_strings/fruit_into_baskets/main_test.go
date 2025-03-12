package fruitintobaskets

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	fruits   []int
	expected int
}

func TestTotalFruits(t *testing.T) {
	testcases := []testcase{
		{
			fruits:   []int{1, 2, 1},
			expected: 3,
		},
		{
			fruits:   []int{0, 1, 2, 2},
			expected: 3,
		},
		{
			fruits:   []int{1, 2, 3, 2, 2},
			expected: 4,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, totalFruit(tc.fruits))
		})
	}
}
