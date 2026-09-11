package fibonacci_number

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestFib(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{name: "example 1", n: 2, want: 1},
		{name: "example 2", n: 3, want: 2},
		{name: "example 3", n: 4, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib(tt.n); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}
