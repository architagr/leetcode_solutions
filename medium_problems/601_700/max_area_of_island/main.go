package maxareaofisland

func maxAreaOfIsland(grid [][]int) int {
	max := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				// Every cell of this island is sunk during the fill, so a
				// later scan can never start a second fill on the same one.
				if area := dfs(i, j, &grid); area > max {
					max = area
				}
			}
		}
	}
	return max
}

// dfs sinks the island reachable from (i, j) and returns how many cells it
// covered. The bounds check and the "already sunk" check share one guard,
// so a caller never has to validate a neighbour before recursing.
func dfs(i, j int, grid *[][]int) int {
	if i < 0 || j < 0 || i >= len(*grid) || j >= len((*grid)[0]) || (*grid)[i][j] != 1 {
		return 0
	}
	// Sink before recursing. Marking on the way in is what stops the four
	// calls below from walking straight back into this cell.
	(*grid)[i][j] = 0

	return 1 + dfs(i+1, j, grid) + dfs(i-1, j, grid) + dfs(i, j+1, grid) + dfs(i, j-1, grid)
}
