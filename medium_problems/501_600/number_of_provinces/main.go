package numberofprovinces

func findCircleNum(isConnected [][]int) int {
	ajList := getAdjeList(isConnected)
	visited := make(map[int]bool, len(isConnected))
	count := 0
	for i := 0; i < len(isConnected); i++ {
		if _, alreadyVisited := visited[i]; !alreadyVisited {
			count++
			findProvince(i, visited, ajList)
		}
	}

	return count
}

func findProvince(start int, visited map[int]bool, ajList map[int][]int) {
	queue := make([]int, 0, len(ajList))
	push := func(node ...int) {
		queue = append(queue, node...)
	}
	pop := func() int {
		x := queue[0]
		queue = queue[1:]
		return x
	}
	push(start)
	for len(queue) > 0 {
		x := pop()
		if _, alreadyVisited := visited[x]; !alreadyVisited {
			push(ajList[x]...)
			visited[x] = true
		}
	}

}

func getAdjeList(isConnected [][]int) map[int][]int {
	ajList := make(map[int][]int, len(isConnected))
	for i := 0; i < len(isConnected); i++ {
		for j := i + 1; j < len(isConnected[i]); j++ {
			if isConnected[i][j] == 1 {
				ajList[i] = append(ajList[i], j)
				ajList[j] = append(ajList[j], i)
			}
		}
	}

	return ajList
}
