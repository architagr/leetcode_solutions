package editdistance

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestMinDistance(t *testing.T) {
	tests := []struct {
		name  string
		word1 string
		word2 string
		want  int
	}{
		{name: "example 1", word1: "horse", word2: "ros", want: 3},
		{name: "example 2", word1: "intention", word2: "execution", want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDistance(tt.word1, tt.word2); got != tt.want {
				t.Errorf("minDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
