package divide_two_integers

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestDivide(t *testing.T) {
	tests := []struct {
		name     string
		dividend int
		divisor  int
		want     int
	}{
		{name: "example 1", dividend: 10, divisor: 3, want: 3},
		{name: "example 2", dividend: 7, divisor: -3, want: -2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Divide(tt.dividend, tt.divisor); got != tt.want {
				t.Errorf("Divide() = %v, want %v", got, tt.want)
			}
		})
	}
}
