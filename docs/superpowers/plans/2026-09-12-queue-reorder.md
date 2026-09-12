# Queue Reorder Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Reshape `challenge/queue.yaml` from 40 consecutive binary-tree days into eight braided topic arcs spanning days 12–81, rewriting every day reference in published-facing files and re-rendering every affected hero image.

**Architecture:** A new `challenge/internal/renumber` package owns the whole operation: it builds an old-day → new-day map keyed by LeetCode number, validates the target ordering against frozen days and `builds_on` chains, rewrites day references across every `*.md` in every queued folder in a single non-chaining pass, re-renders heroes, and writes `queue.yaml` last. A thin `cli.QueueRenumber` wrapper and a `queue-renumber` subcommand expose it. Content for the 30 newly-inserted days is generated afterwards, one arc at a time, by the existing `/leetcode-content` skill.

**Tech Stack:** Go 1.x, `gopkg.in/yaml.v3`, Playwright/Chromium (hero screenshots), existing `challenge/internal/{queue,hero,cli}` packages.

**Spec:** `docs/superpowers/specs/2026-09-12-queue-reorder-design.md`

---

## Findings that amend the spec

The spec was written before the repo was scanned for day references. Three things it got wrong, all confirmed by measurement:

1. **Day references are not confined to five `POST_*.md` files.** There are **435** `Day N` references across queued folders, in seven file types — `POST_LINKEDIN_ARTICLE.md`, `POST_SUBSTACK.md`, `POST_DISCORD.md`, `POST_LINKEDIN.md`, `POST_X.md`, **`INTUITION.md`** and **`SOLUTION.md`**. The last two are not in the spec's list. The rewrite must cover every `*.md` in a queued folder.

2. **78 of those references are free-form prose**, not structured headers or links — "the same one from Day 14", "the opposite trade to Day 15's". They renumber correctly under a mechanical substitution, but only because substitution is simultaneous (see Task 5).

3. **Two references break in a way renumbering cannot fix**, because the reorder turns them into forward references. Both are repaired by hand in Task 9.

One plan amendment was already made as a result: days 49 and 50 were swapped (236 to day 49, 783 to day 50) so that old day 27's "tomorrow: the comparisons use `Val`, not pointer" still lands on 236 LCA-of-a-Binary-Tree. All three `tomorrow` narratives in the repo survive the reorder with this swap.

## File structure

| File | Responsibility |
|---|---|
| `challenge/internal/renumber/renumber.go` | Plan/Change/Result types, day-map construction, validation, apply |
| `challenge/internal/renumber/rewrite.go` | Day-reference rewriting in file content (pure string functions) |
| `challenge/internal/renumber/renumber_test.go` | Validation + apply tests against a temp repo |
| `challenge/internal/renumber/rewrite_test.go` | Rewrite tests (pure, no filesystem) |
| `challenge/internal/queue/queue.go` | Add `StatusPendingContent` constant + `Reorder` method |
| `challenge/internal/cli/cli.go` | Add `QueueRenumber` wrapper |
| `challenge/cmd/leetcodectl/main.go` | Add `queue-renumber` dispatch case |
| `challenge/plans/2026-09-12-reorder.json` | The concrete day mapping this plan applies |

---

## Task 0: Branch and baseline

**Files:** none

- [ ] **Step 1: Cut the feature branch**

```bash
cd /Users/architagarwal/code/leetcode_solutions
git checkout -b queue-reorder-braided-arcs
```

- [ ] **Step 2: Confirm a green baseline before changing anything**

Run: `go test ./challenge/...`
Expected: every package `ok` or `[no test files]`. If anything fails, stop — it failed before this work and must be fixed or understood first.

- [ ] **Step 3: Build the CLI**

Run: `go build -o /tmp/leetcodectl ./challenge/cmd/leetcodectl`
Expected: no output, exit 0.

- [ ] **Step 4: Confirm Chromium is installed for hero rendering**

Run: `npx playwright install chromium`
Expected: either downloads, or reports it is already installed. Task 8 re-renders 40 images and will fail without it.

---

## Task 1: Fix the wrong entry in `number_folder_map.yaml`

`number_folder_map.yaml` records `150: medium_problems/701_800/daily_temperature`. Question 150 is Evaluate Reverse Polish Notation; Daily Temperatures is 739. Both questions are scheduled by this plan (days 62 and 61), so content generation would resolve one of them to the wrong folder.

**Files:**
- Modify: `challenge/number_folder_map.yaml`

- [ ] **Step 1: Confirm the defect**

Run: `grep -n "^    150:\|^    739:" challenge/number_folder_map.yaml`
Expected: shows `150: medium_problems/701_800/daily_temperature` and no 739 entry.

- [ ] **Step 2: Correct the mapping**

Replace the `150:` line and add the missing `739:` line, keeping numeric order:

```yaml
    150: medium_problems/101_200/evaluate_reverse_polish_notation
    739: medium_problems/701_800/daily_temperatures
```

- [ ] **Step 3: Verify both folders exist**

Run:
```bash
ls -d medium_problems/101_200/evaluate_reverse_polish_notation medium_problems/701_800/daily_temperatures
```
Expected: both paths listed, no "No such file".

- [ ] **Step 4: Commit**

```bash
git add challenge/number_folder_map.yaml
git commit -m "fix(challenge): map 150 to Evaluate RPN, not Daily Temperatures

The cache recorded 150 against medium_problems/701_800/daily_temperature.
150 is Evaluate Reverse Polish Notation; Daily Temperatures is 739. Both
are scheduled in the reorder, so resolving either would have picked the
wrong folder."
```

---

## Task 2: Remove the three duplicate folders this plan schedules

Seven solution folders exist in duplicate. Three hold questions this plan schedules, and a duplicate means content generation can resolve to the copy without the content.

**Files:**
- Delete: `medium_problems/701_800/daily_temperature` (keep `daily_temperatures`)
- Delete: `medium_problems/101_200/number_of_islands` (keep `201_300`, matching the repo's number-range convention for 200)
- Delete: `easy_problems/701_800/search_in_a_binary_search_tree` (keep `601_700`, which holds the generated content)

- [ ] **Step 1: Confirm which copy to keep in each pair**

Run:
```bash
for d in medium_problems/701_800/daily_temperature medium_problems/701_800/daily_temperatures \
         medium_problems/101_200/number_of_islands medium_problems/201_300/number_of_islands \
         easy_problems/601_700/search_in_a_binary_search_tree easy_problems/701_800/search_in_a_binary_search_tree; do
  echo "$d -> $(ls "$d" | tr '\n' ' ')"
done
```
Expected: `easy_problems/601_700/search_in_a_binary_search_tree` is the only one of the six holding `HERO.png` and `POST_*.md`. Keep that one.

- [ ] **Step 2: Confirm the queue does not point at a copy being deleted**

Run: `grep -n "daily_temperature$\|101_200/number_of_islands\|701_800/search_in_a_binary" challenge/queue.yaml`
Expected: no output. If any line matches, stop — the queue references a folder about to be deleted, and that entry must be repointed first.

- [ ] **Step 3: Delete the three stale copies**

```bash
git rm -r medium_problems/701_800/daily_temperature \
          medium_problems/101_200/number_of_islands \
          easy_problems/701_800/search_in_a_binary_search_tree
```

- [ ] **Step 4: Confirm the packages still build and test**

Run: `go build ./... && go test ./...`
Expected: pass. A deleted folder held Go code, so a stale import would surface here.

- [ ] **Step 5: Commit**

```bash
git commit -m "chore: drop three duplicate solution folders

Each of these questions had two folders holding real code. 739 Daily
Temperatures, 200 Number of Islands and 700 Search in a BST are all
scheduled in the reorder, and a duplicate lets resolution pick the copy
without the generated content - which for 700 is exactly the copy that
has it.

The remaining four duplicate pairs are not scheduled and are left alone."
```

---

## Task 3: `StatusPendingContent` and `Queue.Reorder`

The 30 newly-inserted days have no content yet. `Entry.HasContent()` already gates posting on `content_ready`/`posted`, so a pending entry is skipped rather than posted empty — but there is no constant naming that state, even though `MarkContentReady`'s doc comment refers to "pending_content".

**Files:**
- Modify: `challenge/internal/queue/queue.go`
- Test: `challenge/internal/queue/queue_test.go`

- [ ] **Step 1: Write the failing test**

Append to `challenge/internal/queue/queue_test.go`:

```go
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
```

- [ ] **Step 2: Run the tests to verify they fail**

Run: `go test ./challenge/internal/queue/ -run 'PendingContent|Reorder' -v`
Expected: FAIL — `undefined: StatusPendingContent` and `q.Reorder undefined`.

- [ ] **Step 3: Add the constant**

In `challenge/internal/queue/queue.go`, extend the status block:

```go
const (
	StatusContentReady = "content_ready"
	StatusPosted       = "posted"

	// StatusPendingContent is a day whose number is reserved but whose
	// content has not been written. HasContent reports false for it, so
	// posting skips it rather than sending an empty day. Reserving the
	// number up front is what lets the whole queue's topic and
	// difficulty shape be decided before any of it is written.
	StatusPendingContent = "pending_content"
)
```

- [ ] **Step 4: Add `Reorder`**

Append to `challenge/internal/queue/queue.go`:

```go
// Reorder reassigns day numbers from a question-number → day map and
// re-sorts the entries so file order matches day order. Entries whose
// number is absent from the map keep the day they have.
//
// It does no validation: callers reaching here have already decided the
// mapping is legal. The renumber package is where frozen days, forward
// builds_on references and missing folders are refused, because those
// checks need the filesystem and the whole intended ordering, not just
// the queue.
func (q *Queue) Reorder(dayByNumber map[int]int) {
	for i := range q.Entries {
		if day, ok := dayByNumber[q.Entries[i].Number]; ok {
			q.Entries[i].Day = day
		}
	}
	sort.Slice(q.Entries, func(i, j int) bool { return q.Entries[i].Day < q.Entries[j].Day })
	highest := 0
	for _, e := range q.Entries {
		if e.Day > highest {
			highest = e.Day
		}
	}
	if q.NextDay <= highest {
		q.NextDay = highest + 1
	}
}
```

- [ ] **Step 5: Run the tests to verify they pass**

Run: `go test ./challenge/internal/queue/ -run 'PendingContent|Reorder' -v`
Expected: both PASS.

- [ ] **Step 6: Run the full package suite**

Run: `go test ./challenge/internal/queue/`
Expected: `ok`.

- [ ] **Step 7: Commit**

```bash
git add challenge/internal/queue/
git commit -m "feat(queue): add pending_content status and Reorder

Reserving a day number before its content exists is how the queue's
topic and difficulty shape gets decided up front, and HasContent already
gated posting on it - but nothing named the state. Reorder reassigns
days from a number map and re-sorts, leaving validation to the caller
that can see the filesystem."
```

---

## Task 4: Day-reference rewriting

The core string operation. It must be **simultaneous**: if day 20 becomes 30 and day 30 becomes 45, rewriting sequentially would turn the first into 45. Go's `regexp.ReplaceAllStringFunc` scans the input once and appends to a new buffer, so replacement output is never rescanned — that is what makes this safe, and the test below is what holds it safe.

**Files:**
- Create: `challenge/internal/renumber/rewrite.go`
- Test: `challenge/internal/renumber/rewrite_test.go`

- [ ] **Step 1: Write the failing test**

Create `challenge/internal/renumber/rewrite_test.go`:

```go
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
```

- [ ] **Step 2: Run the test to verify it fails**

Run: `go test ./challenge/internal/renumber/ -v`
Expected: FAIL — the package does not exist yet.

- [ ] **Step 3: Write the implementation**

Create `challenge/internal/renumber/rewrite.go`:

```go
// Package renumber moves days in the challenge queue and repairs every
// published-facing artifact that carries a day number.
//
// The day number is not private to queue.yaml. It is printed on each
// hero image, written into the header of five post files, and cited in
// prose and cross-reference links throughout a folder's write-ups — 435
// references across the queue as this package was written, in seven
// different file types. Moving a day is therefore a repo-wide edit, and
// doing it by hand is how references go stale.
package renumber

import (
	"fmt"
	"regexp"
	"strconv"
)

// dayRefRe matches every form a day reference takes: the post header
// "Day 12/365", the hero alt text "![Day 12](HERO.png)", a
// cross-reference link "[Day 14: Title](url)", and bare prose such as
// "the same lazy growth as Day 14's".
//
// All four are the same two tokens, so one pattern covers them and
// there is no form-specific branch to forget to update.
var dayRefRe = regexp.MustCompile(`Day (\d+)`)

// RewriteDayRefs replaces every day reference in content according to
// dayByOldDay, leaving references to unmapped days untouched.
//
// Replacement is simultaneous, which is the whole correctness argument:
// ReplaceAllStringFunc scans the input once and appends each replacement
// to a separate buffer, so output is never rescanned. A naive loop of
// per-day string replacements would chain — mapping 20→30 and 30→45
// would carry the original day 20 all the way to 45.
func RewriteDayRefs(content string, dayByOldDay map[int]int) string {
	return dayRefRe.ReplaceAllStringFunc(content, func(match string) string {
		old, err := strconv.Atoi(match[len("Day "):])
		if err != nil {
			return match
		}
		next, ok := dayByOldDay[old]
		if !ok {
			return match
		}
		return fmt.Sprintf("Day %d", next)
	})
}
```

- [ ] **Step 4: Run the test to verify it passes**

Run: `go test ./challenge/internal/renumber/ -v`
Expected: all four tests PASS.

- [ ] **Step 5: Commit**

```bash
git add challenge/internal/renumber/
git commit -m "feat(renumber): rewrite day references in one simultaneous pass

The day number is copied into post headers, hero alt text,
cross-reference links and free prose - 435 references across the queue,
in seven file types. One regexp covers all four forms.

Replacement goes through ReplaceAllStringFunc rather than a loop of
string replacements so that output is never rescanned: mapping 20 to 30
and 30 to 45 in the same pass must not carry day 20 through to 45."
```

---

## Task 5: Plan types, day map, and validation

**Files:**
- Create: `challenge/internal/renumber/renumber.go`
- Test: `challenge/internal/renumber/renumber_test.go`

- [ ] **Step 1: Write the failing test**

Create `challenge/internal/renumber/renumber_test.go`:

```go
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
```

- [ ] **Step 2: Run the test to verify it fails**

Run: `go test ./challenge/internal/renumber/ -run Validate -v`
Expected: FAIL — `undefined: Plan`, `undefined: Validate`, `undefined: DayMap`.

- [ ] **Step 3: Write the implementation**

Create `challenge/internal/renumber/renumber.go`:

```go
package renumber

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"leetcode_solutions/challenge/internal/queue"
)

// Entry is one question's intended position. Entries already in the
// queue are matched by Number; entries absent from it are inserted, so
// one plan both reorders and inserts and the queue is never left half
// reordered.
type Entry struct {
	Day        int    `json:"day"`
	Number     int    `json:"number"`
	Title      string `json:"title"`
	Difficulty string `json:"difficulty"`
	Folder     string `json:"folder"`
	Batch      string `json:"batch"`
	BuildsOn   []int  `json:"buildsOn,omitempty"`
}

// Plan is the full intended ordering. It is complete rather than a
// diff: a partial plan cannot be checked for duplicate days, and the
// whole point of validating up front is that nothing is written until
// the entire target state is known to be legal.
type Plan struct {
	Entries []Entry `json:"entries"`
}

// DayMap returns old day → new day for every entry that actually moves,
// keyed through the LeetCode number, which is the only identity that
// survives a renumber. Days that do not move are absent, so callers can
// treat "in the map" as "needs rewriting".
func DayMap(q *queue.Queue, p Plan) map[int]int {
	oldDayByNumber := make(map[int]int, len(q.Entries))
	for _, e := range q.Entries {
		oldDayByNumber[e.Number] = e.Day
	}
	m := make(map[int]int)
	for _, e := range p.Entries {
		old, ok := oldDayByNumber[e.Number]
		if !ok || old == e.Day {
			continue
		}
		m[old] = e.Day
	}
	return m
}

// Validate reports the first reason the plan must not be applied.
//
// Every check runs before anything is written. A renumber that fails
// halfway leaves post files, hero images and queue.yaml disagreeing
// about what day a question is on, and there is no operation that puts
// that back.
func Validate(q *queue.Queue, p Plan, repoRoot string) error {
	seenDay := map[int]int{}
	seenNumber := map[int]bool{}
	for _, e := range p.Entries {
		if prev, ok := seenDay[e.Day]; ok {
			return fmt.Errorf("day %d is assigned twice, to questions %d and %d", e.Day, prev, e.Number)
		}
		seenDay[e.Day] = e.Number
		if seenNumber[e.Number] {
			return fmt.Errorf("question %d appears twice in the plan", e.Number)
		}
		seenNumber[e.Number] = true
	}

	// A posted day number is public on LinkedIn and Discord. Renumber
	// around it, never through it.
	for _, e := range q.Entries {
		posted := false
		for _, d := range queue.KnownDestinations {
			if e.IsPosted(d) {
				posted = true
				break
			}
		}
		if !posted {
			continue
		}
		for _, pe := range p.Entries {
			if pe.Number == e.Number && pe.Day != e.Day {
				return fmt.Errorf("question %d is posted on day %d and cannot move to day %d", e.Number, e.Day, pe.Day)
			}
		}
	}

	dayByNumber := make(map[int]int, len(p.Entries))
	for _, e := range p.Entries {
		dayByNumber[e.Number] = e.Day
	}
	// Entries outside the plan keep the day they already have, and a
	// plan entry may legitimately build on one of them.
	for _, e := range q.Entries {
		if _, ok := dayByNumber[e.Number]; !ok {
			dayByNumber[e.Number] = e.Day
		}
	}
	for _, e := range p.Entries {
		for _, dep := range e.BuildsOn {
			depDay, ok := dayByNumber[dep]
			if !ok {
				return fmt.Errorf("question %d on day %d has builds_on %d, which is not in the queue", e.Number, e.Day, dep)
			}
			if depDay >= e.Day {
				return fmt.Errorf("question %d on day %d has builds_on %d, which lands on day %d — a forward reference", e.Number, e.Day, dep, depDay)
			}
		}
	}

	for _, e := range p.Entries {
		if e.Folder == "" {
			return fmt.Errorf("question %d on day %d has no folder", e.Number, e.Day)
		}
		info, err := os.Stat(filepath.Join(repoRoot, e.Folder))
		if err != nil || !info.IsDir() {
			return fmt.Errorf("question %d on day %d: folder %s does not exist", e.Number, e.Day, e.Folder)
		}
	}
	return nil
}

// sortedEntries returns the plan's entries in day order.
func sortedEntries(p Plan) []Entry {
	out := append([]Entry(nil), p.Entries...)
	sort.Slice(out, func(i, j int) bool { return out[i].Day < out[j].Day })
	return out
}
```

- [ ] **Step 4: Run the test to verify it passes**

Run: `go test ./challenge/internal/renumber/ -v`
Expected: all tests PASS.

- [ ] **Step 5: Commit**

```bash
git add challenge/internal/renumber/
git commit -m "feat(renumber): plan types, day map and validation

Validation runs entirely before anything is written. A renumber that
fails halfway leaves post files, hero images and queue.yaml disagreeing
about what day a question is on, and nothing puts that back.

Four refusals: moving a day that is already posted somewhere, a
builds_on that would point forward, a folder that does not exist, and a
day or question assigned twice."
```

---

## Task 6: Apply — rewrite files, re-render heroes, write the queue

Order matters: file work first, `queue.yaml` last. A failure partway through then leaves the queue describing the state the files were in before the run, rather than a state that never existed.

**Files:**
- Modify: `challenge/internal/renumber/renumber.go`
- Test: `challenge/internal/renumber/renumber_test.go`

- [ ] **Step 1: Write the failing test**

Append to `challenge/internal/renumber/renumber_test.go`:

```go
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
```

- [ ] **Step 2: Run the tests to verify they fail**

Run: `go test ./challenge/internal/renumber/ -run Apply -v`
Expected: FAIL — `undefined: Apply`.

- [ ] **Step 3: Write the implementation**

Append to `challenge/internal/renumber/renumber.go`:

```go
// Change is one question's move, for the report.
type Change struct {
	Number int    `json:"number"`
	Title  string `json:"title"`
	Folder string `json:"folder"`
	OldDay int    `json:"oldDay,omitempty"`
	NewDay int    `json:"newDay"`
}

// Result reports what a run did, or what a dry run would have done.
type Result struct {
	DryRun         bool     `json:"dryRun"`
	Moved          []Change `json:"moved"`
	Inserted       []Change `json:"inserted"`
	FilesRewritten []string `json:"filesRewritten"`
	HeroesRendered []string `json:"heroesRendered"`
	Warnings       []string `json:"warnings,omitempty"`
}

// HeroRenderer renders a hero image for one day. Apply takes it as a
// function so the tests can run without Chromium — the screenshot path
// depends on an external browser binary and is verified by hand.
type HeroRenderer func(folder string, day int, e Entry) error

// Apply validates the plan, rewrites every day reference in every
// markdown file of every queued folder, re-renders the heroes of moved
// days, and writes queue.yaml last.
//
// heroTemplate empty means skip hero rendering — the tests use that, and
// so does a run that only needs to see the file diff.
func Apply(repoRoot, queuePath string, p Plan, heroTemplate string, dryRun bool) (Result, error) {
	return ApplyWith(repoRoot, queuePath, p, dryRun, heroRendererFor(repoRoot, heroTemplate))
}

// ApplyWith is Apply with an injectable hero renderer.
func ApplyWith(repoRoot, queuePath string, p Plan, dryRun bool, render HeroRenderer) (Result, error) {
	q, err := queue.Load(queuePath)
	if err != nil {
		return Result{}, fmt.Errorf("loading queue %s: %w", queuePath, err)
	}
	if err := Validate(q, p, repoRoot); err != nil {
		return Result{}, err
	}

	res := Result{DryRun: dryRun}
	dayMap := DayMap(q, p)

	inQueue := map[int]queue.Entry{}
	for _, e := range q.Entries {
		inQueue[e.Number] = e
	}
	for _, e := range sortedEntries(p) {
		existing, ok := inQueue[e.Number]
		switch {
		case !ok:
			res.Inserted = append(res.Inserted, Change{Number: e.Number, Title: e.Title, Folder: e.Folder, NewDay: e.Day})
		case existing.Day != e.Day:
			res.Moved = append(res.Moved, Change{Number: e.Number, Title: e.Title, Folder: e.Folder, OldDay: existing.Day, NewDay: e.Day})
		}
	}

	// Rewrite references across every queued folder, not only the moved
	// ones: the stale reference lives in the folder that cites the moved
	// day, which is usually some other question entirely.
	if len(dayMap) > 0 {
		folders := map[string]bool{}
		for _, e := range q.Entries {
			folders[e.Folder] = true
		}
		for _, e := range p.Entries {
			folders[e.Folder] = true
		}
		names := make([]string, 0, len(folders))
		for f := range folders {
			names = append(names, f)
		}
		sort.Strings(names)

		for _, folder := range names {
			matches, err := filepath.Glob(filepath.Join(repoRoot, folder, "*.md"))
			if err != nil {
				return Result{}, err
			}
			sort.Strings(matches)
			for _, path := range matches {
				raw, err := os.ReadFile(path)
				if err != nil {
					return Result{}, fmt.Errorf("reading %s: %w", path, err)
				}
				updated := RewriteDayRefs(string(raw), dayMap)
				if updated == string(raw) {
					continue
				}
				rel, _ := filepath.Rel(repoRoot, path)
				res.FilesRewritten = append(res.FilesRewritten, rel)
				if dryRun {
					continue
				}
				if err := os.WriteFile(path, []byte(updated), 0o644); err != nil {
					return Result{}, fmt.Errorf("writing %s: %w", path, err)
				}
			}
		}
	}

	// Heroes carry the day number on the card, and the palette is
	// derived from the day, so a moved day's image is genuinely wrong
	// until it is re-rendered.
	if render != nil {
		for _, mv := range res.Moved {
			var entry Entry
			for _, e := range p.Entries {
				if e.Number == mv.Number {
					entry = e
					break
				}
			}
			rel := filepath.Join(mv.Folder, "HERO.png")
			if _, err := os.Stat(filepath.Join(repoRoot, rel)); err != nil {
				res.Warnings = append(res.Warnings, fmt.Sprintf("day %d (%d): no HERO.png to re-render", mv.NewDay, mv.Number))
				continue
			}
			res.HeroesRendered = append(res.HeroesRendered, rel)
			if dryRun {
				continue
			}
			if err := render(mv.Folder, mv.NewDay, entry); err != nil {
				return Result{}, fmt.Errorf("re-rendering hero for day %d (%d): %w", mv.NewDay, mv.Number, err)
			}
		}
	}

	if dryRun {
		return res, nil
	}

	dayByNumber := make(map[int]int, len(p.Entries))
	for _, e := range p.Entries {
		dayByNumber[e.Number] = e.Day
	}
	q.Reorder(dayByNumber)
	for _, ins := range res.Inserted {
		var e Entry
		for _, pe := range p.Entries {
			if pe.Number == ins.Number {
				e = pe
				break
			}
		}
		q.Entries = append(q.Entries, queue.Entry{
			Day:        e.Day,
			Number:     e.Number,
			Title:      e.Title,
			Difficulty: e.Difficulty,
			Folder:     e.Folder,
			Batch:      e.Batch,
			Status:     queue.StatusPendingContent,
			BuildsOn:   e.BuildsOn,
			PostedAt:   nil,
		})
	}
	q.Reorder(dayByNumber)
	if err := q.Save(queuePath); err != nil {
		return Result{}, fmt.Errorf("writing queue %s: %w", queuePath, err)
	}
	return res, nil
}
```

- [ ] **Step 4: Add the real hero renderer**

Append to `challenge/internal/renumber/renumber.go`:

```go
// heroRendererFor returns a renderer backed by the hero package, or nil
// when no template was given.
//
// This calls hero.RenderHTML and hero.Screenshot directly rather than
// cli.HeroGenerate, which wraps the same two steps: cli imports this
// package for the queue-renumber subcommand, so importing cli back here
// would be a cycle.
func heroRendererFor(repoRoot, heroTemplate string) HeroRenderer {
	if heroTemplate == "" {
		return nil
	}
	return func(folder string, day int, e Entry) error {
		html, err := hero.RenderHTML(heroTemplate, hero.Data{
			Day:        day,
			Total:      365,
			Topic:      e.Batch,
			Difficulty: e.Difficulty,
			Title:      e.Title,
		})
		if err != nil {
			return fmt.Errorf("rendering hero HTML: %w", err)
		}
		tmp, err := os.CreateTemp("", fmt.Sprintf("hero-day-%d-*.html", day))
		if err != nil {
			return err
		}
		tmpPath := tmp.Name()
		if _, err := tmp.WriteString(html); err != nil {
			tmp.Close()
			os.Remove(tmpPath)
			return err
		}
		if err := tmp.Close(); err != nil {
			os.Remove(tmpPath)
			return err
		}
		if err := hero.Screenshot(tmpPath, filepath.Join(repoRoot, folder, "HERO.png")); err != nil {
			return fmt.Errorf("screenshotting hero (HTML kept at %s): %w", tmpPath, err)
		}
		os.Remove(tmpPath)
		return nil
	}
}
```

Add to the import block:

```go
	"leetcode_solutions/challenge/internal/hero"
```

- [ ] **Step 5: Run the tests to verify they pass**

Run: `go test ./challenge/internal/renumber/ -v`
Expected: all tests PASS.

- [ ] **Step 6: Run the whole suite**

Run: `go test ./challenge/...`
Expected: every package `ok`.

- [ ] **Step 7: Commit**

```bash
git add challenge/internal/renumber/
git commit -m "feat(renumber): apply a reorder across files, heroes and queue

File work first, queue.yaml last. A run that fails partway then leaves
the queue describing the state the files were in before it started,
rather than a state that never existed.

References are rewritten across every queued folder rather than only the
moved ones, because a stale reference lives in the folder that cites the
moved day, not in the day that moved. Only *.md is touched - a day
number in Go source is a comment, not a published artifact.

Inserted entries land as pending_content so posting skips them until
their content is written."
```

---

## Task 7: CLI wiring — `queue-renumber`

**Files:**
- Modify: `challenge/internal/cli/cli.go`
- Modify: `challenge/cmd/leetcodectl/main.go`

- [ ] **Step 1: Add the wrapper**

Append to `challenge/internal/cli/cli.go`:

```go
// QueueRenumber applies a reordering plan: it rewrites every day
// reference in every queued folder's markdown, re-renders the heroes of
// moved days, and writes the queue last. A dry run reports the same
// thing and writes nothing.
func QueueRenumber(repoRoot, queuePath, heroTemplate string, p renumber.Plan, dryRun bool) (renumber.Result, error) {
	return renumber.Apply(repoRoot, queuePath, p, heroTemplate, dryRun)
}
```

Add to the import block: `"leetcode_solutions/challenge/internal/renumber"`

- [ ] **Step 2: Add the dispatch case**

In `challenge/cmd/leetcodectl/main.go`, add after the `queue-append` case:

```go
	case "queue-renumber":
		var in struct {
			RepoRoot     string         `json:"repoRoot"`
			QueuePath    string         `json:"queuePath"`
			HeroTemplate string         `json:"heroTemplate"`
			Plan         renumber.Plan  `json:"plan"`
			DryRun       bool           `json:"dryRun"`
		}
		if err := json.Unmarshal(payload, &in); err != nil {
			return nil, err
		}
		return cli.QueueRenumber(in.RepoRoot, in.QueuePath, in.HeroTemplate, in.Plan, in.DryRun)
```

Add to the import block: `"leetcode_solutions/challenge/internal/renumber"`

- [ ] **Step 3: Build and verify the subcommand is reachable**

Run:
```bash
go build -o /tmp/leetcodectl ./challenge/cmd/leetcodectl && \
/tmp/leetcodectl queue-renumber '{"repoRoot":".","queuePath":"challenge/queue.yaml","plan":{"entries":[]},"dryRun":true}'
```
Expected: JSON with `"dryRun": true` and empty `moved`/`inserted`. An empty plan is legal and does nothing.

- [ ] **Step 4: Document it in the challenge README**

`challenge/README.md` describes what each part of the tooling does. Add to the `Layout` section, after the `queue.yaml` bullet:

```markdown
- `plans/` — reorder mappings fed to `leetcodectl queue-renumber`. Each file is a full
  intended day ordering rather than a diff, so validation can catch a duplicate day or a
  forward `builds_on` before anything is written. Keeping them in the repo means a past
  reshape can be read back: the queue only shows where days ended up, not why.
```

- [ ] **Step 5: Run the whole suite**

Run: `go test ./challenge/...`
Expected: every package `ok`.

- [ ] **Step 6: Commit**

```bash
git add challenge/internal/cli/cli.go challenge/cmd/leetcodectl/main.go challenge/README.md
git commit -m "feat(cli): add queue-renumber subcommand

queue-append was the only queue mutation, so reshaping the queue meant
hand-editing yaml and every file carrying a day number. This exposes the
renumber package the same way every other subcommand is exposed: one
JSON argument in, one JSON result out."
```

---

## Task 8: Write the concrete reorder plan file

The full day 12–81 mapping from the spec, as the JSON `queue-renumber` consumes. Days 1–11 are included and unchanged, so validation checks them rather than assuming.

**Files:**
- Create: `challenge/plans/2026-09-12-reorder.json`

- [ ] **Step 1: Generate the plan JSON from the spec's day map**

The spec's day table is the source. Build the JSON with one entry per day 1–81: `day`, `number`, `title`, `difficulty`, `folder`, `batch`, `buildsOn`.

For days 1–11 and every other day whose question is already queued, copy `title`, `difficulty` and `folder` verbatim from the existing entry in `challenge/queue.yaml` — a mismatch there would rewrite the queue's own metadata as a side effect of renumbering.

For the 30 new days, use these folders, all verified to exist:

| # | Title | Diff | Folder | Batch |
|---|---|---|---|---|
| 206 | Reverse Linked List | easy | `easy_problems/201_300/reverse_linked_list` | linked list |
| 876 | Middle of the Linked List | easy | `easy_problems/801_900/middle_of_the_linked_list` | linked list |
| 141 | Linked List Cycle | easy | `easy_problems/101_200/linked_list_cycle` | linked list |
| 21 | Merge Two Sorted Lists | easy | `easy_problems/1_100/merge_two_sorted_lists` | linked list |
| 234 | Palindrome Linked List | easy | `easy_problems/201_300/palindrome_linked_list` | linked list |
| 19 | Remove Nth Node From End of List | medium | `medium_problems/1_100/remove_nth_node_from_end_of_list` | linked list |
| 92 | Reverse Linked List II | medium | `medium_problems/1_100/reverse_linked_list_ii` | linked list |
| 143 | Reorder List | medium | `medium_problems/101_200/reorder_list` | linked list |
| 1971 | Find if Path Exists in Graph | easy | `easy_problems/1901_2000/find_if_path_exists_in_graph` | graph |
| 200 | Number of Islands | medium | `medium_problems/101_200/number_of_islands` | graph |
| 695 | Max Area of Island | medium | `medium_problems/601_700/max_area_of_island` | graph |
| 547 | Number of Provinces | medium | `medium_problems/501_600/number_of_provinces` | graph |
| 994 | Rotting Oranges | medium | `medium_problems/901_1000/rotting_oranges` | graph |
| 542 | 01 Matrix | medium | `medium_problems/501_600/01_matrix` | graph |
| 261 | Graph Valid Tree | medium | `medium_problems/201_300/graph_valid_tree` | graph |
| 133 | Clone Graph | medium | `medium_problems/101_200/clone_graph` | graph |
| 207 | Course Schedule | medium | `medium_problems/201_300/course_schedule` | graph |
| 20 | Valid Parentheses | easy | `easy_problems/1_100/valid_parentheses` | stack |
| 496 | Next Greater Element I | easy | `easy_problems/401_500/next_greater_element_i` | stack |
| 155 | Min Stack | medium | `medium_problems/101_200/min_stack` | stack |
| 739 | Daily Temperatures | medium | `medium_problems/701_800/daily_temperatures` | stack |
| 150 | Evaluate Reverse Polish Notation | medium | `medium_problems/101_200/evaluate_reverse_polish_notation` | stack |
| 341 | Flatten Nested List Iterator | medium | `medium_problems/301_400/flatten_nested_list_iterator` | stack |
| 636 | Exclusive Time of Functions | medium | `medium_problems/601_700/exclusive_time_of_functions` | stack |
| 1046 | Last Stone Weight | easy | `easy_problems/1001_1100/last_stone_weight` | heap |
| 215 | Kth Largest Element in an Array | medium | `medium_problems/201_300/kth_largest_element_in_an_array` | heap |
| 347 | Top K Frequent Elements | medium | `medium_problems/301_400/top_k_frequent_elements` | heap |
| 692 | Top K Frequent Words | medium | `medium_problems/601_700/top_k_frequent_words` | heap |
| 23 | Merge k Sorted Lists | hard | `hard_problems/1_100/merge_k_sorted_lists` | heap |
| 264 | Ugly Number II | medium | `medium_problems/201_300/ugly_number_ii` | heap |

- [ ] **Step 2: Verify every folder in the plan exists**

Run:
```bash
python3 -c "
import json
p=json.load(open('challenge/plans/2026-09-12-reorder.json'))
import os
missing=[e['folder'] for e in p['entries'] if not os.path.isdir(e['folder'])]
print('missing:', missing or 'none')
print('entries:', len(p['entries']))
days=[e['day'] for e in p['entries']]
print('days 1-81 complete:', sorted(days)==list(range(1,82)))
"
```
Expected: `missing: none`, `entries: 81`, `days 1-81 complete: True`.

- [ ] **Step 3: Commit**

```bash
git add challenge/plans/2026-09-12-reorder.json
git commit -m "chore(challenge): add the day 12-81 reorder mapping"
```

---

## Task 9: Repair the two references a renumber cannot fix

A reorder can turn a backward reference into a forward one, and no string substitution fixes that — the referenced question genuinely has not been published yet at that point in the series. Exactly two such references exist, both found by scanning every `Day N` reference in the repo against the target ordering.

**Files:**
- Modify: `easy_problems/501_600/diameter_of_binary_tree/POST_LINKEDIN_ARTICLE.md:137`
- Modify: `easy_problems/501_600/diameter_of_binary_tree/POST_SUBSTACK.md:144`
- Modify: `medium_problems/301_400/find_leaves_of_binary_tree/INTUITION.md:36`
- Modify: `medium_problems/301_400/find_leaves_of_binary_tree/SOLUTION.md:44`

- [ ] **Step 1: Fix the stale "tomorrow" line in 543 Diameter**

Both files end with `repo. See you tomorrow for Day 13.` This was **already wrong before this reorder**: 543 sat on day 21, so tomorrow was day 22, not 13. It is a leftover from an earlier reshuffle where prose was not updated.

Under the new ordering 543 is day 13 and tomorrow is day 14, 563 Binary Tree Tilt. Replace in both files:

```
repo. See you tomorrow for Day 14.
```

Do this **before** running the renumber in Task 10, so the tool's own rewrite does not renumber the stale `13` into something equally wrong.

- [ ] **Step 2: Fix the forward reference in 366 Find Leaves**

`366 Find Leaves` moves to day 16 but cites day 15's lazy slice growth, which is 102 Level Order Traversal — now day 29, thirteen days later. The reference cannot be renumbered; it has to stop being a reference.

In `INTUITION.md:36`, replace:

```
The outer slice grows lazily, the same way Day 15's did. `if len(*arr) <= index` creates the
```

with:

```
The outer slice grows lazily: `if len(*arr) <= index` creates the
```

In `SOLUTION.md:44`, replace:

```
the first node to reach a round creates its bucket, the same lazy growth as Day 15's
```

with:

```
the first node to reach a round creates its bucket, growing the outer slice on demand
```

- [ ] **Step 3: Confirm no other forward reference exists**

Run:
```bash
python3 - <<'EOF'
import re,os,glob,json
plan=json.load(open('challenge/plans/2026-09-12-reorder.json'))
new_by_num={e['number']:e['day'] for e in plan['entries']}
src=open('challenge/queue.yaml').read()
blocks=re.split(r'\n    - day: ', src)[1:]
d2f={int(b.split('\n')[0]):re.search(r'folder: (.*)',b).group(1).strip() for b in blocks}
d2n={int(b.split('\n')[0]):int(re.search(r'number: (\d+)',b).group(1)) for b in blocks}
o2n={d:new_by_num[d2n[d]] for d in d2f if d2n[d] in new_by_num}
bad=[]
for d,f in sorted(d2f.items()):
    if not os.path.isdir(f): continue
    for p in glob.glob(f+'/*.md'):
        for ln,line in enumerate(open(p),1):
            for m in re.finditer(r'Day (\d+)', line):
                n=int(m.group(1))
                if n in o2n and n!=d and o2n[n]>=o2n[d]:
                    bad.append(f"{p}:{ln}  d{d}->{o2n[d]} cites d{n}->{o2n[n]}")
print("forward refs remaining:", len(bad))
print("\n".join(bad))
EOF
```
Expected: `forward refs remaining: 0`. If any remain, fix them the same way before proceeding — a forward reference published to a reader following in order is a broken link.

- [ ] **Step 4: Commit**

```bash
git add easy_problems/501_600/diameter_of_binary_tree medium_problems/301_400/find_leaves_of_binary_tree
git commit -m "fix(content): repair two day references the reorder would break

543 Diameter ended with 'See you tomorrow for Day 13' while sitting on
day 21 - already wrong before this reorder, a leftover from an earlier
reshuffle. It becomes day 13, so tomorrow is day 14.

366 Find Leaves cited day 15's lazy slice growth. Under the new ordering
366 is day 16 and 102 Level Order is day 29, so the reference points
forward at something the reader has not seen. It now explains the growth
in place instead of citing it."
```

---

## Task 10: Dry run, review, apply

**Files:** none created; this runs the tool against the real repo.

- [ ] **Step 1: Confirm the working tree is clean**

Run: `git status --porcelain`
Expected: no output. The renumber touches ~40 folders, and an unrelated uncommitted change would be impossible to separate from its diff afterwards.

- [ ] **Step 2: Dry run**

Run:
```bash
/tmp/leetcodectl queue-renumber '{
  "repoRoot":".",
  "queuePath":"challenge/queue.yaml",
  "heroTemplate":"challenge/hero_template.html",
  "plan":'"$(cat challenge/plans/2026-09-12-reorder.json)"',
  "dryRun":true
}' > /tmp/renumber-dryrun.json; echo "exit=$?"
```
Expected: `exit=0`. A non-zero exit means validation refused the plan; read the error, fix the plan file, and repeat. Nothing has been written.

- [ ] **Step 3: Read the dry-run report before applying**

Run:
```bash
python3 -c "
import json; r=json.load(open('/tmp/renumber-dryrun.json'))
print('moved   :', len(r['moved']))
print('inserted:', len(r['inserted']))
print('files   :', len(r['filesRewritten']))
print('heroes  :', len(r['heroesRendered']))
print('warnings:', r.get('warnings') or 'none')
for m in r['moved'][:5]: print('  ', m['number'], m['oldDay'], '->', m['newDay'])
"
```
Expected: `moved: 39`, `inserted: 30`, `files` in the low hundreds, `heroes: 39`, `warnings: none`.

39 rather than 40 because 1022 Sum of Root To Leaf is already on day 12 and stays there; the other 39 tree questions move.

If `inserted` is not 30 or `moved` is not 39, stop — the plan file disagrees with the spec and applying it would produce a queue nobody designed.

- [ ] **Step 4: Apply**

Run:
```bash
/tmp/leetcodectl queue-renumber '{
  "repoRoot":".",
  "queuePath":"challenge/queue.yaml",
  "heroTemplate":"challenge/hero_template.html",
  "plan":'"$(cat challenge/plans/2026-09-12-reorder.json)"',
  "dryRun":false
}' > /tmp/renumber-applied.json; echo "exit=$?"
```
Expected: `exit=0`. This re-renders 40 hero images through headless Chromium and takes a few minutes.

- [ ] **Step 5: Commit**

```bash
git add -A
git commit -m "feat(challenge): reorder days 12-81 into braided topic arcs

Forty consecutive binary-tree days become four themed tree arcs
interleaved with linked list, graph, stack and heap arcs drawn from
already-solved problems. The longest unbroken run of one topic drops
from 41 days to 13.

Every arc opens on a question whose technique the previous arc taught,
so the series stays one chain: day 10 flattened a tree into a linked
list, level-order traversal is BFS, and the BST iterator and the nested
list iterator are the same design problem.

Thirty new days are inserted as pending_content - their numbers are
reserved and posting skips them until their write-ups exist."
```

---

## Task 11: Verify the reorder landed correctly

- [ ] **Step 1: Every `Day N/365` header matches its queue entry**

Run:
```bash
python3 - <<'EOF'
import re,glob,os
src=open('challenge/queue.yaml').read()
blocks=re.split(r'\n    - day: ', src)[1:]
bad=[]
for b in blocks:
    d=int(b.split('\n')[0]); f=re.search(r'folder: (.*)',b).group(1).strip()
    for p in glob.glob(f+'/POST_*.md'):
        for ln,line in enumerate(open(p),1):
            for m in re.finditer(r'Day (\d+)/365', line):
                if int(m.group(1))!=d: bad.append(f"{p}:{ln} says Day {m.group(1)}/365, queue says {d}")
print("header mismatches:", len(bad)); print("\n".join(bad[:20]))
EOF
```
Expected: `header mismatches: 0`.

- [ ] **Step 2: No forward `builds_on` in the written queue**

Run:
```bash
python3 - <<'EOF'
import re
src=open('challenge/queue.yaml').read()
blocks=re.split(r'\n    - day: ', src)[1:]
day={}; deps={}
for b in blocks:
    d=int(b.split('\n')[0]); n=int(re.search(r'number: (\d+)',b).group(1)); day[n]=d
    bo=re.search(r'builds_on:\n((?:        - \d+\n)+)',b)
    deps[n]=[int(x) for x in re.findall(r'- (\d+)',bo.group(1))] if bo else []
bad=[f"q{n} d{day[n]} <- q{x} d{day.get(x)}" for n,ds in deps.items() for x in ds if day.get(x,999)>=day[n]]
print("forward builds_on:", len(bad)); print("\n".join(bad))
EOF
```
Expected: `forward builds_on: 0`.

- [ ] **Step 3: Posted days did not move**

Run: `git diff HEAD~1 --stat -- challenge/queue.yaml && grep -A2 "day: 1$" challenge/queue.yaml | head -5`
Expected: days 1–11 still carry their original question numbers (104, 108, 257, 404, 110, 222, 111, 144, 145, 114, 112) and their `posted_at` timestamps are unchanged.

- [ ] **Step 4: Heroes were re-rendered**

Run: `git diff HEAD~1 --stat -- '*HERO.png' | tail -3`
Expected: about 39 `HERO.png` files changed.

- [ ] **Step 5: Spot-check one hero image by eye**

Open `medium_problems/301_400/find_leaves_of_binary_tree/HERO.png` and confirm the card reads **Day 16**. The day number is baked into the image, so this is the one check automation cannot do.

- [ ] **Step 6: Full suite still green**

Run: `go test ./challenge/... && go build ./...`
Expected: all `ok`, build clean.

- [ ] **Step 7: Confirm the next Discord post would be the right day**

Run: `/tmp/leetcodectl linkedin-batch '{"repoRoot":".","queuePath":"challenge/queue.yaml","destination":"discord","count":1,"outPath":"/tmp/next.md"}' 2>/dev/null || true; head -3 /tmp/next.md`
Expected: day 12, question 1022 Sum of Root To Leaf Binary Numbers. If it shows a day above 12, a `pending_content` entry was skipped and the days before it are unwritten — which is correct behaviour only if Task 12 has not run yet.

---

## Task 12: Generate content for the linked-list arc

The tight deadline: 8 questions before day 19, roughly a week out. Run this immediately after Task 11, not later.

- [ ] **Step 1: Generate the arc**

Invoke the existing skill, which fills in `INTUITION.md`, `SOLUTION.md`, the five `POST_*.md` files, `HERO.png` and the companies list per question:

```
/leetcode-content --list-name "linked list"
```

Generate, in day order: 206, 876, 141, 21, 234, 19, 92, 143.

These entries are already in the queue as `pending_content` with their day numbers reserved, so the skill must **not** call `queue-append` for them — appending would assign a second day number to a question that already has one. It should call `mark-content-ready` instead:

```bash
/tmp/leetcodectl mark-content-ready '{"queuePath":"challenge/queue.yaml","numbers":[206,876,141,21,234,19,92,143]}'
```

- [ ] **Step 2: Verify the arc is postable**

Run:
```bash
python3 -c "
import re
src=open('challenge/queue.yaml').read()
for b in re.split(r'\n    - day: ', src)[1:]:
    d=int(b.split('\n')[0])
    if 19<=d<=26:
        print(d, re.search(r'number: (\d+)',b).group(1), re.search(r'status: (.*)',b).group(1).strip())
"
```
Expected: days 19–26 all `content_ready`.

- [ ] **Step 3: Confirm no gap will open before day 19**

Run:
```bash
python3 -c "
import re
src=open('challenge/queue.yaml').read()
pend=[int(b.split('\n')[0]) for b in re.split(r'\n    - day: ', src)[1:] if 'pending_content' in b]
print('first pending day:', min(pend) if pend else 'none')
"
```
Expected: `first pending day: 36` — the graph arc, which is the next batch and is not due yet.

- [ ] **Step 4: Commit**

The `/leetcode-content` skill commits per question. Confirm with `git log --oneline -10`.

### Remaining batches

| Batch | Questions, in day order | Due before |
|---|---|---|
| Graph | 1971, 200, 695, 547, 994, 542, 261, 133, 207 | day 36 |
| Stack & Iterator | 20, 496, 155, 739, 150, 341, 636 | day 58 |
| Heap | 1046, 215, 347, 692, 23, 264 | day 76 |

Each runs exactly as Task 12: generate in day order, `mark-content-ready`, verify no pending day sits before the posting cursor.
