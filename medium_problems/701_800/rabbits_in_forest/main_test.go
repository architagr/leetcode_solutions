package rabbitsinforest

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestNumRabbits(t *testing.T) {
	tests := []struct {
		name    string
		answers []int
		want    int
	}{
		{name: "example 1", answers: []int{1, 1, 2}, want: 5},
		{name: "example 2", answers: []int{10, 10, 10}, want: 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numRabbits(tt.answers); got != tt.want {
				t.Errorf("numRabbits() = %v, want %v", got, tt.want)
			}
		})
	}
}
