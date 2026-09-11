package thetwosneakynumbersofdigitville

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

// normalize sorts the result, since this problem accepts any order.
func normalize(v []int) []int {
	out := append([]int{}, v...)
	sort.Slice(out, func(a, b int) bool { return fmt.Sprint(out[a]) < fmt.Sprint(out[b]) })
	return out
}

// Cases are the worked examples from the problem statement on LeetCode.
func TestGetSneakyNumbers(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{name: "example 1", nums: []int{0, 1, 1, 0}, want: []int{0, 1}},
		{name: "example 2", nums: []int{0, 3, 2, 1, 3, 2}, want: []int{2, 3}},
		{name: "example 3", nums: []int{7, 1, 5, 4, 3, 4, 6, 0, 9, 5, 8, 2}, want: []int{4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSneakyNumbers(tt.nums); !reflect.DeepEqual(normalize(got), normalize(tt.want)) {
				t.Errorf("getSneakyNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
