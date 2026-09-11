package unique_number_of_occurrences

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestUniqueOccurrences(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want bool
	}{
		{name: "example 1", arr: []int{1, 2, 2, 1, 1, 3}, want: true},
		{name: "example 2", arr: []int{1, 2}, want: false},
		{name: "example 3", arr: []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueOccurrences(tt.arr); got != tt.want {
				t.Errorf("UniqueOccurrences() = %v, want %v", got, tt.want)
			}
		})
	}
}
