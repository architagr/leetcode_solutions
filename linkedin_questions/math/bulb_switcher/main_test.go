package bulbswitcher

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBulbSwitcher(t *testing.T) {
	testcases := []struct {
		input    int
		expected int
	}{
		{
			input:    3,
			expected: 1,
		},
	}
	for _, tc := range testcases {
		t.Run(fmt.Sprint(tc.input), func(t *testing.T) {
			assert.Equal(t, tc.expected, bulbSwitch(tc.input))
		})
	}
}
