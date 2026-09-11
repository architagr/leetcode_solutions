package subarraysumsdivisiblebyk

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestSubarraysDivByK(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{name: "example 1", nums: []int{4, 5, 0, -2, -3, 1}, k: 5, want: 7},
		{name: "example 2", nums: []int{5}, k: 9, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarraysDivByK(tt.nums, tt.k); got != tt.want {
				t.Errorf("subarraysDivByK() = %v, want %v", got, tt.want)
			}
		})
	}
}
