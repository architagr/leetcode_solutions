package maxareaofisland

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestMaxAreaOfIsland(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{name: "example 1", grid: [][]int{[]int{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, []int{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}, []int{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0}, []int{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0}, []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, []int{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, []int{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}, want: 6},
		{name: "example 2", grid: [][]int{[]int{0, 0, 0, 0, 0, 0, 0, 0}}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxAreaOfIsland(tt.grid); got != tt.want {
				t.Errorf("maxAreaOfIsland() = %v, want %v", got, tt.want)
			}
		})
	}
}
