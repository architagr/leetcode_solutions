package valid_anagram

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{name: "example 1", s: "anagram", t: "nagaram", want: true},
		{name: "example 2", s: "rat", t: "car", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAnagram(tt.s, tt.t); got != tt.want {
				t.Errorf("IsAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}
