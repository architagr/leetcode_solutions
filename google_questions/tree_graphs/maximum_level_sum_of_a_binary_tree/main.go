package maximumlevelsumofabinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	queue := make([]*TreeNode, 0)
	push := func(node *TreeNode) {
		queue = append(queue, node)
	}
	pop := func() *TreeNode {
		x := queue[0]
		queue = queue[1:]
		return x
	}
	currentLevel := 1
	maxLevel := 1
	maxSum := root.Val
	sum := 0
	push(root)
	push(nil)
	for len(queue) > 0 {
		node := pop()
		if node == nil {
			if sum > maxSum {
				maxLevel = currentLevel
				maxSum = sum
			}
			currentLevel++
			sum = 0
			if len(queue) > 0 {
				push(nil)
			}
			continue
		}
		sum += node.Val
		if node.Left != nil {
			push(node.Left)
		}
		if node.Right != nil {
			push(node.Right)
		}
	}
	return maxLevel
}
