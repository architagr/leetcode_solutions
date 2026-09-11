package xofakindinadeckofcards

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestHasGroupsSizeX(t *testing.T) {
	tests := []struct {
		name string
		deck []int
		want bool
	}{
		{name: "example 1", deck: []int{1, 2, 3, 4, 4, 3, 2, 1}, want: true},
		{name: "example 2", deck: []int{1, 1, 1, 2, 2, 2, 3, 3}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasGroupsSizeX(tt.deck); got != tt.want {
				t.Errorf("hasGroupsSizeX() = %v, want %v", got, tt.want)
			}
		})
	}
}
