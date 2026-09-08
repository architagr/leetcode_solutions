package rotatelist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	input    []int
	k        int
	expected []int
}

func TestRotateRight(t *testing.T) {
	testcases := []testcase{
		{
			input:    []int{},
			k:        2,
			expected: []int{},
		},
		{
			input:    []int{1},
			k:        2,
			expected: []int{1},
		},
		{
			input:    []int{1, 2, 3, 4, 5},
			k:        2,
			expected: []int{4, 5, 1, 2, 3},
		},
		{
			input:    []int{0, 1, 2},
			k:        4,
			expected: []int{2, 0, 1},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := rotateRight(convertToList(tc.input), tc.k)
			assert.Equal(t, tc.expected, convertToArray(got))

		})
	}
}

func convertToList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	head := &ListNode{
		Val: arr[0],
	}
	temp := head
	for _, val := range arr[1:] {
		temp.Next = &ListNode{
			Val: val,
		}
		temp = temp.Next
	}
	return head
}
func convertToArray(head *ListNode) []int {
	temp := head
	arr := make([]int, 0)
	for temp != nil {
		arr = append(arr, temp.Val)
		temp = temp.Next
	}
	return arr
}
