package nth_magical_number

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestNthMagicalNumber(t *testing.T) {
	tests := []struct {
		name string
		n    int
		a    int
		b    int
		want int
	}{
		{name: "example 1", n: 1, a: 2, b: 3, want: 2},
		{name: "example 2", n: 4, a: 2, b: 3, want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NthMagicalNumber(tt.n, tt.a, tt.b); got != tt.want {
				t.Errorf("NthMagicalNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
