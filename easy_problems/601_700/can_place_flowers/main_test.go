package canplaceflowers

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestCanPlaceFlowers(t *testing.T) {
	tests := []struct {
		name      string
		flowerbed []int
		n         int
		want      bool
	}{
		{name: "example 1", flowerbed: []int{1, 0, 0, 0, 1}, n: 1, want: true},
		{name: "example 2", flowerbed: []int{1, 0, 0, 0, 1}, n: 2, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPlaceFlowers(tt.flowerbed, tt.n); got != tt.want {
				t.Errorf("canPlaceFlowers() = %v, want %v", got, tt.want)
			}
		})
	}
}
