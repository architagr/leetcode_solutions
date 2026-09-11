package minimumnumberofarrowstoburstballoons

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestFindMinArrowShots(t *testing.T) {
	tests := []struct {
		name   string
		points [][]int
		want   int
	}{
		{name: "example 1", points: [][]int{[]int{10, 16}, []int{2, 8}, []int{1, 6}, []int{7, 12}}, want: 2},
		{name: "example 2", points: [][]int{[]int{1, 2}, []int{3, 4}, []int{5, 6}, []int{7, 8}}, want: 4},
		{name: "example 3", points: [][]int{[]int{1, 2}, []int{2, 3}, []int{3, 4}, []int{4, 5}}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinArrowShots(tt.points); got != tt.want {
				t.Errorf("findMinArrowShots() = %v, want %v", got, tt.want)
			}
		})
	}
}
