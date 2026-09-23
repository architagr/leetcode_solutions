package pathinzigzaglabelledbinarytree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindLevel(t *testing.T) {
	testcases := []struct{ input, expected int }{
		{input: 1, expected: 1},
		{input: 2, expected: 2},
		{input: 10, expected: 4},
		{input: 15, expected: 4},
		{input: 16, expected: 5},
	}
	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(tb *testing.T) {
			got := findLevel(tc.input)
			if got != tc.expected {
				t.Errorf("test case %d failed: expected %d, got %d", tc.input, tc.expected, got)
			}
		})
	}
}

type testCase struct {
	input    int
	expected []int
}

func TestPathInZigZagTree(t *testing.T) {
	testcases := []testCase{
		{input: 1, expected: []int{1}},
		{input: 16, expected: []int{1, 3, 4, 15, 16}},
		{input: 14, expected: []int{1, 3, 4, 14}},
		{input: 26, expected: []int{1, 2, 6, 10, 26}},
	}
	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(tb *testing.T) {
			got := pathInZigZagTree(tc.input)
			assert.Equal(tb, tc.expected, got)
		})
	}
}
