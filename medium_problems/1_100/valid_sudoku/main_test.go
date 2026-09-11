package valid_sudoku

import "testing"

// board turns the statement's rows of single-character strings into the
// [][]byte the solution takes.
func board(rows []string) [][]byte {
	out := make([][]byte, len(rows))
	for i, r := range rows {
		out[i] = []byte(r)
	}
	return out
}

// Cases are the two worked examples from the problem statement. They
// differ only in the top-left cell, which is what makes the second one
// invalid: a second 8 in the same 3x3 box.
func TestIsValidSuduku(t *testing.T) {
	valid := []string{
		"53..7....",
		"6..195...",
		".98....6.",
		"8...6...3",
		"4..8.3..1",
		"7...2...6",
		".6....28.",
		"...419..5",
		"....8..79",
	}
	invalid := append([]string{"83..7...."}, valid[1:]...)

	tests := []struct {
		name string
		rows []string
		want bool
	}{
		{name: "example 1", rows: valid, want: true},
		{name: "example 2", rows: invalid, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidSuduku(board(tt.rows)); got != tt.want {
				t.Errorf("IsValidSuduku() = %v, want %v", got, tt.want)
			}
		})
	}
}
