package findmodeinbinarysarchtree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	data := make(map[int]int)
	data = getCnt(root, data)

	res := make([]int, 0)
	cnt := 0
	for _, v := range data {
		if v > cnt {
			cnt = v
		}
	}

	for k, v := range data {
		if v == cnt {
			res = append(res, k)
		}
	}
	return res
}

func getCnt(root *TreeNode, data map[int]int) map[int]int {
	if root == nil {
		return nil
	}
	data[root.Val]++
	getCnt(root.Left, data)
	getCnt(root.Right, data)
	return data

}
