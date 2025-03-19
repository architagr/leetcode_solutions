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
	push(nil)
	var prev *Node
	current := root

	for len(q) > 0 {
		current = pop()
		if current == nil {
			if len(q) > 0 {
				push(nil)
			}
			prev = current
			continue
		}
		current.Next = prev
		prev = current
		if current.Right != nil {
			push(current.Right)
		}

		if current.Left != nil {
			push(current.Left)
		}
	}
}
