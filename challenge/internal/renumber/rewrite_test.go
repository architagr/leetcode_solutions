package renumber

import "testing"

func TestRewriteDayRefsIsSimultaneous(t *testing.T) {
	// 20→30 and 30→45 in one pass. A sequential rewrite would turn the
	// original "Day 20" into "Day 45" by rescanning its own output.
	got := RewriteDayRefs("Day 20/365 builds on Day 30.", map[int]int{20: 30, 30: 45})
	want := "Day 30/365 builds on Day 45."
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestRewriteDayRefsCoversEveryForm(t *testing.T) {
	in := "**365 Days of LeetCode Challenge — Day 12/365**\n" +
		"![Day 12](HERO.png)\n" +
		"- [Day 14: Average of Levels](https://example.com/x/) — why\n" +
		"the same lazy growth as Day 14's did\n"
	want := "**365 Days of LeetCode Challenge — Day 40/365**\n" +
		"![Day 40](HERO.png)\n" +
		"- [Day 27: Average of Levels](https://example.com/x/) — why\n" +
		"the same lazy growth as Day 27's did\n"
	if got := RewriteDayRefs(in, map[int]int{12: 40, 14: 27}); got != want {
		t.Fatalf("got:\n%s\nwant:\n%s", got, want)
	}
}

func TestRewriteDayRefsLeavesUnmappedDaysAlone(t *testing.T) {
	got := RewriteDayRefs("Day 8 taught preorder; Day 12 uses it.", map[int]int{12: 40})
	want := "Day 8 taught preorder; Day 40 uses it."
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestRewriteDayRefsIgnoresNonDayNumbers(t *testing.T) {
	// "365" in the denominator and a bare year must never be touched.
	got := RewriteDayRefs("Day 12/365, written in 2026, 12 nodes deep.", map[int]int{12: 40, 365: 9, 2026: 1})
	want := "Day 40/365, written in 2026, 12 nodes deep."
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}
