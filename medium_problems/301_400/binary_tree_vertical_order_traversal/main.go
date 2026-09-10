package binarytreeverticalordertraversal

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// customNode pairs a node with its column index so the queue can carry
// both. It embeds TreeNode by value, so constructing one copies the node -
// harmless here, since only Val and the child pointers are read and those
// pointers still refer to the real children, but embedding a *TreeNode
// would avoid the copy.
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
	// push records the value into its column AND enqueues the node, so the
	// map is built at discovery time rather than on pop.
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

	// Breadth-first is required here, unlike Day 15's depth-first level
	// order. This problem needs each column ordered top to bottom, and a
	// BFS visits by increasing depth, so appending as it goes is already
	// correct. A DFS would drive one branch to the bottom first and could
	// append a deep node to a column ahead of a shallower one.
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
	// Leans on the constraint of at most 100 nodes, which bounds any column
	// index to [-100, 100]. Collecting the map's keys and sorting them
	// would not depend on that constraint holding.
	for i := -101; i <= 101; i++ {
		if l, ok := m[i]; ok {
			result = append(result, l)
		}
	}
	return result
}
