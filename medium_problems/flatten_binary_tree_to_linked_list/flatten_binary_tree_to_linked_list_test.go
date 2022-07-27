package flatten_binary_tree_to_linked_list

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	input    *TreeNode
	expected []int
}

func TestFlatten(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    GetTestCase1(),
		expected: []int{1, 2, 3, 4, 5, 6},
	})
	testCases = append(testCases, TestCase{
		input:    nil,
		expected: []int{},
	})
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			Flatten(testCase.input)
			x := GetLitListToArray(testCase.input)

			if !reflect.DeepEqual(x, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}

func GetTestCase3() *TreeNode {

	var node51 *TreeNode = new(TreeNode)
	node51.Val = 10

	var node41 *TreeNode = new(TreeNode)
	node41.Val = 8

	var node42 *TreeNode = new(TreeNode)
	node42.Val = 9
	node42.Right = node51

	var node31 *TreeNode = new(TreeNode)
	node31.Val = 4
	node31.Left = node41

	var node32 *TreeNode = new(TreeNode)
	node32.Val = 5
	node31.Right = node42

	var node33 *TreeNode = new(TreeNode)
	node33.Val = 6

	var node34 *TreeNode = new(TreeNode)
	node34.Val = 7

	var node21 *TreeNode = new(TreeNode)
	node21.Val = 2
	node21.Right = node32
	node21.Left = node31

	var node22 *TreeNode = new(TreeNode)
	node22.Val = 3
	node22.Right = node34
	node22.Left = node33

	var node1 *TreeNode = new(TreeNode)
	node1.Val = 1
	node1.Right = node22
	node1.Left = node21

	return node1
}

func GetTestCase1() *TreeNode {

	var node31 *TreeNode = new(TreeNode)
	node31.Val = 3

	var node32 *TreeNode = new(TreeNode)
	node32.Val = 4

	var node33 *TreeNode = new(TreeNode)
	node33.Val = 6

	var node21 *TreeNode = new(TreeNode)
	node21.Val = 2
	node21.Left = node31
	node21.Right = node32

	var node22 *TreeNode = new(TreeNode)
	node22.Val = 5
	node22.Right = node33

	var node1 *TreeNode = new(TreeNode)
	node1.Val = 1
	node1.Right = node22
	node1.Left = node21

	return node1
}

func GetLitListToArray(head *TreeNode) []int {
	arr := make([]int, 0)

	for head != nil {
		arr = append(arr, head.Val)
		head = head.Right
	}
	return arr
}
