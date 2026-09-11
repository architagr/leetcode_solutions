package dungeongame

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestCalculateMinimumHP(t *testing.T) {
	tests := []struct {
		name    string
		dungeon [][]int
		want    int
	}{
		{name: "example 1", dungeon: [][]int{[]int{-2, -3, 3}, []int{-5, -10, 1}, []int{10, 30, -5}}, want: 7},
		{name: "example 2", dungeon: [][]int{[]int{0}}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateMinimumHP(tt.dungeon); got != tt.want {
				t.Errorf("calculateMinimumHP() = %v, want %v", got, tt.want)
			}
		})
	}
}
