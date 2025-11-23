package addstrings

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	num1, num2, extected string
}

func TestAddStrings(t *testing.T) {
	testcases := []testcase{
		{
			num1:     "11",
			num2:     "123",
			extected: "134",
		},
		{
			num1:     "1",
			num2:     "9",
			extected: "10",
		},
		{
			num1:     "456",
			num2:     "77",
			extected: "533",
		},
		{
			num1:     "0",
			num2:     "0",
			extected: "0",
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.extected, addStrings(tc.num1, tc.num2))
		})
	}
}
