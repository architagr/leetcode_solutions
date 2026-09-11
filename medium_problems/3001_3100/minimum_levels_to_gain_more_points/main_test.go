package minimumlevelstogainmorepoints

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestMinimumLevels(t *testing.T) {
	tests := []struct {
		name     string
		possible []int
		want     int
	}{
		{name: "example 1", possible: []int{1, 0, 1, 0}, want: 1},
		{name: "example 2", possible: []int{1, 1, 1, 1, 1}, want: 3},
		{name: "example 3", possible: []int{0, 0}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumLevels(tt.possible); got != tt.want {
				t.Errorf("minimumLevels() = %v, want %v", got, tt.want)
			}
		})
	}
}
