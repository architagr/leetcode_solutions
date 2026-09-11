package reverse_integer

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestReverse(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want int
	}{
		{name: "example 1", x: 123, want: 321},
		{name: "example 2", x: -123, want: -321},
		{name: "example 3", x: 120, want: 21},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.x); got != tt.want {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
