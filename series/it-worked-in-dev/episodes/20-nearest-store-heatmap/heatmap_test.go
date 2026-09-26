package neareststoreheatmap

import (
	"math/rand"
	"reflect"
	"testing"
)

// scatter places k distinct stores on a rows x cols grid, the same ones for
// the same seed.
func scatter(rows, cols, k int, seed int64) []Point {
	r := rand.New(rand.NewSource(seed))
	seen := map[Point]bool{}
	out := make([]Point, 0, k)
	for len(out) < k {
		p := Point{r.Intn(rows), r.Intn(cols)}
		if !seen[p] {
			seen[p] = true
			out = append(out, p)
		}
	}
	return out
}

func TestAllThreeAgree(t *testing.T) {
	type tc struct {
		rows, cols int
		stores     []Point
	}
	cases := []tc{
		{1, 1, []Point{{0, 0}}},
		{1, 9, []Point{{0, 4}}},
		{7, 1, []Point{{0, 0}, {6, 0}}},
		{5, 5, []Point{{2, 2}}},
		{5, 5, []Point{{0, 0}, {4, 4}}},
		{40, 60, scatter(40, 60, 1, 1)},
		{40, 60, scatter(40, 60, 30, 2)},
		{40, 60, scatter(40, 60, 2400, 3)}, // every block is a store
	}
	for _, s := range shapes {
		if len(s.stores) <= 100 {
			cases = append(cases, tc{s.rows, s.cols, s.stores})
		}
	}
	for _, x := range cases {
		want := HeatmapByScan(x.rows, x.cols, x.stores)
		if got := HeatmapByBFS(x.rows, x.cols, x.stores); !reflect.DeepEqual(got, want) {
			t.Fatalf("%dx%d, %d stores: BFS disagrees with the scan", x.rows, x.cols, len(x.stores))
		}
		if got := HeatmapBySweeps(x.rows, x.cols, x.stores); !reflect.DeepEqual(got, want) {
			t.Fatalf("%dx%d, %d stores: sweeps disagree with the scan", x.rows, x.cols, len(x.stores))
		}
	}
}

// Two rivers, each with one bridge, at opposite ends. From the store at the
// top left, the bottom-right block is reached by walking right, crossing, walking
// back left, crossing again, and walking right: a zigzag.
//
//	S . . . .
//	~ ~ ~ ~ .
//	. . . . .
//	. ~ ~ ~ ~
//	. . . . X
//
// The first sweep only looks up and left; the second only down and right. The
// walk turns back on itself twice, so neither sweep ever sees the leg along
// row 2 feeding the leg down column 0.
func TestSweepsBreakWhenAWalkMustDoubleBack(t *testing.T) {
	rows, cols := 5, 5
	wall := make([][]bool, rows)
	for r := range wall {
		wall[r] = make([]bool, cols)
	}
	for c := 0; c < 4; c++ {
		wall[1][c] = true
		wall[3][c+1] = true
	}
	stores := []Point{{0, 0}}
	bfs := HeatmapByBFSWithWalls(rows, cols, stores, wall)
	sweeps := HeatmapBySweepsWithWalls(rows, cols, stores, wall)
	scan := HeatmapByScan(rows, cols, stores)
	t.Logf("bottom-right block: BFS %d, two sweeps %d, straight-line scan %d",
		bfs[4][4], sweeps[4][4], scan[4][4])
	if bfs[4][4] != 16 {
		t.Fatalf("BFS says %d, the walk is 16", bfs[4][4])
	}
	if sweeps[4][4] == bfs[4][4] {
		t.Fatal("the sweeps got the rivers right; the episode says they do not")
	}
}
