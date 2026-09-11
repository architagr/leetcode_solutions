package valid_perfect_square

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestIsPerfectSquare(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want bool
	}{
		{name: "example 1", num: 16, want: true},
		{name: "example 2", num: 14, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPerfectSquare(tt.num); got != tt.want {
				t.Errorf("IsPerfectSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
