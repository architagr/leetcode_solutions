package check_every_row_column_contains_all_numbers

func CheckValid(matrix [][]int) bool {
	n := len(matrix)
	for i := 0; i < n; i++ {
		m := make([]bool, n)
		for j := 0; j < n; j++ {
			if m[matrix[i][j]-1] {
				return false
			}
			m[matrix[i][j]-1] = true
		}

	}
	for i := 0; i < n; i++ {
		m := make([]bool, n)
		for j := 0; j < n; j++ {
			if m[matrix[j][i]-1] {
				return false
			}
			m[matrix[j][i]-1] = true
		}
	}
	return true
}
