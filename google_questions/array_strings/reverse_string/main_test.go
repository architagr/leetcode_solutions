package reverse_string

import (
	"reflect"
	"testing"
)

// ReverseString works in place, so the case compares the argument after the call.
// Cases are the worked examples from the problem statement on LeetCode.
func TestReverseString(t *testing.T) {
	tests := []struct {
		name string
		s    []byte
		want []byte
	}{
		{name: "example 1", s: []byte{'h', 'e', 'l', 'l', 'o'}, want: []byte{'o', 'l', 'l', 'e', 'h'}},
		{name: "example 2", s: []byte{'H', 'a', 'n', 'n', 'a', 'h'}, want: []byte{'h', 'a', 'n', 'n', 'a', 'H'}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arg := append([]byte{}, tt.s...)
			ReverseString(arg)
			got := arg
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseString() = %v, want %v", got, tt.want)
			}
		})
	}
}
