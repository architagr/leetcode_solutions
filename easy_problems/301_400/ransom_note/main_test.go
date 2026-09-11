package ransom_note

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestCanConstruct(t *testing.T) {
	tests := []struct {
		name       string
		ransomNote string
		magazine   string
		want       bool
	}{
		{name: "example 1", ransomNote: "a", magazine: "b", want: false},
		{name: "example 2", ransomNote: "aa", magazine: "ab", want: false},
		{name: "example 3", ransomNote: "aa", magazine: "aab", want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanConstruct(tt.ransomNote, tt.magazine); got != tt.want {
				t.Errorf("CanConstruct() = %v, want %v", got, tt.want)
			}
		})
	}
}
