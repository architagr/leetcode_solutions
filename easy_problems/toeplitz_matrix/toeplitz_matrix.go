package toeplitz_matrix

func IsToePlitzMatrix(matrix [][]int) bool {

	//top row
	for i := 0; i < len(matrix[0]); i++ {
		for k, l := 1, i+1; k < len(matrix) && l < len(matrix[0]); k, l = k+1, l+1 {
			if matrix[k-1][l-1] != matrix[k][l] {
				return false
			}
		}
	}
	//column
	for i := 0; i < len(matrix); i++ {
		for k, l := i+1, 1; k < len(matrix) && l < len(matrix[0]); k, l = k+1, l+1 {
			if matrix[k-1][l-1] != matrix[k][l] {
				return false
			}
		}
	}
	return true
}
