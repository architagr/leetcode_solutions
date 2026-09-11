package find_the_highest_altitude0

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestLargestAltitude(t *testing.T) {
	tests := []struct {
		name string
		gain []int
		want int
	}{
		{name: "example 1", gain: []int{-5, 1, 5, 0, -7}, want: 1},
		{name: "example 2", gain: []int{-4, -3, -2, -1, 4, 3, 2}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LargestAltitude(tt.gain); got != tt.want {
				t.Errorf("LargestAltitude() = %v, want %v", got, tt.want)
			}
		})
	}
}
