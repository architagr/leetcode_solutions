package reverse_string_ii

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestReverseStr(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want string
	}{
		{name: "example 1", s: "abcdefg", k: 2, want: "bacdfeg"},
		{name: "example 2", s: "abcd", k: 2, want: "bacd"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseStr(tt.s, tt.k); got != tt.want {
				t.Errorf("ReverseStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
