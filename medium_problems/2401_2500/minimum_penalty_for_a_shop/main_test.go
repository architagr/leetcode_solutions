package minimumpenaltyforashop

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestBestClosingTime(t *testing.T) {
	tests := []struct {
		name      string
		customers string
		want      int
	}{
		{name: "example 1", customers: "YYNY", want: 2},
		{name: "example 2", customers: "NNNNN", want: 0},
		{name: "example 3", customers: "YYYY", want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bestClosingTime(tt.customers); got != tt.want {
				t.Errorf("bestClosingTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
