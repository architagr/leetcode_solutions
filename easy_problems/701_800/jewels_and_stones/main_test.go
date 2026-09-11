package jewels_and_stones

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestCountJewels(t *testing.T) {
	tests := []struct {
		name   string
		jewels string
		stones string
		want   int
	}{
		{name: "example 1", jewels: "aA", stones: "aAAbbbb", want: 3},
		{name: "example 2", jewels: "z", stones: "ZZ", want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountJewels(tt.jewels, tt.stones); got != tt.want {
				t.Errorf("CountJewels() = %v, want %v", got, tt.want)
			}
		})
	}
}
