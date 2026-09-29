package findbottomlefttreevalue

import "errors"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	var result *TreeNode
	queue := []*TreeNode{root, nil}
	push := func(n *TreeNode) {
		if n == nil && len(queue) == 0 {
			return
		}
		queue = append(queue, n)
	}
	pop := func() (*TreeNode, error) {
		if len(queue) == 0 {
			return nil, errors.New("queue is empty")
		}
		n := queue[0]
		queue = queue[1:]
		return n, nil
	}
	for len(queue) > 0 {
		n, err := pop()
		if err != nil {
			break
		}
		if n == nil {
			push(n)
			if len(queue) > 0 {
				result = nil
			}
			continue
		}
		if n.Left != nil {
			push(n.Left)
		}
		if n.Right != nil {
			push(n.Right)
		}
		if result == nil {
			result = n
		}
	}
	return result.Val
}
