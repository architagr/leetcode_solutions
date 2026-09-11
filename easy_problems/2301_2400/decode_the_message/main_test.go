package decode_the_message

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestDecodeMessage(t *testing.T) {
	tests := []struct {
		name    string
		key     string
		message string
		want    string
	}{
		{name: "example 1", key: "the quick brown fox jumps over the lazy dog", message: "vkbs bs t suepuv", want: "this is a secret"},
		{name: "example 2", key: "eljuxhpwnyrdgtqkviszcfmabo", message: "zwx hnfx lqantp mnoeius ycgk vcnjrdb", want: "the five boxing wizards jump quickly"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecodeMessage(tt.key, tt.message); got != tt.want {
				t.Errorf("DecodeMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
