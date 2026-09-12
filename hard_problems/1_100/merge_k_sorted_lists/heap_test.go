package merge_k_sorted_lists

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func build(vals []int) *ListNode {
	d := &ListNode{}
	c := d
	for _, v := range vals {
		c.Next = &ListNode{Val: v}
		c = c.Next
	}
	return d.Next
}

// drain walks the result with a step cap, so a cycle fails the test instead
// of hanging it. Splicing existing nodes makes a cycle a real possibility if
// the merged list is not terminated.
func drain(t *testing.T, h *ListNode, cap int) []int {
	t.Helper()
	out := []int{}
	for n, steps := h, 0; n != nil; n, steps = n.Next, steps+1 {
		if steps > cap {
			t.Fatalf("cycle: walked past %d nodes", cap)
		}
		out = append(out, n.Val)
	}
	return out
}

func TestMergeKListsEdgeCases(t *testing.T) {
	cases := []struct {
		name string
		in   [][]int
		want []int
	}{
		{"no lists", nil, []int{}},
		{"one empty list", [][]int{nil}, []int{}},
		{"all lists empty", [][]int{nil, nil, nil}, []int{}},
		{"single list", [][]int{{1, 2, 3}}, []int{1, 2, 3}},
		{"empties interleaved", [][]int{nil, {2}, nil, {1, 3}}, []int{1, 2, 3}},
		{"all equal", [][]int{{1, 1}, {1}, {1, 1}}, []int{1, 1, 1, 1, 1}},
		{"one long one short", [][]int{{1, 2, 3, 4, 5}, {6}}, []int{1, 2, 3, 4, 5, 6}},
		{"negatives", [][]int{{-3, -1}, {-2, 0}}, []int{-3, -2, -1, 0}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			lists := make([]*ListNode, len(c.in))
			total := 0
			for i, v := range c.in {
				lists[i] = build(v)
				total += len(v)
			}
			got := drain(t, MergeKLists(lists), total+10)
			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("got %v, want %v", got, c.want)
			}
		})
	}
}

// The merged output is every input value in sorted order, which makes a
// reference to check against on random input.
func TestMergeKListsAgainstSort(t *testing.T) {
	r := rand.New(rand.NewSource(11))
	for trial := 0; trial < 1000; trial++ {
		k := r.Intn(5)
		lists := make([]*ListNode, k)
		want := []int{}
		for i := 0; i < k; i++ {
			n := r.Intn(5)
			vals := make([]int, n)
			for j := range vals {
				vals[j] = r.Intn(10) - 5
			}
			sort.Ints(vals)
			want = append(want, vals...)
			lists[i] = build(vals)
		}
		sort.Ints(want)
		if len(want) == 0 {
			want = []int{}
		}
		got := drain(t, MergeKLists(lists), len(want)+10)
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("trial %d: got %v, want %v", trial, got, want)
		}
	}
}

// Splicing means the result must reuse the input nodes, not copy them.
func TestMergeKListsSplicesRatherThanCopies(t *testing.T) {
	a := build([]int{1, 3})
	b := build([]int{2})
	originals := map[*ListNode]bool{a: true, a.Next: true, b: true}

	for n := MergeKLists([]*ListNode{a, b}); n != nil; n = n.Next {
		if !originals[n] {
			t.Fatal("result contains a node that was not in the input: values were copied")
		}
	}
}
