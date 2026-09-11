package findthedifference

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestFindTheDifference(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want byte
	}{
		{name: "example 1", s: "abcd", t: "abcde", want: 'e'},
		{name: "example 2", s: "", t: "y", want: 'y'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTheDifference(tt.s, tt.t); got != tt.want {
				t.Errorf("findTheDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
