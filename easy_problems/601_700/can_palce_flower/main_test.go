package canpalceflower

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	flowerBed []int
	n         int
	expected  bool
}

func TestCanPlaceFlowers(t *testing.T) {
	testcases := []testcase{
		{
			flowerBed: []int{1, 0, 0, 0, 1},
			n:         1,
			expected:  true,
		},
		{
			flowerBed: []int{1, 0, 0, 0, 1},
			n:         2,
			expected:  false,
		},
		{
			flowerBed: []int{1, 0, 0, 0, 0, 1},
			n:         2,
			expected:  false,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, canPlaceFlowers(tc.flowerBed, tc.n))
		})
	}
}
