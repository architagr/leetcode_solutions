package matrix_diagonal_sum

func DIagomalSum(mat [][]int) int {
	n, sum := len(mat), 0
	m := n - 1
	for i := 0; i < n; i++ {
		sum += mat[i][i] + mat[i][m-i]
	}

	if n%2 != 0 {
		x := n / 2
		sum -= mat[x][x]
	}

	return sum
}
