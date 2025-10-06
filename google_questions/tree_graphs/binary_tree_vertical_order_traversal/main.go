package binarytreeverticalordertraversal

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type customNode struct {
	TreeNode
	order int
}

func verticalOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	m := make(map[int][]int)
	q := make([]*customNode, 0)
	push := func(n *TreeNode, o int) {
		l, ok := m[o]
		if !ok {
			l = make([]int, 0)
		}
		l = append(l, n.Val)
		m[o] = l
		q = append(q, &customNode{TreeNode: *n, order: o})
	}
	pop := func() *customNode {
		n := q[0]
		q = q[1:]
		return n
	}
	push(root, 0)

	for len(q) > 0 {
		n := pop()
		if n.Left != nil {
			push(n.Left, n.order-1)
		}
		if n.Right != nil {
			push(n.Right, n.order+1)
		}
	}

	result := make([][]int, 0)
	for i := -101; i <= 101; i++ {
		if l, ok := m[i]; ok {
			result = append(result, l)
		}
	}
	return result
}
