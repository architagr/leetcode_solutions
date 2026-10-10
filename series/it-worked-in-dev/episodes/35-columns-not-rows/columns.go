package columnsnotrows

import "sort"

// The model viewer draws a decision tree in columns: every node goes in the
// column of its x position (a left child one to the left of its parent, a
// right child one to the right), columns from left to right, and inside a
// column from the top down - and left to right where two nodes share a row.

// Node is one split or leaf.
type Node struct {
	ID          int
	Left, Right *Node
}

type placed struct{ id, col, depth int }

// ColumnsBySort is what I would write: work out every node's column and
// depth, sort by column then depth, and cut the sorted list into columns.
func ColumnsBySort(root *Node) [][]int {
	var all []placed
	var walk func(n *Node, col, depth int)
	walk = func(n *Node, col, depth int) {
		if n == nil {
			return
		}
		all = append(all, placed{n.ID, col, depth})
		walk(n.Left, col-1, depth+1)
		walk(n.Right, col+1, depth+1)
	}
	walk(root, 0, 0)
	// Stable, so two nodes in the same row and column keep the walk's
	// order, which is left to right.
	sort.SliceStable(all, func(i, j int) bool {
		if all[i].col != all[j].col {
			return all[i].col < all[j].col
		}
		return all[i].depth < all[j].depth
	})
	var out [][]int
	for i, p := range all {
		if i == 0 || p.col != all[i-1].col {
			out = append(out, nil)
		}
		out[len(out)-1] = append(out[len(out)-1], p.id)
	}
	return out
}

// ColumnsByLevel walks the tree a row at a time. Rows come out top to bottom
// and each row left to right, so appending a node to its column the moment it
// is reached puts every column in order already.
func ColumnsByLevel(root *Node) [][]int {
	if root == nil {
		return nil
	}
	type item struct {
		n   *Node
		col int
	}
	var left, right [][]int // columns -1, -2, ... and 0, 1, ...
	queue := []item{{root, 0}}
	for head := 0; head < len(queue); head++ {
		it := queue[head]
		if it.col >= 0 {
			for len(right) <= it.col {
				right = append(right, nil)
			}
			right[it.col] = append(right[it.col], it.n.ID) // arrives in order
		} else {
			c := -it.col - 1
			for len(left) <= c {
				left = append(left, nil)
			}
			left[c] = append(left[c], it.n.ID)
		}
		if it.n.Left != nil {
			queue = append(queue, item{it.n.Left, it.col - 1})
		}
		if it.n.Right != nil {
			queue = append(queue, item{it.n.Right, it.col + 1})
		}
	}
	out := make([][]int, 0, len(left)+len(right))
	for i := len(left) - 1; i >= 0; i-- {
		out = append(out, left[i])
	}
	return append(out, right...)
}
