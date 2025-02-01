package validnumber

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNumberPositive(t *testing.T) {
	testcases := []string{
		"2",
		"0089",
		"-0.1",
		"+3.14",
		"4.",
		"-.9",
		"2e10",
		"-90E3",
		"3e+7",
		"+6e-1",
		"53.5e93",
		"-123.456e789",
		"0",
	}
	for index, tc := range testcases {
		t.Run(fmt.Sprintf("%d-%s", index, tc), func(t *testing.T) {
			assert.True(t, isNumber(tc))
		})
	}
}

func TestIsNumberNegative(t *testing.T) {
	testcases := []string{
		"abc", "1a", "1e", "e3", "99e2.5", "--6", "-+3", "95a54e53", "e", ".", "6+1",
	}
	for index, tc := range testcases {
		t.Run(fmt.Sprintf("%d-%s", index, tc), func(t *testing.T) {
			assert.False(t, isNumber(tc))
		})
	}
}
