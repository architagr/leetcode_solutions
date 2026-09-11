package trapping_rain_water

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestTrap(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{name: "example 1", height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, want: 6},
		{name: "example 2", height: []int{4, 2, 0, 3, 2, 5}, want: 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Trap(tt.height); got != tt.want {
				t.Errorf("Trap() = %v, want %v", got, tt.want)
			}
		})
	}
}
