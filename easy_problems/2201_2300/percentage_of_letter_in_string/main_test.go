package percentage_of_letter_in_string

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestPercentageLetter(t *testing.T) {
	tests := []struct {
		name   string
		s      string
		letter byte
		want   int
	}{
		{name: "example 1", s: "foobar", letter: 'o', want: 33},
		{name: "example 2", s: "jjjj", letter: 'k', want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := percentageLetter(tt.s, tt.letter); got != tt.want {
				t.Errorf("percentageLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}
