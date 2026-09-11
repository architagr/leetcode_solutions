package xoroperationinanarray

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestXorOperation(t *testing.T) {
	tests := []struct {
		name  string
		n     int
		start int
		want  int
	}{
		{name: "example 1", n: 5, start: 0, want: 8},
		{name: "example 2", n: 4, start: 3, want: 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := xorOperation(tt.n, tt.start); got != tt.want {
				t.Errorf("xorOperation() = %v, want %v", got, tt.want)
			}
		})
	}
}
