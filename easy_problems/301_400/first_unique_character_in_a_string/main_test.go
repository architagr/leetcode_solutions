package first_unique_character_in_string

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestFirstUniqChar(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{name: "example 1", s: "leetcode", want: 0},
		{name: "example 2", s: "loveleetcode", want: 2},
		{name: "example 3", s: "aabb", want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstUniqChar(tt.s); got != tt.want {
				t.Errorf("firstUniqChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
