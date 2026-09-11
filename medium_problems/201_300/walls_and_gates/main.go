package wallsandgates

type cell struct {
	x, y, cost int
}

func wallsAndGates(rooms [][]int) {
	inf := 2147483647
	n := len(rooms)
	m := len(rooms[0])
	q := make([]cell, 0, n*m)
	push := func(x, y, cost int) {
		if x < 0 || x >= n || y < 0 || y >= m || rooms[x][y] != inf {
			return
		}
		rooms[x][y] = cost
		q = append(q, cell{x: x, y: y, cost: cost})
	}
	pop := func() cell {
		x := q[0]
		q = q[1:]
		return x
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if rooms[i][j] != 0 {
				continue
			}
			q = append(q, cell{x: i, y: j, cost: 0})
		}
	}

	for len(q) > 0 {
		x := pop()
		push(x.x+1, x.y, x.cost+1)
		push(x.x-1, x.y, x.cost+1)

		push(x.x, x.y+1, x.cost+1)
		push(x.x, x.y-1, x.cost+1)
	}
}
