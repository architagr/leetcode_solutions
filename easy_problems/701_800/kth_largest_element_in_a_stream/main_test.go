package kthlargestelementinastream

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	k         int
	result    []int
	initArray []int
	pushArray []int
}

func TestAdd(t *testing.T) {

	testcases := []testCase{
		// {
		// 	k:         3,
		// 	initArray: []int{4, 5, 8, 2},
		// 	pushArray: []int{3, 5, 10, 9, 4},
		// 	result:    []int{4, 5, 5, 8, 8},
		// },
		// {
		// 	k:         4,
		// 	initArray: []int{7, 7, 7, 7, 8, 3},
		// 	pushArray: []int{2, 10, 9, 9},
		// 	result:    []int{7, 7, 7, 8},
		// },
		{
			k:         3,
			initArray: []int{5, -1},          // -1, 5
			pushArray: []int{2, 1, -1, 3, 4}, // -1, 2, 5 | 1, 2,5
			result:    []int{-1, 1, 1, 2, 3},
		},
	}

	for i, tc := range testcases {

		t.Run(fmt.Sprint(i), func(tb *testing.T) {
			obj := Constructor(tc.k, tc.initArray)
			got := make([]int, len(tc.result))
			for i, num := range tc.pushArray {
				got[i] = obj.Add(num)
			}
			assert.Equal(tb, tc.result, got)

		})
	}
}
