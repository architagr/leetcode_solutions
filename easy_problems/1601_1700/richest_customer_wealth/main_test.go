package richestcutomerwealth

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestMaximumWealth(t *testing.T) {
	tests := []struct {
		name     string
		accounts [][]int
		want     int
	}{
		{name: "example 1", accounts: [][]int{[]int{1, 2, 3}, []int{3, 2, 1}}, want: 6},
		{name: "example 2", accounts: [][]int{[]int{1, 5}, []int{7, 3}, []int{3, 5}}, want: 10},
		{name: "example 3", accounts: [][]int{[]int{2, 8, 7}, []int{7, 1, 3}, []int{1, 9, 5}}, want: 17},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumWealth(tt.accounts); got != tt.want {
				t.Errorf("maximumWealth() = %v, want %v", got, tt.want)
			}
		})
	}
}
