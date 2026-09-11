package maximumnumberofvowelsinasubstringofgivenlength

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestMaxVowels(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want int
	}{
		{name: "example 1", s: "abciiidef", k: 3, want: 3},
		{name: "example 2", s: "aeiou", k: 2, want: 2},
		{name: "example 3", s: "leetcode", k: 3, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxVowels(tt.s, tt.k); got != tt.want {
				t.Errorf("maxVowels() = %v, want %v", got, tt.want)
			}
		})
	}
}
