package minimumsizesubarraysum

import (
	"fmt"
	"testing"
)

type TestCase struct {
	target   int
	nums     []int
	expected int
}

func TestMinSubArrayeLen(t *testing.T) {

	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		target:   7,
		nums:     []int{2, 3, 1, 2, 4, 3},
		expected: 2,
	})
	testCases = append(testCases, TestCase{
		target:   7,
		nums:     []int{2, 3, 4, 3, 1, 2},
		expected: 2,
	})
	testCases = append(testCases, TestCase{
		target:   4,
		nums:     []int{1, 4, 4},
		expected: 1,
	})
	testCases = append(testCases, TestCase{
		target:   11,
		nums:     []int{1, 1, 1, 1, 1},
		expected: 0,
	})

	for i, ut := range testCases {
		t.Run(fmt.Sprintf("Running %d", i), func(tb *testing.T) {
			got := minSubArrayeLen(ut.target, ut.nums)
			if got != ut.expected {
				tb.Errorf("expected %d but got %d", ut.expected, got)
			}
		})

	}
}
