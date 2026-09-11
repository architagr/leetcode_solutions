package lowestcommonancestorofdeepestleaves

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lcaDeepestLeaves(root *TreeNode) *TreeNode {

}

func findLca(list1, list2 []*TreeNode) *TreeNode {
	var res *TreeNode = list1[0]

	l := len(list1)
	if len(list2) < l {
		l = len(list2)
	}

	for i := 1; i < l; i++ {
		if list1[i] == list2[i] {
			res = list1[i]
		}
	}
	return res
}
