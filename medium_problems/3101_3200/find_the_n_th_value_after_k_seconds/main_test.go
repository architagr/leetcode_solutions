package findthenthvalueafterkseconds

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestValueAfterKSeconds(t *testing.T) {
	tests := []struct {
		name string
		n    int
		k    int
		want int
	}{
		{name: "example 1", n: 4, k: 5, want: 56},
		{name: "example 2", n: 5, k: 3, want: 35},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := valueAfterKSeconds(tt.n, tt.k); got != tt.want {
				t.Errorf("valueAfterKSeconds() = %v, want %v", got, tt.want)
			}
		})
	}
}
