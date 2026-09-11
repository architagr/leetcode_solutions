package clonegraph

// Definition for a Node.
type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	visited := make(map[int]*Node)
	return clone(node, visited)
}
func clone(node *Node, visited map[int]*Node) *Node {
	if node == nil {
		return node
	}
	newNode, ok := visited[node.Val]
	if ok {
		return newNode
	}
	newNode = new(Node)
	newNode.Val = node.Val
	visited[node.Val] = newNode

	if len(node.Neighbors) > 0 {
		newNode.Neighbors = make([]*Node, len(node.Neighbors))
		for i, nNode := range node.Neighbors {
			newNode.Neighbors[i] = clone(nNode, visited)
		}
	}

	return newNode
}
