package find_if_path_exists_in_graph

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestValidPath(t *testing.T) {
	tests := []struct {
		name        string
		n           int
		edges       [][]int
		source      int
		destination int
		want        bool
	}{
		{name: "example 1", n: 3, edges: [][]int{[]int{0, 1}, []int{1, 2}, []int{2, 0}}, source: 0, destination: 2, want: true},
		{name: "example 2", n: 6, edges: [][]int{[]int{0, 1}, []int{0, 2}, []int{3, 5}, []int{5, 4}, []int{4, 3}}, source: 0, destination: 5, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validPath(tt.n, tt.edges, tt.source, tt.destination); got != tt.want {
				t.Errorf("validPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
