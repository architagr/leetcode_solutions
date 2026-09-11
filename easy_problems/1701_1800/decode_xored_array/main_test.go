package decodexoredarray

import (
	"reflect"
	"testing"
)

// Cases are the worked examples from the problem statement on LeetCode.
func TestDecode(t *testing.T) {
	tests := []struct {
		name    string
		encoded []int
		first   int
		want    []int
	}{
		{name: "example 1", encoded: []int{1, 2, 3}, first: 1, want: []int{1, 0, 2, 1}},
		{name: "example 2", encoded: []int{6, 2, 7, 3}, first: 4, want: []int{4, 2, 0, 7, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decode(tt.encoded, tt.first); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
