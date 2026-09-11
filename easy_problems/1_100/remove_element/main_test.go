package removeelement

import (
	"reflect"
	"sort"
	"testing"
)

// The statement judges only the first k entries of nums, in any order,
// so the case sorts that prefix before comparing and ignores the rest.
func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name      string
		nums      []int
		val       int
		want      int
		remaining []int
	}{
		{name: "example 1", nums: []int{3, 2, 2, 3}, val: 3, want: 2, remaining: []int{2, 2}},
		{name: "example 2", nums: []int{0, 1, 2, 2, 3, 0, 4, 2}, val: 2, want: 5, remaining: []int{0, 0, 1, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := append([]int{}, tt.nums...)
			k := removeElement(nums, tt.val)
			if k != tt.want {
				t.Fatalf("removeElement() = %v, want %v", k, tt.want)
			}
			got := append([]int{}, nums[:k]...)
			sort.Ints(got)
			if !reflect.DeepEqual(got, tt.remaining) {
				t.Errorf("nums[:k] = %v, want %v", got, tt.remaining)
			}
		})
	}
}
