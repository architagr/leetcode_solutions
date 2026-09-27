package closestpricelookup

import "sort"

// The same price index as episode 23: a binary search tree keyed by price in
// paise. Left is cheaper, right is dearer.
type Node struct {
	Price       int
	Left, Right *Node
}

// Sorted is the function the listing page already had: every price, in
// order. Written the way day 50 writes it.
func Sorted(n *Node) []int {
	if n == nil {
		return []int{}
	}
	return append(Sorted(n.Left), append([]int{n.Price}, Sorted(n.Right)...)...)
}

// ClosestBySorting is what I would write. "Show me the product nearest ₹X":
// take the sorted prices, binary search for X, and compare the two either side.
// Ties go to the cheaper one.
func ClosestBySorting(root *Node, target int) int {
	prices := Sorted(root)
	if len(prices) == 0 {
		return -1
	}
	i := sort.SearchInts(prices, target)
	switch {
	case i == 0:
		return prices[0]
	case i == len(prices):
		return prices[len(prices)-1]
	}
	lo, hi := prices[i-1], prices[i]
	if target-lo <= hi-target {
		return lo
	}
	return hi
}

// ClosestByWalk is day 47's walk: in order, keeping the previous price, and
// stopping as soon as it reaches a price at or above the target - the answer
// is that price or the one just before it.
func ClosestByWalk(root *Node, target int) int {
	prev, answer, done := -1, -1, false
	var walk func(n *Node)
	walk = func(n *Node) {
		if n == nil || done {
			return
		}
		walk(n.Left)
		if done {
			return
		}
		if n.Price >= target {
			done = true
			if prev >= 0 && target-prev <= n.Price-target {
				answer = prev
			} else {
				answer = n.Price
			}
			return
		}
		prev = n.Price
		walk(n.Right)
	}
	walk(root)
	if !done {
		return prev // every price is below the target
	}
	return answer
}

// ClosestByDescent walks down from the root once. The two candidates are the
// last price seen below the target and the last seen at or above it, and both
// are on the search path to the target, so nothing off that path is read.
func ClosestByDescent(root *Node, target int) int {
	below, above := -1, -1
	for n := root; n != nil; {
		if n.Price < target {
			below = n.Price // the best "just under" so far; anything closer is to the right
			n = n.Right
		} else {
			above = n.Price // the best "at or over" so far; anything closer is to the left
			n = n.Left
		}
	}
	switch {
	case below < 0:
		return above
	case above < 0:
		return below
	case target-below <= above-target:
		return below
	}
	return above
}

// SortedFlat builds the same list appending into one slice, so the only
// cost left is the list itself: every price read and stored, once.
func SortedFlat(n *Node, out []int) []int {
	if n == nil {
		return out
	}
	out = SortedFlat(n.Left, out)
	out = append(out, n.Price)
	return SortedFlat(n.Right, out)
}

// ClosestBySortedFlat is ClosestBySorting on the flat list: the fairest
// version of "build the sorted prices, then binary search".
func ClosestBySortedFlat(root *Node, target int) int {
	prices := SortedFlat(root, nil)
	if len(prices) == 0 {
		return -1
	}
	i := sort.SearchInts(prices, target)
	switch {
	case i == 0:
		return prices[0]
	case i == len(prices):
		return prices[len(prices)-1]
	}
	lo, hi := prices[i-1], prices[i]
	if target-lo <= hi-target {
		return lo
	}
	return hi
}
