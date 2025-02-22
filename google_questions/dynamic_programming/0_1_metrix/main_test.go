package metrix

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	mat    [][]int
	result [][]int
}

func TestUpdateMetrix(t *testing.T) {
	testcases := []testcase{
		{
			mat: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0},
			},
			result: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0},
			},
		},
		{
			mat: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{1, 1, 1},
			},
			result: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{1, 2, 1},
			},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.result, updateMatrix(tc.mat))
		})
	}
}
