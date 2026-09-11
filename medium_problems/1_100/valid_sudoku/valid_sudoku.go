package valid_sudoku

func IsValidSuduku(board [][]byte) bool {
	if !checkColumn(board) {
		return false
	}
	if !checkRow(board) {
		return false
	}

	if !checkAllGrids(board) {
		return false
	}

	return true
}

func checkColumn(board [][]byte) bool {
	n := 9
	for i := 0; i < n; i++ {
		m := make([]bool, n+1)
		for j := 0; j < n; j++ {
			if board[j][i] != '.' {
				if m[board[j][i]-'0'] {
					return false
				}
				m[board[j][i]-'0'] = true
			}
		}
	}
	return true
}
func checkRow(board [][]byte) bool {
	n := 9
	for i := 0; i < n; i++ {
		m := make([]bool, n+1)
		for j := 0; j < n; j++ {
			if board[i][j] != '.' {
				if m[board[i][j]-'0'] {
					return false
				}
				m[board[i][j]-'0'] = true
			}
		}
	}
	return true
}
func checkAllGrids(board [][]byte) bool {

	if !checkGrids(board, 0, 0) {
		return false
	}
	if !checkGrids(board, 0, 3) {
		return false
	}
	if !checkGrids(board, 0, 6) {
		return false
	}
	if !checkGrids(board, 3, 0) {
		return false
	}
	if !checkGrids(board, 3, 3) {
		return false
	}
	if !checkGrids(board, 3, 6) {
		return false
	}
	if !checkGrids(board, 6, 0) {
		return false
	}
	if !checkGrids(board, 6, 3) {
		return false
	}
	if !checkGrids(board, 6, 6) {
		return false
	}
	return true
}
func checkGrids(board [][]byte, col, row int) bool {
	m := make([]bool, 10)
	for i := row; i < row+3; i++ {

		for j := col; j < col+3; j++ {
			if board[i][j] != '.' {
				if m[board[i][j]-'0'] {
					return false
				}
				m[board[i][j]-'0'] = true
			}
		}
	}
	return true
}
