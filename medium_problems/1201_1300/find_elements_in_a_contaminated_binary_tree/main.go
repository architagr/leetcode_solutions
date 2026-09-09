package findelementsinacontaminatedbinarytree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type FindElements struct {
	nodes map[int]struct{}
	tree  *TreeNode
}

func Constructor(root *TreeNode) FindElements {
	root.Val = 0
	data := make(map[int]struct{})
	data[root.Val] = struct{}{}
	updateNode(root, data)

	return FindElements{
		tree:  root,
		nodes: data,
	}
}

func updateNode(node *TreeNode, data map[int]struct{}) {
	if node == nil {
		return
	}
	if node.Left != nil {
		node.Left.Val = (2 * node.Val) + 1
		data[node.Left.Val] = struct{}{}
		updateNode(node.Left, data)
	}
	if node.Right != nil {
		node.Right.Val = (2 * node.Val) + 2
		data[node.Right.Val] = struct{}{}
		updateNode(node.Right, data)
	}
}

func (this *FindElements) Find(target int) bool {
	_, ok := this.nodes[target]
	return ok
}

/**
 * Your FindElements object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Find(target);
 */
