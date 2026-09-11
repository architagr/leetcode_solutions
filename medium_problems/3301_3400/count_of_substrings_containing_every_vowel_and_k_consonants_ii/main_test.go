package countofsubstringscontainingeveryvowelandkconsonantsii

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestCountOfSubstrings(t *testing.T) {
	tests := []struct {
		name string
		word string
		k    int
		want int64
	}{
		{name: "example 1", word: "aeioqq", k: 1, want: 0},
		{name: "example 2", word: "aeiou", k: 0, want: 1},
		{name: "example 3", word: "ieaouqqieaouqq", k: 1, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countOfSubstrings(tt.word, tt.k); got != tt.want {
				t.Errorf("countOfSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
