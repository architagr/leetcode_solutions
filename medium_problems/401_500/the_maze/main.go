package themaze

var (
	dirX = []int{0, 1, 0, -1}
	dirY = []int{1, 0, -1, 0}
)

func hasPath(maze [][]int, start, destination []int) bool {
	visited := make([][]bool, len(maze))
	for i := 0; i < len(maze); i++ {
		visited[i] = make([]bool, len(maze[i]))
	}
	return dfs(maze, start, destination, visited)
}

func dfs(maze [][]int, start, destination []int, visited [][]bool) bool {
	if start[0] == destination[0] && start[1] == destination[1] {
		return true
	}

	if visited[start[0]][start[1]] {
		return false
	}
	visited[start[0]][start[1]] = true
	for i := 0; i < 4; i++ {
		r := start[0]
		c := start[1]
		for r >= 0 && c >= 0 && r < len(maze) && c < len(maze[0]) && maze[r][c] == 0 {
			r += dirX[i]
			c += dirY[i]
		}
		if dfs(maze, []int{r - dirX[i], c - dirY[i]}, destination, visited) {
			return true
		}
	}
	return false
}
