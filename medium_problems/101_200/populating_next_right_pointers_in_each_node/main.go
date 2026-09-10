package populatingnextrightpointersineachnode

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	bfs(root)
	return root
}

func bfs(root *Node) {
	if root == nil {
		return
	}
	q := make([]*Node, 0)
	push := func(n *Node) {
		q = append(q, n)
	}
	pop := func() *Node {
		x := q[0]
		q = q[1:]
		return x
	}
	push(root)
	// Level marker, as in Day 16: everything ahead of it is one level.
	push(nil)
	var prev *Node
	current := root

	for len(q) > 0 {
		current = pop()
		if current == nil {
			// Guarded, or the final sentinel is re-pushed onto an empty
			// queue and the loop spins forever.
			if len(q) > 0 {
				push(nil)
			}
			// current is nil here, so this resets prev. The next node
			// popped is the RIGHTMOST of the new level and gets Next = nil,
			// which is exactly the rule for it - reached by the general
			// assignment below rather than by a special case.
			prev = current
			continue
		}
		// Because the walk goes right to left, the node to current's right
		// is simply the node visited just before it. No lookahead, and
		// nothing is written into a node already passed.
		current.Next = prev
		prev = current
		// RIGHT before LEFT. This one inversion is the whole idea: it makes
		// the BFS visit each level backwards, which is what turns "my next
		// right node" into "the node I just visited".
		if current.Right != nil {
			push(current.Right)
		}

		if current.Left != nil {
			push(current.Left)
		}
	}
}
