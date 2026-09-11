package reverse_linked_list

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	inputA *ListNode

	expected []int
}

func TestReverseList(t *testing.T) {
	testCases := make([]TestCase, 0)
	a := GetTestCase1()
	testCases = append(testCases, TestCase{
		inputA:   a,
		expected: []int{4, 2, 1},
	})
	testCases = append(testCases, TestCase{
		inputA:   nil,
		expected: []int{},
	})
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := ReverseList(testCase.inputA)
			x := GetLitListToArray(got)

			if !reflect.DeepEqual(x, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}

func GetTestCase1() *ListNode {
	var node3 *ListNode = new(ListNode)
	node3.Val = 4

	var node2 *ListNode = new(ListNode)
	node2.Val = 2
	node2.Next = node3

	var node1 *ListNode = new(ListNode)
	node1.Val = 1
	node1.Next = node2

	return node1
}

func GetLitListToArray(head *ListNode) []int {
	arr := make([]int, 0)

	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	return arr
}
