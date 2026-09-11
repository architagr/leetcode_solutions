package balanceabinarysearchtree

import (
	"reflect"
	"sort"
	"testing"
)

// buildTree rebuilds a tree from the level-order form the examples use,
// where nil marks a missing child.
func buildTree(vals []any) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}
	root := &TreeNode{Val: vals[0].(int)}
	queue := []*TreeNode{root}
	i := 1
	for len(queue) > 0 && i < len(vals) {
		node := queue[0]
		queue = queue[1:]
		if i < len(vals) {
			if vals[i] != nil {
				node.Left = &TreeNode{Val: vals[i].(int)}
				queue = append(queue, node.Left)
			}
			i++
		}
		if i < len(vals) {
			if vals[i] != nil {
				node.Right = &TreeNode{Val: vals[i].(int)}
				queue = append(queue, node.Right)
			}
			i++
		}
	}
	return root
}

func inorder(n *TreeNode, out *[]int) {
	if n == nil {
		return
	}
	inorder(n.Left, out)
	*out = append(*out, n.Val)
	inorder(n.Right, out)
}

// height returns -1 when the subtree is not height-balanced.
func height(n *TreeNode) int {
	if n == nil {
		return 0
	}
	l, r := height(n.Left), height(n.Right)
	if l < 0 || r < 0 || l-r > 1 || r-l > 1 {
		return -1
	}
	if l > r {
		return l + 1
	}
	return r + 1
}

// The statement says any balanced BST holding the same values is
// accepted, so compare the properties rather than the shape.
func TestBalanceBST(t *testing.T) {
	tests := []struct {
		name string
		root []any
	}{
		{name: "example 1", root: []any{1, nil, 2, nil, 3, nil, 4, nil, nil}},
		{name: "example 2", root: []any{2, 1, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var want []int
			inorder(buildTree(tt.root), &want)
			sort.Ints(want)

			got := balanceBST(buildTree(tt.root))

			var vals []int
			inorder(got, &vals)
			if !sort.IntsAreSorted(vals) {
				t.Errorf("result is not a BST, inorder = %v", vals)
			}
			if !reflect.DeepEqual(vals, want) {
				t.Errorf("values = %v, want %v", vals, want)
			}
			if height(got) < 0 {
				t.Errorf("result is not height-balanced, inorder = %v", vals)
			}
		})
	}
}
