package add_two_numbers_ll

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	inputA, inputB *ListNode

	expected []int
}

func TestAddTwoNumbers(t *testing.T) {
	testCases := make([]TestCase, 0)
	a, b := GetTestCase1()
	testCases = append(testCases, TestCase{
		inputA:   a,
		inputB:   b,
		expected: []int{2, 4, 8},
	})

	a2, b2 := GetTestCase2()
	testCases = append(testCases, TestCase{
		inputA:   a2,
		inputB:   b2,
		expected: []int{2, 4, 2, 5, 5, 6},
	})

	a3, b3 := GetTestCase3()
	testCases = append(testCases, TestCase{
		inputA:   a3,
		inputB:   b3,
		expected: []int{0, 2, 0, 1},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := addTwoNumbers(testCase.inputA, testCase.inputB)
			x := GetLitListToArray(got)

			if !reflect.DeepEqual(x, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got.Val)
			}
		})
	}
}
func GetTestCase3() (*ListNode, *ListNode) {
	var node3 *ListNode = new(ListNode)
	node3.Val = 9

	var node2 *ListNode = new(ListNode)
	node2.Val = 9
	node2.Next = node3

	var node1 *ListNode = new(ListNode)
	node1.Val = 9
	node1.Next = node2

	var node21 *ListNode = new(ListNode)
	node21.Val = 2

	var node11 *ListNode = new(ListNode)
	node11.Val = 1
	node11.Next = node21

	return node1, node11
}

func GetTestCase2() (*ListNode, *ListNode) {

	var node3 *ListNode = new(ListNode)
	node3.Val = 9

	var node2 *ListNode = new(ListNode)
	node2.Val = 2
	node2.Next = node3

	var node1 *ListNode = new(ListNode)
	node1.Val = 1
	node1.Next = node2

	var node6 *ListNode = new(ListNode)
	node6.Val = 6

	var node5 *ListNode = new(ListNode)
	node5.Val = 5
	node5.Next = node6

	var node4 *ListNode = new(ListNode)
	node4.Val = 4
	node4.Next = node5

	var node31 *ListNode = new(ListNode)
	node31.Val = 3
	node31.Next = node4

	var node21 *ListNode = new(ListNode)
	node21.Val = 2
	node21.Next = node31

	var node11 *ListNode = new(ListNode)
	node11.Val = 1
	node11.Next = node21

	return node1, node11
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

	var node31 *ListNode = new(ListNode)
	node31.Val = 4

	var node21 *ListNode = new(ListNode)
	node21.Val = 2
	node21.Next = node31

	var node11 *ListNode = new(ListNode)
	node11.Val = 1
	node11.Next = node21

	return node1, node11
}

func GetLitListToArray(head *ListNode) []int {
	arr := make([]int, 0)

	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	return arr
}
