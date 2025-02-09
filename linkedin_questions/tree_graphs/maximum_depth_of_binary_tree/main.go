package maximumdepthofbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := 0
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
	for len(queue) > 0 {
		node := pop()
		if node == nil {
			if len(queue) > 0 {
				push(nil)
			}
			max++
			continue
		}
		if node.Left != nil {
			push(node.Left)
		}
		if node.Right != nil {
			push(node.Right)
		}
	}

	return max
}
