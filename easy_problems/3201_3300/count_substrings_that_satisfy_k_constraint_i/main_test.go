package countsubstringsthatsatisfykconstrainti

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestCountKConstraintSubstrings(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want int
	}{
		{name: "example 1", s: "10101", k: 1, want: 12},
		{name: "example 2", s: "1010101", k: 2, want: 25},
		{name: "example 3", s: "11111", k: 1, want: 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countKConstraintSubstrings(tt.s, tt.k); got != tt.want {
				t.Errorf("countKConstraintSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
