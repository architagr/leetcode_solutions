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

func writeFile(t *testing.T, root, rel, content string) {
	t.Helper()
	p := filepath.Join(root, rel)
	if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(p, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}

func readFile(t *testing.T, root, rel string) string {
	t.Helper()
	b, err := os.ReadFile(filepath.Join(root, rel))
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

func TestApplyRewritesEveryMarkdownFileInAQueuedFolder(t *testing.T) {
	root := testRepo(t)
	writeFile(t, root, "b/POST_DISCORD.md", "**Day 2/365**\n")
	writeFile(t, root, "b/INTUITION.md", "Same trick as Day 3's.\n")
	writeFile(t, root, "b/SOLUTION.md", "Unlike Day 3, this one recurses.\n")
	writeFile(t, root, "b/main.go", "// Day 2 must not be touched in code\n")
	writeFile(t, root, "c/POST_DISCORD.md", "**Day 3/365**\n")
	writeFile(t, root, "a/POST_DISCORD.md", "**Day 1/365**\n")
	queuePath := filepath.Join(root, "queue.yaml")

	q := baseQueue()
	if err := q.Save(queuePath); err != nil {
		t.Fatal(err)
	}
	// Swap days 2 and 3; day 1 is posted and stays.
	p := Plan{Entries: []Entry{
		{Day: 1, Number: 104, Title: "Max Depth", Folder: "a"},
		{Day: 2, Number: 257, Title: "Binary Tree Paths", Folder: "c"},
		{Day: 3, Number: 108, Title: "Sorted Array to BST", Folder: "b"},
	}}

	res, err := Apply(root, queuePath, p, "", true)
	if err != nil {
		t.Fatalf("dry run failed: %v", err)
	}
	if readFile(t, root, "b/POST_DISCORD.md") != "**Day 2/365**\n" {
		t.Fatal("dry run must not write files")
	}
	if len(res.FilesRewritten) == 0 {
		t.Fatal("dry run must still report which files would change")
	}

	if _, err := Apply(root, queuePath, p, "", false); err != nil {
		t.Fatalf("apply failed: %v", err)
	}
	if got := readFile(t, root, "b/POST_DISCORD.md"); got != "**Day 3/365**\n" {
		t.Fatalf("POST_DISCORD.md not renumbered: %q", got)
	}
	if got := readFile(t, root, "b/INTUITION.md"); got != "Same trick as Day 2's.\n" {
		t.Fatalf("INTUITION.md prose not renumbered: %q", got)
	}
	if got := readFile(t, root, "b/SOLUTION.md"); got != "Unlike Day 2, this one recurses.\n" {
		t.Fatalf("SOLUTION.md prose not renumbered: %q", got)
	}
	if got := readFile(t, root, "b/main.go"); got != "// Day 2 must not be touched in code\n" {
		t.Fatalf("non-markdown file was rewritten: %q", got)
	}
	if got := readFile(t, root, "a/POST_DISCORD.md"); got != "**Day 1/365**\n" {
		t.Fatalf("unmoved posted day was rewritten: %q", got)
	}

	reloaded, err := queue.Load(queuePath)
	if err != nil {
		t.Fatal(err)
	}
	byNum := map[int]int{}
	for _, e := range reloaded.Entries {
		byNum[e.Number] = e.Day
	}
	if byNum[108] != 3 || byNum[257] != 2 {
		t.Fatalf("queue not reordered: %v", byNum)
	}
}

func TestApplyInsertsNewEntriesAsPendingContent(t *testing.T) {
	root := testRepo(t, "a", "b", "c", "d")
	writeFile(t, root, "a/POST_DISCORD.md", "**Day 1/365**\n")
	queuePath := filepath.Join(root, "queue.yaml")
	q := baseQueue()
	if err := q.Save(queuePath); err != nil {
		t.Fatal(err)
	}
	p := Plan{Entries: []Entry{
		{Day: 1, Number: 104, Title: "Max Depth", Folder: "a"},
		{Day: 2, Number: 108, Title: "Sorted Array to BST", Folder: "b"},
		{Day: 3, Number: 257, Title: "Binary Tree Paths", Folder: "c"},
		{Day: 4, Number: 206, Title: "Reverse Linked List", Difficulty: "easy", Folder: "d", Batch: "linked list", BuildsOn: []int{104}},
	}}
	res, err := Apply(root, queuePath, p, "", false)
	if err != nil {
		t.Fatalf("apply failed: %v", err)
	}
	if len(res.Inserted) != 1 || res.Inserted[0].Number != 206 {
		t.Fatalf("expected one insertion of 206, got %+v", res.Inserted)
	}
	reloaded, _ := queue.Load(queuePath)
	var found *queue.Entry
	for i := range reloaded.Entries {
		if reloaded.Entries[i].Number == 206 {
			found = &reloaded.Entries[i]
		}
	}
	if found == nil {
		t.Fatal("206 was not inserted")
	}
	if found.Status != queue.StatusPendingContent {
		t.Fatalf("inserted entry must be pending_content, got %q", found.Status)
	}
	if found.Day != 4 || found.Batch != "linked list" || len(found.BuildsOn) != 1 {
		t.Fatalf("inserted entry is wrong: %+v", found)
	}
}

func TestApplyIsIdempotent(t *testing.T) {
	root := testRepo(t, "a", "b", "c")
	writeFile(t, root, "b/POST_DISCORD.md", "**Day 2/365**\n")
	writeFile(t, root, "c/POST_DISCORD.md", "**Day 3/365**\n")
	queuePath := filepath.Join(root, "queue.yaml")
	q := baseQueue()
	if err := q.Save(queuePath); err != nil {
		t.Fatal(err)
	}
	p := Plan{Entries: []Entry{
		{Day: 1, Number: 104, Folder: "a"},
		{Day: 2, Number: 257, Folder: "c"},
		{Day: 3, Number: 108, Folder: "b"},
	}}
	if _, err := Apply(root, queuePath, p, "", false); err != nil {
		t.Fatal(err)
	}
	first := readFile(t, root, "b/POST_DISCORD.md")
	res, err := Apply(root, queuePath, p, "", false)
	if err != nil {
		t.Fatal(err)
	}
	if len(res.Moved) != 0 {
		t.Fatalf("second run must move nothing, got %+v", res.Moved)
	}
	if second := readFile(t, root, "b/POST_DISCORD.md"); second != first {
		t.Fatalf("second run changed the file: %q then %q", first, second)
	}
}

func TestApplyRefusesAnInvalidPlanWithoutWriting(t *testing.T) {
	root := testRepo(t, "a", "b", "c")
	writeFile(t, root, "b/POST_DISCORD.md", "**Day 2/365**\n")
	queuePath := filepath.Join(root, "queue.yaml")
	q := baseQueue()
	if err := q.Save(queuePath); err != nil {
		t.Fatal(err)
	}
	// 104 is posted on day 1 and may not move.
	p := Plan{Entries: []Entry{
		{Day: 1, Number: 108, Folder: "b"},
		{Day: 2, Number: 104, Folder: "a"},
	}}
	if _, err := Apply(root, queuePath, p, "", false); err == nil {
		t.Fatal("expected apply to refuse the plan")
	}
	if got := readFile(t, root, "b/POST_DISCORD.md"); got != "**Day 2/365**\n" {
		t.Fatalf("a refused plan must write nothing, but the file changed: %q", got)
	}
}
