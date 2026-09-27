package createbinarytreefromdescriptions

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createBinaryTree(descriptions [][]int) *TreeNode {
	// value -> node, created the first time a value appears on either side,
	// so edges can arrive in any order. c collects every child value.
	m := make(map[int]*TreeNode, 2*len(descriptions))
	c := make(map[int]struct{}, 2*len(descriptions))
	for i := range descriptions {
		node, nodeFound := m[descriptions[i][0]]
		childNode, childNodeFound := m[descriptions[i][1]]
		if !nodeFound {
			node = &TreeNode{
				Val: descriptions[i][0],
			}
			m[node.Val] = node
		}
		if !childNodeFound {
			childNode = &TreeNode{
				Val: descriptions[i][1],
			}
			m[childNode.Val] = childNode
		}
		c[childNode.Val] = struct{}{}
		if descriptions[i][2] == 1 {
			node.Left = childNode
		} else {
			node.Right = childNode
		}
	}
	// The root is the only node that is never a child.
	for v := range c {
		delete(m, v)
	}
	var root *TreeNode

	for _, n := range m {
		root = n
	}
	return root
}
