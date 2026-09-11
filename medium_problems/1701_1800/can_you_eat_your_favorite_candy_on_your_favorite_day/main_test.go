package canyoueatyourfavoritecandyonyourfavoriteday

import (
	"reflect"
	"testing"
)

// Cases are the worked examples from the problem statement on LeetCode.
func TestCanEat(t *testing.T) {
	tests := []struct {
		name         string
		candiesCount []int
		queries      [][]int
		want         []bool
	}{
		{name: "example 1", candiesCount: []int{7, 4, 5, 3, 8}, queries: [][]int{[]int{0, 2, 2}, []int{4, 2, 4}, []int{2, 13, 1000000000}}, want: []bool{true, false, true}},
		{name: "example 2", candiesCount: []int{5, 2, 6, 4, 1}, queries: [][]int{[]int{3, 1, 2}, []int{4, 10, 3}, []int{3, 10, 100}, []int{4, 100, 30}, []int{1, 3, 1}}, want: []bool{false, true, true, false, false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canEat(tt.candiesCount, tt.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("canEat() = %v, want %v", got, tt.want)
			}
		})
	}
}
