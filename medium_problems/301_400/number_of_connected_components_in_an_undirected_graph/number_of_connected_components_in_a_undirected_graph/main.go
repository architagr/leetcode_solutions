package numberofconnectedcomponentsinaundirectedgraph

func countComponents(n int, edges [][]int) int {
	visitedArr := buildVisitedArray(n)
	adjencyMap := buildAdjencyMap(n, edges)
	count := 0
	for i := 0; i < n; i++ {
		if !visitedArr[i] {
			count++
			visitedArr[i] = true
			bfs(&visitedArr, adjencyMap, []int{i})
		}
	}
	return count
}

func bfs(visitedArr *[]bool, adjencyMap map[int][]int, queue []int) {
	for len(queue) > 0 {
		x := queue[0]
		queue = queue[1:]

		connectNodes := adjencyMap[x]
		for _, node := range connectNodes {
			if !(*visitedArr)[node] {
				queue = append(queue, node)
				(*visitedArr)[node] = true
			}
		}
	}
}

func buildVisitedArray(n int) []bool {
	return make([]bool, n)
}

func buildAdjencyMap(n int, edges [][]int) map[int][]int {
	adjencyMap := make(map[int][]int, n)

	for _, edge := range edges {
		a, b := edge[0], edge[1]
		connectedToA, existA := adjencyMap[a]
		if !existA {
			connectedToA = make([]int, 0)
		}
		connectedToA = append(connectedToA, b)
		connectedToB, existb := adjencyMap[b]
		if !existb {
			connectedToB = make([]int, 0)
		}
		connectedToB = append(connectedToB, a)
		adjencyMap[a] = connectedToA
		adjencyMap[b] = connectedToB
	}
	return adjencyMap
}
