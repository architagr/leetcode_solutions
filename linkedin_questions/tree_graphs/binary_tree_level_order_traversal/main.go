package binarytreelevelordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	result := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root, nil)
	pop := func() *TreeNode {
		node := queue[0]
		queue = queue[1:]
		return node
	}
	push := func(node *TreeNode) {
		queue = append(queue, node)
	}
	arr := make([]int, 0)
	for len(queue) > 0 {
		node := pop()
		if node == nil {
			if len(queue) > 0 {
				push(nil)
			}
			x := make([]int, len(arr))
			copy(x, arr)
			result = append(result, x)
			arr = make([]int, 0)
			continue
		}
		arr = append(arr, node.Val)
		if node.Left != nil {
			push(node.Left)
		}
		if node.Right != nil {
			push(node.Right)
		}
	}

	return result
}
