package binarysearchtreeiterator

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BSTIterator precomputes the whole in-order sequence in the constructor,
// so Next and HasNext never touch the tree. That makes the two repeatedly
// called methods trivial, at the cost of O(n) memory held for the
// iterator's lifetime and a full walk before the caller asks for anything.
//
// The problem's follow-up wants O(h) memory instead, using a stack holding
// the leftmost spine. Both are reasonable; this one is the right trade
// when the caller will consume most of the tree.
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

// inOrder appends left-node-right into this.inorder. It is a method
// rather than a free function so it can append onto the struct field
// directly, instead of threading the slice through as a return value.
func (this *BSTIterator) inOrder(node *TreeNode) {
	if node == nil {
		return
	}
	this.inOrder(node.Left)
	this.inorder = append(this.inorder, node.Val)
	this.inOrder(node.Right)
}

// Next reads and advances. No bounds check, which is safe only because
// the problem guarantees Next is called when a value exists.
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
