package possiblebipartition

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestPossibleBipartition(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		dislikes [][]int
		want     bool
	}{
		{name: "example 1", n: 4, dislikes: [][]int{[]int{1, 2}, []int{1, 3}, []int{2, 4}}, want: true},
		{name: "example 2", n: 3, dislikes: [][]int{[]int{1, 2}, []int{1, 3}, []int{2, 3}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := possibleBipartition(tt.n, tt.dislikes); got != tt.want {
				t.Errorf("possibleBipartition() = %v, want %v", got, tt.want)
			}
		})
	}
}
