package power_of_two

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestIsPowerOfTwo(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{name: "example 1", n: 1, want: true},
		{name: "example 2", n: 16, want: true},
		{name: "example 3", n: 3, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPowerOfTwo(tt.n); got != tt.want {
				t.Errorf("IsPowerOfTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
