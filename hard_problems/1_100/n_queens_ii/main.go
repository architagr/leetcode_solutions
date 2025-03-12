package nqueensii

var (
	available     = true
	not_available = false
)

func totalNQueens(n int) int {
	if n == 1 {
		return 1
	}
	col := generateDefaultCol(n)
	d1, d2 := generateDefaultD1(n), generateDefaultD2(n)
	rowIndex := 0
	result := 0
	for colIndex := 0; colIndex < n; colIndex++ {
		setNotAvailable(rowIndex, colIndex, col, d1, d2)
		generateBoard(n, rowIndex+1, col, d1, d2, &result)
		setAvailable(rowIndex, colIndex, col, d1, d2)

	}
	return result
}
func setNotAvailable(rowIndex, colIndex int, col, d1, d2 map[int]bool) {
	col[colIndex] = not_available
	d1[d1Index(rowIndex, colIndex)] = not_available
	d2[d2Index(rowIndex, colIndex)] = not_available
}
func setAvailable(rowIndex, colIndex int, col, d1, d2 map[int]bool) {
	col[colIndex] = available
	d1[d1Index(rowIndex, colIndex)] = available
	d2[d2Index(rowIndex, colIndex)] = available
}

func generateBoard(n, rowIndex int, col, d1, d2 map[int]bool, result *int) {
	if rowIndex >= n {
		(*result)++
		return
	}

	for colIndex := 0; colIndex < n; colIndex++ {
		if col[colIndex] == available && d1[d1Index(rowIndex, colIndex)] == available && d2[d2Index(rowIndex, colIndex)] == available {
			setNotAvailable(rowIndex, colIndex, col, d1, d2)
			generateBoard(n, rowIndex+1, col, d1, d2, result)
			setAvailable(rowIndex, colIndex, col, d1, d2)
		}
	}

}
func d1Index(rowIndex, colIndex int) int {
	return colIndex - rowIndex
}
func d2Index(rowIndex, colIndex int) int {
	return colIndex + rowIndex
}

func generateDefaultD1(n int) (d1 map[int]bool) {
	d1 = make(map[int]bool)
	for i := -1 * (n - 1); i < n; i++ {
		d1[i] = available
	}
	return
}
func generateDefaultD2(n int) (d1 map[int]bool) {
	d1 = make(map[int]bool)
	for i := 0; i <= 2*(n-1); i++ {
		d1[i] = available
	}
	return
}
func generateDefaultCol(n int) (col map[int]bool) {
	col = make(map[int]bool)
	for i := 0; i < n; i++ {
		col[i] = available
	}
	return
}
