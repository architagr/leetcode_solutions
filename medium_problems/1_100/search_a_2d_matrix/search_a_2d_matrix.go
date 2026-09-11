package search_a_2d_matrix

func SearchMatrix(matrix [][]int, target int) bool {
	rows, cols := len(matrix), len(matrix[0])
	i, j := rows-1, 0
	for i >= 0 && j < cols {
		if matrix[i][j] < target {
			j++
		} else if matrix[i][j] > target {
			i--
		} else {
			return true
		}
	}
	return false
}
