package flattennestedlistiterator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	list     []*NestedInteger
	expected []int
}

func TestFlatternNestedListIterator(t *testing.T) {
	testcases := []testcase{
		{
			list: []*NestedInteger{
				{
					list: []*NestedInteger{
						{n: 1},
						{n: 1},
					},
				},
				{
					n: 2,
				},
				{
					list: []*NestedInteger{
						{n: 1},
						{n: 1},
					},
				},
			},
			expected: []int{1, 1, 2, 1, 1},
		},
		{
			list: []*NestedInteger{
				{
					n: 1,
				},
				{
					list: []*NestedInteger{
						{n: 4},
						{
							list: []*NestedInteger{
								{n: 6},
							},
						},
					},
				},
			},
			expected: []int{1, 4, 6},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			obj := Constructor(tc.list)
			arr := []int{}
			for obj.HasNext() {
				arr = append(arr, obj.Next())
			}
			assert.Equal(t, tc.expected, arr)
		})
	}
}
