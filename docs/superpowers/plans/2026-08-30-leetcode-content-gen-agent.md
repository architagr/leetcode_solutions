# LeetCode Content-Gen Agent Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Build the `/leetcode-content` Claude Code skill and its supporting `leetcodectl` Go CLI, so a user can feed it a LeetCode problem-list URL + difficulty filter and have it generate teaching content (intuition/solution docs, commented code, companies, hero image, LinkedIn/Discord post drafts) for every already-solved question in that list, in a resumable, ordered queue.

**Architecture:** All deterministic logic (slug/path math, git-log scanning, queue bookkeeping, git-mv reorg, LeetCode GraphQL fetch, hero-image templating) lives in small, independently-tested Go packages under `challenge/internal/`, wired together by a single JSON-in/JSON-out CLI (`challenge/cmd/leetcodectl`). A Claude Code project skill (`.claude/skills/leetcode-content/SKILL.md`) orchestrates calls to that CLI and performs the parts that need an LLM (reading existing solution code and writing `INTUITION.md`/`SOLUTION.md`/post drafts in the user's voice).

**Tech Stack:** Go (matches existing repo), `gopkg.in/yaml.v3` (already an indirect dependency) for the two YAML state files, standard library `net/http`/`encoding/json` for the LeetCode GraphQL client, `text/template` for the hero-image HTML, `npx playwright` (external, invoked via `os/exec`) for the HTML→PNG screenshot step.

**Spec:** `docs/superpowers/specs/2026-08-30-leetcode-content-gen-agent-design.md`

---

## Before you start

Run these once, in the repo root, before Task 1:

```bash
go env GOFLAGS
```

No specific output expected — just confirming `go` is on PATH and the module resolves. All Go commands below assume you're in the repo root (`leetcode_solutions` module).

---

### Task 1: `slugutil` — slug/path math

**Files:**
- Create: `challenge/internal/slugutil/slugutil.go`
- Test: `challenge/internal/slugutil/slugutil_test.go`

- [ ] **Step 1: Write the failing test**

```go
package slugutil

import "testing"

func TestFolderName(t *testing.T) {
	cases := map[string]string{
		"univalued-binary-tree": "univalued_binary_tree",
		"two-sum":               "two_sum",
		"already_snake":         "already_snake",
	}
	for slug, want := range cases {
		if got := FolderName(slug); got != want {
			t.Errorf("FolderName(%q) = %q, want %q", slug, got, want)
		}
	}
}

func TestRangeBucket(t *testing.T) {
	cases := map[int]string{
		1:    "1_100",
		55:   "1_100",
		100:  "1_100",
		101:  "101_200",
		965:  "901_1000",
		1001: "1001_1100",
	}
	for number, want := range cases {
		if got := RangeBucket(number); got != want {
			t.Errorf("RangeBucket(%d) = %q, want %q", number, got, want)
		}
	}
}
```

- [ ] **Step 2: Run test to verify it fails**

Run: `go test ./challenge/internal/slugutil/...`
Expected: FAIL — `undefined: FolderName` (package doesn't exist yet)

- [ ] **Step 3: Write minimal implementation**

```go
// Package slugutil converts LeetCode question slugs and numbers into this
// repo's folder-naming conventions (snake_case folder names, 100-wide
// number-range buckets like "901_1000").
package slugutil

import (
	"fmt"
	"strings"
)

// FolderName converts a LeetCode URL slug (dash-case, e.g.
// "univalued-binary-tree") into this repo's folder naming convention
// (snake_case, e.g. "univalued_binary_tree").
func FolderName(slug string) string {
	return strings.ReplaceAll(slug, "-", "_")
}

// RangeBucket returns the 100-wide folder bucket name for a LeetCode
// question number, matching the existing repo convention
// (easy_problems/901_1000, easy_problems/1_100, ...).
func RangeBucket(number int) string {
	start := ((number-1)/100)*100 + 1
	end := start + 99
	return fmt.Sprintf("%d_%d", start, end)
}
```

- [ ] **Step 4: Run test to verify it passes**

Run: `go test ./challenge/internal/slugutil/...`
Expected: PASS

- [ ] **Step 5: Commit**

```bash
git add challenge/internal/slugutil
git commit -m "feat(challenge): add slugutil package for slug/range math"
```

---

### Task 2: `queue` — the shared ordered queue state

**Files:**
- Create: `challenge/internal/queue/queue.go`
- Test: `challenge/internal/queue/queue_test.go`

- [ ] **Step 1: Write the failing test**

```go
package queue

import (
	"os"
	"path/filepath"
	"testing"
)

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
```

(`NextDay` is 3, not 2, because `Append` increments it — `Queue{NextDay: 2}` + one `Append` = `NextDay: 3`, consistent with `TestAppendAssignsDayAndIncrements` above.)

- [ ] **Step 2: Run test to verify it fails**

Run: `go test ./challenge/internal/queue/...`
Expected: FAIL — package doesn't exist yet

- [ ] **Step 3: Write minimal implementation**

```go
// Package queue manages challenge/queue.yaml, the ordered, resumable
// record of which already-solved LeetCode questions have had content
// generated and/or posted, in the fixed "365 Days" order.
package queue

import (
	"os"

	"gopkg.in/yaml.v3"
)

const (
	StatusContentReady = "content_ready"
	StatusPosted        = "posted"
)

// Entry is one question's record in the challenge queue.
type Entry struct {
	Day        int     `yaml:"day"`
	Number     int     `yaml:"number"`
	Title      string  `yaml:"title"`
	Difficulty string  `yaml:"difficulty"`
	Folder     string  `yaml:"folder"`
	Batch      string  `yaml:"batch"`
	Status     string  `yaml:"status"`
	PostedAt   *string `yaml:"posted_at"`
}

// Queue is the full contents of challenge/queue.yaml.
type Queue struct {
	NextDay int     `yaml:"next_day"`
	Entries []Entry `yaml:"entries"`
}

// Load reads a Queue from path. A missing file is not an error — it
// returns an empty Queue starting at day 1, so a fresh repo works with
// no setup.
func Load(path string) (*Queue, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return &Queue{NextDay: 1}, nil
		}
		return nil, err
	}
	var q Queue
	if err := yaml.Unmarshal(data, &q); err != nil {
		return nil, err
	}
	if q.NextDay == 0 {
		q.NextDay = 1
	}
	return &q, nil
}

// Save writes the Queue to path as YAML.
func (q *Queue) Save(path string) error {
	data, err := yaml.Marshal(q)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}

// Has reports whether an entry for the given question number already
// exists in the queue.
func (q *Queue) Has(number int) bool {
	for _, e := range q.Entries {
		if e.Number == number {
			return true
		}
	}
	return false
}

// Append assigns the next Day number to e, sets its status to
// content_ready, appends it, and returns the assigned day.
func (q *Queue) Append(e Entry) int {
	if q.NextDay == 0 {
		q.NextDay = 1
	}
	e.Day = q.NextDay
	e.Status = StatusContentReady
	q.Entries = append(q.Entries, e)
	q.NextDay++
	return e.Day
}
```

- [ ] **Step 4: Run test to verify it passes**

Run: `go test ./challenge/internal/queue/...`
Expected: PASS

- [ ] **Step 5: Add yaml.v3 as a direct dependency and commit**

```bash
go mod tidy
git add challenge/internal/queue go.mod go.sum
git commit -m "feat(challenge): add queue package for challenge/queue.yaml"
```

---

### Task 3: `companies` — best-effort company lookup

**Files:**
- Create: `challenge/internal/companies/companies.go`
- Create: `challenge/companies_dataset.json`
- Create: `challenge/companies_dataset.md`
- Test: `challenge/internal/companies/companies_test.go`

- [ ] **Step 1: Write the failing test**

```go
package companies

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadMissingFileReturnsEmptyDataset(t *testing.T) {
	ds, err := Load(filepath.Join(t.TempDir(), "nope.json"))
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if len(ds.Lookup(1)) != 0 {
		t.Errorf("Lookup on empty dataset = %v, want empty", ds.Lookup(1))
	}
}

func TestLoadAndLookup(t *testing.T) {
	path := filepath.Join(t.TempDir(), "dataset.json")
	if err := os.WriteFile(path, []byte(`{"965": ["Google", "Amazon"]}`), 0o644); err != nil {
		t.Fatal(err)
	}
	ds, err := Load(path)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	got := ds.Lookup(965)
	if len(got) != 2 || got[0] != "Google" || got[1] != "Amazon" {
		t.Errorf("Lookup(965) = %v", got)
	}
	if len(ds.Lookup(1)) != 0 {
		t.Errorf("Lookup(1) = %v, want empty", ds.Lookup(1))
	}
}
```

- [ ] **Step 2: Run test to verify it fails**

Run: `go test ./challenge/internal/companies/...`
Expected: FAIL — package doesn't exist yet

- [ ] **Step 3: Write minimal implementation**

```go
// Package companies looks up which companies are known to have asked a
// given LeetCode question, from a vendored, manually-refreshed dataset
// snapshot (LeetCode's own company-tag data is Premium-only with no
// public API).
package companies

import (
	"encoding/json"
	"os"
	"strconv"
)

// Dataset maps a LeetCode question number (as a string key) to the list
// of companies known to have asked it.
type Dataset map[string][]string

// Load reads a Dataset from a JSON file. A missing file is not an error
// — it returns an empty Dataset, so lookups are simply best-effort.
func Load(path string) (Dataset, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return Dataset{}, nil
		}
		return nil, err
	}
	var ds Dataset
	if err := json.Unmarshal(data, &ds); err != nil {
		return nil, err
	}
	return ds, nil
}

// Lookup returns the companies known for a question number, or an empty
// (non-nil) slice if there's no entry.
func (d Dataset) Lookup(number int) []string {
	if companies, ok := d[strconv.Itoa(number)]; ok {
		return companies
	}
	return []string{}
}
```

- [ ] **Step 4: Run test to verify it passes**

Run: `go test ./challenge/internal/companies/...`
Expected: PASS

- [ ] **Step 5: Seed the vendored dataset file and its refresh notes**

`challenge/companies_dataset.json`:
```json
{}
```

`challenge/companies_dataset.md`:
```markdown
# Companies dataset

`companies_dataset.json` is a manually-vendored snapshot mapping LeetCode
question numbers to companies known to have asked them (as JSON object
keys are strings, question numbers are quoted: `"965": ["Google", "Amazon"]`).

LeetCode's own company-tag data is Premium-only and has no public API, so
this file starts empty and is refreshed by hand — periodically replace it
wholesale with a fresh export from a public community dataset (e.g. a
GitHub repo tracking company-wise LeetCode question lists). The
content-gen agent's `COMPANIES.md` output is best-effort: if a question
number has no entry here, that file is simply not generated.
```

- [ ] **Step 6: Commit**

```bash
git add challenge/internal/companies challenge/companies_dataset.json challenge/companies_dataset.md
git commit -m "feat(challenge): add companies lookup package and vendored dataset seed"
```

---

### Task 4: `gitmap` — number→folder map from commit history

**Files:**
- Create: `challenge/internal/gitmap/gitmap.go`
- Test: `challenge/internal/gitmap/gitmap_test.go`

- [ ] **Step 1: Write the failing test**

```go
package gitmap

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestParseCommitSubject(t *testing.T) {
	cases := map[string]struct {
		number int
		ok     bool
	}{
		"Easy(965) Univalued Binary Tree":    {965, true},
		"medium(3652)":                       {3652, true},
		"Easy (168)":                         {168, true},
		"Binary Tree Vertical Order Traversal": {0, false},
		"upgrade go and testify":             {0, false},
	}
	for subject, want := range cases {
		number, ok := ParseCommitSubject(subject)
		if ok != want.ok || (ok && number != want.number) {
			t.Errorf("ParseCommitSubject(%q) = (%d, %v), want (%d, %v)", subject, number, ok, want.number, want.ok)
		}
	}
}

func initTestRepo(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	run := func(args ...string) {
		cmd := exec.Command("git", args...)
		cmd.Dir = dir
		if out, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("git %v: %v\n%s", args, err, out)
		}
	}
	run("init", "-q")
	run("config", "user.email", "test@example.com")
	run("config", "user.name", "Test")
	return dir
}

func commitFile(t *testing.T, dir, relPath, subject string) string {
	t.Helper()
	full := filepath.Join(dir, relPath)
	if err := os.MkdirAll(filepath.Dir(full), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(full, []byte("package p\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command("git", "-C", dir, "add", relPath)
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("git add: %v\n%s", err, out)
	}
	cmd = exec.Command("git", "-C", dir, "commit", "-q", "-m", subject)
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("git commit: %v\n%s", err, out)
	}
	cmd = exec.Command("git", "-C", dir, "rev-parse", "HEAD")
	out, err := cmd.Output()
	if err != nil {
		t.Fatal(err)
	}
	return string(out[:len(out)-1]) // trim trailing newline
}

func TestScanNewCommitsFromScratch(t *testing.T) {
	dir := initTestRepo(t)
	commitFile(t, dir, "easy_problems/1_100/two_sum/main.go", "Easy(1) Two Sum")
	newest := commitFile(t, dir, "easy_problems/901_1000/univalued_binary_tree/main.go", "Easy(965) Univalued Binary Tree")

	entries, newestSHA, err := ScanNewCommits(dir, "")
	if err != nil {
		t.Fatalf("ScanNewCommits: %v", err)
	}
	if newestSHA != newest {
		t.Errorf("newestSHA = %q, want %q", newestSHA, newest)
	}
	if entries[1] != "easy_problems/1_100/two_sum" {
		t.Errorf("entries[1] = %q", entries[1])
	}
	if entries[965] != "easy_problems/901_1000/univalued_binary_tree" {
		t.Errorf("entries[965] = %q", entries[965])
	}
}

func TestScanNewCommitsIncremental(t *testing.T) {
	dir := initTestRepo(t)
	first := commitFile(t, dir, "easy_problems/1_100/two_sum/main.go", "Easy(1) Two Sum")
	commitFile(t, dir, "easy_problems/901_1000/univalued_binary_tree/main.go", "Easy(965) Univalued Binary Tree")

	entries, _, err := ScanNewCommits(dir, first)
	if err != nil {
		t.Fatalf("ScanNewCommits: %v", err)
	}
	if _, ok := entries[1]; ok {
		t.Errorf("entries should not include commit 1 (already scanned): %v", entries)
	}
	if entries[965] != "easy_problems/901_1000/univalued_binary_tree" {
		t.Errorf("entries[965] = %q", entries[965])
	}
}

func TestSaveLoadRoundTrip(t *testing.T) {
	path := filepath.Join(t.TempDir(), "number_folder_map.yaml")
	m := &Map{LastScannedCommit: "abc123", Entries: map[int]string{1: "easy_problems/1_100/two_sum"}}
	if err := m.Save(path); err != nil {
		t.Fatalf("Save: %v", err)
	}
	loaded, err := Load(path)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if loaded.LastScannedCommit != "abc123" || loaded.Entries[1] != "easy_problems/1_100/two_sum" {
		t.Fatalf("loaded = %+v", loaded)
	}
}

func TestLoadMissingFile(t *testing.T) {
	m, err := Load(filepath.Join(t.TempDir(), "nope.yaml"))
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if m.Entries == nil || len(m.Entries) != 0 {
		t.Errorf("Entries = %v, want empty map", m.Entries)
	}
}
```

- [ ] **Step 2: Run test to verify it fails**

Run: `go test ./challenge/internal/gitmap/...`
Expected: FAIL — package doesn't exist yet

- [ ] **Step 3: Write minimal implementation**

```go
// Package gitmap builds and maintains a LeetCode question number →
// solution folder map by scanning git commit history, following this
// repo's "Type(Number) Title" commit message convention (e.g.
// "Easy(965) Univalued Binary Tree").
package gitmap

import (
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

var commitSubjectRe = regexp.MustCompile(`(?i)^\s*(easy|medium|hard)\s*\(\s*(\d+)\s*\)`)

// ParseCommitSubject extracts the LeetCode question number from a commit
// subject line following the "Type(Number) Title" convention. ok is
// false if the subject doesn't match (many historical commits don't).
func ParseCommitSubject(subject string) (number int, ok bool) {
	m := commitSubjectRe.FindStringSubmatch(subject)
	if m == nil {
		return 0, false
	}
	n, err := strconv.Atoi(m[2])
	if err != nil {
		return 0, false
	}
	return n, true
}

// Map is the persisted contents of challenge/number_folder_map.yaml.
type Map struct {
	LastScannedCommit string         `yaml:"last_scanned_commit"`
	Entries           map[int]string `yaml:"map"`
}

// Load reads a Map from path. A missing file returns an empty Map, so a
// fresh repo needs no setup.
func Load(path string) (*Map, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return &Map{Entries: map[int]string{}}, nil
		}
		return nil, err
	}
	var m Map
	if err := yaml.Unmarshal(data, &m); err != nil {
		return nil, err
	}
	if m.Entries == nil {
		m.Entries = map[int]string{}
	}
	return &m, nil
}

// Save writes the Map to path as YAML.
func (m *Map) Save(path string) error {
	data, err := yaml.Marshal(m)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}

// ScanNewCommits walks git log in repoDir for commits after sinceSHA
// (exclusive; empty sinceSHA scans full history from the beginning),
// extracting question number → folder for every commit whose subject
// matches ParseCommitSubject and which touched at least one non-test
// .go file. It returns the new entries found and the newest commit SHA
// seen (for persisting as the next LastScannedCommit).
func ScanNewCommits(repoDir, sinceSHA string) (entries map[int]string, newestSHA string, err error) {
	rangeArg := "HEAD"
	if sinceSHA != "" {
		rangeArg = sinceSHA + "..HEAD"
	}
	cmd := exec.Command("git", "-C", repoDir, "log", "--reverse", "--pretty=format:%H%x1f%s", rangeArg)
	out, err := cmd.Output()
	if err != nil {
		return nil, sinceSHA, err
	}
	entries = map[int]string{}
	newestSHA = sinceSHA
	trimmed := strings.TrimSpace(string(out))
	if trimmed == "" {
		return entries, newestSHA, nil
	}
	for _, line := range strings.Split(trimmed, "\n") {
		parts := strings.SplitN(line, "\x1f", 2)
		if len(parts) != 2 {
			continue
		}
		sha, subject := parts[0], parts[1]
		newestSHA = sha
		number, ok := ParseCommitSubject(subject)
		if !ok {
			continue
		}
		folder, ok := folderForCommit(repoDir, sha)
		if !ok {
			continue
		}
		entries[number] = folder
	}
	return entries, newestSHA, nil
}

func folderForCommit(repoDir, sha string) (string, bool) {
	cmd := exec.Command("git", "-C", repoDir, "show", "--name-only", "--pretty=format:", sha)
	out, err := cmd.Output()
	if err != nil {
		return "", false
	}
	for _, f := range strings.Split(strings.TrimSpace(string(out)), "\n") {
		if strings.HasSuffix(f, ".go") && !strings.HasSuffix(f, "_test.go") {
			return filepath.Dir(f), true
		}
	}
	return "", false
}
```

- [ ] **Step 4: Run test to verify it passes**

Run: `go test ./challenge/internal/gitmap/...`
Expected: PASS

- [ ] **Step 5: Commit**

```bash
git add challenge/internal/gitmap
git commit -m "feat(challenge): add gitmap package to derive number->folder from commit history"
```

---

### Task 5: `resolver` — locate an already-solved question on disk

**Files:**
- Create: `challenge/internal/resolver/resolver.go`
- Test: `challenge/internal/resolver/resolver_test.go`

- [ ] **Step 1: Write the failing test**

```go
package resolver

import (
	"os"
	"path/filepath"
	"testing"
)

func mkdirs(t *testing.T, root string, dirs ...string) {
	t.Helper()
	for _, d := range dirs {
		if err := os.MkdirAll(filepath.Join(root, d), 0o755); err != nil {
			t.Fatal(err)
		}
	}
}

func TestCandidatePathsOrderAndSources(t *testing.T) {
	root := t.TempDir()
	mkdirs(t, root, "google_questions/design/min_stack")

	candidates := CandidatePaths(root, 155, "easy", "min-stack")
	if len(candidates) < 4 {
		t.Fatalf("expected at least 4 static candidates, got %d: %+v", len(candidates), candidates)
	}
	if candidates[0].Path != filepath.Join("easy_problems", "101_200", "min_stack") || candidates[0].Source != "canonical" {
		t.Errorf("candidates[0] = %+v", candidates[0])
	}
	if candidates[1].Path != filepath.Join("easy_problems", "min_stack") || candidates[1].Source != "straggler" {
		t.Errorf("candidates[1] = %+v", candidates[1])
	}
	found := false
	for _, c := range candidates {
		if c.Path == filepath.Join("google_questions", "design", "min_stack") && c.Source == "google_questions" {
			found = true
		}
	}
	if !found {
		t.Errorf("expected a google_questions candidate, got %+v", candidates)
	}
}

func TestFindPicksFirstExisting(t *testing.T) {
	root := t.TempDir()
	mkdirs(t, root, "easy_problems/min_stack") // only the straggler location exists

	candidates := CandidatePaths(root, 155, "easy", "min-stack")
	loc, ok := Find(root, candidates)
	if !ok {
		t.Fatal("Find: expected a match")
	}
	if loc.Source != "straggler" {
		t.Errorf("Find matched %+v, want straggler", loc)
	}
}

func TestFindNoMatch(t *testing.T) {
	root := t.TempDir()
	candidates := CandidatePaths(root, 155, "easy", "min-stack")
	if _, ok := Find(root, candidates); ok {
		t.Error("Find: expected no match on empty repo")
	}
}

func TestFuzzyFind(t *testing.T) {
	root := t.TempDir()
	// Folder name predates a LeetCode slug change, but content matches.
	mkdirs(t, root, "medium_problems/701_800/Min_Stack")

	loc, ok := FuzzyFind(root, "min-stack")
	if !ok {
		t.Fatal("FuzzyFind: expected a match")
	}
	if loc.Path != filepath.Join("medium_problems", "701_800", "Min_Stack") || loc.Source != "fuzzy" {
		t.Errorf("FuzzyFind matched %+v", loc)
	}
}
```

- [ ] **Step 2: Run test to verify it fails**

Run: `go test ./challenge/internal/resolver/...`
Expected: FAIL — package doesn't exist yet

- [ ] **Step 3: Write minimal implementation**

```go
// Package resolver locates where (if anywhere) an already-solved
// LeetCode question lives in the repo, checking the canonical location
// first and falling back through known non-canonical layouts.
package resolver

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"leetcode_solutions/challenge/internal/slugutil"
)

// Location is one place a solved question might live in the repo, with
// Path relative to the repo root.
type Location struct {
	Path      string
	Canonical bool
	Source    string // canonical | straggler | google_questions | linkedin_questions | gitmap | fuzzy
}

// CandidatePaths returns the static, filesystem-derivable places a
// solved question might live, in priority order. It does not consult
// git history or do fuzzy name matching — see FuzzyFind for that.
func CandidatePaths(repoRoot string, number int, difficulty, slug string) []Location {
	folder := slugutil.FolderName(slug)
	diffRoot := difficulty + "_problems"

	locs := []Location{
		{Path: filepath.Join(diffRoot, slugutil.RangeBucket(number), folder), Canonical: true, Source: "canonical"},
		{Path: filepath.Join(diffRoot, folder), Canonical: false, Source: "straggler"},
	}
	for _, topicRoot := range []string{"google_questions", "linkedin_questions"} {
		matches, _ := filepath.Glob(filepath.Join(repoRoot, topicRoot, "*", folder))
		for _, m := range matches {
			rel, err := filepath.Rel(repoRoot, m)
			if err != nil {
				continue
			}
			locs = append(locs, Location{Path: rel, Canonical: false, Source: topicRoot})
		}
	}
	return locs
}

// Find returns the first candidate that actually exists as a directory
// on disk.
func Find(repoRoot string, candidates []Location) (Location, bool) {
	for _, c := range candidates {
		if info, err := os.Stat(filepath.Join(repoRoot, c.Path)); err == nil && info.IsDir() {
			return c, true
		}
	}
	return Location{}, false
}

// FuzzyFind is the last-resort match: it walks every folder under the
// three difficulty roots (any depth) looking for a name matching slug's
// folder-normalized form, case-insensitively. Used when a repo folder
// predates a LeetCode title/slug change and none of the static
// candidates line up.
func FuzzyFind(repoRoot, slug string) (Location, bool) {
	target := strings.ToLower(slugutil.FolderName(slug))
	var found Location
	ok := false
	for _, root := range []string{"easy_problems", "medium_problems", "hard_problems"} {
		start := filepath.Join(repoRoot, root)
		if _, err := os.Stat(start); err != nil {
			continue
		}
		filepath.WalkDir(start, func(path string, d fs.DirEntry, err error) error {
			if err != nil || ok || !d.IsDir() {
				return nil
			}
			if strings.ToLower(d.Name()) == target {
				rel, relErr := filepath.Rel(repoRoot, path)
				if relErr == nil {
					found = Location{Path: rel, Canonical: false, Source: "fuzzy"}
					ok = true
				}
			}
			return nil
		})
		if ok {
			break
		}
	}
	return found, ok
}
```

- [ ] **Step 4: Run test to verify it passes**

Run: `go test ./challenge/internal/resolver/...`
Expected: PASS

- [ ] **Step 5: Commit**

```bash
git add challenge/internal/resolver
git commit -m "feat(challenge): add resolver package to locate solved questions on disk"
```

---

### Task 6: `reorg` — move a non-canonical solution into place

**Files:**
- Create: `challenge/internal/reorg/reorg.go`
- Test: `challenge/internal/reorg/reorg_test.go`

- [ ] **Step 1: Write the failing test**

```go
package reorg

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestTargetPath(t *testing.T) {
	got := TargetPath("easy", 965, "univalued-binary-tree")
	want := filepath.Join("easy_problems", "901_1000", "univalued_binary_tree")
	if got != want {
		t.Errorf("TargetPath = %q, want %q", got, want)
	}
}

func initTestRepo(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	run := func(args ...string) {
		cmd := exec.Command("git", args...)
		cmd.Dir = dir
		if out, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("git %v: %v\n%s", args, err, out)
		}
	}
	run("init", "-q")
	run("config", "user.email", "test@example.com")
	run("config", "user.name", "Test")
	return dir
}

func TestMoveRelocatesAndPreservesHistory(t *testing.T) {
	dir := initTestRepo(t)
	from := "google_questions/design/min_stack"
	if err := os.MkdirAll(filepath.Join(dir, from), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, from, "main.go"), []byte("package minstack\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	run := func(args ...string) {
		cmd := exec.Command("git", args...)
		cmd.Dir = dir
		if out, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("git %v: %v\n%s", args, err, out)
		}
	}
	run("add", ".")
	run("commit", "-q", "-m", "Easy(155) Min Stack")

	newPath, err := Move(dir, from, "easy", 155, "min-stack")
	if err != nil {
		t.Fatalf("Move: %v", err)
	}
	want := filepath.Join("easy_problems", "101_200", "min_stack")
	if newPath != want {
		t.Errorf("Move returned %q, want %q", newPath, want)
	}
	if _, err := os.Stat(filepath.Join(dir, from)); !os.IsNotExist(err) {
		t.Errorf("old path %q should no longer exist", from)
	}
	if _, err := os.Stat(filepath.Join(dir, want, "main.go")); err != nil {
		t.Errorf("new path missing main.go: %v", err)
	}
}

func TestMoveNoOpWhenAlreadyCanonical(t *testing.T) {
	dir := initTestRepo(t)
	canonical := filepath.Join("easy_problems", "101_200", "min_stack")
	newPath, err := Move(dir, canonical, "easy", 155, "min-stack")
	if err != nil {
		t.Fatalf("Move: %v", err)
	}
	if newPath != canonical {
		t.Errorf("Move = %q, want no-op %q", newPath, canonical)
	}
}
```

- [ ] **Step 2: Run test to verify it fails**

Run: `go test ./challenge/internal/reorg/...`
Expected: FAIL — package doesn't exist yet

- [ ] **Step 3: Write minimal implementation**

```go
// Package reorg relocates an already-solved question's folder from a
// non-canonical location (a root straggler, or one of the old
// topic-organized directories) into the canonical
// {difficulty}_problems/{range}/{slug} layout, preserving git history
// via `git mv`.
package reorg

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"leetcode_solutions/challenge/internal/slugutil"
)

// TargetPath returns the canonical destination path (relative to the
// repo root) a solved question should live at.
func TargetPath(difficulty string, number int, slug string) string {
	return filepath.Join(difficulty+"_problems", slugutil.RangeBucket(number), slugutil.FolderName(slug))
}

// Move git-mv's fromPath (relative to repoRoot) to the canonical
// location for (difficulty, number, slug), creating parent directories
// as needed, and returns the new relative path. If fromPath is already
// the canonical location, it's a no-op.
func Move(repoRoot, fromPath, difficulty string, number int, slug string) (string, error) {
	target := TargetPath(difficulty, number, slug)
	if fromPath == target {
		return target, nil
	}
	if err := os.MkdirAll(filepath.Join(repoRoot, filepath.Dir(target)), 0o755); err != nil {
		return "", err
	}
	cmd := exec.Command("git", "-C", repoRoot, "mv", fromPath, target)
	if out, err := cmd.CombinedOutput(); err != nil {
		return "", fmt.Errorf("git mv %s -> %s: %w: %s", fromPath, target, err, out)
	}
	return target, nil
}
```

- [ ] **Step 4: Run test to verify it passes**

Run: `go test ./challenge/internal/reorg/...`
Expected: PASS

- [ ] **Step 5: Commit**

```bash
git add challenge/internal/reorg
git commit -m "feat(challenge): add reorg package to git-mv solutions into canonical layout"
```

---

### Task 7: `leetcode` — GraphQL client for problem list + statement

**Files:**
- Create: `challenge/internal/leetcode/leetcode.go`
- Test: `challenge/internal/leetcode/leetcode_test.go`

- [ ] **Step 1: Write the failing test**

```go
package leetcode

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestParseFavoriteSlug(t *testing.T) {
	cases := map[string]string{
		"https://leetcode.com/problem-list/d0wm7eej/":  "d0wm7eej",
		"https://leetcode.com/problem-list/d0wm7eej":   "d0wm7eej",
		"https://leetcode.com/problem-list/abc123/?x=1": "abc123",
	}
	for url, want := range cases {
		got, err := ParseFavoriteSlug(url)
		if err != nil {
			t.Fatalf("ParseFavoriteSlug(%q): %v", url, err)
		}
		if got != want {
			t.Errorf("ParseFavoriteSlug(%q) = %q, want %q", url, got, want)
		}
	}
	if _, err := ParseFavoriteSlug("https://leetcode.com/problems/two-sum/"); err == nil {
		t.Error("expected error for a non-problem-list URL")
	}
}

func TestFetchProblemList(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Variables struct {
				FavoriteSlug string `json:"favoriteSlug"`
			} `json:"variables"`
		}
		json.NewDecoder(r.Body).Decode(&req)
		if req.Variables.FavoriteSlug != "d0wm7eej" {
			t.Errorf("favoriteSlug = %q", req.Variables.FavoriteSlug)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"favoriteQuestionList":{"questions":[
			{"questionFrontendId":"965","title":"Univalued Binary Tree","titleSlug":"univalued-binary-tree","difficulty":"Easy"},
			{"questionFrontendId":"1022","title":"Sum of Root To Leaf Binary Numbers","titleSlug":"sum-of-root-to-leaf-binary-numbers","difficulty":"Easy"}
		]}}}`))
	}))
	defer server.Close()

	client := &Client{BaseURL: server.URL, HTTPClient: server.Client()}
	problems, err := client.FetchProblemList("d0wm7eej", "easy")
	if err != nil {
		t.Fatalf("FetchProblemList: %v", err)
	}
	if len(problems) != 2 {
		t.Fatalf("got %d problems, want 2", len(problems))
	}
	if problems[0].Number != 965 || problems[0].Slug != "univalued-binary-tree" || problems[0].Difficulty != "easy" {
		t.Errorf("problems[0] = %+v", problems[0])
	}
}

func TestFetchQuestionContent(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"question":{"title":"Univalued Binary Tree","titleSlug":"univalued-binary-tree","difficulty":"Easy","content":"<p>A binary tree is univalued...</p>"}}}`))
	}))
	defer server.Close()

	client := &Client{BaseURL: server.URL, HTTPClient: server.Client()}
	q, err := client.FetchQuestionContent("univalued-binary-tree")
	if err != nil {
		t.Fatalf("FetchQuestionContent: %v", err)
	}
	if q.Title != "Univalued Binary Tree" || q.Difficulty != "easy" || q.ContentHTML == "" {
		t.Errorf("q = %+v", q)
	}
}
```

- [ ] **Step 2: Run test to verify it fails**

Run: `go test ./challenge/internal/leetcode/...`
Expected: FAIL — package doesn't exist yet

- [ ] **Step 3: Write minimal implementation**

```go
// Package leetcode is a minimal client for LeetCode's public (unofficial,
// undocumented) GraphQL API: fetching a custom problem list and a single
// question's statement. No authentication — Premium-only fields (like
// official company tags) are not available through this client.
package leetcode

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Problem is one entry from a LeetCode custom problem list.
type Problem struct {
	Number     int
	Title      string
	Slug       string
	Difficulty string
}

// Question is a single question's statement content.
type Question struct {
	Title       string
	Slug        string
	Difficulty  string
	ContentHTML string
}

// Client talks to LeetCode's GraphQL endpoint.
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

// NewClient returns a Client pointed at LeetCode's real GraphQL endpoint.
func NewClient() *Client {
	return &Client{BaseURL: "https://leetcode.com/graphql", HTTPClient: http.DefaultClient}
}

// ParseFavoriteSlug extracts the list id from a LeetCode problem-list
// URL, e.g. "https://leetcode.com/problem-list/d0wm7eej/" -> "d0wm7eej".
func ParseFavoriteSlug(rawURL string) (string, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	parts := strings.Split(strings.Trim(u.Path, "/"), "/")
	for i, p := range parts {
		if p == "problem-list" && i+1 < len(parts) && parts[i+1] != "" {
			return parts[i+1], nil
		}
	}
	return "", fmt.Errorf("not a problem-list URL: %s", rawURL)
}

// FetchProblemList fetches every question in a LeetCode custom problem
// list, optionally filtered by difficulty ("" for all difficulties).
func (c *Client) FetchProblemList(favoriteSlug, difficulty string) ([]Problem, error) {
	query := `query favoriteQuestionList($favoriteSlug: String!, $filter: FavoriteQuestionFilterInput) {
	  favoriteQuestionList(favoriteSlug: $favoriteSlug, filter: $filter) {
	    questions {
	      questionFrontendId
	      title
	      titleSlug
	      difficulty
	    }
	  }
	}`
	filter := map[string]any{}
	if difficulty != "" {
		filter["difficulty"] = strings.ToUpper(difficulty)
	}
	variables := map[string]any{"favoriteSlug": favoriteSlug, "filter": filter}

	var resp struct {
		Data struct {
			FavoriteQuestionList struct {
				Questions []struct {
					QuestionFrontendID string `json:"questionFrontendId"`
					Title              string `json:"title"`
					TitleSlug          string `json:"titleSlug"`
					Difficulty         string `json:"difficulty"`
				} `json:"questions"`
			} `json:"favoriteQuestionList"`
		} `json:"data"`
	}
	if err := c.doGraphQL(query, variables, &resp); err != nil {
		return nil, err
	}
	problems := make([]Problem, 0, len(resp.Data.FavoriteQuestionList.Questions))
	for _, q := range resp.Data.FavoriteQuestionList.Questions {
		n, err := strconv.Atoi(q.QuestionFrontendID)
		if err != nil {
			continue
		}
		problems = append(problems, Problem{
			Number:     n,
			Title:      q.Title,
			Slug:       q.TitleSlug,
			Difficulty: strings.ToLower(q.Difficulty),
		})
	}
	return problems, nil
}

// FetchQuestionContent fetches a single question's statement by slug.
func (c *Client) FetchQuestionContent(slug string) (Question, error) {
	query := `query questionContent($titleSlug: String!) {
	  question(titleSlug: $titleSlug) {
	    title
	    titleSlug
	    difficulty
	    content
	  }
	}`
	variables := map[string]any{"titleSlug": slug}

	var resp struct {
		Data struct {
			Question struct {
				Title      string `json:"title"`
				TitleSlug  string `json:"titleSlug"`
				Difficulty string `json:"difficulty"`
				Content    string `json:"content"`
			} `json:"question"`
		} `json:"data"`
	}
	if err := c.doGraphQL(query, variables, &resp); err != nil {
		return Question{}, err
	}
	q := resp.Data.Question
	return Question{
		Title:       q.Title,
		Slug:        q.TitleSlug,
		Difficulty:  strings.ToLower(q.Difficulty),
		ContentHTML: q.Content,
	}, nil
}

func (c *Client) doGraphQL(query string, variables map[string]any, out any) error {
	body, err := json.Marshal(map[string]any{"query": query, "variables": variables})
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, c.BaseURL, bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("leetcode graphql: unexpected status %d", resp.StatusCode)
	}
	return json.NewDecoder(resp.Body).Decode(out)
}
```

- [ ] **Step 4: Run test to verify it passes**

Run: `go test ./challenge/internal/leetcode/...`
Expected: PASS

- [ ] **Step 5: Commit**

```bash
git add challenge/internal/leetcode
git commit -m "feat(challenge): add leetcode GraphQL client for problem list + question content"
```

---

### Task 8: `hero` — hero-image HTML templating

**Files:**
- Create: `challenge/internal/hero/hero.go`
- Create: `challenge/hero_template.html`
- Test: `challenge/internal/hero/hero_test.go`

- [ ] **Step 1: Write the failing test**

```go
package hero

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRenderHTML(t *testing.T) {
	tmplPath := filepath.Join(t.TempDir(), "template.html")
	tmpl := `<h1>Day {{.Day}} / {{.Total}}</h1><h2>{{.Title}}</h2><p>{{.Topic}} - {{.Difficulty}}</p>`
	if err := os.WriteFile(tmplPath, []byte(tmpl), 0o644); err != nil {
		t.Fatal(err)
	}

	html, err := RenderHTML(tmplPath, Data{
		Day: 23, Total: 365, Topic: "Binary Tree", Difficulty: "Easy", Title: "Univalued Binary Tree",
	})
	if err != nil {
		t.Fatalf("RenderHTML: %v", err)
	}
	for _, want := range []string{"Day 23 / 365", "Univalued Binary Tree", "Binary Tree - Easy"} {
		if !strings.Contains(html, want) {
			t.Errorf("rendered HTML missing %q, got: %s", want, html)
		}
	}
}

func TestRenderHTMLRealTemplate(t *testing.T) {
	html, err := RenderHTML("../../hero_template.html", Data{
		Day: 1, Total: 365, Topic: "Arrays", Difficulty: "Easy", Title: "Two Sum",
	})
	if err != nil {
		t.Fatalf("RenderHTML with real template: %v", err)
	}
	for _, want := range []string{"1", "365", "Arrays", "Easy", "Two Sum"} {
		if !strings.Contains(html, want) {
			t.Errorf("rendered real template missing %q", want)
		}
	}
}
```

- [ ] **Step 2: Run test to verify it fails**

Run: `go test ./challenge/internal/hero/...`
Expected: FAIL — package and template file don't exist yet

- [ ] **Step 3: Write the implementation**

```go
// Package hero renders the "365 Days of LeetCode Challenge" hero image
// shown at the top of each LinkedIn/Discord post: an HTML template is
// filled with the day counter, topic, difficulty and title, then
// screenshotted to a PNG by an external headless browser.
package hero

import (
	"bytes"
	"os/exec"
	"path/filepath"
	"text/template"
)

// Data is the set of placeholders the hero template fills in.
type Data struct {
	Day        int
	Total      int
	Topic      string
	Difficulty string
	Title      string
}

// RenderHTML fills the template at templatePath with data and returns
// the rendered HTML.
func RenderHTML(templatePath string, data Data) (string, error) {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// Screenshot renders htmlPath (a local file) to a PNG at outPath using a
// headless Chromium instance via Playwright's CLI. Requires
// `npx playwright install chromium` to have been run once on the
// machine — see challenge/README.md. Not covered by automated tests
// since it depends on an external browser binary; verify manually.
func Screenshot(htmlPath, outPath string) error {
	absHTML, err := filepath.Abs(htmlPath)
	if err != nil {
		return err
	}
	cmd := exec.Command("npx", "playwright", "screenshot",
		"--viewport-size=1200,630",
		"file://"+absHTML, outPath)
	return cmd.Run()
}
```

- [ ] **Step 4: Write the hero template**

`challenge/hero_template.html`:
```html
<!doctype html>
<html>
<head>
<meta charset="utf-8">
<style>
  body { margin: 0; width: 1200px; height: 630px; font-family: -apple-system, "Segoe UI", Helvetica, Arial, sans-serif; }
  .card {
    width: 1200px; height: 630px; box-sizing: border-box; padding: 64px;
    background: linear-gradient(135deg, #1a1a2e 0%, #16213e 100%);
    color: #ffffff; display: flex; flex-direction: column; justify-content: space-between;
  }
  .brand { display: flex; align-items: center; gap: 16px; }
  .badge {
    width: 56px; height: 56px; border-radius: 12px; background: #ffa116;
    display: flex; align-items: center; justify-content: center;
    font-weight: 800; font-size: 22px; color: #1a1a2e;
  }
  .challenge-name { font-size: 20px; letter-spacing: 2px; text-transform: uppercase; color: #ffa116; }
  .day { font-size: 96px; font-weight: 800; margin: 24px 0 0 0; }
  .day small { font-size: 40px; font-weight: 500; color: #cccccc; }
  .meta { font-size: 28px; color: #ffa116; text-transform: uppercase; letter-spacing: 1px; margin-bottom: 12px; }
  .title { font-size: 48px; font-weight: 700; line-height: 1.2; }
</style>
</head>
<body>
  <div class="card">
    <div class="brand">
      <div class="badge">LC</div>
      <div class="challenge-name">365 Days of LeetCode Challenge</div>
    </div>
    <div>
      <div class="day">Day {{.Day}}<small> / {{.Total}}</small></div>
      <div class="meta" style="margin-top: 32px;">{{.Topic}} &middot; {{.Difficulty}}</div>
      <div class="title">{{.Title}}</div>
    </div>
  </div>
</body>
</html>
```

Note: the badge is a stylized "LC" mark rather than a reproduction of LeetCode's actual logo artwork, to avoid vendoring a third-party trademarked image asset. Swap in the real logo file later if desired by adding an `<img>` tag pointing at a vendored asset.

- [ ] **Step 5: Run tests to verify they pass**

Run: `go test ./challenge/internal/hero/...`
Expected: PASS

- [ ] **Step 6: Commit**

```bash
git add challenge/internal/hero challenge/hero_template.html
git commit -m "feat(challenge): add hero package and hero-image HTML template"
```

---

### Task 9: `cli` — orchestration functions behind the CLI

**Files:**
- Create: `challenge/internal/cli/cli.go`
- Test: `challenge/internal/cli/cli_test.go`

- [ ] **Step 1: Write the failing test**

```go
package cli

import (
	"os"
	"path/filepath"
	"testing"

	"leetcode_solutions/challenge/internal/queue"
)

func TestResolveFindsCanonical(t *testing.T) {
	root := t.TempDir()
	target := filepath.Join(root, "easy_problems", "901_1000", "univalued_binary_tree")
	if err := os.MkdirAll(target, 0o755); err != nil {
		t.Fatal(err)
	}
	mapPath := filepath.Join(root, "challenge", "number_folder_map.yaml")

	result, err := Resolve(root, mapPath, 965, "easy", "univalued-binary-tree")
	if err != nil {
		t.Fatalf("Resolve: %v", err)
	}
	if result.Status != "resolved" || result.Source != "canonical" {
		t.Errorf("result = %+v", result)
	}
}

func TestResolveUnresolved(t *testing.T) {
	root := t.TempDir()
	mapPath := filepath.Join(root, "challenge", "number_folder_map.yaml")

	result, err := Resolve(root, mapPath, 999999, "easy", "does-not-exist")
	if err != nil {
		t.Fatalf("Resolve: %v", err)
	}
	if result.Status != "unresolved" {
		t.Errorf("result = %+v, want unresolved", result)
	}
}

func TestQueueAppendAndHas(t *testing.T) {
	queuePath := filepath.Join(t.TempDir(), "queue.yaml")

	has, err := QueueHas(queuePath, 965)
	if err != nil {
		t.Fatalf("QueueHas: %v", err)
	}
	if has {
		t.Error("QueueHas on empty queue = true, want false")
	}

	day, err := QueueAppend(queuePath, queue.Entry{
		Number: 965, Title: "Univalued Binary Tree", Difficulty: "easy",
		Folder: "easy_problems/901_1000/univalued_binary_tree", Batch: "binary tree easy problems",
	})
	if err != nil {
		t.Fatalf("QueueAppend: %v", err)
	}
	if day != 1 {
		t.Errorf("day = %d, want 1", day)
	}

	has, err = QueueHas(queuePath, 965)
	if err != nil {
		t.Fatalf("QueueHas: %v", err)
	}
	if !has {
		t.Error("QueueHas after append = false, want true")
	}
}

func TestCompaniesLookup(t *testing.T) {
	path := filepath.Join(t.TempDir(), "dataset.json")
	if err := os.WriteFile(path, []byte(`{"965":["Google"]}`), 0o644); err != nil {
		t.Fatal(err)
	}
	got, err := CompaniesLookup(path, 965)
	if err != nil {
		t.Fatalf("CompaniesLookup: %v", err)
	}
	if len(got) != 1 || got[0] != "Google" {
		t.Errorf("got = %v", got)
	}
}

func TestHeroRenderHTML(t *testing.T) {
	tmplPath := filepath.Join(t.TempDir(), "t.html")
	if err := os.WriteFile(tmplPath, []byte(`Day {{.Day}}: {{.Title}}`), 0o644); err != nil {
		t.Fatal(err)
	}
	outPath := filepath.Join(t.TempDir(), "out.html")

	if err := HeroRenderHTML(tmplPath, heroDataFixture(), outPath); err != nil {
		t.Fatalf("HeroRenderHTML: %v", err)
	}
	data, err := os.ReadFile(outPath)
	if err != nil {
		t.Fatalf("read output: %v", err)
	}
	if string(data) != "Day 1: Two Sum" {
		t.Errorf("output = %q", data)
	}
}
```

Add this small fixture helper at the bottom of the same test file:

```go
func heroDataFixture() (d struct {
	Day        int
	Total      int
	Topic      string
	Difficulty string
	Title      string
}) {
	d.Day = 1
	d.Title = "Two Sum"
	return d
}
```

- [ ] **Step 2: Run test to verify it fails**

Run: `go test ./challenge/internal/cli/...`
Expected: FAIL — package doesn't exist yet, and the anonymous struct fixture won't match `hero.Data`'s type

Fix the fixture before implementing: replace the anonymous-struct `heroDataFixture` with one that returns `hero.Data` directly, so it's usable in `HeroRenderHTML`'s call:

```go
func heroDataFixture() hero.Data {
	return hero.Data{Day: 1, Title: "Two Sum"}
}
```

(add `"leetcode_solutions/challenge/internal/hero"` to the test file's imports)

- [ ] **Step 3: Write the implementation**

```go
// Package cli holds the orchestration logic behind each leetcodectl
// subcommand, kept separate from cmd/leetcodectl/main.go so it's
// testable without spawning a subprocess.
package cli

import (
	"os"

	"leetcode_solutions/challenge/internal/companies"
	"leetcode_solutions/challenge/internal/gitmap"
	"leetcode_solutions/challenge/internal/hero"
	"leetcode_solutions/challenge/internal/leetcode"
	"leetcode_solutions/challenge/internal/queue"
	"leetcode_solutions/challenge/internal/reorg"
	"leetcode_solutions/challenge/internal/resolver"
)

// ResolveResult is the outcome of trying to locate a solved question.
type ResolveResult struct {
	Status    string `json:"status"` // resolved | unresolved
	Path      string `json:"path,omitempty"`
	Canonical bool   `json:"canonical,omitempty"`
	Source    string `json:"source,omitempty"`
}

// Resolve locates a solved question on disk: static candidates first,
// then the git-history-derived number map, then a fuzzy name match.
//
// Note: resolver.Location.Source is a resolver.Source, a defined string
// type — assigning it to ResolveResult's plain string field requires an
// explicit conversion (string(loc.Source)), it will not compile as a bare
// assignment.
func Resolve(repoRoot, mapPath string, number int, difficulty, slug string) (ResolveResult, error) {
	candidates := resolver.CandidatePaths(repoRoot, number, difficulty, slug)
	if loc, ok := resolver.Find(repoRoot, candidates); ok {
		return ResolveResult{Status: "resolved", Path: loc.Path, Canonical: loc.Canonical, Source: string(loc.Source)}, nil
	}

	m, err := gitmap.Load(mapPath)
	if err != nil {
		return ResolveResult{}, err
	}
	if folder, ok := m.Entries[number]; ok {
		if info, statErr := os.Stat(repoRoot + "/" + folder); statErr == nil && info.IsDir() {
			return ResolveResult{Status: "resolved", Path: folder, Source: string(resolver.SourceGitmap)}, nil
		}
	}

	if loc, ok := resolver.FuzzyFind(repoRoot, slug); ok {
		return ResolveResult{Status: "resolved", Path: loc.Path, Source: string(loc.Source)}, nil
	}

	return ResolveResult{Status: "unresolved"}, nil
}

// GitmapUpdate incrementally rescans git history for new
// "Type(Number) Title" commits and persists them into mapPath.
func GitmapUpdate(repoRoot, mapPath string) (updated int, newestSHA string, err error) {
	m, err := gitmap.Load(mapPath)
	if err != nil {
		return 0, "", err
	}
	newEntries, newest, err := gitmap.ScanNewCommits(repoRoot, m.LastScannedCommit)
	if err != nil {
		return 0, "", err
	}
	for number, folder := range newEntries {
		m.Entries[number] = folder
	}
	if newest != "" {
		m.LastScannedCommit = newest
	}
	if err := m.Save(mapPath); err != nil {
		return 0, "", err
	}
	return len(newEntries), m.LastScannedCommit, nil
}

// Reorg moves a non-canonically-located solution into its canonical
// folder.
func Reorg(repoRoot, fromPath, difficulty string, number int, slug string) (string, error) {
	return reorg.Move(repoRoot, fromPath, difficulty, number, slug)
}

// QueueHas reports whether the queue already has an entry for number.
func QueueHas(queuePath string, number int) (bool, error) {
	q, err := queue.Load(queuePath)
	if err != nil {
		return false, err
	}
	return q.Has(number), nil
}

// QueueAppend appends e to the queue, assigning it the next day number,
// and returns that day number.
func QueueAppend(queuePath string, e queue.Entry) (int, error) {
	q, err := queue.Load(queuePath)
	if err != nil {
		return 0, err
	}
	day := q.Append(e)
	if err := q.Save(queuePath); err != nil {
		return 0, err
	}
	return day, nil
}

// CompaniesLookup returns the companies known for a question number
// from the vendored dataset (empty slice if none known).
func CompaniesLookup(datasetPath string, number int) ([]string, error) {
	ds, err := companies.Load(datasetPath)
	if err != nil {
		return nil, err
	}
	return ds.Lookup(number), nil
}

// HeroRenderHTML renders the hero template with data and writes it to
// outPath.
func HeroRenderHTML(templatePath string, data hero.Data, outPath string) error {
	html, err := hero.RenderHTML(templatePath, data)
	if err != nil {
		return err
	}
	return os.WriteFile(outPath, []byte(html), 0o644)
}

// FetchList fetches a LeetCode problem list by URL, filtered by
// difficulty. Thin pass-through over the leetcode package (which has
// its own tests against a fake server) — not independently unit tested.
func FetchList(listURL, difficulty string) ([]leetcode.Problem, error) {
	favoriteSlug, err := leetcode.ParseFavoriteSlug(listURL)
	if err != nil {
		return nil, err
	}
	return leetcode.NewClient().FetchProblemList(favoriteSlug, difficulty)
}

// FetchQuestion fetches a single question's statement by slug. Thin
// pass-through, see FetchList's note.
func FetchQuestion(slug string) (leetcode.Question, error) {
	return leetcode.NewClient().FetchQuestionContent(slug)
}
```

- [ ] **Step 4: Run tests to verify they pass**

Run: `go test ./challenge/internal/cli/...`
Expected: PASS

- [ ] **Step 5: Commit**

```bash
git add challenge/internal/cli
git commit -m "feat(challenge): add cli package orchestrating resolve/reorg/queue/hero/fetch"
```

---

### Task 10: `leetcodectl` binary

**Files:**
- Create: `challenge/cmd/leetcodectl/main.go`

- [ ] **Step 1: Write the implementation**

```go
// Command leetcodectl is the deterministic-logic half of the
// /leetcode-content Claude Code skill: each subcommand takes a single
// JSON argument and prints a single JSON result, so the skill can drive
// it without any argv-flag parsing.
package main

import (
	"encoding/json"
	"fmt"
	"os"

	"leetcode_solutions/challenge/internal/cli"
	"leetcode_solutions/challenge/internal/hero"
	"leetcode_solutions/challenge/internal/queue"
)

func main() {
	if len(os.Args) < 2 {
		fail(fmt.Errorf("usage: leetcodectl <command> [json-args]"))
	}
	cmd := os.Args[1]
	payload := []byte("{}")
	if len(os.Args) > 2 {
		payload = []byte(os.Args[2])
	}

	result, err := dispatch(cmd, payload)
	if err != nil {
		fail(err)
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(result); err != nil {
		fail(err)
	}
}

func dispatch(cmd string, payload []byte) (any, error) {
	switch cmd {
	case "resolve":
		var in struct {
			RepoRoot   string `json:"repoRoot"`
			MapPath    string `json:"mapPath"`
			Number     int    `json:"number"`
			Difficulty string `json:"difficulty"`
			Slug       string `json:"slug"`
		}
		if err := json.Unmarshal(payload, &in); err != nil {
			return nil, err
		}
		return cli.Resolve(in.RepoRoot, in.MapPath, in.Number, in.Difficulty, in.Slug)

	case "reorg":
		var in struct {
			RepoRoot   string `json:"repoRoot"`
			FromPath   string `json:"fromPath"`
			Difficulty string `json:"difficulty"`
			Number     int    `json:"number"`
			Slug       string `json:"slug"`
		}
		if err := json.Unmarshal(payload, &in); err != nil {
			return nil, err
		}
		newPath, err := cli.Reorg(in.RepoRoot, in.FromPath, in.Difficulty, in.Number, in.Slug)
		if err != nil {
			return nil, err
		}
		return map[string]string{"path": newPath}, nil

	case "queue-has":
		var in struct {
			QueuePath string `json:"queuePath"`
			Number    int    `json:"number"`
		}
		if err := json.Unmarshal(payload, &in); err != nil {
			return nil, err
		}
		has, err := cli.QueueHas(in.QueuePath, in.Number)
		if err != nil {
			return nil, err
		}
		return map[string]bool{"has": has}, nil

	case "queue-append":
		var in struct {
			QueuePath string      `json:"queuePath"`
			Entry     queue.Entry `json:"entry"`
		}
		if err := json.Unmarshal(payload, &in); err != nil {
			return nil, err
		}
		day, err := cli.QueueAppend(in.QueuePath, in.Entry)
		if err != nil {
			return nil, err
		}
		return map[string]int{"day": day}, nil

	case "companies-lookup":
		var in struct {
			DatasetPath string `json:"datasetPath"`
			Number      int    `json:"number"`
		}
		if err := json.Unmarshal(payload, &in); err != nil {
			return nil, err
		}
		list, err := cli.CompaniesLookup(in.DatasetPath, in.Number)
		if err != nil {
			return nil, err
		}
		return map[string][]string{"companies": list}, nil

	case "gitmap-update":
		var in struct {
			RepoRoot string `json:"repoRoot"`
			MapPath  string `json:"mapPath"`
		}
		if err := json.Unmarshal(payload, &in); err != nil {
			return nil, err
		}
		updated, newest, err := cli.GitmapUpdate(in.RepoRoot, in.MapPath)
		if err != nil {
			return nil, err
		}
		return map[string]any{"updated": updated, "newestCommit": newest}, nil

	case "hero-render-html":
		var in struct {
			TemplatePath string    `json:"templatePath"`
			OutPath      string    `json:"outPath"`
			Data         hero.Data `json:"data"`
		}
		if err := json.Unmarshal(payload, &in); err != nil {
			return nil, err
		}
		if err := cli.HeroRenderHTML(in.TemplatePath, in.Data, in.OutPath); err != nil {
			return nil, err
		}
		return map[string]string{"outPath": in.OutPath}, nil

	case "hero-screenshot":
		var in struct {
			HTMLPath string `json:"htmlPath"`
			OutPath  string `json:"outPath"`
		}
		if err := json.Unmarshal(payload, &in); err != nil {
			return nil, err
		}
		if err := hero.Screenshot(in.HTMLPath, in.OutPath); err != nil {
			return nil, err
		}
		return map[string]string{"outPath": in.OutPath}, nil

	case "fetch-list":
		var in struct {
			URL        string `json:"url"`
			Difficulty string `json:"difficulty"`
		}
		if err := json.Unmarshal(payload, &in); err != nil {
			return nil, err
		}
		return cli.FetchList(in.URL, in.Difficulty)

	case "fetch-question":
		var in struct {
			Slug string `json:"slug"`
		}
		if err := json.Unmarshal(payload, &in); err != nil {
			return nil, err
		}
		return cli.FetchQuestion(in.Slug)

	default:
		return nil, fmt.Errorf("unknown command %q", cmd)
	}
}

func fail(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
```

- [ ] **Step 2: Build it**

Run: `go build -o /tmp/leetcodectl ./challenge/cmd/leetcodectl`
Expected: no output, exit code 0

- [ ] **Step 3: Manually verify a couple of subcommands end-to-end**

Run: `/tmp/leetcodectl resolve '{"repoRoot":".","mapPath":"challenge/number_folder_map.yaml","number":965,"difficulty":"easy","slug":"univalued-binary-tree"}'`
Expected (note `canonical: true` — the `omitempty` tag only drops it when `false`):
```json
{
  "status": "resolved",
  "path": "easy_problems/901_1000/univalued_binary_tree",
  "canonical": true,
  "source": "canonical"
}
```

Run: `/tmp/leetcodectl queue-has '{"queuePath":"challenge/queue.yaml","number":1}'`
Expected: `{"has": false}` (queue.yaml doesn't exist yet — created in Task 11)

- [ ] **Step 4: Commit**

```bash
git add challenge/cmd/leetcodectl
git commit -m "feat(challenge): add leetcodectl CLI binary wiring cli package to JSON subcommands"
```

---

### Task 11: Seed the shared state files

**Files:**
- Create: `challenge/queue.yaml`
- Create: `challenge/number_folder_map.yaml`

- [ ] **Step 1: Create the empty queue**

`challenge/queue.yaml`:
```yaml
next_day: 1
entries: []
```

- [ ] **Step 2: Create the empty git-history cache**

`challenge/number_folder_map.yaml`:
```yaml
last_scanned_commit: ""
map: {}
```

- [ ] **Step 3: Verify `leetcodectl` reads them correctly**

Run: `/tmp/leetcodectl queue-has '{"queuePath":"challenge/queue.yaml","number":1}'`
Expected: `{"has": false}`

Run: `/tmp/leetcodectl gitmap-update '{"repoRoot":".","mapPath":"challenge/number_folder_map.yaml"}'`
Expected: JSON with `"updated"` equal to however many `Type(Number) Title`-style commits exist in this repo's history, and a non-empty `"newestCommit"`.

- [ ] **Step 4: Commit**

```bash
git add challenge/queue.yaml challenge/number_folder_map.yaml
git commit -m "feat(challenge): seed empty queue.yaml and number_folder_map.yaml"
```

---

### Task 12: `/leetcode-content` Claude Code skill

**Files:**
- Create: `.claude/skills/leetcode-content/SKILL.md`

- [ ] **Step 1: Write the skill**

`.claude/skills/leetcode-content/SKILL.md`:
```markdown
---
name: leetcode-content
description: 'Generate teaching content (intuition/solution docs, commented code, companies, hero image, LinkedIn/Discord post drafts) for already-solved LeetCode questions from a problem-list URL, in resumable "365 Days of LeetCode Challenge" order. Trigger: /leetcode-content'
---

# LeetCode Content-Gen Agent

## Invocation

```
/leetcode-content <leetcode-problem-list-url> --filter difficulty=<easy|medium|hard> --list-name "<batch name>"
```

Design doc: `docs/superpowers/specs/2026-08-30-leetcode-content-gen-agent-design.md`

All `leetcodectl` calls below assume it's built once per session:

```bash
go build -o /tmp/leetcodectl ./challenge/cmd/leetcodectl
```

Every subcommand takes one JSON argument and prints one JSON result to stdout, e.g.:

```bash
/tmp/leetcodectl resolve '{"repoRoot":".","mapPath":"challenge/number_folder_map.yaml","number":965,"difficulty":"easy","slug":"univalued-binary-tree"}'
```

## Steps

1. **Refresh the git-history cache, and commit the refresh immediately, unconditionally —
   don't wait for the first per-question commit in step 3 to sweep it up.** If every
   question in this run turns out to be already-queued or not-yet-solved, the loop in
   step 3 never reaches a commit at all, and a rescanned `number_folder_map.yaml` would
   otherwise sit as an uncommitted working-tree change with nothing to notice it:
   ```bash
   /tmp/leetcodectl gitmap-update '{"repoRoot":".","mapPath":"challenge/number_folder_map.yaml"}'
   git add challenge/number_folder_map.yaml
   git diff --cached --quiet challenge/number_folder_map.yaml || git commit -m "chore(challenge): refresh git-history cache"
   ```

2. **Fetch the problem list:**
   ```bash
   /tmp/leetcodectl fetch-list '{"url":"<the URL>","difficulty":"<the difficulty value parsed out of --filter, e.g. --filter difficulty=easy -> \"easy\">"}'
   ```
   Sort the returned problems by `Number` ascending.

3. **For each problem, in that sorted order:**

   If any `leetcodectl` call in this loop (other than the two failure modes explicitly
   handled below — unresolved `resolve`, failed `hero-screenshot`) exits non-zero or
   returns unparseable output: stop processing THIS problem only, note it in the
   end-of-run summary with the error message, and continue to the next problem in the
   list. Never let one problem's failure abort the whole batch, and never silently
   proceed past a failed call as if it had succeeded (e.g. don't fabricate a folder path
   or day number when a call that was supposed to produce one failed).

   **Post-`queue-append` failures need special handling, since `queue-append` (step i)
   is the point of no return.** Once it succeeds, `queue-has` will report this number as
   already queued on every future run — there is no "un-append" operation. So a failure
   in any of steps (j) through (m) — hero rendering (other than the already-handled
   `hero-screenshot` case, which is non-fatal by design), writing the post drafts, or the
   final `git commit` — leaves a permanently "claimed" queue entry with incomplete or
   missing content, which a future re-run will silently skip forever rather than retry.
   If this happens: do NOT treat it as a normal per-problem skip. Call it out prominently
   and separately in the end-of-run summary (e.g. "Day N / question NUMBER was queued but
   its commit failed — needs manual follow-up: <error>"), so the user knows to
   investigate and finish that entry by hand rather than assuming a clean re-run will
   pick it up.

   a. Skip it if `queue-has` returns `{"has": true}` (the subcommand's JSON result, not a bare `true`):
      ```bash
      /tmp/leetcodectl queue-has '{"queuePath":"challenge/queue.yaml","number":<number>}'
      ```

   b. Resolve its location:
      ```bash
      /tmp/leetcodectl resolve '{"repoRoot":".","mapPath":"challenge/number_folder_map.yaml","number":<number>,"difficulty":"<difficulty>","slug":"<slug>"}'
      ```
      - If `"status": "unresolved"` — this question isn't solved yet (or genuinely can't be
        found). Skip it silently and move to the next problem; it'll be picked up on a future
        re-run once solved. Do NOT ask the user about every unresolved question — only ask if
        you have a specific, well-founded suspicion it IS solved but the tooling missed it
        (e.g. you can see a matching file via other means).
      - If `"source"` is anything other than `"canonical"` — the question was found in a
        non-canonical location (a root straggler, or under `google_questions`/`linkedin_questions`,
        or via the gitmap/fuzzy fallback). Note: `ResolveResult.Canonical` is tagged
        `json:"canonical,omitempty"`, so a non-canonical result never actually contains a
        `"canonical":false` key — the key is simply absent. Check `source`, not the
        presence/value of `canonical`. Reorganize it:
        ```bash
        /tmp/leetcodectl reorg '{"repoRoot":".","fromPath":"<result.path>","difficulty":"<difficulty>","number":<number>,"slug":"<slug>"}'
        ```
        Use the returned `path` as the folder for the rest of this step.

   c. Read the existing solution at `<folder>/main.go` (and `main_test.go` if present) —
      this is the source of truth for INTUITION.md, SOLUTION.md, and the inline comments.
      Never rewrite the algorithm; only add commentary.

   d. If `<folder>/README.md` is missing, fetch the statement and write it:
      ```bash
      /tmp/leetcodectl fetch-question '{"slug":"<slug>"}'
      ```
      Write `<folder>/README.md` from the returned `Title`/`Difficulty`/`ContentHTML`,
      converting the HTML statement to clean Markdown (headings, code blocks for
      examples, a Constraints list) — follow the style of existing README.md files
      elsewhere in the repo (e.g. `easy_problems/1_100/climbing_stairs/README.md`).

      `ContentHTML` may contain `<img>` tags (tree/graph diagrams, matrix illustrations,
      etc. — common on LeetCode statements). For each one: download it to
      `<folder>/images/<n>.<ext>` (`<ext>` from the URL or `Content-Type`, `<n>` a
      1-based index in the order the images appear) with
      `curl -sL '<image-url>' -o '<folder>/images/<n>.<ext>'`, and rewrite that image's
      markdown reference to the local relative path (`images/<n>.<ext>`) instead of the
      remote URL. Do not leave a remote LeetCode CDN URL in the committed README — this
      repo should be self-contained and not depend on LeetCode's CDN staying up. If a
      download fails, skip that one image (note it in the end-of-run summary) rather than
      failing the whole README.

   e. Write `<folder>/INTUITION.md`: a plain-language walkthrough of the approach used in
      the existing code — the key insight, why this technique applies, time/space
      complexity. Written for a reader who's seen the problem but not the solution.

   f. Write `<folder>/SOLUTION.md`: a narrated walkthrough of the ACTUAL code in
      `main.go` — reference real function/variable names, explain each meaningful step
      in the order they appear in the code. This is not a generic solution write-up; it
      must describe this specific implementation.

   g. Edit `<folder>/main.go` in place to add inline `//` comments at non-obvious steps
      (loop invariants, why a particular data structure, edge cases handled). Do not
      change any logic, formatting style, or the test file.

   h. Look up companies:
      ```bash
      /tmp/leetcodectl companies-lookup '{"datasetPath":"challenge/companies_dataset.json","number":<number>}'
      ```
      If the returned `companies` array is non-empty, write `<folder>/COMPANIES.md` listing
      them. If empty, don't create the file at all.

   i. Append to the queue FIRST — this assigns and returns the authoritative Day number
      that every remaining step in this iteration needs (hero image, both post drafts):
      ```bash
      /tmp/leetcodectl queue-append '{"queuePath":"challenge/queue.yaml","entry":{"number":<number>,"title":"<title>","difficulty":"<difficulty>","folder":"<folder>","batch":"<list-name>"}}'
      ```
      Use the returned `day` for steps (j)-(m) below. Do not try to predict or read
      `next_day` yourself before calling this — `queue-append` is the only source of
      truth for which day number a question gets, and calling it exactly once per
      question, before rendering anything that embeds the day number, is what keeps a
      multi-question batch's day numbers correct. (This ordering matters specifically
      because there is no read-only way to peek `next_day` without mutating it — the
      only two things that touch it are `queue-has`, which doesn't return it, and
      `queue-append`, which advances it. Calling `queue-append` first removes any need
      to predict its value.)

   j. Render and screenshot the hero image, using the real `day` from step (i):
      ```bash
      /tmp/leetcodectl hero-render-html '{"templatePath":"challenge/hero_template.html","outPath":"<folder>/HERO.html","data":{"Day":<day>,"Total":365,"Topic":"<list-name>","Difficulty":"<difficulty>","Title":"<title>"}}'
      /tmp/leetcodectl hero-screenshot '{"htmlPath":"<folder>/HERO.html","outPath":"<folder>/HERO.png"}'
      ```
      If `hero-screenshot` fails (e.g. Playwright/Chromium isn't installed — see
      `challenge/README.md`), leave `HERO.html` in place, note the failure to the user at
      the end of the run, and continue with the rest of the pipeline; don't block the batch
      on it.

   k. Write `<folder>/POST_LINKEDIN.md`: header "365 Days of LeetCode Challenge — Day
      <day>/365", question title + LeetCode link, a 2-3 line intuition hook (not the full
      writeup — a teaser), a short code snippet or a link to the repo file, and the
      companies list if `COMPANIES.md` was written. Do NOT include a Discord-join call to
      action in the body — that's posted separately as a comment by the (future) poster
      subsystem.

   l. Write `<folder>/POST_DISCORD.md`: a shorter version of the same content, using
      Discord markdown (`**bold**`, `` ``` `` code fences), same Day N header.

   m. Commit everything for this question in one commit:
      ```bash
      git add <folder>
      git add challenge/queue.yaml challenge/number_folder_map.yaml
      git commit -m "Content(<number>) <title> - Day <day>"
      ```

4. **At the end of the run**, summarize for the user: how many questions were processed,
   how many were skipped as not-yet-solved, any folders reorganized, any hero-image
   generation failures, any per-image download failures in README generation, and any
   questions skipped mid-pipeline due to an unexpected `leetcodectl` failure. Give any
   post-`queue-append` failure (steps j-m, per the note above) its OWN separate, clearly
   flagged line — e.g. "⚠ Day N / question NUMBER was queued but incomplete — needs
   manual follow-up: <error>" — do not fold it into the general skipped-questions bullet,
   since it cannot be silently retried on a future run.

## Notes

- This skill only produces content and updates `challenge/queue.yaml` — it never posts
  anything to LinkedIn or Discord. Posting is a separate, not-yet-built subsystem.
- Re-running this skill with the same URL/list is always safe — already-queued questions
  are skipped, so partial batches (e.g. "10 of 30 done this week") resume cleanly.
```

- [ ] **Step 2: Commit**

```bash
git add .claude/skills/leetcode-content/SKILL.md
git commit -m "feat: add /leetcode-content skill orchestrating the content-gen pipeline"
```

---

### Task 13: `challenge/README.md` — setup + usage docs

**Files:**
- Create: `challenge/README.md`

- [ ] **Step 1: Write the docs**

`challenge/README.md`:
```markdown
# challenge/ — 365 Days of LeetCode Challenge tooling

Supports the `/leetcode-content` Claude Code skill (`.claude/skills/leetcode-content/SKILL.md`).
Design doc: `docs/superpowers/specs/2026-08-30-leetcode-content-gen-agent-design.md`.

## Layout

- `cmd/leetcodectl/` — CLI binary: deterministic logic (matching, reorg, queue, hero
  rendering, LeetCode fetch) behind JSON-in/JSON-out subcommands.
- `internal/` — the packages behind that CLI, each independently tested
  (`slugutil`, `queue`, `gitmap`, `resolver`, `reorg`, `leetcode`, `companies`, `hero`, `cli`).
- `queue.yaml` — the ordered, resumable record of which questions have content generated
  and/or have been posted. Read/written by `leetcodectl`; don't hand-edit unless fixing a
  mistake.
- `number_folder_map.yaml` — incremental cache mapping LeetCode question numbers to repo
  folders, derived from commit history. Safe to delete; it'll rebuild (slowly, via a full
  history scan) on the next `gitmap-update`.
- `companies_dataset.json` / `companies_dataset.md` — vendored, manually-refreshed
  company-tag data (see `companies_dataset.md` for how to refresh it).
- `hero_template.html` — the HTML template rendered + screenshotted into each post's
  `HERO.png`.

## One-time setup

```bash
go build -o /tmp/leetcodectl ./challenge/cmd/leetcodectl   # or wherever you keep local tools
npx playwright install chromium                             # needed for hero-image screenshots
```

## Running the tests

```bash
go test ./challenge/...
```
```

- [ ] **Step 2: Commit**

```bash
git add challenge/README.md
git commit -m "docs(challenge): add setup and usage README"
```

---

## Final check

Run the full suite once more from the repo root to confirm nothing regressed:

```bash
go build ./...
go test ./...
```

Expected: all packages build, all tests pass (existing solution tests plus the new
`challenge/...` tests).
