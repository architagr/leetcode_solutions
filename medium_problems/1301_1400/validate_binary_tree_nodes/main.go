package validatebinarytreenodes

func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	root := getRoot(n, leftChild, rightChild)
	if root == -1 {
		return false
	}

	visited := map[int]bool{root: true}
	q := []int{root}

	for len(q) != 0 {
		cur := q[0]
		q = q[1:]

		nei := [2]int{
			leftChild[cur], rightChild[cur],
		}

		for _, node := range nei {
			if node == -1 {
				continue
			}
			if _, ok := visited[node]; ok {
				return false
			}

			q = append(q, node)
			visited[node] = true
		}

	}

	return len(visited) == n
}

func getRoot(n int, left []int, right []int) int {
	seenMap := map[int]bool{}

	for _, node := range left {
		seenMap[node] = true
	}

	for _, node := range right {
		seenMap[node] = true
	}

	for i := range n {
		if !seenMap[i] {
			return i
		}
	}

	return -1
}
