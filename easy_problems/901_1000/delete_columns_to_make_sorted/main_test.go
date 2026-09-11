package deletecolumnstomakesorted

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestMinDeletionSize(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want int
	}{
		{name: "example 1", strs: []string{"cba", "daf", "ghi"}, want: 1},
		{name: "example 2", strs: []string{"a", "b"}, want: 0},
		{name: "example 3", strs: []string{"zyx", "wvu", "tsr"}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDeletionSize(tt.strs); got != tt.want {
				t.Errorf("minDeletionSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
