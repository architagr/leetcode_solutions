package reduce_array_size_to_half

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestMinSetSize(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{name: "example 1", arr: []int{3, 3, 3, 3, 5, 5, 5, 2, 2, 7}, want: 2},
		{name: "example 2", arr: []int{7, 7, 7, 7, 7, 7}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinSetSize(tt.arr); got != tt.want {
				t.Errorf("MinSetSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
