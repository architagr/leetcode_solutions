package longestnicesubstring

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestLongestNiceSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{name: "example 1", s: "YazaAay", want: "aAa"},
		{name: "example 2", s: "Bb", want: "Bb"},
		{name: "example 3", s: "c", want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestNiceSubstring(tt.s); got != tt.want {
				t.Errorf("longestNiceSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
