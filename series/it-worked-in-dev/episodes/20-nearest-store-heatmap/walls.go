package neareststoreheatmap

// The version with a river in it. Blocks marked in wall cannot be walked
// through, so the walking distance stops being |dRow| + |dCol|. It is here for
// the one test that shows where the sweeps stop being correct; it is not
// benchmarked.

// HeatmapByBFSWithWalls is HeatmapByBFS that refuses to step onto a wall.
// Unreachable blocks, and walls, stay -1.
func HeatmapByBFSWithWalls(rows, cols int, stores []Point, wall [][]bool) [][]int {
	out := make([][]int, rows)
	for r := range out {
		out[r] = make([]int, cols)
		for c := range out[r] {
			out[r][c] = -1
		}
	}
	queue := []Point{}
	for _, s := range stores {
		if out[s.Row][s.Col] < 0 {
			out[s.Row][s.Col] = 0
			queue = append(queue, s)
		}
	}
	dr := [4]int{-1, 1, 0, 0}
	dc := [4]int{0, 0, -1, 1}
	for i := 0; i < len(queue); i++ {
		p := queue[i]
		for k := 0; k < 4; k++ {
			r, c := p.Row+dr[k], p.Col+dc[k]
			if r >= 0 && r < rows && c >= 0 && c < cols && !wall[r][c] && out[r][c] < 0 {
				out[r][c] = out[p.Row][p.Col] + 1
				queue = append(queue, Point{r, c})
			}
		}
	}
	return out
}

// HeatmapBySweepsWithWalls is the two-pass version taught to skip walls. It is
// wrong: with walls, a shortest walk can need to go down and then up, and no
// single pass sees both halves of it.
func HeatmapBySweepsWithWalls(rows, cols int, stores []Point, wall [][]bool) [][]int {
	far := rows * cols
	out := make([][]int, rows)
	for r := range out {
		out[r] = make([]int, cols)
		for c := range out[r] {
			out[r][c] = far
		}
	}
	for _, s := range stores {
		out[s.Row][s.Col] = 0
	}
	relax := func(r, c, nr, nc int) {
		if !wall[nr][nc] && out[nr][nc]+1 < out[r][c] {
			out[r][c] = out[nr][nc] + 1
		}
	}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if wall[r][c] {
				continue
			}
			if r > 0 {
				relax(r, c, r-1, c)
			}
			if c > 0 {
				relax(r, c, r, c-1)
			}
		}
	}
	for r := rows - 1; r >= 0; r-- {
		for c := cols - 1; c >= 0; c-- {
			if wall[r][c] {
				continue
			}
			if r < rows-1 {
				relax(r, c, r+1, c)
			}
			if c < cols-1 {
				relax(r, c, r, c+1)
			}
		}
	}
	for r := range out {
		for c := range out[r] {
			if wall[r][c] || out[r][c] == far {
				out[r][c] = -1
			}
		}
	}
	return out
}
