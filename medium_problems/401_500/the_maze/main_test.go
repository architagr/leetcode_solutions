package themaze

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestHasPath(t *testing.T) {
	maze := [][]int{
		{0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0},
		{1, 1, 0, 1, 1},
		{0, 0, 0, 0, 0},
	}
	tests := []struct {
		name               string
		start, destination []int
		want               bool
	}{
		{name: "example 1", start: []int{0, 4}, destination: []int{4, 4}, want: true},
		{name: "example 2", start: []int{0, 4}, destination: []int{3, 2}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rows := make([][]int, len(maze))
			for i := range maze {
				rows[i] = append([]int{}, maze[i]...)
			}
			if got := hasPath(rows, tt.start, tt.destination); got != tt.want {
				t.Errorf("hasPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
