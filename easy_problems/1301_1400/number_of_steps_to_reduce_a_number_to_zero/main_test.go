package numberofstepstoreduceanumbertozero

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestNumberOfSteps(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want int
	}{
		{name: "example 1", num: 14, want: 6},
		{name: "example 2", num: 8, want: 4},
		{name: "example 3", num: 123, want: 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfSteps(tt.num); got != tt.want {
				t.Errorf("numberOfSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
