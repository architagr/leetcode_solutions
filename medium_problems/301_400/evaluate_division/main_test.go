package evaluatedivision

import (
	"reflect"
	"testing"
)

// Cases are the worked examples from the problem statement on LeetCode.
func TestCalcEquation(t *testing.T) {
	tests := []struct {
		name      string
		equations [][]string
		values    []float64
		queries   [][]string
		want      []float64
	}{
		{name: "example 1", equations: [][]string{[]string{"a", "b"}, []string{"b", "c"}}, values: []float64{2.0, 3.0}, queries: [][]string{[]string{"a", "c"}, []string{"b", "a"}, []string{"a", "e"}, []string{"a", "a"}, []string{"x", "x"}}, want: []float64{6.0, 0.5, -1.0, 1.0, -1.0}},
		{name: "example 2", equations: [][]string{[]string{"a", "b"}, []string{"b", "c"}, []string{"bc", "cd"}}, values: []float64{1.5, 2.5, 5.0}, queries: [][]string{[]string{"a", "c"}, []string{"c", "b"}, []string{"bc", "cd"}, []string{"cd", "bc"}}, want: []float64{3.75, 0.4, 5.0, 0.2}},
		{name: "example 3", equations: [][]string{[]string{"a", "b"}}, values: []float64{0.5}, queries: [][]string{[]string{"a", "b"}, []string{"b", "a"}, []string{"a", "c"}, []string{"x", "y"}}, want: []float64{0.5, 2.0, -1.0, -1.0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcEquation(tt.equations, tt.values, tt.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calcEquation() = %v, want %v", got, tt.want)
			}
		})
	}
}
