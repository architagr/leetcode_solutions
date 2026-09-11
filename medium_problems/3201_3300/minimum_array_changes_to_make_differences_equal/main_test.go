package minimumarraychangestomakedifferencesequal

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestMinChanges(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{name: "example 1", nums: []int{1, 0, 1, 2, 4, 3}, k: 4, want: 2},
		{name: "example 2", nums: []int{0, 1, 2, 3, 3, 6, 5, 4}, k: 6, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minChanges(tt.nums, tt.k); got != tt.want {
				t.Errorf("minChanges() = %v, want %v", got, tt.want)
			}
		})
	}
}
