package detectcapital

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestDetectCapitalUse(t *testing.T) {
	tests := []struct {
		name string
		word string
		want bool
	}{
		{name: "example 1", word: "USA", want: true},
		{name: "example 2", word: "FlaG", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCapitalUse(tt.word); got != tt.want {
				t.Errorf("detectCapitalUse() = %v, want %v", got, tt.want)
			}
		})
	}
}
