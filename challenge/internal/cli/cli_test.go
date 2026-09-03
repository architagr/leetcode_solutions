package cli

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"leetcode_solutions/challenge/internal/gitmap"
	"leetcode_solutions/challenge/internal/hero"
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

func TestResolveFallsBackToGitmap(t *testing.T) {
	root := t.TempDir()
	// Actual folder lives somewhere resolver's static candidates and
	// FuzzyFind's easy/medium/hard walk never look (outside all three
	// difficulty roots), so only the gitmap entry can find it.
	folder := filepath.Join("archive", "965_univalued_binary_tree")
	if err := os.MkdirAll(filepath.Join(root, folder), 0o755); err != nil {
		t.Fatal(err)
	}
	mapPath := filepath.Join(root, "challenge", "number_folder_map.yaml")
	if err := os.MkdirAll(filepath.Dir(mapPath), 0o755); err != nil {
		t.Fatal(err)
	}
	m := &gitmap.Map{Entries: map[int]string{965: folder}}
	if err := m.Save(mapPath); err != nil {
		t.Fatal(err)
	}

	result, err := Resolve(root, mapPath, 965, "easy", "univalued-binary-tree")
	if err != nil {
		t.Fatalf("Resolve: %v", err)
	}
	if result.Status != "resolved" || result.Source != "gitmap" || result.Path != folder {
		t.Errorf("result = %+v", result)
	}
}

func TestResolveFallsBackToFuzzy(t *testing.T) {
	root := t.TempDir()
	mapPath := filepath.Join(root, "challenge", "number_folder_map.yaml")
	// Folder name predates a LeetCode slug change, but content matches
	// case-insensitively; the range bucket (701_800) doesn't match
	// RangeBucket(155) (101_200), so it can't be found as a static
	// candidate either.
	if err := os.MkdirAll(filepath.Join(root, "medium_problems", "701_800", "Min_Stack"), 0o755); err != nil {
		t.Fatal(err)
	}

	result, err := Resolve(root, mapPath, 155, "medium", "min-stack")
	if err != nil {
		t.Fatalf("Resolve: %v", err)
	}
	wantPath := filepath.Join("medium_problems", "701_800", "Min_Stack")
	if result.Status != "resolved" || result.Source != "fuzzy" || result.Path != wantPath {
		t.Errorf("result = %+v", result)
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

func heroDataFixture() hero.Data {
	return hero.Data{Day: 1, Title: "Two Sum"}
}

func TestPostDiscordPostsOldestUnpostedAndMarksIt(t *testing.T) {
	repoRoot := t.TempDir()
	folder := "easy_problems/101_200/two_sum"
	if err := os.MkdirAll(filepath.Join(repoRoot, folder), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(repoRoot, folder, "POST_DISCORD.md"), []byte("day one message"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(repoRoot, folder, "HERO.png"), []byte("fake-png"), 0o644); err != nil {
		t.Fatal(err)
	}

	queuePath := filepath.Join(repoRoot, "queue.yaml")
	q := &queue.Queue{NextDay: 1}
	q.Append(queue.Entry{Number: 1, Title: "Two Sum", Difficulty: "easy", Folder: folder, Batch: "arrays"})
	if err := q.Save(queuePath); err != nil {
		t.Fatal(err)
	}

	var gotBody string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, _ := io.ReadAll(r.Body)
		gotBody = string(b)
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	res, err := PostDiscord(repoRoot, queuePath, server.URL)
	if err != nil {
		t.Fatalf("PostDiscord: %v", err)
	}
	if !res.Posted || res.Day != 1 || res.Number != 1 {
		t.Errorf("result = %+v, want a posted day 1 / number 1", res)
	}
	if !strings.Contains(gotBody, "day one message") {
		t.Errorf("webhook body missing the message: %s", gotBody)
	}

	reloaded, err := queue.Load(queuePath)
	if err != nil {
		t.Fatal(err)
	}
	if !reloaded.Entries[0].IsPosted(queue.DestinationDiscord) {
		t.Error("entry should be marked posted to discord after a successful post")
	}
	if reloaded.Entries[0].IsPosted(queue.DestinationLinkedInMain) {
		t.Error("posting to discord must not mark linkedin")
	}
}

func TestPostDiscordNothingToPostIsNotAnError(t *testing.T) {
	repoRoot := t.TempDir()
	queuePath := filepath.Join(repoRoot, "queue.yaml")
	q := &queue.Queue{NextDay: 2, Entries: []queue.Entry{{Day: 1, Number: 1, Folder: "f"}}}
	q.MarkPosted(1, queue.DestinationDiscord, time.Now())
	if err := q.Save(queuePath); err != nil {
		t.Fatal(err)
	}

	res, err := PostDiscord(repoRoot, queuePath, "http://example.invalid")
	if err != nil {
		t.Fatalf("an idle day must not be an error: %v", err)
	}
	if res.Posted {
		t.Errorf("result = %+v, want Posted false", res)
	}
}

func TestPostDiscordFailedPostLeavesQueueUnchanged(t *testing.T) {
	repoRoot := t.TempDir()
	folder := "easy_problems/101_200/two_sum"
	if err := os.MkdirAll(filepath.Join(repoRoot, folder), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(repoRoot, folder, "POST_DISCORD.md"), []byte("msg"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(repoRoot, folder, "HERO.png"), []byte("png"), 0o644); err != nil {
		t.Fatal(err)
	}
	queuePath := filepath.Join(repoRoot, "queue.yaml")
	q := &queue.Queue{NextDay: 1}
	q.Append(queue.Entry{Number: 1, Folder: folder})
	if err := q.Save(queuePath); err != nil {
		t.Fatal(err)
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	if _, err := PostDiscord(repoRoot, queuePath, server.URL); err == nil {
		t.Fatal("a failed webhook must be an error")
	}
	reloaded, err := queue.Load(queuePath)
	if err != nil {
		t.Fatal(err)
	}
	if reloaded.Entries[0].IsPosted(queue.DestinationDiscord) {
		t.Error("a failed post must not mark the entry, so the day can be retried")
	}
}

func TestPostDiscordMissingContentFileIsAnError(t *testing.T) {
	repoRoot := t.TempDir()
	queuePath := filepath.Join(repoRoot, "queue.yaml")
	q := &queue.Queue{NextDay: 1}
	q.Append(queue.Entry{Number: 1, Folder: "easy_problems/101_200/gone"})
	if err := q.Save(queuePath); err != nil {
		t.Fatal(err)
	}

	_, err := PostDiscord(repoRoot, queuePath, "http://example.invalid")
	if err == nil {
		t.Fatal("a missing POST_DISCORD.md must be an error, not a skip to the next day")
	}
	if !strings.Contains(err.Error(), "POST_DISCORD.md") {
		t.Errorf("error should name the missing file, got: %v", err)
	}
}

func writeLinkedInFolder(t *testing.T, repoRoot, folder string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Join(repoRoot, folder), 0o755); err != nil {
		t.Fatal(err)
	}
	for name, body := range map[string]string{
		"POST_LINKEDIN.md":         "teaser for " + folder,
		"POST_LINKEDIN_ARTICLE.md": "article for " + folder,
		"HERO.png":                 "png",
	} {
		if err := os.WriteFile(filepath.Join(repoRoot, folder, name), []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
	}
}

func TestLinkedInBatchWritesRequestedDaysInOrder(t *testing.T) {
	repoRoot := t.TempDir()
	writeLinkedInFolder(t, repoRoot, "f1")
	writeLinkedInFolder(t, repoRoot, "f2")
	writeLinkedInFolder(t, repoRoot, "f3")

	queuePath := filepath.Join(repoRoot, "queue.yaml")
	q := &queue.Queue{NextDay: 1}
	q.Append(queue.Entry{Number: 104, Title: "One", Folder: "f1"})
	q.Append(queue.Entry{Number: 108, Title: "Two", Folder: "f2"})
	q.Append(queue.Entry{Number: 257, Title: "Three", Folder: "f3"})
	if err := q.Save(queuePath); err != nil {
		t.Fatal(err)
	}

	outPath := filepath.Join(repoRoot, "batch.md")
	res, err := LinkedInBatch(repoRoot, queuePath, queue.DestinationLinkedInMain, 2, outPath)
	if err != nil {
		t.Fatalf("LinkedInBatch: %v", err)
	}
	if len(res.Days) != 2 || res.Days[0] != 1 || res.Days[1] != 2 {
		t.Errorf("Days = %v, want [1 2]", res.Days)
	}
	if len(res.Numbers) != 2 || res.Numbers[0] != 104 {
		t.Errorf("Numbers = %v, want [104 108]", res.Numbers)
	}

	body, err := os.ReadFile(outPath)
	if err != nil {
		t.Fatalf("batch file not written: %v", err)
	}
	text := string(body)
	for _, want := range []string{"teaser for f1", "article for f1", "teaser for f2", "HERO.png", "Day 1", "Day 2"} {
		if !strings.Contains(text, want) {
			t.Errorf("batch file missing %q", want)
		}
	}
	if strings.Contains(text, "f3") {
		t.Error("batch file should stop at the requested count, day 3 leaked in")
	}
	// Preparing content must not claim the days — the user may only get
	// through some of them in LinkedIn's scheduler.
	reloaded, err := queue.Load(queuePath)
	if err != nil {
		t.Fatal(err)
	}
	if reloaded.Entries[0].IsPosted(queue.DestinationLinkedInMain) {
		t.Error("LinkedInBatch must not mark days posted; marking is an explicit follow-up")
	}
	if !strings.Contains(res.MarkCommand, "104") || !strings.Contains(res.MarkCommand, "mark-posted") {
		t.Errorf("MarkCommand should be a ready-to-run mark-posted command, got %q", res.MarkCommand)
	}
}

func TestLinkedInBatchSkipsAlreadyPostedAndReportsEmpty(t *testing.T) {
	repoRoot := t.TempDir()
	writeLinkedInFolder(t, repoRoot, "f1")
	queuePath := filepath.Join(repoRoot, "queue.yaml")
	q := &queue.Queue{NextDay: 1}
	q.Append(queue.Entry{Number: 104, Title: "One", Folder: "f1"})
	q.MarkPosted(104, queue.DestinationLinkedInMain, time.Now())
	if err := q.Save(queuePath); err != nil {
		t.Fatal(err)
	}

	res, err := LinkedInBatch(repoRoot, queuePath, queue.DestinationLinkedInMain, 5, filepath.Join(repoRoot, "batch.md"))
	if err != nil {
		t.Fatalf("an empty batch must not be an error: %v", err)
	}
	if len(res.Days) != 0 {
		t.Errorf("Days = %v, want empty", res.Days)
	}
}

func TestLinkedInBatchMissingContentIsAnError(t *testing.T) {
	repoRoot := t.TempDir()
	queuePath := filepath.Join(repoRoot, "queue.yaml")
	q := &queue.Queue{NextDay: 1}
	q.Append(queue.Entry{Number: 104, Title: "One", Folder: "gone"})
	if err := q.Save(queuePath); err != nil {
		t.Fatal(err)
	}

	_, err := LinkedInBatch(repoRoot, queuePath, queue.DestinationLinkedInMain, 1, filepath.Join(repoRoot, "batch.md"))
	if err == nil {
		t.Fatal("a missing POST_LINKEDIN.md must be an error")
	}
	if !strings.Contains(err.Error(), "POST_LINKEDIN.md") {
		t.Errorf("error should name the missing file, got %v", err)
	}
}

func TestMarkPostedMarksListedNumbersOnly(t *testing.T) {
	repoRoot := t.TempDir()
	queuePath := filepath.Join(repoRoot, "queue.yaml")
	q := &queue.Queue{NextDay: 1}
	q.Append(queue.Entry{Number: 104, Folder: "f1"})
	q.Append(queue.Entry{Number: 108, Folder: "f2"})
	if err := q.Save(queuePath); err != nil {
		t.Fatal(err)
	}

	res, err := MarkPosted(queuePath, queue.DestinationLinkedInMain, []int{104, 999})
	if err != nil {
		t.Fatalf("MarkPosted: %v", err)
	}
	if len(res.Marked) != 1 || res.Marked[0] != 104 {
		t.Errorf("Marked = %v, want [104]", res.Marked)
	}
	if len(res.NotFound) != 1 || res.NotFound[0] != 999 {
		t.Errorf("NotFound = %v, want [999]", res.NotFound)
	}

	reloaded, err := queue.Load(queuePath)
	if err != nil {
		t.Fatal(err)
	}
	if !reloaded.Entries[0].IsPosted(queue.DestinationLinkedInMain) {
		t.Error("104 should be marked")
	}
	if reloaded.Entries[1].IsPosted(queue.DestinationLinkedInMain) {
		t.Error("108 was not in the list and must stay unmarked")
	}
	if reloaded.Entries[0].IsPosted(queue.DestinationDiscord) {
		t.Error("marking linkedin must not touch discord")
	}
}

func TestMarkPostedRejectsUnknownDestination(t *testing.T) {
	repoRoot := t.TempDir()
	queuePath := filepath.Join(repoRoot, "queue.yaml")
	q := &queue.Queue{NextDay: 1}
	q.Append(queue.Entry{Number: 104, Folder: "f1"})
	if err := q.Save(queuePath); err != nil {
		t.Fatal(err)
	}

	if _, err := MarkPosted(queuePath, "twitter", []int{104}); err == nil {
		t.Fatal("an unknown destination should be rejected, not silently written")
	}
}
