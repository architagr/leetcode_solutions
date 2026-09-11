package merge_sorted_list

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	inputA, inputB *ListNode

	expected []int
}

func TestMergeTwoLists(t *testing.T) {
	testCases := make([]TestCase, 0)
	a, b := GetTestCase1()
	a2, b2 := GetTestCase2()
	testCases = append(testCases, TestCase{
		inputA:   a,
		inputB:   b,
		expected: []int{1, 1, 2, 3, 4, 4},
	})
	testCases = append(testCases, TestCase{
		inputA:   a2,
		inputB:   b2,
		expected: []int{0, 0},
	})
	testCases = append(testCases, TestCase{
		inputA:   nil,
		inputB:   nil,
		expected: []int{},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := MergeTwoLists(testCase.inputA, testCase.inputB)
			x := GetLitListToArray(got)

			if !reflect.DeepEqual(x, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}

func GetTestCase3() (*ListNode, *ListNode) {
	return nil, nil
}

func GetTestCase2() (*ListNode, *ListNode) {
	return new(ListNode), new(ListNode)
}

func GetTestCase1() (*ListNode, *ListNode) {
	var node3 *ListNode = new(ListNode)
	node3.Val = 4

	var node2 *ListNode = new(ListNode)
	node2.Val = 2
	node2.Next = node3

	var node1 *ListNode = new(ListNode)
	node1.Val = 1
	node1.Next = node2

	var node23 *ListNode = new(ListNode)
	node23.Val = 4

	var node22 *ListNode = new(ListNode)
	node22.Val = 3
	node22.Next = node23

	var node21 *ListNode = new(ListNode)
	node21.Val = 1
	node21.Next = node22

	return node1, node21
}

func GetLitListToArray(head *ListNode) []int {
	arr := make([]int, 0)

	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	return arr
}
