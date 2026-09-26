package neareststoreheatmap

// A city as a grid of blocks. The heatmap is, for every block, how many blocks
// away the nearest open store is, walking along streets: up, down, left, right.
type Point struct{ Row, Col int }

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// HeatmapByScan is what I would write. With no obstacles, the walking distance
// between two blocks is |dRow| + |dCol|, so every block checks every store and
// keeps the closest.
//
// It is the definition, it needs no data structure, and it is exactly right.
func HeatmapByScan(rows, cols int, stores []Point) [][]int {
	out := make([][]int, rows)
	for r := range out {
		out[r] = make([]int, cols)
		for c := range out[r] {
			best := rows + cols
			for _, s := range stores {
				if d := abs(r-s.Row) + abs(c-s.Col); d < best {
					best = d
				}
			}
			out[r][c] = best
		}
	}
	return out
}

// HeatmapByBFS seeds one queue with every store at distance 0 and spreads
// outward, so each block is reached first by its nearest store. Episode 19's
// fix, on a grid.
func HeatmapByBFS(rows, cols int, stores []Point) [][]int {
	out := make([][]int, rows)
	for r := range out {
		out[r] = make([]int, cols)
		for c := range out[r] {
			out[r][c] = -1
		}
	}
	queue := make([]Point, 0, rows*cols)
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
			if r >= 0 && r < rows && c >= 0 && c < cols && out[r][c] < 0 {
				out[r][c] = out[p.Row][p.Col] + 1
				queue = append(queue, Point{r, c})
			}
		}
	}
	return out
}

// HeatmapBySweeps fills the grid in two passes and keeps no queue at all.
//
// A shortest walk leaves a block going up, down, left or right. The first
// pass, top-left to bottom-right, finds the best walk whose first step is up
// or left: both of those neighbours are already final when it gets here. The
// second pass, backwards, does down and right, and keeps whichever is smaller.
func HeatmapBySweeps(rows, cols int, stores []Point) [][]int {
	far := rows + cols // further than any real distance on this grid
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
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if r > 0 && out[r-1][c]+1 < out[r][c] {
				out[r][c] = out[r-1][c] + 1
			}
			if c > 0 && out[r][c-1]+1 < out[r][c] {
				out[r][c] = out[r][c-1] + 1
			}
		}
	}
	for r := rows - 1; r >= 0; r-- {
		for c := cols - 1; c >= 0; c-- {
			// A minimum, not an assignment: the first pass's answer stands
			// whenever the nearest store was up or to the left.
			if r < rows-1 && out[r+1][c]+1 < out[r][c] {
				out[r][c] = out[r+1][c] + 1
			}
			if c < cols-1 && out[r][c+1]+1 < out[r][c] {
				out[r][c] = out[r][c+1] + 1
			}
		}
	}
	return out
}
