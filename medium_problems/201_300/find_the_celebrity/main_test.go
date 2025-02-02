package findthecelebrity

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	knows    func(a int, b int) bool
	n        int
	expected int
}

func TestSolution(t *testing.T) {
	testcases := []testcase{
		{
			knows: func(a int, b int) bool {
				arr := [][]int{
					{1, 0, 0, 1, 0, 1},
					{0, 1, 1, 0, 0, 0},
					{1, 1, 1, 0, 0, 1},
					{1, 1, 0, 1, 0, 1},
					{0, 1, 0, 0, 1, 0},
					{1, 1, 1, 0, 0, 1},
				}
				return arr[a][b] == 1
			},
			n:        6,
			expected: -1,
		},
		{
			knows: func(a int, b int) bool {
				arr := [][]int{
					{1, 0},
					{0, 1},
				}
				return arr[a][b] == 1
			},
			n:        2,
			expected: -1,
		},
		{
			knows: func(a int, b int) bool {
				arr := [][]int{
					{1, 1, 0},
					{0, 1, 0},
					{1, 1, 1},
				}
				return arr[a][b] == 1
			},
			n:        3,
			expected: 1,
		},
		{
			knows: func(a int, b int) bool {
				arr := [][]int{
					{1, 0, 1},
					{1, 1, 0},
					{0, 1, 1},
				}
				return arr[a][b] == 1
			},
			n:        3,
			expected: -1,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			f := solution(tc.knows)
			assert.Equal(t, tc.expected, f(tc.n))
		})
	}
}
