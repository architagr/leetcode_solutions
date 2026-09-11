package shiftingletters

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestShiftingLetters(t *testing.T) {
	tests := []struct {
		name   string
		s      string
		shifts []int
		want   string
	}{
		{name: "example 1", s: "abc", shifts: []int{3, 5, 9}, want: "rpl"},
		{name: "example 2", s: "aaa", shifts: []int{1, 2, 3}, want: "gfd"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shiftingLetters(tt.s, tt.shifts); got != tt.want {
				t.Errorf("shiftingLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}
