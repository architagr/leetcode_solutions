package constructbinaryserchtreefrompreordertraversal

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	index := findNextLargestIndex(preorder)
	return &TreeNode{
		Val:   preorder[0],
		Left:  bstFromPreorder(preorder[1:index]),
		Right: bstFromPreorder(preorder[index:]),
	}

}
func findNextLargestIndex(arr []int) int {
	if len(arr) > 1 {
		for i := 1; i < len(arr); i++ {
			if arr[i] > arr[0] {
				return i
			}
		}
	}

	return len(arr)
}
