package groupanagrams

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

// normalize sorts the result, since this problem accepts any order.
func normalize(v [][]string) [][]string {
	out := append([][]string{}, v...)
	for i := range out {
		sort.Slice(out[i], func(a, b int) bool { return fmt.Sprint(out[i][a]) < fmt.Sprint(out[i][b]) })
	}
	sort.Slice(out, func(a, b int) bool { return fmt.Sprint(out[a]) < fmt.Sprint(out[b]) })
	return out
}

// Cases are the worked examples from the problem statement on LeetCode.
func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want [][]string
	}{
		{name: "example 1", strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"}, want: [][]string{[]string{"bat"}, []string{"nat", "tan"}, []string{"ate", "eat", "tea"}}},
		{name: "example 2", strs: []string{""}, want: [][]string{[]string{""}}},
		{name: "example 3", strs: []string{"a"}, want: [][]string{[]string{"a"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupAnagrams(tt.strs); !reflect.DeepEqual(normalize(got), normalize(tt.want)) {
				t.Errorf("groupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
