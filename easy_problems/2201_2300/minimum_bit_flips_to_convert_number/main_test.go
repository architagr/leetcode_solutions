package minimumbitflipstoconvertnumber

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestMinBitFlips(t *testing.T) {
	tests := []struct {
		name  string
		start int
		goal  int
		want  int
	}{
		{name: "example 1", start: 10, goal: 7, want: 3},
		{name: "example 2", start: 3, goal: 4, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minBitFlips(tt.start, tt.goal); got != tt.want {
				t.Errorf("minBitFlips() = %v, want %v", got, tt.want)
			}
		})
	}
}
