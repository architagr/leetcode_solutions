package matrix_diagonal_sum

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestDIagomalSum(t *testing.T) {
	tests := []struct {
		name string
		mat  [][]int
		want int
	}{
		{name: "example 3", mat: [][]int{[]int{5}}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DIagomalSum(tt.mat); got != tt.want {
				t.Errorf("DIagomalSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
