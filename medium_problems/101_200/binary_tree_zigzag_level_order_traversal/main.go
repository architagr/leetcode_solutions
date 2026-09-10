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
	// The nil is a level marker: everything ahead of it in the queue
	// belongs to the current level. Popping it means that level just
	// drained, which is the boundary a depth-first walk never learns.
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
			// Guarded: re-pushing the sentinel onto an otherwise empty
			// queue would spin the loop forever on it.
			if len(queue) > 0 {
				push(nil)
			}
			// The copy is required, not defensive. arr is reused for the
			// next level and reverseArr mutates in place, so appending arr
			// itself would alias one backing array into every result row.
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

// reverseArr reverses in place. l <= r rather than l < r lets an
// odd-length slice swap its middle element with itself, which is a no-op
// and keeps the condition simple.
func reverseArr(x []int) []int {
	l, r := 0, len(x)-1
	for l <= r {
		x[l], x[r] = x[r], x[l]
		l++
		r--
	}
	return x
}
