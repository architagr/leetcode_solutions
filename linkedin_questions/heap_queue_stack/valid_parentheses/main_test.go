package valid_parentheses

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestIsValid(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{name: "example 1", s: "()", want: true},
		{name: "example 2", s: "()[]{}", want: true},
		{name: "example 3", s: "(]", want: false},
		{name: "example 4", s: "([])", want: true},
		{name: "example 5", s: "([)]", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValid(tt.s); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
