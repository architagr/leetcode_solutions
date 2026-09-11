package ugly_number_II

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestNthUglyNumber(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{name: "example 1", n: 10, want: 12},
		{name: "example 2", n: 1, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NthUglyNumber(tt.n); got != tt.want {
				t.Errorf("NthUglyNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
