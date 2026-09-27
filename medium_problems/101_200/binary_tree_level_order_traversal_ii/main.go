package binarytreelevelordertraversalii

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	q := make([]*TreeNode, 0)

	push := func(node *TreeNode) {
		q = append(q, node)
	}
	pop := func() *TreeNode {
		x := q[0]
		q = q[1:]
		return x
	}
	push(root)
	push(nil)
	result := make([][]int, 0)
	level := 0
	result = append(result, []int{})
	for len(q) > 0 {
		n := pop()
		// The nil marker ends a level. If nodes remain, they are the whole
		// next level: open a new slot and mark its end.
		if n == nil {
			if len(q) > 0 {
				result = append(result, []int{})
				level++
				push(nil)
			}
			continue
		}
		result[level] = append(result[level], n.Val)
		if n.Left != nil {
			push(n.Left)
		}
		if n.Right != nil {
			push(n.Right)
		}
	}
	// Bottom-up is top-down with the levels flipped. Reversing once is
	// O(levels); prepending each level as it's found would copy every time.
	return reverse(result)
}
func reverse(arr [][]int) [][]int {
	i, j := 0, len(arr)-1
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	return arr
}
