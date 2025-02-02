package integertoenglishwords

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	num      int
	expected string
}

func TestNumberToWordsForUnder1000(t *testing.T) {
	testcases := []testcase{
		{
			num:      0,
			expected: "Zero",
		},
		{
			num:      1,
			expected: "One",
		},
		{
			num:      11,
			expected: "Eleven",
		},
		{
			num:      13,
			expected: "Thirteen",
		},
		{
			num:      17,
			expected: "Seventeen",
		},
		{
			num:      21,
			expected: "Twenty One",
		},
		{
			num:      28,
			expected: "Twenty Eight",
		},
		{
			num:      99,
			expected: "Ninety Nine",
		},
		{
			num:      123,
			expected: "One Hundred Twenty Three",
		},
		{
			num:      119,
			expected: "One Hundred Nineteen",
		},
		{
			num:      111,
			expected: "One Hundred Eleven",
		},
		{
			num:      110,
			expected: "One Hundred Ten",
		},
		{
			num:      12345,
			expected: "Twelve Thousand Three Hundred Forty Five",
		},
		{
			num:      12000,
			expected: "Twelve Thousand",
		},

		{
			num:      1234567,
			expected: "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven",
		},
		{
			num:      1234000,
			expected: "One Million Two Hundred Thirty Four Thousand",
		},
		{
			num:      1234019,
			expected: "One Million Two Hundred Thirty Four Thousand Nineteen",
		},
		{
			num:      11234019,
			expected: "Eleven Million Two Hundred Thirty Four Thousand Nineteen",
		},
		{
			num:      19234019,
			expected: "Nineteen Million Two Hundred Thirty Four Thousand Nineteen",
		},
		{
			num:      49234019,
			expected: "Forty Nine Million Two Hundred Thirty Four Thousand Nineteen",
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprint(tc.num), func(t *testing.T) {
			assert.Equal(t, tc.expected, numberToWords(tc.num))
		})
	}
}
