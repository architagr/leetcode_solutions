package exclusivetimeoffunctions

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	n        int
	logs     []string
	expected []int
}

func TestExclusiveTime(t *testing.T) {
	testcases := []testcase{
		{
			n:        2,
			logs:     []string{"0:start:0", "0:start:2", "0:end:5", "1:start:7", "1:end:7", "0:end:8"},
			expected: []int{8, 1},
		},
		{
			n:        1,
			logs:     []string{"0:start:0", "0:start:2", "0:end:5", "0:start:6", "0:end:6", "0:end:7"},
			expected: []int{8},
		},
		{
			n:        2,
			logs:     []string{"0:start:0", "0:start:2", "0:end:5", "1:start:6", "1:end:6", "0:end:7"},
			expected: []int{7, 1},
		},
		{
			n:        2,
			logs:     []string{"0:start:0", "1:start:2", "1:end:5", "0:end:6"},
			expected: []int{3, 4},
		},
	}
	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := exclusiveTime(tc.n, tc.logs)
			assert.Equal(t, tc.expected, got)
		})
	}
}
