package queue

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestMarkPostedSetsOnlyThatDestination(t *testing.T) {
	q := &Queue{Entries: []Entry{{Number: 104, Day: 1}, {Number: 108, Day: 2}}}
	at := time.Date(2026, 9, 3, 9, 0, 0, 0, time.UTC)

	if !q.MarkPosted(104, DestinationDiscord, at) {
		t.Fatal("MarkPosted(104) = false, want true")
	}
	got := q.Entries[0].PostedAt[DestinationDiscord]
	if got == nil || *got != "2026-09-03T09:00:00Z" {
		t.Errorf("PostedAt[discord] = %v, want 2026-09-03T09:00:00Z", got)
	}
	if q.Entries[0].PostedAt[DestinationLinkedInMain] != nil {
		t.Error("marking discord must not touch other destinations")
	}
	if len(q.Entries[1].PostedAt) != 0 {
		t.Errorf("other entries must be untouched, got %v", q.Entries[1].PostedAt)
	}
	if q.MarkPosted(999, DestinationDiscord, at) {
		t.Error("MarkPosted for an unknown number = true, want false")
	}
}

func TestIsPosted(t *testing.T) {
	ts := "2026-09-03T09:00:00Z"
	e := Entry{PostedAt: map[string]*string{
		DestinationDiscord:      &ts,
		DestinationLinkedInMain: nil,
	}}
	if !e.IsPosted(DestinationDiscord) {
		t.Error("IsPosted(discord) = false, want true")
	}
	if e.IsPosted(DestinationLinkedInMain) {
		t.Error("IsPosted(linkedin_main_account) = true (explicit null), want false")
	}
	if e.IsPosted(DestinationLinkedInGroup) {
		t.Error("IsPosted for an absent key = true, want false")
	}
	if (Entry{}).IsPosted(DestinationDiscord) {
		t.Error("IsPosted on a nil map = true, want false")
	}
}

func TestLoadNormalizesLegacyPostedAt(t *testing.T) {
	path := filepath.Join(t.TempDir(), "queue.yaml")
	legacy := `next_day: 3
entries:
    - day: 1
      number: 104
      title: Maximum Depth of Binary Tree
      difficulty: easy
      folder: easy_problems/101_200/maximum_depth_of_binary_tree
      batch: binary tree
      status: content_ready
      posted_at: null
    - day: 2
      number: 108
      title: Convert Sorted Array to Binary Search Tree
      difficulty: easy
      folder: easy_problems/101_200/convert_sorted_array_to_binary_search_tree
      batch: binary tree
      status: content_ready
      posted_at: 2026-09-01T10:00:00Z
`
	if err := os.WriteFile(path, []byte(legacy), 0o644); err != nil {
		t.Fatal(err)
	}

	q, err := Load(path)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if q.Entries[0].IsPosted(DestinationDiscord) {
		t.Error("legacy posted_at: null should mean not posted anywhere")
	}
	// A legacy bare timestamp predates per-destination tracking; it can only
	// have meant the one destination that existed, so it maps to discord.
	if !q.Entries[1].IsPosted(DestinationDiscord) {
		t.Error("legacy bare timestamp should normalize to a discord timestamp")
	}
}

func TestSaveLoadRoundTripPreservesDestinations(t *testing.T) {
	path := filepath.Join(t.TempDir(), "queue.yaml")
	q := &Queue{NextDay: 1}
	q.Append(Entry{Number: 104, Title: "Maximum Depth of Binary Tree", Difficulty: "easy", Folder: "f", Batch: "binary tree"})
	q.MarkPosted(104, DestinationDiscord, time.Date(2026, 9, 3, 9, 0, 0, 0, time.UTC))
	if err := q.Save(path); err != nil {
		t.Fatalf("Save: %v", err)
	}

	loaded, err := Load(path)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if !loaded.Entries[0].IsPosted(DestinationDiscord) {
		t.Error("discord timestamp lost in round trip")
	}
	if loaded.Entries[0].IsPosted(DestinationLinkedInMain) {
		t.Error("linkedin must still be unposted after round trip")
	}
}

func TestLoadMissingFileReturnsEmptyQueue(t *testing.T) {
	q, err := Load(filepath.Join(t.TempDir(), "queue.yaml"))
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if q.NextDay != 1 {
		t.Errorf("NextDay = %d, want 1", q.NextDay)
	}
	if len(q.Entries) != 0 {
		t.Errorf("Entries = %v, want empty", q.Entries)
	}
}

func TestAppendAssignsDayAndIncrements(t *testing.T) {
	q := &Queue{NextDay: 23}
	day := q.Append(Entry{
		Number:     965,
		Title:      "Univalued Binary Tree",
		Difficulty: "easy",
		Folder:     "easy_problems/901_1000/univalued_binary_tree",
		Batch:      "binary tree easy problems",
	})
	if day != 23 {
		t.Errorf("Append returned day %d, want 23", day)
	}
	if q.NextDay != 24 {
		t.Errorf("NextDay after append = %d, want 24", q.NextDay)
	}
	if len(q.Entries) != 1 || q.Entries[0].Status != StatusContentReady {
		t.Fatalf("Entries = %+v", q.Entries)
	}
}

func TestHas(t *testing.T) {
	q := &Queue{Entries: []Entry{{Number: 965}}}
	if !q.Has(965) {
		t.Error("Has(965) = false, want true")
	}
	if q.Has(1) {
		t.Error("Has(1) = true, want false")
	}
}

func TestSaveLoadRoundTrip(t *testing.T) {
	path := filepath.Join(t.TempDir(), "queue.yaml")
	q := &Queue{NextDay: 2}
	q.Append(Entry{Number: 1, Title: "Two Sum", Difficulty: "easy", Folder: "easy_problems/1_100/two_sum", Batch: "arrays"})
	if err := q.Save(path); err != nil {
		t.Fatalf("Save: %v", err)
	}
	if _, err := os.Stat(path); err != nil {
		t.Fatalf("file not written: %v", err)
	}
	loaded, err := Load(path)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if loaded.NextDay != 3 || len(loaded.Entries) != 1 || loaded.Entries[0].Number != 1 {
		t.Fatalf("loaded = %+v", loaded)
	}
}

func TestNextUnpostedReturnsOldestFirstUpToLimit(t *testing.T) {
	posted := "2026-09-01T09:00:00Z"
	q := &Queue{Entries: []Entry{
		{Day: 3, Number: 257, Status: StatusContentReady},
		{Day: 1, Number: 104, Status: StatusContentReady, PostedAt: map[string]*string{DestinationLinkedInMain: &posted}},
		{Day: 2, Number: 108, Status: StatusContentReady},
		{Day: 4, Number: 404, Status: StatusContentReady},
	}}

	got := q.NextUnposted(DestinationLinkedInMain, 2)
	if len(got) != 2 {
		t.Fatalf("got %d entries, want 2", len(got))
	}
	if got[0].Day != 2 || got[1].Day != 3 {
		t.Errorf("got days %d,%d — want 2,3 (oldest unposted first, day 1 already posted)", got[0].Day, got[1].Day)
	}

	if all := q.NextUnposted(DestinationLinkedInMain, 99); len(all) != 3 {
		t.Errorf("a limit past the end should return all 3 unposted, got %d", len(all))
	}
	if none := q.NextUnposted(DestinationLinkedInMain, 0); len(none) != 0 {
		t.Errorf("limit 0 should return nothing, got %d", len(none))
	}
}

func TestMarkContentReadyFlipsOnlyThatEntry(t *testing.T) {
	q := &Queue{Entries: []Entry{
		{Number: 114, Day: 10, Status: "pending_content"},
		{Number: 145, Day: 9, Status: "pending_content"},
	}}

	if !q.MarkContentReady(114) {
		t.Fatal("MarkContentReady(114) = false, want true")
	}
	if q.Entries[0].Status != StatusContentReady {
		t.Errorf("entry 114 status = %q, want %q", q.Entries[0].Status, StatusContentReady)
	}
	if q.Entries[1].Status != "pending_content" {
		t.Errorf("entry 145 status = %q, want it untouched", q.Entries[1].Status)
	}
	if q.MarkContentReady(999) {
		t.Error("MarkContentReady on an absent number = true, want false")
	}
}

// A posted entry has already gone out; nothing about writing content
// later should rewind that.
func TestMarkContentReadyLeavesPostedEntriesAlone(t *testing.T) {
	q := &Queue{Entries: []Entry{{Number: 104, Day: 1, Status: StatusPosted}}}
	if !q.MarkContentReady(104) {
		t.Fatal("MarkContentReady on a posted entry = false, want true")
	}
	if q.Entries[0].Status != StatusPosted {
		t.Errorf("status = %q, want it left as %q", q.Entries[0].Status, StatusPosted)
	}
}

// Running the same flip twice is how a re-run of a partly-finished batch
// behaves, so it has to be harmless.
func TestMarkContentReadyIsIdempotent(t *testing.T) {
	q := &Queue{Entries: []Entry{{Number: 114, Day: 10, Status: "pending_content"}}}
	q.MarkContentReady(114)
	q.MarkContentReady(114)
	if q.Entries[0].Status != StatusContentReady {
		t.Errorf("status = %q, want %q", q.Entries[0].Status, StatusContentReady)
	}
}

// builds_on has to round-trip through a save/load cycle. It did not
// before this field existed: Unmarshal dropped the key and Marshal wrote
// the file back without it, so every queue write deleted the chain.
func TestBuildsOnSurvivesASaveLoadCycle(t *testing.T) {
	path := filepath.Join(t.TempDir(), "queue.yaml")
	q := &Queue{NextDay: 3, Entries: []Entry{
		{Day: 1, Number: 104, Status: StatusContentReady},
		{Day: 2, Number: 222, Status: StatusContentReady, BuildsOn: []int{104, 110}},
	}}
	if err := q.Save(path); err != nil {
		t.Fatal(err)
	}

	reloaded, err := Load(path)
	if err != nil {
		t.Fatal(err)
	}
	got := reloaded.Entries[1].BuildsOn
	if len(got) != 2 || got[0] != 104 || got[1] != 110 {
		t.Errorf("BuildsOn = %v, want [104 110]", got)
	}
	// An empty chain is a real answer for a question that introduces
	// something new, and should not litter the file with empty keys.
	if reloaded.Entries[0].BuildsOn != nil {
		t.Errorf("BuildsOn = %v, want nil for an entry without a chain", reloaded.Entries[0].BuildsOn)
	}
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Count(string(raw), "builds_on") != 1 {
		t.Errorf("builds_on should appear once, for the entry that has one:\n%s", raw)
	}
}

func TestPendingContentEntryIsNotSelectedForPosting(t *testing.T) {
	q := &Queue{NextDay: 3, Entries: []Entry{
		{Day: 1, Number: 104, Status: StatusContentReady, PostedAt: map[string]*string{}},
		{Day: 2, Number: 206, Status: StatusPendingContent, PostedAt: map[string]*string{}},
	}}
	got := q.NextUnposted(DestinationDiscord, 10)
	if len(got) != 1 || got[0].Number != 104 {
		t.Fatalf("pending_content entry must not be selected; got %+v", got)
	}
}

func TestReorderAssignsDaysAndKeepsNextDayAhead(t *testing.T) {
	q := &Queue{NextDay: 3, Entries: []Entry{
		{Day: 1, Number: 104, Status: StatusContentReady, PostedAt: map[string]*string{}},
		{Day: 2, Number: 108, Status: StatusContentReady, PostedAt: map[string]*string{}},
	}}
	q.Reorder(map[int]int{104: 2, 108: 1})
	byNum := map[int]int{}
	for _, e := range q.Entries {
		byNum[e.Number] = e.Day
	}
	if byNum[104] != 2 || byNum[108] != 1 {
		t.Fatalf("days not reassigned: %v", byNum)
	}
	if q.Entries[0].Day != 1 {
		t.Fatalf("entries must be sorted by day after reorder, got day %d first", q.Entries[0].Day)
	}
	if q.NextDay != 3 {
		t.Fatalf("NextDay must stay one past the highest day, got %d", q.NextDay)
	}
}
