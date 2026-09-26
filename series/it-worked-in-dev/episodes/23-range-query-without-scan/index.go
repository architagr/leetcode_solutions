package rangequerywithoutscan

// A price index: a binary search tree keyed by price, in paise. Everything in
// a node's left subtree is cheaper, everything in its right subtree dearer.
type Node struct {
	Price       int
	Left, Right *Node
}

// RangeByScan is what I would write. Walk the whole index in order - left,
// node, right, so the prices come out sorted - and keep the ones in range.
//
// It is the definition of "every price between lo and hi, in order".
func RangeByScan(root *Node, lo, hi int) []int {
	var out []int
	var walk func(n *Node)
	walk = func(n *Node) {
		if n == nil {
			return
		}
		walk(n.Left)
		if n.Price >= lo && n.Price <= hi {
			out = append(out, n.Price)
		}
		walk(n.Right)
	}
	walk(root)
	return out
}

// RangeByPruning is the same in-order walk, except it does not go down a side
// that cannot contain anything in range.
func RangeByPruning(root *Node, lo, hi int) []int {
	var out []int
	var walk func(n *Node)
	walk = func(n *Node) {
		if n == nil {
			return
		}
		// Everything on the left is cheaper than n. If n is already below
		// lo, so is all of it.
		if n.Price > lo {
			walk(n.Left)
		}
		if n.Price >= lo && n.Price <= hi {
			out = append(out, n.Price)
		}
		// Everything on the right is dearer than n. If n is already above
		// hi, so is all of it.
		if n.Price < hi {
			walk(n.Right)
		}
	}
	walk(root)
	return out
}
