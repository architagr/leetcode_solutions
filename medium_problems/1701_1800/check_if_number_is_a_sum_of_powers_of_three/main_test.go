package check_if_number_is_a_sum_of_powers_of_three

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestCheckPowersOfThree(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{name: "example 1", n: 12, want: true},
		{name: "example 2", n: 91, want: true},
		{name: "example 3", n: 21, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckPowersOfThree(tt.n); got != tt.want {
				t.Errorf("CheckPowersOfThree() = %v, want %v", got, tt.want)
			}
		})
	}
}
