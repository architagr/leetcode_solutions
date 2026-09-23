package collapsedtreeview

import (
	"fmt"
	"reflect"
	"testing"
)

// thread builds a reply tree `depth` deep where every comment that is not a
// leaf has `branch` replies. IDs carry the depth so a wrong level shows up as
// a wrong ID rather than as a silent reordering.
func thread(depth, branch int) *Comment {
	n := 0
	var build func(level int) *Comment
	build = func(level int) *Comment {
		n++
		c := &Comment{ID: fmt.Sprintf("d%d-%d", level, n)}
		if level == depth {
			return c
		}
		c.Replies = make([]*Comment, 0, branch)
		for i := 0; i < branch; i++ {
			c.Replies = append(c.Replies, build(level+1))
		}
		return c
	}
	return build(0)
}

// chain is one person replying to themselves: the shape a thread takes when
// two people argue.
func chain(depth int) *Comment { return thread(depth, 1) }

// wide is a popular post: one comment, and everybody replies to it directly.
func wide(width int) *Comment {
	root := &Comment{ID: "d0-1"}
	for i := 0; i < width; i++ {
		root.Replies = append(root.Replies, &Comment{ID: fmt.Sprintf("d1-%d", i+1)})
	}
	return root
}

func count(c *Comment) int {
	if c == nil {
		return 0
	}
	n := 1
	for _, r := range c.Replies {
		n += count(r)
	}
	return n
}

func TestAllThreeCorrectVersionsAgree(t *testing.T) {
	for _, tr := range []*Comment{
		thread(0, 3), thread(1, 4), thread(3, 3), thread(5, 2),
		chain(40), wide(500), lopsided(),
	} {
		want := VisibleByAllLevels(tr)
		if got := VisibleByDepthMap(tr); !reflect.DeepEqual(got, want) {
			t.Fatalf("depth map disagrees:\n got %v\nwant %v", got, want)
		}
		if got := VisibleByFirstArrival(tr); !reflect.DeepEqual(got, want) {
			t.Fatalf("first arrival disagrees:\n got %v\nwant %v", got, want)
		}
	}
}

// lopsided is the thread the tempting version gets wrong: the newest reply to
// the root has no replies of its own, while an older sibling has a long tail.
//
//	root
//	├── old      <- has the deep replies
//	│   └── old-1
//	│       └── old-2
//	└── new      <- newest, and a dead end
func lopsided() *Comment {
	return &Comment{ID: "root", Replies: []*Comment{
		{ID: "old", Replies: []*Comment{
			{ID: "old-1", Replies: []*Comment{{ID: "old-2"}}},
		}},
		{ID: "new"},
	}}
}

func TestNewestBranchMissesDeeperSiblings(t *testing.T) {
	tr := lopsided()
	want := VisibleByAllLevels(tr)
	got := VisibleByNewestBranch(tr)
	if reflect.DeepEqual(got, want) {
		t.Fatal("the newest-branch version agreed here; the episode says it does not")
	}
	t.Logf("following the newest reply down gives %v", got)
	t.Logf("the collapsed view actually shows      %v", want)
	if len(got) >= len(want) {
		t.Fatalf("expected the newest branch to run out early: %d against %d",
			len(got), len(want))
	}
}

// One entry per level, and the entry is the last comment at that level in
// reading order. Checked against the level-order walk rather than by eye.
func TestOneEntryPerLevelAndItIsTheLast(t *testing.T) {
	tr := thread(4, 3)
	levels := AllLevels(tr)
	got := VisibleByFirstArrival(tr)
	if len(got) != len(levels) {
		t.Fatalf("%d visible entries for %d levels", len(got), len(levels))
	}
	for i, level := range levels {
		if got[i] != level[len(level)-1] {
			t.Fatalf("level %d: visible is %q, last in the level is %q",
				i, got[i], level[len(level)-1])
		}
	}
	t.Logf("%d comments, %d levels, %d visible", count(tr), len(levels), len(got))
}

// What each version holds while it works. The answer is h entries in every
// case; the question is what has to exist alongside it.
func TestHeldWhileWorking(t *testing.T) {
	tr := thread(6, 4)
	levels := AllLevels(tr)
	held := 0
	for _, l := range levels {
		held += len(l)
	}
	t.Logf("%d comments in %d levels: building every level holds %d IDs, "+
		"the visible answer is %d", count(tr), len(levels), held, len(levels))
}

// The two preview versions return the same lines. What differs is how many
// previews had to be built to get them.
func TestPreviewsAgreeAndCountBuilds(t *testing.T) {
	for _, s := range shapes {
		a, builtA := PreviewsByDepthMap(s.root)
		b, builtB := PreviewsByFirstArrival(s.root)
		if !reflect.DeepEqual(a, b) {
			t.Fatalf("%s: the two preview versions disagree", s.name)
		}
		t.Logf("%-11s %6d comments, %2d visible  |  previews built: "+
			"last writer wins %6d, newest first %3d",
			s.name, count(s.root), len(a), builtA, builtB)
	}
}
