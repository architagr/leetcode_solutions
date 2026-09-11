package shiftinglettersii

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestShiftingLetters(t *testing.T) {
	tests := []struct {
		name   string
		s      string
		shifts [][]int
		want   string
	}{
		{name: "example 1", s: "abc", shifts: [][]int{[]int{0, 1, 0}, []int{1, 2, 1}, []int{0, 2, 1}}, want: "ace"},
		{name: "example 2", s: "dztz", shifts: [][]int{[]int{0, 0, 0}, []int{1, 1, 1}}, want: "catz"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shiftingLetters(tt.s, tt.shifts); got != tt.want {
				t.Errorf("shiftingLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}
