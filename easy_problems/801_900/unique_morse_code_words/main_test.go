package unique_morse_code_words

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestUniqueMorseRepresentations(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  int
	}{
		{name: "example 1", words: []string{"gin", "zen", "gig", "msg"}, want: 2},
		{name: "example 2", words: []string{"a"}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueMorseRepresentations(tt.words); got != tt.want {
				t.Errorf("UniqueMorseRepresentations() = %v, want %v", got, tt.want)
			}
		})
	}
}
