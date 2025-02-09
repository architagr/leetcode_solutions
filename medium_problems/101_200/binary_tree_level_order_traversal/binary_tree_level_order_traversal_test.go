package binary_tree_level_order_traversal

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	input    *TreeNode
	expected [][]int
}

func TestLevelOrder(t *testing.T) {
	testCases := make([]TestCase, 0)
	a := GetTestCase1()
	testCases = append(testCases, TestCase{
		input:    a,
		expected: [][]int{{3}, {9, 20}, {15, 7}},
	})

	testCases = append(testCases, TestCase{
		input:    nil,
		expected: [][]int{},
	})
	b := GetTestCase3()
	testCases = append(testCases, TestCase{
		input:    b,
		expected: [][]int{{1}, {2, 3}, {4, 5, 6, 7}, {8, 9}, {10}},
	})
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := LevelOrder(testCase.input)

			if !reflect.DeepEqual(got, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
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
	node31.Val = 15

	var node32 *TreeNode = new(TreeNode)
	node32.Val = 7

	var node21 *TreeNode = new(TreeNode)
	node21.Val = 9

	var node22 *TreeNode = new(TreeNode)
	node22.Val = 20
	node22.Left = node31
	node22.Right = node32

	var node1 *TreeNode = new(TreeNode)
	node1.Val = 3
	node1.Right = node22
	node1.Left = node21

	return node1
}
