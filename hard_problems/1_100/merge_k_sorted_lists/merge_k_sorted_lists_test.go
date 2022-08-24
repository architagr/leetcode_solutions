package merge_k_sorted_lists

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	input    []*ListNode
	expected []int
}

func TestMergeKLists(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    GetTestCase1(),
		expected: []int{1, 1, 2, 3, 4, 4, 5, 6},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := MergeKLists(testCase.input)
			x := GetLitListToArray(got)

			if !reflect.DeepEqual(x, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}

func GetTestCase1() []*ListNode {
	var node13 *ListNode = new(ListNode)
	node13.Val = 5

	var node12 *ListNode = new(ListNode)
	node12.Val = 4
	node12.Next = node13

	var node11 *ListNode = new(ListNode)
	node11.Val = 1
	node11.Next = node12

	var node23 *ListNode = new(ListNode)
	node23.Val = 4

	var node22 *ListNode = new(ListNode)
	node22.Val = 3
	node22.Next = node23

	var node21 *ListNode = new(ListNode)
	node21.Val = 1
	node21.Next = node22

	var node32 *ListNode = new(ListNode)
	node32.Val = 6

	var node31 *ListNode = new(ListNode)
	node31.Val = 2
	node31.Next = node32

	return []*ListNode{node11, node21, node31}
}

func GetLitListToArray(head *ListNode) []int {
	arr := make([]int, 0)

	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	return arr
}
