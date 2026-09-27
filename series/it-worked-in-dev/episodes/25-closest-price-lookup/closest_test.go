package closestpricelookup

import (
	"math/rand"
	"testing"
)

func insert(root *Node, p int) *Node {
	if root == nil {
		return &Node{Price: p}
	}
	n := root
	for {
		if p < n.Price {
			if n.Left == nil {
				n.Left = &Node{Price: p}
				return root
			}
			n = n.Left
		} else {
			if n.Right == nil {
				n.Right = &Node{Price: p}
				return root
			}
			n = n.Right
		}
	}
}

// catalog is n distinct prices, spaced unevenly so there are real gaps to be
// closest across, inserted in random order.
func catalog(n int, seed int64) (*Node, []int) {
	r := rand.New(rand.NewSource(seed))
	prices := make([]int, n)
	p := 0
	for i := range prices {
		p += 1 + r.Intn(50)
		prices[i] = p
	}
	var root *Node
	for _, i := range r.Perm(n) {
		root = insert(root, prices[i])
	}
	return root, prices
}

func TestAllThreeAgree(t *testing.T) {
	root, prices := catalog(3000, 1)
	targets := []int{-5, 0, prices[0], prices[0] + 1, prices[len(prices)-1], prices[len(prices)-1] + 999}
	r := rand.New(rand.NewSource(2))
	for i := 0; i < 500; i++ {
		targets = append(targets, r.Intn(prices[len(prices)-1]+100))
	}
	// the exact midpoint between two neighbours: a tie, which goes cheaper
	targets = append(targets, (prices[100]+prices[101])/2)
	for _, tg := range targets {
		a, b, c := ClosestBySorting(root, tg), ClosestByWalk(root, tg), ClosestByDescent(root, tg)
		if d := ClosestBySortedFlat(root, tg); d != a {
			t.Fatalf("target %d: sorting %d, flat %d", tg, a, d)
		}
		if a != b || b != c {
			t.Fatalf("target %d: sorting %d, walk %d, descent %d", tg, a, b, c)
		}
	}
	single := &Node{Price: 40}
	for _, tg := range []int{0, 40, 99} {
		if ClosestByDescent(single, tg) != 40 || ClosestByWalk(single, tg) != 40 {
			t.Fatal("a one-product index should always answer that product")
		}
	}
}
