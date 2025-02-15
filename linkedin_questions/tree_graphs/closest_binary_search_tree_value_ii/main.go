package closestbinarysearchtreevalueii

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	minIndex int
	min      float64
)

func closestKValues(root *TreeNode, target float64, k int) []int {
	// get inorder of the tree while getting inorder get the node and index with min diff
	// use merge sort logic to merge 2 half of the array till we have k elements
	min = math.MaxInt64
	minIndex = -1
	inOrderArr := make([]int, 0)
	inOrderTraversal(root, target, &inOrderArr)
	l := 0
	l++
	i, j := minIndex-1, minIndex+1
	for i >= 0 && j < len(inOrderArr) && l < k {
		if absDiff(inOrderArr[i], target) < absDiff(inOrderArr[j], target) {
			i--
		} else {
			j++
		}
		l++
	}
	for ; i >= 0 && l < k; l++ {
		i--
	}
	for ; j < len(inOrderArr) && l < k; l++ {
		j++
	}

	return inOrderArr[i+1 : j]
}

func inOrderTraversal(root *TreeNode, target float64, arr *[]int) {
	if root == nil {
		return
	}
	inOrderTraversal(root.Left, target, arr)
	diff := absDiff(root.Val, target)
	if min > diff {
		min = diff
		minIndex = len(*arr)
	}
	(*arr) = append((*arr), root.Val)
	inOrderTraversal(root.Right, target, arr)
}

func absDiff(a int, target float64) float64 {
	x := float64(a) - target
	if x < 0 {
		x *= -1
	}

	return x
}
