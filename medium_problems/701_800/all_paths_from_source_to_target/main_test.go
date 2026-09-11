package allpathsfromsourcetotarget

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

// normalize sorts the result, since this problem accepts any order.
func normalize(v [][]int) [][]int {
	out := append([][]int{}, v...)
	for i := range out {
		sort.Slice(out[i], func(a, b int) bool { return fmt.Sprint(out[i][a]) < fmt.Sprint(out[i][b]) })
	}
	sort.Slice(out, func(a, b int) bool { return fmt.Sprint(out[a]) < fmt.Sprint(out[b]) })
	return out
}

// Cases are the worked examples from the problem statement on LeetCode.
func TestAllPathsSourceTarget(t *testing.T) {
	tests := []struct {
		name  string
		graph [][]int
		want  [][]int
	}{
		{name: "example 1", graph: [][]int{[]int{1, 2}, []int{3}, []int{3}, []int{}}, want: [][]int{[]int{0, 1, 3}, []int{0, 2, 3}}},
		{name: "example 2", graph: [][]int{[]int{4, 3, 1}, []int{3, 2, 4}, []int{3}, []int{4}, []int{}}, want: [][]int{[]int{0, 4}, []int{0, 3, 4}, []int{0, 1, 3, 4}, []int{0, 1, 2, 3, 4}, []int{0, 1, 4}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := allPathsSourceTarget(tt.graph); !reflect.DeepEqual(normalize(got), normalize(tt.want)) {
				t.Errorf("allPathsSourceTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
