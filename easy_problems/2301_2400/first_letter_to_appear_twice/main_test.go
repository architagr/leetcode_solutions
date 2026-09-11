package first_letter_to_appear_twice

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestRepeatedCharacter(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want byte
	}{
		{name: "example 1", s: "abccbaacz", want: 'c'},
		{name: "example 2", s: "abcdd", want: 'd'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RepeatedCharacter(tt.s); got != tt.want {
				t.Errorf("RepeatedCharacter() = %v, want %v", got, tt.want)
			}
		})
	}
}
