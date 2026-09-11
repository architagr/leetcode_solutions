package nthtribonaccinumber

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestTribonacci(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{name: "example 1", n: 4, want: 4},
		{name: "example 2", n: 25, want: 1389537},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tribonacci(tt.n); got != tt.want {
				t.Errorf("tribonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}
