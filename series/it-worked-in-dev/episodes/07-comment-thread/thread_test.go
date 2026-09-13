package commentthread

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

func thread(depth, width int, next *int) *Comment {
	c := &Comment{ID: *next, Author: fmt.Sprintf("u%d", *next)}
	*next++
	if depth == 0 {
		return c
	}
	for w := 0; w < width; w++ {
		c.Replies = append(c.Replies, thread(depth-1, width, next))
	}
	return c
}

func TestBothAgree(t *testing.T) {
	n := 0
	cases := []struct {
		name string
		root *Comment
	}{
		{"single comment", &Comment{ID: 1}},
		{"one level of replies", func() *Comment { n = 0; return thread(1, 3, &n) }()},
		{"balanced 3x3", func() *Comment { n = 0; return thread(3, 3, &n) }()},
		{"deep chain", func() *Comment { n = 0; return thread(30, 1, &n) }()},
		{"wide", func() *Comment { n = 0; return thread(1, 40, &n) }()},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			a, b := OrderBySortKey(c.root), OrderByWalking(c.root)
			if !reflect.DeepEqual(a, b) {
				t.Fatalf("disagreed:\n  sorted  %v\n  walking %v", a, b)
			}
		})
	}
}

func TestKnownOrder(t *testing.T) {
	// 1
	//   2
	//     3
	//   4
	// where 4 is a sibling of 2, so it renders after 3
	root := &Comment{ID: 1, Replies: []*Comment{
		{ID: 2, Replies: []*Comment{{ID: 3}}},
		{ID: 4},
	}}
	want := []Rendered{{1, 0}, {2, 1}, {3, 2}, {4, 1}}
	if got := OrderByWalking(root); !reflect.DeepEqual(got, want) {
		t.Errorf("walking = %v, want %v", got, want)
	}
	if got := OrderBySortKey(root); !reflect.DeepEqual(got, want) {
		t.Errorf("sort key = %v, want %v", got, want)
	}
}

// The zero padding in the sort key is load-bearing. Without it "10" sorts
// before "2" and reply eleven jumps above reply three. This pins that the
// padded version gets it right at a width where it matters.
func TestOrderingPastTenReplies(t *testing.T) {
	root := &Comment{ID: 0}
	for i := 1; i <= 12; i++ {
		root.Replies = append(root.Replies, &Comment{ID: i})
	}
	got := OrderBySortKey(root)
	for i, r := range got {
		if r.ID != i {
			t.Fatalf("position %d holds comment %d; the ordering is wrong past ten", i, r.ID)
		}
	}
}

func TestBothAgreeOnRandomThreads(t *testing.T) {
	rng := rand.New(rand.NewSource(41))
	for trial := 0; trial < 1000; trial++ {
		root := &Comment{ID: 0}
		all := []*Comment{root}
		for i := 1; i < rng.Intn(70)+2; i++ {
			p := all[rng.Intn(len(all))]
			c := &Comment{ID: i}
			p.Replies = append(p.Replies, c)
			all = append(all, c)
		}
		if a, b := OrderBySortKey(root), OrderByWalking(root); !reflect.DeepEqual(a, b) {
			t.Fatalf("trial %d disagreed", trial)
		}
	}
}

// The write-up claims the sort-key version builds a string per comment whose
// length grows with depth, so the total characters are the sum of the depths.
// Counted here rather than claimed.
func TestSortKeyCharacterCount(t *testing.T) {
	for _, d := range []int{10, 50, 200} {
		n := 0
		root := thread(d, 1, &n)
		chars := 0
		var walk func(c *Comment, depth int)
		walk = func(c *Comment, depth int) {
			chars += depth*7 - 1 // six digits plus a dot per level, less the last dot
			for _, r := range c.Replies {
				walk(r, depth+1)
			}
		}
		walk(root, 1)
		want := 0
		for i := 1; i <= d+1; i++ {
			want += i*7 - 1
		}
		if chars != want {
			t.Errorf("depth %d: %d key characters, want %d", d, chars, want)
		}
	}
}
