package reverse_vowels_of_a_string

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestReverseVowels(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{name: "example 1", s: "IceCreAm", want: "AceCreIm"},
		{name: "example 2", s: "leetcode", want: "leotcede"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseVowels(tt.s); got != tt.want {
				t.Errorf("ReverseVowels() = %v, want %v", got, tt.want)
			}
		})
	}
}
