package missingranges

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	nums    []int
	lower   int
	upper   int
	expeted [][]int
}

func TestFindMissingRange(t *testing.T) {

	testcases := []testcase{
		{
			lower: 0,
			upper: 99,
			nums:  []int{0, 1, 3, 50, 75},
			expeted: [][]int{
				{2, 2},
				{4, 49},
				{51, 74},
				{76, 99},
			},
		},
		{
			lower:   -1,
			upper:   -1,
			nums:    []int{-1},
			expeted: [][]int{},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.ElementsMatch(t, findMissingRanges(tc.nums, tc.lower, tc.upper), tc.expeted)
		})
	}
}
