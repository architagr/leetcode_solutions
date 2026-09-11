package equal_Row_and_column_pairs

func EqualPairs(grid [][]int) int {
	count, n := 0, len(grid)
	transposes := getTranspose(grid)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			flag := true
			for k := 0; k < n; k++ {
				if grid[i][k] != transposes[j][k] {
					flag = false
					break
				}
			}
			if flag {
				count++
			}
		}

	}
	return count
}

func getTranspose(grid [][]int) [][]int {
	n := len(grid)
	newGrid := make([][]int, n)
	for i := 0; i < n; i++ {
		row := make([]int, n)
		for j := 0; j < n; j++ {
			row[j] = grid[j][i]
		}
		newGrid[i] = row
	}
	return newGrid
}
