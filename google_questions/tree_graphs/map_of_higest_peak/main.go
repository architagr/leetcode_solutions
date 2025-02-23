package mapofhigestpeak

type state struct {
	i, j  int
	count int
}

func highestPeak(isWater [][]int) [][]int {

	queue := make([]state, 0, len(isWater))
	result := make([][]int, len(isWater))
	for i := 0; i < len(isWater); i++ {
		result[i] = make([]int, len(isWater[i]))
		for j := 0; j < len(isWater[i]); j++ {
			if isWater[i][j] == 1 {
				result[i][j] = 0
				queue = append(queue, state{i: i, j: j, count: 0})
			} else {
				result[i][j] = -1
			}
		}
	}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// update north
		if current.i > 0 && result[current.i-1][current.j] == -1 {
			result[current.i-1][current.j] = current.count + 1
			queue = append(queue, state{i: current.i - 1, j: current.j, count: current.count + 1})
		}
		// update south
		if current.i < len(isWater)-1 && result[current.i+1][current.j] == -1 {
			result[current.i+1][current.j] = current.count + 1
			queue = append(queue, state{i: current.i + 1, j: current.j, count: current.count + 1})
		}

		// update west
		if current.j > 0 && result[current.i][current.j-1] == -1 {
			result[current.i][current.j-1] = current.count + 1
			queue = append(queue, state{i: current.i, j: current.j - 1, count: current.count + 1})
		}
		// update east
		if current.j < len(isWater[current.i])-1 && result[current.i][current.j+1] == -1 {
			result[current.i][current.j+1] = current.count + 1
			queue = append(queue, state{i: current.i, j: current.j + 1, count: current.count + 1})
		}
	}
	return result
}
