package construct_b_tree_preorder_traversal

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	inputA, inputB []int
	expected       []int
}

func TestBuildBinaryTree(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		inputA:   []int{1, 2, 4, 3, 5, 7, 6, 8, 9},
		inputB:   []int{4, 2, 1, 7, 5, 3, 8, 6, 9},
		expected: []int{1, 2, 4, 3, 5, 7, 6, 8, 9},
	})
	testCases = append(testCases, TestCase{
		inputA:   []int{1, 2, 3, 4, 5},
		inputB:   []int{1, 2, 3, 4, 5},
		expected: []int{1, 2, 3, 4, 5},
	})
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := BuildBinaryTree(testCase.inputA, testCase.inputB)
			x := PreorderTraversal(got)
			if !reflect.DeepEqual(x, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}

func PreorderTraversal(A *TreeNode) []int {
	arr := make([]int, 0)
	return traversal(A, arr)
}

func traversal(A *TreeNode, arr []int) []int {
	if A == nil {
		return arr
	}
	arr = append(arr, A.Val)
	arr = traversal(A.Left, arr)
	arr = traversal(A.Right, arr)

	return arr
}
