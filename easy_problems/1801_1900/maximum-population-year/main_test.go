package maximumpopulationyear

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestMaximumPopulation(t *testing.T) {
	tests := []struct {
		name string
		logs [][]int
		want int
	}{
		{name: "example 1", logs: [][]int{[]int{1993, 1999}, []int{2000, 2010}}, want: 1993},
		{name: "example 2", logs: [][]int{[]int{1950, 1961}, []int{1960, 1971}, []int{1970, 1981}}, want: 1960},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumPopulation(tt.logs); got != tt.want {
				t.Errorf("maximumPopulation() = %v, want %v", got, tt.want)
			}
		})
	}
}
