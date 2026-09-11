package circularsentence

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestIsCircularSentence(t *testing.T) {
	tests := []struct {
		name     string
		sentence string
		want     bool
	}{
		{name: "example 1", sentence: "leetcode exercises sound delightful", want: true},
		{name: "example 2", sentence: "eetcode", want: true},
		{name: "example 3", sentence: "Leetcode is cool", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isCircularSentence(tt.sentence); got != tt.want {
				t.Errorf("isCircularSentence() = %v, want %v", got, tt.want)
			}
		})
	}
}
