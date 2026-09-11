package minimumamountoftimetocollectgarbage

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestGarbageCollection(t *testing.T) {
	tests := []struct {
		name    string
		garbage []string
		travel  []int
		want    int
	}{
		{name: "example 1", garbage: []string{"G", "P", "GP", "GG"}, travel: []int{2, 4, 3}, want: 21},
		{name: "example 2", garbage: []string{"MMM", "PGM", "GP"}, travel: []int{3, 10}, want: 37},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := garbageCollection(tt.garbage, tt.travel); got != tt.want {
				t.Errorf("garbageCollection() = %v, want %v", got, tt.want)
			}
		})
	}
}
