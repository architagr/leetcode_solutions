package findthepivotinteger

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestPivotInteger(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{name: "example 1", n: 8, want: 6},
		{name: "example 2", n: 1, want: 1},
		{name: "example 3", n: 4, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pivotInteger(tt.n); got != tt.want {
				t.Errorf("pivotInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}
