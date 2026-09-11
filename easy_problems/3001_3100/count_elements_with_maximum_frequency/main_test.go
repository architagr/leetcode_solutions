package countelementswithmaximumfrequency

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestMaxFrequencyElements(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "example 1", nums: []int{1, 2, 2, 3, 1, 4}, want: 4},
		{name: "example 2", nums: []int{1, 2, 3, 4, 5}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxFrequencyElements(tt.nums); got != tt.want {
				t.Errorf("maxFrequencyElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
