package happy_number

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestIsHappyNum(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{name: "example 1", n: 19, want: true},
		{name: "example 2", n: 2, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isHappyNum(tt.n); got != tt.want {
				t.Errorf("isHappyNum() = %v, want %v", got, tt.want)
			}
		})
	}
}
