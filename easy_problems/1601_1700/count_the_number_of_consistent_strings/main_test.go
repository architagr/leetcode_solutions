package count_number_of_consistent_strings

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestCountConsistentStrings(t *testing.T) {
	tests := []struct {
		name    string
		allowed string
		words   []string
		want    int
	}{
		{name: "example 1", allowed: "ab", words: []string{"ad", "bd", "aaab", "baa", "badab"}, want: 2},
		{name: "example 2", allowed: "abc", words: []string{"a", "b", "c", "ab", "ac", "bc", "abc"}, want: 7},
		{name: "example 3", allowed: "cad", words: []string{"cc", "acd", "b", "ba", "bac", "bad", "ac", "d"}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountConsistentStrings(tt.allowed, tt.words); got != tt.want {
				t.Errorf("CountConsistentStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
