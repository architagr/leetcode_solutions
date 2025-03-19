package shortestpathinbinarymatrix

const (
	visiting  = 2
	completed = 3
)

type node struct {
	row, col int
	cost     int
}

func shortestPathBinaryMatrix(grid [][]int) int {

	n := len(grid)
	q := make([]node, 0, n*n)
	push := func(row, col, cost int) {
		if row >= len(grid) || col >= len(grid) || row < 0 || col < 0 || grid[row][col] != 0 {
			return
		}
		grid[row][col] = 1
		q = append(q, node{row: row, col: col, cost: cost})
	}
	pop := func() node {
		x := q[0]
		q = q[1:]
		return x
	}

	push(0, 0, 1)

	for len(q) > 0 {
		currNode := pop()
		if currNode.row == len(grid)-1 && currNode.col == len(grid)-1 {
			return currNode.cost
		}
		push(currNode.row+1, currNode.col, currNode.cost+1)
		push(currNode.row-1, currNode.col, currNode.cost+1)
		push(currNode.row, currNode.col+1, currNode.cost+1)
		push(currNode.row, currNode.col-1, currNode.cost+1)
		push(currNode.row+1, currNode.col+1, currNode.cost+1)
		push(currNode.row+1, currNode.col-1, currNode.cost+1)
		push(currNode.row-1, currNode.col+1, currNode.cost+1)
		push(currNode.row-1, currNode.col-1, currNode.cost+1)
	}
	return -1
}
