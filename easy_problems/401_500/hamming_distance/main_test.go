package hamming_distance

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestHammingDistance(t *testing.T) {
	tests := []struct {
		name string
		x    int
		y    int
		want int
	}{
		{name: "example 1", x: 1, y: 4, want: 2},
		{name: "example 2", x: 3, y: 1, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HammingDistance(tt.x, tt.y); got != tt.want {
				t.Errorf("HammingDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
