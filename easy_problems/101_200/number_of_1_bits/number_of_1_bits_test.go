package number_of_1_bits

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
// The statement now states n in decimal; this solution keeps the older
// uint32 signature, so the cases are written in that type.
func TestHammingWeight(t *testing.T) {
	tests := []struct {
		name string
		num  uint32
		want int
	}{
		{name: "example 1", num: 11, want: 3},
		{name: "example 2", num: 128, want: 1},
		{name: "example 3", num: 2147483645, want: 30},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HammingWeight(tt.num); got != tt.want {
				t.Errorf("HammingWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
