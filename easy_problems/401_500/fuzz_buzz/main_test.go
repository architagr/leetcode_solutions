package fuzzbuzz

import (
	"reflect"
	"testing"
)

// Cases are the worked examples from the problem statement on LeetCode.
func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []string
	}{
		{name: "example 1", n: 3, want: []string{"1", "2", "Fizz"}},
		{name: "example 2", n: 5, want: []string{"1", "2", "Fizz", "4", "Buzz"}},
		{name: "example 3", n: 15, want: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fizzBuzz(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fizzBuzz() = %v, want %v", got, tt.want)
			}
		})
	}
}
