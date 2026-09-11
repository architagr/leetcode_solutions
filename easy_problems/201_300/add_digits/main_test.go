package add_digits

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestAddDigits(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want int
	}{
		{name: "example 1", num: 38, want: 2},
		{name: "example 2", num: 0, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddDigits(tt.num); got != tt.want {
				t.Errorf("AddDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
