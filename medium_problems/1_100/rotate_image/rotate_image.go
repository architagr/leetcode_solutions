package rotate_image

func Rotate(matrix [][]int) {
	matrix = transpose(matrix)
	matrix = reverse2dArray(matrix)
}

func transpose(matrix [][]int) [][]int {
	n := len(matrix)

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	return matrix
}
func reverse2dArray(arr [][]int) [][]int {
	n := len(arr)
	for i := 0; i < n; i++ {
		arr[i] = reverse1dArray(arr[i])
	}
	return arr
}

func reverse1dArray(a []int) []int {
	i, j := 0, len(a)-1

	for i <= j {
		a[j], a[i] = a[i], a[j]
		i++
		j--
	}
	return a
}
