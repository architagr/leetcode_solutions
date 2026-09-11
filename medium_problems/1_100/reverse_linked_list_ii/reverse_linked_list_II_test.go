package reverse_linked_list_II

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	inputA      *ListNode
	left, right int
	expected    []int
}

func TestReverseList(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		inputA:   GetTestCase1(),
		left:     2,
		right:    4,
		expected: []int{1, 4, 3, 2, 5},
	})
	testCases = append(testCases, TestCase{
		inputA:   GetTestCase2(),
		left:     2,
		right:    4,
		expected: []int{1, 4, 3, 2, 5, 6},
	})
	testCases = append(testCases, TestCase{
		inputA:   GetTestCase3(),
		left:     1,
		right:    2,
		expected: []int{5, 3},
	})
	testCases = append(testCases, TestCase{
		inputA: &ListNode{
			Val:  5,
			Next: nil,
		},
		left:     1,
		right:    1,
		expected: []int{5},
	})
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := ReverseBetween(testCase.inputA, testCase.left, testCase.right)
			x := GetLitListToArray(got)

			if !reflect.DeepEqual(x, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}

func GetTestCase1() *ListNode {
	var node5 *ListNode = new(ListNode)
	node5.Val = 5

	var node4 *ListNode = new(ListNode)
	node4.Val = 4
	node4.Next = node5

	var node3 *ListNode = new(ListNode)
	node3.Val = 3
	node3.Next = node4

	var node2 *ListNode = new(ListNode)
	node2.Val = 2
	node2.Next = node3

	var node1 *ListNode = new(ListNode)
	node1.Val = 1
	node1.Next = node2

	return node1
}

func GetTestCase2() *ListNode {
	var node6 *ListNode = new(ListNode)
	node6.Val = 6

	var node5 *ListNode = new(ListNode)
	node5.Val = 5
	node5.Next = node6

	var node4 *ListNode = new(ListNode)
	node4.Val = 4
	node4.Next = node5

	var node3 *ListNode = new(ListNode)
	node3.Val = 3
	node3.Next = node4

	var node2 *ListNode = new(ListNode)
	node2.Val = 2
	node2.Next = node3

	var node1 *ListNode = new(ListNode)
	node1.Val = 1
	node1.Next = node2

	return node1
}

func GetTestCase3() *ListNode {

	var node2 *ListNode = new(ListNode)
	node2.Val = 5

	var node1 *ListNode = new(ListNode)
	node1.Val = 3
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
