package countvowelstringsinranges

import (
	"reflect"
	"testing"
)

// Cases are the worked examples from the problem statement on LeetCode.
func TestVowelStrings(t *testing.T) {
	tests := []struct {
		name    string
		words   []string
		queries [][]int
		want    []int
	}{
		{name: "example 1", words: []string{"aba", "bcb", "ece", "aa", "e"}, queries: [][]int{[]int{0, 2}, []int{1, 4}, []int{1, 1}}, want: []int{2, 3, 0}},
		{name: "example 2", words: []string{"a", "e", "i"}, queries: [][]int{[]int{0, 2}, []int{0, 1}, []int{2, 2}}, want: []int{3, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := vowelStrings(tt.words, tt.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("vowelStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
