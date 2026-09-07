package flatten_binary_tree_to_linked_list

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Flatten(root *TreeNode) {
	if root == nil {
		return
	}
	temp := root
	arr := make([]TreeNode, 0)
	getPreOrderArray(temp, &arr)

	root.Val = arr[0].Val
	temp = root
	temp.Left = nil

	for i := 1; i < len(arr); i++ {
		y := new(TreeNode)
		y.Val = arr[i].Val
		temp.Right = y
		temp.Left = nil
		temp = temp.Right
	}

}

func getPreOrderArray(root *TreeNode, arr *[]TreeNode) {

	if root == nil {
		return
	}

	(*arr) = append((*arr), TreeNode{
		Val: root.Val,
	})
	getPreOrderArray(root.Left, arr)
	getPreOrderArray(root.Right, arr)

}
