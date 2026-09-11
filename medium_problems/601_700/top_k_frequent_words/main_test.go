package topkfrequentwords

import (
	"reflect"
	"testing"
)

// Cases are the worked examples from the problem statement on LeetCode.
func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		k     int
		want  []string
	}{
		{name: "example 1", words: []string{"i", "love", "leetcode", "i", "love", "coding"}, k: 2, want: []string{"i", "love"}},
		{name: "example 2", words: []string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}, k: 4, want: []string{"the", "is", "sunny", "day"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := topKFrequent(tt.words, tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topKFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}
