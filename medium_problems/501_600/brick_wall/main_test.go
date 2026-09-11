package brickwall

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestLeastBricks(t *testing.T) {
	tests := []struct {
		name string
		wall [][]int
		want int
	}{
		{name: "example 1", wall: [][]int{[]int{1, 2, 2, 1}, []int{3, 1, 2}, []int{1, 3, 2}, []int{2, 4}, []int{3, 1, 2}, []int{1, 3, 1, 1}}, want: 2},
		{name: "example 2", wall: [][]int{[]int{1}, []int{1}, []int{1}}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leastBricks(tt.wall); got != tt.want {
				t.Errorf("leastBricks() = %v, want %v", got, tt.want)
			}
		})
	}
}
