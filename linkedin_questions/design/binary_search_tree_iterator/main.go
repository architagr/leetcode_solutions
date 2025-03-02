package binarysearchtreeiterator

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	inorder []int
	root    *TreeNode
	index   int
}

func Constructor(root *TreeNode) BSTIterator {
	obj := &BSTIterator{
		root:    root,
		index:   0,
		inorder: make([]int, 0),
	}
	obj.inOrder(root)
	return *obj
}

func (this *BSTIterator) inOrder(node *TreeNode) {
	if node == nil {
		return
	}
	this.inOrder(node.Left)
	this.inorder = append(this.inorder, node.Val)
	this.inOrder(node.Right)
}

func (this *BSTIterator) Next() int {
	val := this.inorder[this.index]
	this.index++
	return val
}

func (this *BSTIterator) HasNext() bool {
	return this.index < len(this.inorder)
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
