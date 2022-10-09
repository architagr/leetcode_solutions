package path_sum

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
	testCases = append(testCases, GetTestCase1())

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := hasPathSum(testCase.input, testCase.target)

			if got != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}

func GetTestCase1() TestCase {
	var node41 *TreeNode = new(TreeNode)
	node41.Val = 7

	var node42 *TreeNode = new(TreeNode)
	node42.Val = 2

	var node43 *TreeNode = new(TreeNode)
	node43.Val = 1

	var node31 *TreeNode = new(TreeNode)
	node31.Val = 11
	node31.Right = node41
	node31.Left = node42

	var node32 *TreeNode = nil

	var node33 *TreeNode = new(TreeNode)
	node33.Val = 13
	var node34 *TreeNode = new(TreeNode)
	node34.Val = 4
	node34.Left = node43

	var node21 *TreeNode = new(TreeNode)
	node21.Val = 4
	node21.Left = node31
	node21.Right = node32

	var node22 *TreeNode = new(TreeNode)
	node22.Val = 8
	node22.Left = node33
	node22.Right = node34

	var node1 *TreeNode = new(TreeNode)
	node1.Val = 5
	node1.Left = node21
	node1.Right = node22

	return TestCase{
		input:    node1,
		target:   22,
		expected: true,
	}
}
