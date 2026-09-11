package addtwointegers

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestSum(t *testing.T) {
	tests := []struct {
		name string
		num1 int
		num2 int
		want int
	}{
		{name: "example 1", num1: 12, num2: 5, want: 17},
		{name: "example 2", num1: -10, num2: 4, want: -6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sum(tt.num1, tt.num2); got != tt.want {
				t.Errorf("sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
