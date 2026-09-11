package countdistinctnumbersonboard

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestDistinctIntegers(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{name: "example 1", n: 5, want: 4},
		{name: "example 2", n: 3, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distinctIntegers(tt.n); got != tt.want {
				t.Errorf("distinctIntegers() = %v, want %v", got, tt.want)
			}
		})
	}
}
