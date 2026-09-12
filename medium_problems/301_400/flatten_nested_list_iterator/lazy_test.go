package flattennestedlistiterator

import (
	"reflect"
	"testing"
)

func drain(l []*NestedInteger) []int {
	it := Constructor(l)
	out := []int{}
	for it.HasNext() {
		out = append(out, it.Next())
	}
	return out
}

func TestEdgeCases(t *testing.T) {
	cases := []struct {
		name string
		in   []*NestedInteger
		want []int
	}{
		{"empty outer", []*NestedInteger{}, []int{}},
		{"flat", []*NestedInteger{{n: 1}, {n: 2}, {n: 3}}, []int{1, 2, 3}},
		{"deep left spine", []*NestedInteger{
			{list: []*NestedInteger{{list: []*NestedInteger{{list: []*NestedInteger{{n: 7}}}}}}},
			{n: 8},
		}, []int{7, 8}},
		{"trailing nested", []*NestedInteger{
			{n: 1},
			{list: []*NestedInteger{{n: 2}, {n: 3}}},
		}, []int{1, 2, 3}},
	}
	for _, c := range cases {
		if got := drain(c.in); !reflect.DeepEqual(got, c.want) {
			t.Errorf("%s: got %v want %v", c.name, got, c.want)
		}
	}
}

// The point of the rewrite: construction must not walk the structure.
func TestConstructionIsLazy(t *testing.T) {
	in := []*NestedInteger{
		{list: []*NestedInteger{{n: 1}, {n: 2}}},
		{n: 3},
	}
	it := Constructor(in)
	if len(it.stack) != 1 {
		t.Fatalf("constructor should hold exactly one frame, got %d", len(it.stack))
	}
	if it.stack[0].i != 0 {
		t.Fatalf("constructor should not have advanced, cursor at %d", it.stack[0].i)
	}
}

// HasNext must be idempotent: calling it repeatedly cannot consume anything.
func TestHasNextIsIdempotent(t *testing.T) {
	it := Constructor([]*NestedInteger{{list: []*NestedInteger{{n: 5}}}, {n: 6}})
	for i := 0; i < 10; i++ {
		if !it.HasNext() {
			t.Fatal("HasNext went false while values remained")
		}
	}
	if v := it.Next(); v != 5 {
		t.Fatalf("got %d, want 5", v)
	}
	if v := it.Next(); v != 6 {
		t.Fatalf("got %d, want 6", v)
	}
	if it.HasNext() {
		t.Fatal("HasNext true after the last value")
	}
}
