package intervalsbetweenidenticalelements

import (
	"reflect"
	"testing"
)

// Cases are the worked examples from the problem statement on LeetCode.
func TestGetDistances(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want []int64
	}{
		{name: "example 1", arr: []int{2, 1, 3, 1, 2, 3, 3}, want: []int64{4, 2, 7, 2, 4, 4, 5}},
		{name: "example 2", arr: []int{10, 5, 10, 10}, want: []int64{5, 0, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDistances(tt.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getDistances() = %v, want %v", got, tt.want)
			}
		})
	}
}
