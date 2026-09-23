package cousinsinbinarytreeii

import "errors"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func replaceValueInTree(root *TreeNode) *TreeNode {
	levelOrderSumArr := levelOrderSum(root)
	root.Val = 0
	parse(root, 0, levelOrderSumArr)
	return root
}

func parse(node *TreeNode, level int, sumArr []int) {
	if node.Left == nil && node.Right == nil {
		return
	}
	currSum := 0
	if node.Left != nil {
		currSum += node.Left.Val
	}
	if node.Right != nil {
		currSum += node.Right.Val
	}

	if node.Left != nil {
		node.Left.Val = sumArr[level+1] - currSum
		parse(node.Left, level+1, sumArr)
	}
	if node.Right != nil {
		node.Right.Val = sumArr[level+1] - currSum
		parse(node.Right, level+1, sumArr)
	}
}

func levelOrderSum(root *TreeNode) []int {
	res := make([]int, 0)

	queue := make([]*TreeNode, 0)
	queue = append(queue, root, nil)
	push := func(n *TreeNode) {
		if n == nil && len(queue) == 0 {
			return
		}

		queue = append(queue, n)
	}
	pop := func() (*TreeNode, error) {
		if len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]
			return node, nil
		}
		return nil, errors.New("queue is empty")
	}

	sum := 0
	for len(queue) > 0 {
		node, err := pop()
		if err != nil {
			break
		}
		if node == nil {
			res = append(res, sum)
			sum = 0
			push(nil)
			continue
		}
		sum = sum + node.Val
		if node.Left != nil {
			push(node.Left)
		}
		if node.Right != nil {
			push(node.Right)
		}
	}
	return res
}
