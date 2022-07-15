package validate_binary_search_tree

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputA *TreeNode

	expected bool
}

func TestIsValidBst(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		inputA:   GetTestCase1(),
		expected: false,
	})
	testCases = append(testCases, TestCase{
		inputA:   GetTestCase2(),
		expected: true,
	})
	testCases = append(testCases, TestCase{
		inputA:   GetTestCase3(),
		expected: false,
	})
	testCases = append(testCases, TestCase{
		inputA:   GetTestCase4(),
		expected: false,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := IsValidBst(testCase.inputA)

			if got != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}

func GetTestCase1() *TreeNode {
	var node31 *TreeNode = new(TreeNode)
	node31.Val = 3

	var node32 *TreeNode = new(TreeNode)
	node32.Val = 6

	var node21 *TreeNode = new(TreeNode)
	node21.Val = 1

	var node22 *TreeNode = new(TreeNode)
	node22.Val = 4
	node22.Left = node31
	node22.Right = node32

	var node1 *TreeNode = new(TreeNode)
	node1.Val = 5
	node1.Left = node21
	node1.Right = node22

	return node1
}

func GetTestCase2() *TreeNode {
	var node21 *TreeNode = new(TreeNode)
	node21.Val = 1

	var node22 *TreeNode = new(TreeNode)
	node22.Val = 3

	var node1 *TreeNode = new(TreeNode)
	node1.Val = 2
	node1.Left = node21
	node1.Right = node22

	return node1
}

func GetTestCase3() *TreeNode {
	var node31 *TreeNode = new(TreeNode)
	node31.Val = 3

	var node32 *TreeNode = new(TreeNode)
	node32.Val = 7

	var node21 *TreeNode = new(TreeNode)
	node21.Val = 4

	var node22 *TreeNode = new(TreeNode)
	node22.Val = 6
	node22.Left = node31
	node22.Right = node32

	var node1 *TreeNode = new(TreeNode)
	node1.Val = 5
	node1.Left = node21
	node1.Right = node22

	return node1
}

func GetTestCase4() *TreeNode {
	var node41 *TreeNode = new(TreeNode)
	node41.Val = 27

	var node31 *TreeNode = new(TreeNode)
	node31.Val = 19
	node31.Right = node41

	var node32 *TreeNode = new(TreeNode)
	node32.Val = 56

	var node21 *TreeNode = new(TreeNode)
	node21.Val = 26
	node21.Left = node31

	var node22 *TreeNode = new(TreeNode)
	node22.Val = 56
	node22.Right = node32

	var node1 *TreeNode = new(TreeNode)
	node1.Val = 32
	node1.Left = node21
	node1.Right = node22

	return node1
}
