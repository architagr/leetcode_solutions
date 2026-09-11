package transpose_matrix

func Transpose(matrix [][]int) [][]int {
	n := len(matrix)
	m := len(matrix[0])
	arr := make([][]int, m)
	for i := 0; i < m; i++ {
		arr[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			arr[j][i] = matrix[i][j]
		}
	}
	return arr
}

func Transposea(matrix [][]int) [][]int {
	tr := make([][]int, len(matrix[0]))
	for r := range tr {
		tr[r] = make([]int, len(matrix))
	}
	for r := range matrix {
		for c := range matrix[r] {
			tr[c][r] = matrix[r][c]
		}
	}
	return tr
}
