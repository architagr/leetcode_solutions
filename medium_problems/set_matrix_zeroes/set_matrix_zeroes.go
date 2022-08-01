package set_matrix_zeroes

func SetZeroes(matrix [][]int) {
	row := make(map[int]bool)
	col := make(map[int]bool)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			_, okr := row[i]
			_, okc := col[j]
			if okc || okr {
				matrix[i][j] = 0
			}
		}
	}
}
