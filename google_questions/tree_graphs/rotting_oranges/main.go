package rottingoranges

type node struct {
	row, col int
	t        int
}

func orangesRotting(grid [][]int) int {
	q := make([]node, 0, 100)
	push := func(row, col, t int) {
		if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) || grid[row][col] != 1 {
			return
		}
		grid[row][col] = 2
		q = append(q, node{row: row, col: col, t: t})
	}
	pop := func() node {
		x := q[0]
		q = q[1:]
		return x
	}
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 2 {
				q = append(q, node{row: i, col: j, t: 0})
			}
		}
	}

	for len(q) > 0 {
		x := pop()
		ans = x.t
		push(x.row+1, x.col, x.t+1)
		push(x.row-1, x.col, x.t+1)
		push(x.row, x.col+1, x.t+1)
		push(x.row, x.col-1, x.t+1)
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}
	return ans
}
