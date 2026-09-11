package countsubstringswithkfrequencycharactersi

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestNumberOfSubstrings(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want int
	}{
		{name: "example 1", s: "abacb", k: 2, want: 4},
		{name: "example 2", s: "abcde", k: 1, want: 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfSubstrings(tt.s, tt.k); got != tt.want {
				t.Errorf("numberOfSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
