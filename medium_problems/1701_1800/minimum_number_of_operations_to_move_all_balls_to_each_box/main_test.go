package minimumnumberofoperationstomoveallballstoeachbox

import (
	"reflect"
	"testing"
)

// Cases are the worked examples from the problem statement on LeetCode.
func TestMinOperations(t *testing.T) {
	tests := []struct {
		name  string
		boxes string
		want  []int
	}{
		{name: "example 1", boxes: "110", want: []int{1, 1, 3}},
		{name: "example 2", boxes: "001011", want: []int{11, 8, 5, 4, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations(tt.boxes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
