package partitionlabels

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	s        string
	expected []int
}

func TestPartitionLabels(t *testing.T) {
	testcases := []testcase{
		{
			s:        "ababcbacadefegdehijhklij",
			expected: []int{9, 7, 8},
		},
		{
			s:        "eccbbbbdec",
			expected: []int{10},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, partitionLabels(tc.s))
		})
	}
}
