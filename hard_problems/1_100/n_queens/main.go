package nqueens

var (
	available     = true
	not_available = false
)

func solveNQueens(n int) [][]string {
	if n == 1 {
		return [][]string{{"Q"}}
	}
	col := generateDefaultCol(n)
	d1, d2 := generateDefaultD1(n), generateDefaultD2(n)
	result := make([][]string, 0)

	board := generateDefaultBoard(n)
	rowIndex := 0
	for colIndex := 0; colIndex < n; colIndex++ {
		setNotAvailable(rowIndex, colIndex, col, d1, d2, &board)
		generateBoard(n, rowIndex+1, col, d1, d2, &board, &result)
		setAvailable(rowIndex, colIndex, col, d1, d2, &board)

	}
	return result
}
func setNotAvailable(rowIndex, colIndex int, col, d1, d2 map[int]bool, board *[][]byte) {
	col[colIndex] = not_available
	d1[d1Index(rowIndex, colIndex)] = not_available
	d2[d2Index(rowIndex, colIndex)] = not_available
	(*board)[rowIndex][colIndex] = 'Q'
}
func setAvailable(rowIndex, colIndex int, col, d1, d2 map[int]bool, board *[][]byte) {
	col[colIndex] = available
	d1[d1Index(rowIndex, colIndex)] = available
	d2[d2Index(rowIndex, colIndex)] = available
	(*board)[rowIndex][colIndex] = '.'
}

func copyBoard(board *[][]byte) []string {
	result := make([]string, len(*board))
	for i, row := range *board {
		result[i] = string(row)
	}
	return result
}

func generateBoard(n, rowIndex int, col, d1, d2 map[int]bool, board *[][]byte, result *[][]string) {
	if rowIndex >= n {
		*result = append((*result), copyBoard(board))
		return
	}

	for colIndex := 0; colIndex < n; colIndex++ {
		if col[colIndex] == available && d1[d1Index(rowIndex, colIndex)] == available && d2[d2Index(rowIndex, colIndex)] == available {
			setNotAvailable(rowIndex, colIndex, col, d1, d2, board)
			generateBoard(n, rowIndex+1, col, d1, d2, board, result)
			setAvailable(rowIndex, colIndex, col, d1, d2, board)
		}
	}

}
func d1Index(rowIndex, colIndex int) int {
	return colIndex - rowIndex
}
func d2Index(rowIndex, colIndex int) int {
	return colIndex + rowIndex
}
func generateDefaultBoard(n int) [][]byte {
	result := make([][]byte, n)
	for i := 0; i < n; i++ {
		data := make([]byte, n)
		for j := 0; j < n; j++ {
			data[j] = '.'
		}
		result[i] = data
	}
	return result
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
