package excelsheetcolumntitle

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	input    int
	expected string
}

func TestConvertToTitle(t *testing.T) {
	tests := []testcase{
		{input: 1, expected: "A"},
		{input: 28, expected: "AB"},
		{input: 701, expected: "ZY"},
		{input: 52, expected: "AZ"},
		{input: 703, expected: "AAA"},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, convertToTitle(tc.input))
		})
	}
}
