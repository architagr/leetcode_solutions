package binarytreezigzaglevelordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
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
	leftToRight := true
	arr := make([]int, 0)
	for len(queue) > 0 {
		node := pop()
		if node == nil {
			if len(queue) > 0 {
				push(nil)
			}
			x := make([]int, len(arr))
			copy(x, arr)
			if !leftToRight {
				x = reverseArr(x)
			}
			result = append(result, x)
			arr = make([]int, 0)
			leftToRight = !leftToRight
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

func reverseArr(x []int) []int {
	l, r := 0, len(x)-1
	for l <= r {
		x[l], x[r] = x[r], x[l]
		l++
		r--
	}
	return x
}
