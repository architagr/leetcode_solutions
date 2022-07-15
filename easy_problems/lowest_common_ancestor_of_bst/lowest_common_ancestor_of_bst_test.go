package lowest_common_ancestor_of_bst

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input  *TreeNode
	inputp *TreeNode
	inputq *TreeNode

	expected *TreeNode
}

func TestLowestCommonAncestor(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, GetTestCase1())
	testCases = append(testCases, GetTestCase2())
	testCases = append(testCases, GetTestCase3())
	testCases = append(testCases, GetTestCase4())

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := LowestCommonAncestor(testCase.input, testCase.inputp, testCase.inputq)

			if got.Val != testCase.expected.Val {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}

func GetTestCase1() TestCase {
	var node41 *TreeNode = new(TreeNode)
	node41.Val = 3

	var node42 *TreeNode = new(TreeNode)
	node42.Val = 5

	var node31 *TreeNode = new(TreeNode)
	node31.Val = 6

	var node32 *TreeNode = new(TreeNode)
	node32.Val = 4
	node32.Left = node41
	node32.Right = node42

	var node33 *TreeNode = new(TreeNode)
	node33.Val = 7
	var node34 *TreeNode = new(TreeNode)
	node34.Val = 9

	var node21 *TreeNode = new(TreeNode)
	node21.Val = 2
	node21.Left = node31
	node21.Right = node32

	var node22 *TreeNode = new(TreeNode)
	node22.Val = 8
	node22.Left = node33
	node22.Right = node34

	var node1 *TreeNode = new(TreeNode)
	node1.Val = 6
	node1.Left = node21
	node1.Right = node22

	return TestCase{
		input:    node1,
		inputp:   node21,
		inputq:   node22,
		expected: node1,
	}
}

func GetTestCase2() TestCase {
	var node41 *TreeNode = new(TreeNode)
	node41.Val = 3

	var node42 *TreeNode = new(TreeNode)
	node42.Val = 5

	var node31 *TreeNode = new(TreeNode)
	node31.Val = 6

	var node32 *TreeNode = new(TreeNode)
	node32.Val = 4
	node32.Left = node41
	node32.Right = node42

	var node33 *TreeNode = new(TreeNode)
	node33.Val = 7
	var node34 *TreeNode = new(TreeNode)
	node34.Val = 9

	var node21 *TreeNode = new(TreeNode)
	node21.Val = 2
	node21.Left = node31
	node21.Right = node32

	var node22 *TreeNode = new(TreeNode)
	node22.Val = 8
	node22.Left = node33
	node22.Right = node34

	var node1 *TreeNode = new(TreeNode)
	node1.Val = 6
	node1.Left = node21
	node1.Right = node22

	return TestCase{
		input:    node1,
		inputp:   node21,
		inputq:   node32,
		expected: node21,
	}
}

func GetTestCase3() TestCase {

	var node21 *TreeNode = new(TreeNode)
	node21.Val = 1

	var node1 *TreeNode = new(TreeNode)
	node1.Val = 2
	node1.Left = node21

	return TestCase{
		input:    node1,
		inputp:   node1,
		inputq:   node21,
		expected: node1,
	}
}

func GetTestCase4() TestCase {

	var node21 *TreeNode = new(TreeNode)
	node21.Val = 1

	var node22 *TreeNode = new(TreeNode)
	node22.Val = 3

	var node1 *TreeNode = new(TreeNode)
	node1.Val = 2
	node1.Left = node21
	node1.Right = node22

	return TestCase{
		input:    node1,
		inputp:   node22,
		inputq:   node21,
		expected: node1,
	}
}
