package renumber

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"leetcode_solutions/challenge/internal/queue"
)

func postedAt(stamp string) map[string]*string {
	if stamp == "" {
		return map[string]*string{queue.DestinationDiscord: nil}
	}
	return map[string]*string{queue.DestinationDiscord: &stamp}
}

// testRepo builds a throwaway repo with one folder per entry so
// validation has real directories to find.
func testRepo(t *testing.T, folders ...string) string {
	t.Helper()
	root := t.TempDir()
	for _, f := range folders {
		if err := os.MkdirAll(filepath.Join(root, f), 0o755); err != nil {
			t.Fatal(err)
		}
	}
	return root
}

func baseQueue() *queue.Queue {
	return &queue.Queue{NextDay: 4, Entries: []queue.Entry{
		{Day: 1, Number: 104, Title: "Max Depth", Folder: "a", Status: queue.StatusContentReady, PostedAt: postedAt("2026-09-03T06:16:03Z")},
		{Day: 2, Number: 108, Title: "Sorted Array to BST", Folder: "b", Status: queue.StatusContentReady, PostedAt: postedAt("")},
		{Day: 3, Number: 257, Title: "Binary Tree Paths", Folder: "c", Status: queue.StatusContentReady, PostedAt: postedAt("")},
	}}
}

func TestValidateRefusesMovingAPostedDay(t *testing.T) {
	root := testRepo(t, "a", "b", "c")
	p := Plan{Entries: []Entry{
		{Day: 1, Number: 108, Folder: "b"},
		{Day: 2, Number: 104, Folder: "a"}, // 104 is posted on day 1
		{Day: 3, Number: 257, Folder: "c"},
	}}
	err := Validate(baseQueue(), p, root)
	if err == nil || !strings.Contains(err.Error(), "104") {
		t.Fatalf("expected refusal naming question 104, got %v", err)
	}
}

func TestValidateRefusesForwardBuildsOn(t *testing.T) {
	root := testRepo(t, "a", "b", "c")
	p := Plan{Entries: []Entry{
		{Day: 1, Number: 104, Folder: "a"},
		{Day: 2, Number: 257, Folder: "c", BuildsOn: []int{108}}, // 108 lands on day 3
		{Day: 3, Number: 108, Folder: "b"},
	}}
	err := Validate(baseQueue(), p, root)
	if err == nil || !strings.Contains(err.Error(), "builds_on") {
		t.Fatalf("expected a builds_on refusal, got %v", err)
	}
}

func TestValidateRefusesMissingFolder(t *testing.T) {
	root := testRepo(t, "a", "b") // "c" deliberately absent
	p := Plan{Entries: []Entry{
		{Day: 1, Number: 104, Folder: "a"},
		{Day: 2, Number: 108, Folder: "b"},
		{Day: 3, Number: 257, Folder: "c"},
	}}
	err := Validate(baseQueue(), p, root)
	if err == nil || !strings.Contains(err.Error(), "c") {
		t.Fatalf("expected a missing-folder refusal naming c, got %v", err)
	}
}

func TestValidateRefusesDuplicateDayOrNumber(t *testing.T) {
	root := testRepo(t, "a", "b", "c")
	dupDay := Plan{Entries: []Entry{
		{Day: 1, Number: 104, Folder: "a"},
		{Day: 1, Number: 108, Folder: "b"},
	}}
	if err := Validate(baseQueue(), dupDay, root); err == nil {
		t.Fatal("expected a duplicate-day refusal")
	}
	dupNum := Plan{Entries: []Entry{
		{Day: 1, Number: 104, Folder: "a"},
		{Day: 2, Number: 104, Folder: "a"},
	}}
	if err := Validate(baseQueue(), dupNum, root); err == nil {
		t.Fatal("expected a duplicate-number refusal")
	}
}

func TestValidateAcceptsAPlanThatKeepsPostedDaysPut(t *testing.T) {
	root := testRepo(t, "a", "b", "c")
	p := Plan{Entries: []Entry{
		{Day: 1, Number: 104, Folder: "a"},
		{Day: 2, Number: 257, Folder: "c"},
		{Day: 3, Number: 108, Folder: "b", BuildsOn: []int{104}},
	}}
	if err := Validate(baseQueue(), p, root); err != nil {
		t.Fatalf("expected the plan to validate, got %v", err)
	}
}

func TestDayMapSkipsUnmovedDays(t *testing.T) {
	p := Plan{Entries: []Entry{
		{Day: 1, Number: 104},
		{Day: 5, Number: 108},
	}}
	m := DayMap(baseQueue(), p)
	if _, ok := m[1]; ok {
		t.Fatal("day 1 does not move; it must be absent from the map")
	}
	if m[2] != 5 {
		t.Fatalf("expected old day 2 to map to 5, got %d", m[2])
	}
}
