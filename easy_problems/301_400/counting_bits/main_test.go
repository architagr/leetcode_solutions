package counting_bits

import (
	"reflect"
	"testing"
)

// Cases are the worked examples from the problem statement on LeetCode.
func TestCountBits(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []int
	}{
		{name: "example 1", n: 2, want: []int{0, 1, 1}},
		{name: "example 2", n: 5, want: []int{0, 1, 1, 2, 1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountBits(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
