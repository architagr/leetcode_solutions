package two_sum_iv_input_is_a_bst

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    *TreeNode
	target   int
	expected bool
}

func TestLowestCommonAncestor(t *testing.T) {
	testCases := make([]TestCase, 0)
	// testCases = append(testCases, GetTestCase1())
	testCases = append(testCases, GetTestCase2())

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := findTarget(testCase.input, testCase.target)

			if got != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}

func GetTestCase1() TestCase {

	var node31 *TreeNode = new(TreeNode)
	node31.Val = 2

	var node32 *TreeNode = new(TreeNode)
	node32.Val = 4

	var node33 *TreeNode = new(TreeNode)
	node33.Val = 7

	var node21 *TreeNode = new(TreeNode)
	node21.Val = 3
	node21.Left = node31
	node21.Right = node32

	var node22 *TreeNode = new(TreeNode)
	node22.Val = 6
	node22.Right = node33

	var node1 *TreeNode = new(TreeNode)
	node1.Val = 5
	node1.Left = node21
	node1.Right = node22

	return TestCase{
		input:    node1,
		target:   28,
		expected: false,
	}
}

func GetTestCase2() TestCase {

	var node31 *TreeNode = new(TreeNode)
	node31.Val = 2

	var node32 *TreeNode = new(TreeNode)
	node32.Val = 4

	var node33 *TreeNode = new(TreeNode)
	node33.Val = 7

	var node21 *TreeNode = new(TreeNode)
	node21.Val = 3
	node21.Left = node31
	node21.Right = node32

	var node22 *TreeNode = new(TreeNode)
	node22.Val = 6
	node22.Right = node33

	var node1 *TreeNode = new(TreeNode)
	node1.Val = 5
	node1.Left = node21
	node1.Right = node22

	return TestCase{
		input:    node1,
		target:   9,
		expected: true,
	}
}
