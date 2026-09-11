package poor_pigs

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestPoorPigs(t *testing.T) {
	tests := []struct {
		name          string
		buckets       int
		minutesToDie  int
		minutesToTest int
		want          int
	}{
		{name: "example 1", buckets: 4, minutesToDie: 15, minutesToTest: 15, want: 2},
		{name: "example 2", buckets: 4, minutesToDie: 15, minutesToTest: 30, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PoorPigs(tt.buckets, tt.minutesToDie, tt.minutesToTest); got != tt.want {
				t.Errorf("PoorPigs() = %v, want %v", got, tt.want)
			}
		})
	}
}
