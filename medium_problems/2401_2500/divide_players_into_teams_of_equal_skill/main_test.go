package divideplayersintoteamsofequalskill

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestDividePlayers(t *testing.T) {
	tests := []struct {
		name  string
		skill []int
		want  int64
	}{
		{name: "example 1", skill: []int{3, 2, 5, 1, 3, 4}, want: 22},
		{name: "example 2", skill: []int{3, 4}, want: 12},
		{name: "example 3", skill: []int{1, 1, 2, 3}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dividePlayers(tt.skill); got != tt.want {
				t.Errorf("dividePlayers() = %v, want %v", got, tt.want)
			}
		})
	}
}
