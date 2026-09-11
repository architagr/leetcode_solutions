package reshape_the_matrix

func MatrixReshape(mat [][]int, r, c int) [][]int {
	arr := make([][]int, r)
	row := make([]int, c)
	if r*c != len(mat)*len(mat[0]) {
		return mat
	}
	k, l := 0, 0

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if l == 0 {
				row = make([]int, c)
			}
			row[l] = mat[i][j]
			l++
			if l == c {
				arr[k] = row
				k++
				l = 0
			}
		}
	}
	return arr
}
