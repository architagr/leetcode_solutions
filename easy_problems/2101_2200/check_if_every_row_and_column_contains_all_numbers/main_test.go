package check_every_row_column_contains_all_numbers

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestCheckValid(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   bool
	}{
		{name: "example 1", matrix: [][]int{[]int{1, 2, 3}, []int{3, 1, 2}, []int{2, 3, 1}}, want: true},
		{name: "example 2", matrix: [][]int{[]int{1, 1, 1}, []int{1, 2, 3}, []int{1, 2, 3}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckValid(tt.matrix); got != tt.want {
				t.Errorf("CheckValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
