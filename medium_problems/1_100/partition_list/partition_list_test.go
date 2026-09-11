package partition_list

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	inputA   *ListNode
	x        int
	expected []int
}

func TestPartitionList(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		inputA:   GetTestCase1(),
		x:        3,
		expected: []int{1, 2, 2, 4, 3, 5},
	})
	testCases = append(testCases, TestCase{
		inputA:   GetTestCase2(),
		x:        2,
		expected: []int{1, 2},
	})
	testCases = append(testCases, TestCase{
		inputA: &ListNode{
			Val: 1,
		},
		x:        0,
		expected: []int{1},
	})
	testCases = append(testCases, TestCase{
		inputA: &ListNode{
			Val: -1,
		},
		x:        0,
		expected: []int{-1},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := PartitionList(testCase.inputA, testCase.x)
			x := GetLitListToArray(got)
			if !reflect.DeepEqual(x, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}

func GetTestCase2() *ListNode {
	var node2 *ListNode = new(ListNode)
	node2.Val = 1

	var node1 *ListNode = new(ListNode)
	node1.Val = 2
	node1.Next = node2

	return node1
}

func GetTestCase1() *ListNode {
	var node6 *ListNode = new(ListNode)
	node6.Val = 2

	var node5 *ListNode = new(ListNode)
	node5.Val = 5
	node5.Next = node6

	var node4 *ListNode = new(ListNode)
	node4.Val = 2
	node4.Next = node5

	var node3 *ListNode = new(ListNode)
	node3.Val = 3
	node3.Next = node4

	var node2 *ListNode = new(ListNode)
	node2.Val = 4
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
