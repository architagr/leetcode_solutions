package min_max_game

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestMinMaxGame(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "example 1", nums: []int{1, 3, 5, 2, 4, 8, 2, 2}, want: 1},
		{name: "example 2", nums: []int{3}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinMaxGame(tt.nums); got != tt.want {
				t.Errorf("MinMaxGame() = %v, want %v", got, tt.want)
			}
		})
	}
}
