package closestbinarysearchtreevalueii

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Package-level rather than parameters, which works only because
// closestKValues resets both at the top of every call. It does make the
// function non-reentrant: two goroutines calling it concurrently would
// corrupt each other's search. Threading them through as pointers, the way
// arr already is, would remove that.
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
	// Starts at 1, not 0: the closest element itself is already part of
	// the answer, so only k-1 more are needed.
	l := 0
	l++
	// Expand outward from the closest value, always consuming whichever
	// neighbour is nearer the target. Greedy is safe because distance
	// grows monotonically as you move away in either direction - this is
	// a merge step run outward from a centre.
	//
	// It works at all because the k closest values in a SORTED array are
	// always contiguous: any answer that skipped a nearer value for a
	// further one could swap them and be no worse.
	i, j := minIndex-1, minIndex+1
	for i >= 0 && j < len(inOrderArr) && l < k {
		if absDiff(inOrderArr[i], target) < absDiff(inOrderArr[j], target) {
			i--
		} else {
			j++
		}
		l++
	}
	// One side ran out before k values were collected, so the other
	// supplies the rest. Only one of these two loops can ever execute.
	for ; i >= 0 && l < k; l++ {
		i--
	}
	for ; j < len(inOrderArr) && l < k; l++ {
		j++
	}

	// Both pointers sit one step outside the collected range, so i+1 and j
	// are exactly the inclusive-exclusive bounds of the window.
	return inOrderArr[i+1 : j]
}

func inOrderTraversal(root *TreeNode, target float64, arr *[]int) {
	if root == nil {
		return
	}
	inOrderTraversal(root.Left, target, arr)
	// minIndex is captured before the append, which is exactly the index
	// this value is about to occupy. The flatten and the closest-value
	// search are therefore the same pass.
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
