package reverse_words_in_a_string_iii

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestReverseWords(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{name: "example 2", s: "Mr Ding", want: "rM gniD"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseWords(tt.s); got != tt.want {
				t.Errorf("ReverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
