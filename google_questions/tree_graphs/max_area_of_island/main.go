package maxareaofisland

func maxAreaOfIsland(grid [][]int) int {
	max := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				visted := map[node]bool{}
				bfs(i, j, visted, &grid)
				if len(visted) > max {
					max = len(visted)
				}

			}
		}
	}
	return max
}

type node struct {
	i, j int
}

func bfs(i, j int, visited map[node]bool, grid *[][]int) {
	if i < 0 || j < 0 || i >= len(*grid) || j >= len((*grid)[0]) || (*grid)[i][j] != 1 {
		return
	}
	visited[node{i: i, j: j}] = true
	(*grid)[i][j] = 0
	bfs(i+1, j, visited, grid)
	bfs(i-1, j, visited, grid)
	bfs(i, j+1, visited, grid)
	bfs(i, j-1, visited, grid)
}
