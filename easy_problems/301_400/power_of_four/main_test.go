package power_of_four

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestIsPowerOfFour(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{name: "example 1", n: 16, want: true},
		{name: "example 2", n: 5, want: false},
		{name: "example 3", n: 1, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPowerOfFour(tt.n); got != tt.want {
				t.Errorf("IsPowerOfFour() = %v, want %v", got, tt.want)
			}
		})
	}
}
