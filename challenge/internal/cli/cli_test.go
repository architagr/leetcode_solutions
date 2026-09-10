package cli

import (
	"errors"
	"fmt"
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
	"leetcode_solutions/challenge/internal/xpost"
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

// postDiscordFixture writes a repo with two queued days, day 1 already on
// Discord at postedAt, and returns the repo root and queue path.
func postDiscordFixture(t *testing.T, postedAt time.Time) (repoRoot, queuePath string) {
	t.Helper()
	repoRoot = t.TempDir()
	for _, folder := range []string{"p/one", "p/two"} {
		if err := os.MkdirAll(filepath.Join(repoRoot, folder), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(repoRoot, folder, "POST_DISCORD.md"), []byte("msg "+folder), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	queuePath = filepath.Join(repoRoot, "queue.yaml")
	q := &queue.Queue{NextDay: 1}
	q.Append(queue.Entry{Number: 1, Folder: "p/one"})
	q.Append(queue.Entry{Number: 2, Folder: "p/two"})
	q.MarkPosted(1, queue.DestinationDiscord, postedAt)
	if err := q.Save(queuePath); err != nil {
		t.Fatal(err)
	}
	return repoRoot, queuePath
}

func TestPostDiscordSkipsWhenADayAlreadyWentOutToday(t *testing.T) {
	repoRoot, queuePath := postDiscordFixture(t, time.Now().UTC())

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("a second run the same day must not call the webhook")
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	res, err := PostDiscord(repoRoot, queuePath, server.URL)
	if err != nil {
		t.Fatalf("a same-day rerun must not be an error: %v", err)
	}
	if res.Posted {
		t.Errorf("result = %+v, want Posted false", res)
	}
	if !strings.Contains(res.Message, "already") {
		t.Errorf("message = %q, want it to explain the skip", res.Message)
	}

	reloaded, err := queue.Load(queuePath)
	if err != nil {
		t.Fatal(err)
	}
	if reloaded.Entries[1].IsPosted(queue.DestinationDiscord) {
		t.Error("day 2 must stay unposted after a skipped run")
	}
}

func TestPostDiscordPostsAgainOnTheNextDay(t *testing.T) {
	repoRoot, queuePath := postDiscordFixture(t, time.Now().UTC().AddDate(0, 0, -1))

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	res, err := PostDiscord(repoRoot, queuePath, server.URL)
	if err != nil {
		t.Fatalf("PostDiscord: %v", err)
	}
	if !res.Posted || res.Day != 2 {
		t.Errorf("result = %+v, want a posted day 2", res)
	}
}

func TestPostDiscordAllowSameDayOverridesTheGuard(t *testing.T) {
	repoRoot, queuePath := postDiscordFixture(t, time.Now().UTC())

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	res, err := PostDiscord(repoRoot, queuePath, server.URL, AllowSameDay())
	if err != nil {
		t.Fatalf("PostDiscord: %v", err)
	}
	if !res.Posted || res.Day != 2 {
		t.Errorf("result = %+v, want the guard overridden and day 2 posted", res)
	}
}

func TestHeroGenerateLeavesNoHTMLBehind(t *testing.T) {
	folder := t.TempDir()
	tmplPath := filepath.Join(t.TempDir(), "t.html")
	if err := os.WriteFile(tmplPath, []byte(`Day {{.Day}} {{.Palette.Name}}`), 0o644); err != nil {
		t.Fatal(err)
	}
	outPath := filepath.Join(folder, "HERO.png")

	var gotHTML string
	restore := stubScreenshot(func(htmlPath, out string) error {
		// The HTML must still exist while the browser needs it.
		b, err := os.ReadFile(htmlPath)
		if err != nil {
			return err
		}
		gotHTML = string(b)
		return os.WriteFile(out, []byte("fake-png"), 0o644)
	})
	defer restore()

	res, err := HeroGenerate(tmplPath, hero.Data{Day: 5, Title: "Balanced Binary Tree"}, outPath)
	if err != nil {
		t.Fatalf("HeroGenerate: %v", err)
	}
	if res.OutPath != outPath {
		t.Errorf("OutPath = %q, want %q", res.OutPath, outPath)
	}
	if want := "Day 5 " + hero.PaletteFor(5).Name; gotHTML != want {
		t.Errorf("screenshotted HTML = %q, want %q", gotHTML, want)
	}
	if _, err := os.Stat(outPath); err != nil {
		t.Errorf("PNG not written: %v", err)
	}

	entries, err := os.ReadDir(folder)
	if err != nil {
		t.Fatal(err)
	}
	for _, e := range entries {
		if strings.HasSuffix(e.Name(), ".html") {
			t.Errorf("%s left behind in the solution folder; only HERO.png belongs there", e.Name())
		}
	}
}

func TestHeroGenerateKeepsTheHTMLWhenTheScreenshotFails(t *testing.T) {
	folder := t.TempDir()
	tmplPath := filepath.Join(t.TempDir(), "t.html")
	if err := os.WriteFile(tmplPath, []byte(`Day {{.Day}}`), 0o644); err != nil {
		t.Fatal(err)
	}

	restore := stubScreenshot(func(htmlPath, out string) error {
		return fmt.Errorf("playwright is not installed")
	})
	defer restore()

	_, err := HeroGenerate(tmplPath, hero.Data{Day: 5}, filepath.Join(folder, "HERO.png"))
	if err == nil {
		t.Fatal("HeroGenerate: err = nil, want the screenshot failure")
	}
	// The rendered HTML is the only way to debug a screenshot failure, so
	// it survives — but outside the solution folder, which stays clean.
	kept := heroHTMLPathFromError(t, err)
	if _, statErr := os.Stat(kept); statErr != nil {
		t.Errorf("HTML kept for debugging is missing: %v", statErr)
	}
	if strings.HasPrefix(kept, folder) {
		t.Errorf("kept HTML %q is inside the solution folder", kept)
	}
}

// heroHTMLPathFromError pulls the retained HTML path out of the error
// message, which is the only place a caller learns about it.
func heroHTMLPathFromError(t *testing.T, err error) string {
	t.Helper()
	for _, field := range strings.Fields(err.Error()) {
		field = strings.TrimRight(field, "):,.")
		if strings.HasSuffix(field, ".html") {
			return field
		}
	}
	t.Fatalf("error %q names no .html file to debug with", err)
	return ""
}

// stubScreenshot swaps in a fake browser for one test and returns the
// restore func.
func stubScreenshot(fn func(htmlPath, outPath string) error) func() {
	prev := screenshot
	screenshot = fn
	return func() { screenshot = prev }
}

// fakeX swaps the X client for the duration of a test, recording what it
// was asked to post and returning postID (or err, when set).
func fakeX(t *testing.T, postID string, err error) *struct {
	Calls    int
	Text     string
	HeroPath string
	Creds    xpost.Credentials
} {
	t.Helper()
	rec := &struct {
		Calls    int
		Text     string
		HeroPath string
		Creds    xpost.Credentials
	}{}
	old := postToX
	postToX = func(creds xpost.Credentials, text, heroPath string) (string, error) {
		rec.Calls++
		rec.Text, rec.HeroPath, rec.Creds = text, heroPath, creds
		return postID, err
	}
	t.Cleanup(func() { postToX = old })
	return rec
}

func xCreds() xpost.Credentials {
	return xpost.Credentials{ConsumerKey: "ck", ConsumerSecret: "cs", AccessToken: "at", AccessSecret: "as"}
}

// xFixture writes a repo with one queued day whose POST_X.md and
// HERO.png exist, and returns the repo root and queue path.
func xFixture(t *testing.T) (repoRoot, queuePath, folder string) {
	t.Helper()
	repoRoot = t.TempDir()
	folder = "easy_problems/101_200/two_sum"
	if err := os.MkdirAll(filepath.Join(repoRoot, folder), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(repoRoot, folder, "POST_X.md"), []byte("day one on X\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(repoRoot, folder, "HERO.png"), []byte("fake-png"), 0o644); err != nil {
		t.Fatal(err)
	}
	queuePath = filepath.Join(repoRoot, "queue.yaml")
	q := &queue.Queue{NextDay: 1}
	q.Append(queue.Entry{Number: 1, Title: "Two Sum", Difficulty: "easy", Folder: folder, Batch: "arrays"})
	if err := q.Save(queuePath); err != nil {
		t.Fatal(err)
	}
	return repoRoot, queuePath, folder
}

func TestPostXPostsOldestUnpostedAndMarksIt(t *testing.T) {
	repoRoot, queuePath, folder := xFixture(t)
	rec := fakeX(t, "1799999999", nil)

	res, err := PostX(repoRoot, queuePath, "@architagr", xCreds())
	if err != nil {
		t.Fatalf("PostX: %v", err)
	}
	if !res.Posted || res.Day != 1 || res.Number != 1 {
		t.Errorf("result = %+v, want a posted day 1 / number 1", res)
	}
	// The handle carries an @ in the queue and in most places a person
	// writes it; the URL must not.
	if want := "https://x.com/architagr/status/1799999999"; res.URL != want {
		t.Errorf("URL = %q, want %q", res.URL, want)
	}
	if rec.Text != "day one on X" {
		t.Errorf("posted text = %q, want the file contents with the trailing newline trimmed", rec.Text)
	}
	if rec.HeroPath != filepath.Join(repoRoot, folder, "HERO.png") {
		t.Errorf("hero path = %q, want the day's HERO.png", rec.HeroPath)
	}

	reloaded, err := queue.Load(queuePath)
	if err != nil {
		t.Fatal(err)
	}
	if !reloaded.Entries[0].IsPosted(queue.DestinationX) {
		t.Error("entry should be marked posted to x after a successful post")
	}
	if reloaded.Entries[0].IsPosted(queue.DestinationDiscord) {
		t.Error("posting to x must not mark discord")
	}
}

// The queue is the record of what is owed. Marking a day that never went
// out skips it permanently, so a failed post must leave the queue alone.
func TestPostXFailedPostLeavesQueueUnchanged(t *testing.T) {
	repoRoot, queuePath, _ := xFixture(t)
	fakeX(t, "", errors.New("X returned 403: duplicate content"))

	if _, err := PostX(repoRoot, queuePath, "architagr", xCreds()); err == nil {
		t.Fatal("a rejected post must be an error")
	}

	reloaded, err := queue.Load(queuePath)
	if err != nil {
		t.Fatal(err)
	}
	if reloaded.Entries[0].IsPosted(queue.DestinationX) {
		t.Error("a failed post must not mark the entry, or the day is skipped forever")
	}
}

func TestPostXMissingContentFileIsAnError(t *testing.T) {
	repoRoot, queuePath, folder := xFixture(t)
	if err := os.Remove(filepath.Join(repoRoot, folder, "POST_X.md")); err != nil {
		t.Fatal(err)
	}
	rec := fakeX(t, "1", nil)

	if _, err := PostX(repoRoot, queuePath, "architagr", xCreds()); err == nil {
		t.Fatal("a missing POST_X.md must be an error")
	}
	if rec.Calls != 0 {
		t.Error("nothing should have been posted when the content file is missing")
	}
}

// A day whose hero screenshot failed should still be able to go out as
// text rather than blocking the day.
func TestPostXWithoutAHeroImagePostsTextOnly(t *testing.T) {
	repoRoot, queuePath, folder := xFixture(t)
	if err := os.Remove(filepath.Join(repoRoot, folder, "HERO.png")); err != nil {
		t.Fatal(err)
	}
	rec := fakeX(t, "1", nil)

	res, err := PostX(repoRoot, queuePath, "architagr", xCreds())
	if err != nil {
		t.Fatalf("PostX: %v", err)
	}
	if !res.Posted {
		t.Error("a missing hero must not stop the post")
	}
	if rec.HeroPath != "" {
		t.Errorf("hero path = %q, want empty when there is no HERO.png", rec.HeroPath)
	}
}

// A missing secret must fail loudly. Reported as an idle day it would
// look like the challenge had simply run out of content.
func TestPostXWithIncompleteCredentialsIsAnErrorNotAnIdleDay(t *testing.T) {
	repoRoot, queuePath, _ := xFixture(t)
	rec := fakeX(t, "1", nil)

	res, err := PostX(repoRoot, queuePath, "architagr", xpost.Credentials{ConsumerKey: "ck"})
	if err == nil {
		t.Fatalf("incomplete credentials must be an error, got %+v", res)
	}
	if !strings.Contains(err.Error(), "X_API_SECRET") {
		t.Errorf("error should name the missing variables, got %v", err)
	}
	if rec.Calls != 0 {
		t.Error("nothing should have been posted with incomplete credentials")
	}
}

func TestPostXNothingToPostIsNotAnError(t *testing.T) {
	repoRoot := t.TempDir()
	queuePath := filepath.Join(repoRoot, "queue.yaml")
	q := &queue.Queue{NextDay: 2, Entries: []queue.Entry{{Day: 1, Number: 1, Folder: "f"}}}
	q.MarkPosted(1, queue.DestinationX, time.Now().Add(-48*time.Hour))
	if err := q.Save(queuePath); err != nil {
		t.Fatal(err)
	}
	fakeX(t, "1", nil)

	res, err := PostX(repoRoot, queuePath, "architagr", xCreds())
	if err != nil {
		t.Fatalf("an idle day must not be an error: %v", err)
	}
	if res.Posted {
		t.Errorf("result = %+v, want Posted false", res)
	}
}

// A cron GitHub delayed past midnight, then firing again on schedule,
// must not burn two days in one day.
func TestPostXSkipsWhenADayAlreadyWentOutToday(t *testing.T) {
	repoRoot, queuePath, _ := xFixture(t)
	q, err := queue.Load(queuePath)
	if err != nil {
		t.Fatal(err)
	}
	day2 := "medium_problems/1_100/add_two_numbers"
	if err := os.MkdirAll(filepath.Join(repoRoot, day2), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(repoRoot, day2, "POST_X.md"), []byte("day two on X"), 0o644); err != nil {
		t.Fatal(err)
	}
	q.Append(queue.Entry{Number: 2, Title: "Add Two Numbers", Folder: day2, Batch: "lists"})
	q.MarkPosted(1, queue.DestinationX, time.Now().UTC())
	if err := q.Save(queuePath); err != nil {
		t.Fatal(err)
	}
	rec := fakeX(t, "1", nil)

	res, err := PostX(repoRoot, queuePath, "architagr", xCreds())
	if err != nil {
		t.Fatalf("PostX: %v", err)
	}
	if res.Posted {
		t.Errorf("result = %+v, want Posted false — day 1 already went out today", res)
	}
	if rec.Calls != 0 {
		t.Error("the same-day guard must run before any network call")
	}

	// The manual catch-up override still gets through.
	if res, err := PostX(repoRoot, queuePath, "architagr", xCreds(), AllowSameDay()); err != nil || !res.Posted {
		t.Errorf("AllowSameDay = (%+v, %v), want a posted day", res, err)
	}
}

// An empty handle is not an error — the post still goes out, there is
// just no URL to report.
func TestPostXWithoutAHandleOmitsTheURL(t *testing.T) {
	repoRoot, queuePath, _ := xFixture(t)
	fakeX(t, "1799999999", nil)

	res, err := PostX(repoRoot, queuePath, "", xCreds())
	if err != nil {
		t.Fatalf("PostX: %v", err)
	}
	if !res.Posted {
		t.Error("a missing handle must not stop the post")
	}
	if res.URL != "" {
		t.Errorf("URL = %q, want empty without a handle", res.URL)
	}
}

// batchFixture writes a repo with one queued day carrying every
// long-form content file, and returns the repo root and queue path.
func batchFixture(t *testing.T, articleBody, substackBody string) (repoRoot, queuePath string) {
	t.Helper()
	repoRoot = t.TempDir()
	folder := "easy_problems/101_200/two_sum"
	if err := os.MkdirAll(filepath.Join(repoRoot, folder), 0o755); err != nil {
		t.Fatal(err)
	}
	files := map[string]string{
		"POST_LINKEDIN.md":         "short teaser",
		"POST_LINKEDIN_ARTICLE.md": articleBody,
		"POST_SUBSTACK.md":         substackBody,
		"HERO.png":                 "fake-png",
	}
	for name, body := range files {
		if err := os.WriteFile(filepath.Join(repoRoot, folder, name), []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	queuePath = filepath.Join(repoRoot, "queue.yaml")
	q := &queue.Queue{NextDay: 1}
	q.Append(queue.Entry{Number: 1, Title: "Two Sum", Difficulty: "easy", Folder: folder, Batch: "arrays"})
	if err := q.Save(queuePath); err != nil {
		t.Fatal(err)
	}
	return repoRoot, queuePath
}

const articleWithMeta = `---
meta_title: A title that fits
meta_description: A description that fits comfortably.
---

# The article

Body text.
`

// LinkedIn asks for the SEO title and description in their own fields at
// publish time, so the batch has to surface them separately rather than
// leaving them buried at the top of the article.
func TestLinkedInBatchLiftsArticleMetaIntoItsOwnBlock(t *testing.T) {
	repoRoot, queuePath := batchFixture(t, articleWithMeta, "substack body")
	outPath := filepath.Join(t.TempDir(), "batch.md")

	res, err := LinkedInBatch(repoRoot, queuePath, queue.DestinationLinkedInMain, 7, outPath)
	if err != nil {
		t.Fatal(err)
	}
	if len(res.Warnings) != 0 {
		t.Errorf("warnings = %v, want none for well-formed meta", res.Warnings)
	}

	doc, err := os.ReadFile(outPath)
	if err != nil {
		t.Fatal(err)
	}
	got := string(doc)
	for _, want := range []string{
		"Meta title (17/60): A title that fits",
		"Meta description (36/155): A description that fits comfortably.",
		"# The article",
	} {
		if !strings.Contains(got, want) {
			t.Errorf("batch document is missing %q:\n%s", want, got)
		}
	}
	// The front matter must not also appear inline, or it gets pasted
	// into the article body as visible YAML.
	if strings.Contains(got, "meta_title:") {
		t.Errorf("raw front matter leaked into the batch body:\n%s", got)
	}
}

// Days written before the meta fields existed must still prepare, with
// the gap reported rather than raised.
func TestLinkedInBatchReportsMissingMetaAsAWarning(t *testing.T) {
	repoRoot, queuePath := batchFixture(t, "# An older article\n\nNo front matter.\n", "substack body")
	outPath := filepath.Join(t.TempDir(), "batch.md")

	res, err := LinkedInBatch(repoRoot, queuePath, queue.DestinationLinkedInMain, 7, outPath)
	if err != nil {
		t.Fatalf("an article without front matter must still prepare: %v", err)
	}
	if len(res.Warnings) != 2 {
		t.Errorf("warnings = %v, want one each for the missing title and description", res.Warnings)
	}
	for _, w := range res.Warnings {
		if !strings.Contains(w, "day 1") || !strings.Contains(w, "POST_LINKEDIN_ARTICLE.md") {
			t.Errorf("warning should name the day and the file: %q", w)
		}
	}

	doc, err := os.ReadFile(outPath)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(doc), "Meta title (0/60): (not set)") {
		t.Errorf("batch document should show the gap explicitly:\n%s", doc)
	}
}

func TestSubstackBatchPreparesPostsAndTracksItsOwnDestination(t *testing.T) {
	repoRoot, queuePath := batchFixture(t, articleWithMeta, "---\nmeta_title: Substack title\n---\n\nSubstack body.\n")
	outPath := filepath.Join(t.TempDir(), "substack.md")

	res, err := SubstackBatch(repoRoot, queuePath, 4, outPath)
	if err != nil {
		t.Fatal(err)
	}
	if len(res.Days) != 1 || res.Days[0] != 1 {
		t.Errorf("days = %v, want [1]", res.Days)
	}
	if !strings.Contains(res.MarkCommand, queue.DestinationSubstack) {
		t.Errorf("mark command should target substack: %s", res.MarkCommand)
	}

	doc, err := os.ReadFile(outPath)
	if err != nil {
		t.Fatal(err)
	}
	got := string(doc)
	for _, want := range []string{"Substack batch", "Meta title (14/60): Substack title", "Substack body."} {
		if !strings.Contains(got, want) {
			t.Errorf("substack document is missing %q:\n%s", want, got)
		}
	}
	// LinkedIn's files must not bleed into a Substack batch.
	if strings.Contains(got, "short teaser") {
		t.Errorf("substack document included LinkedIn content:\n%s", got)
	}
}

// Preparing a batch must never mark anything: scheduling by hand is
// partial, and a claimed day that was never published is skipped forever.
func TestSubstackBatchDoesNotMarkTheQueue(t *testing.T) {
	repoRoot, queuePath := batchFixture(t, articleWithMeta, "substack body")

	if _, err := SubstackBatch(repoRoot, queuePath, 4, filepath.Join(t.TempDir(), "out.md")); err != nil {
		t.Fatal(err)
	}
	reloaded, err := queue.Load(queuePath)
	if err != nil {
		t.Fatal(err)
	}
	if reloaded.Entries[0].IsPosted(queue.DestinationSubstack) {
		t.Error("preparing a batch must not mark the day as published")
	}
}

func TestSubstackBatchMissingContentFileIsAnError(t *testing.T) {
	repoRoot, queuePath := batchFixture(t, articleWithMeta, "substack body")
	if err := os.Remove(filepath.Join(repoRoot, "easy_problems/101_200/two_sum", "POST_SUBSTACK.md")); err != nil {
		t.Fatal(err)
	}
	if _, err := SubstackBatch(repoRoot, queuePath, 4, filepath.Join(t.TempDir(), "out.md")); err == nil {
		t.Fatal("a missing POST_SUBSTACK.md must be an error")
	}
}

// Unclosed front matter would otherwise be pasted into the published
// post as visible YAML.
func TestBatchRejectsBrokenFrontMatter(t *testing.T) {
	repoRoot, queuePath := batchFixture(t, "---\nmeta_title: oops\n\n# Article\n", "substack body")
	_, err := LinkedInBatch(repoRoot, queuePath, queue.DestinationLinkedInMain, 7, filepath.Join(t.TempDir(), "out.md"))
	if err == nil {
		t.Fatal("front matter that is never closed must be an error")
	}
	if !strings.Contains(err.Error(), "POST_LINKEDIN_ARTICLE.md") {
		t.Errorf("error should name the file: %v", err)
	}
}

func TestXBatchPreparesPostsForHandPosting(t *testing.T) {
	repoRoot, queuePath, folder := xFixture(t)
	outPath := filepath.Join(t.TempDir(), "x.md")

	res, err := XBatch(repoRoot, queuePath, 7, outPath)
	if err != nil {
		t.Fatal(err)
	}
	if len(res.Days) != 1 || res.Days[0] != 1 {
		t.Errorf("days = %v, want [1]", res.Days)
	}
	if !strings.Contains(res.MarkCommand, queue.DestinationX) {
		t.Errorf("mark command should target x: %s", res.MarkCommand)
	}

	doc, err := os.ReadFile(outPath)
	if err != nil {
		t.Fatal(err)
	}
	got := string(doc)
	if !strings.Contains(got, "day one on X") {
		t.Errorf("x document is missing the post body:\n%s", got)
	}
	// Posting by hand means attaching the image by hand, so the document
	// has to say which file.
	if !strings.Contains(got, filepath.Join(folder, "HERO.png")) {
		t.Errorf("x document should name the image to attach:\n%s", got)
	}
}

// The same rule as every other batch: preparing claims nothing, because
// a day marked but never posted is skipped forever.
func TestXBatchDoesNotMarkTheQueue(t *testing.T) {
	repoRoot, queuePath, _ := xFixture(t)

	if _, err := XBatch(repoRoot, queuePath, 7, filepath.Join(t.TempDir(), "out.md")); err != nil {
		t.Fatal(err)
	}
	reloaded, err := queue.Load(queuePath)
	if err != nil {
		t.Fatal(err)
	}
	if reloaded.Entries[0].IsPosted(queue.DestinationX) {
		t.Error("preparing a batch must not mark the day as posted")
	}
}
