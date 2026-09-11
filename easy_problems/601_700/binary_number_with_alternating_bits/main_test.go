package binary_number_with_alternating_bits

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestHasAlternatingBits(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{name: "example 1", n: 5, want: true},
		{name: "example 2", n: 7, want: false},
		{name: "example 3", n: 11, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasAlternatingBits(tt.n); got != tt.want {
				t.Errorf("HasAlternatingBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
