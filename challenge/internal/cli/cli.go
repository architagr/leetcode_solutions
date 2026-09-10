// Package cli holds the orchestration logic behind each leetcodectl
// subcommand, kept separate from cmd/leetcodectl/main.go so it's
// testable without spawning a subprocess.
package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
	"unicode/utf8"

	"leetcode_solutions/challenge/internal/companies"
	"leetcode_solutions/challenge/internal/discordpost"
	"leetcode_solutions/challenge/internal/gitmap"
	"leetcode_solutions/challenge/internal/hero"
	"leetcode_solutions/challenge/internal/leetcode"
	"leetcode_solutions/challenge/internal/postmeta"
	"leetcode_solutions/challenge/internal/queue"
	"leetcode_solutions/challenge/internal/reorg"
	"leetcode_solutions/challenge/internal/resolver"
	"leetcode_solutions/challenge/internal/xpost"
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
		return ResolveResult{}, fmt.Errorf("load gitmap %s: %w", mapPath, err)
	}
	if folder, ok := m.Entries[number]; ok {
		if info, statErr := os.Stat(filepath.Join(repoRoot, folder)); statErr == nil && info.IsDir() {
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
		return 0, "", fmt.Errorf("load gitmap %s: %w", mapPath, err)
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
		return false, fmt.Errorf("load queue %s: %w", queuePath, err)
	}
	return q.Has(number), nil
}

// QueueAppend appends e to the queue, assigning it the next day number,
// and returns that day number.
func QueueAppend(queuePath string, e queue.Entry) (int, error) {
	q, err := queue.Load(queuePath)
	if err != nil {
		return 0, fmt.Errorf("load queue %s: %w", queuePath, err)
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
		return nil, fmt.Errorf("load companies dataset %s: %w", datasetPath, err)
	}
	return ds.Lookup(number), nil
}

// screenshot is a seam so tests can exercise HeroGenerate's file
// handling without a real browser. Production always uses the real one.
var screenshot = hero.Screenshot

// HeroGenerateResult reports where the hero image landed.
type HeroGenerateResult struct {
	OutPath string `json:"outPath"`
	Palette string `json:"palette"`
}

// HeroGenerate renders the hero template and screenshots it to outPath
// in one step, rendering the intermediate HTML to a temp file rather
// than next to the PNG.
//
// A solution folder is read by people browsing the repo, and a 65KB
// HTML file full of base64 logos next to the image it produced is noise:
// once the PNG exists the HTML has no readers. On failure the HTML is
// kept and named in the error, since it's the only way to debug what the
// browser was given — but in the temp dir, not the solution folder.
func HeroGenerate(templatePath string, data hero.Data, outPath string) (HeroGenerateResult, error) {
	html, err := hero.RenderHTML(templatePath, data)
	if err != nil {
		return HeroGenerateResult{}, fmt.Errorf("rendering hero HTML: %w", err)
	}

	tmp, err := os.CreateTemp("", fmt.Sprintf("hero-day-%d-*.html", data.Day))
	if err != nil {
		return HeroGenerateResult{}, fmt.Errorf("creating temp hero HTML: %w", err)
	}
	tmpPath := tmp.Name()
	if _, err := tmp.WriteString(html); err != nil {
		tmp.Close()
		os.Remove(tmpPath)
		return HeroGenerateResult{}, fmt.Errorf("writing temp hero HTML: %w", err)
	}
	if err := tmp.Close(); err != nil {
		os.Remove(tmpPath)
		return HeroGenerateResult{}, fmt.Errorf("closing temp hero HTML: %w", err)
	}

	if err := screenshot(tmpPath, outPath); err != nil {
		return HeroGenerateResult{}, fmt.Errorf("screenshotting hero (rendered HTML kept at %s): %w", tmpPath, err)
	}
	os.Remove(tmpPath)

	return HeroGenerateResult{OutPath: outPath, Palette: hero.PaletteFor(data.Day).Name}, nil
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

// PostOption tweaks a post-discord run. The zero set of options is the
// scheduled behaviour; every option here exists for a hand-run override.
type PostOption func(*postConfig)

type postConfig struct {
	allowSameDay bool
}

// AllowSameDay lets a run post even though a day already went out today.
// Only for a deliberate manual catch-up after a missed day — the daily
// cron must never pass it.
func AllowSameDay() PostOption {
	return func(c *postConfig) { c.allowSameDay = true }
}

// PostDiscordResult reports what a post-discord run did. Posted is false
// on an idle day, when every queued entry has already gone to Discord.
type PostDiscordResult struct {
	Posted   bool   `json:"posted"`
	Day      int    `json:"day,omitempty"`
	Number   int    `json:"number,omitempty"`
	Title    string `json:"title,omitempty"`
	PostedAt string `json:"postedAt,omitempty"`
	Message  string `json:"message"`
}

// PostDiscord posts the oldest queue entry not yet sent to Discord, then
// records the timestamp against that entry's discord destination.
//
// The queue is only written after Discord accepts the message, so a
// failed post leaves the day unclaimed and the next run retries it
// rather than skipping ahead — the day sequence is the whole point of
// the queue.
func PostDiscord(repoRoot, queuePath, webhookURL string, opts ...PostOption) (PostDiscordResult, error) {
	var cfg postConfig
	for _, opt := range opts {
		opt(&cfg)
	}

	q, err := queue.Load(queuePath)
	if err != nil {
		return PostDiscordResult{}, fmt.Errorf("load queue %s: %w", queuePath, err)
	}

	// One post a day is the whole shape of the challenge, so a run that
	// finds today's day already sent is a no-op rather than the next day
	// going out early. Without this, a cron GitHub delayed past midnight
	// and then fired again on schedule would burn two days in one day.
	if !cfg.allowSameDay {
		if already, ok := queue.PostedOn(q, queue.DestinationDiscord, time.Now()); ok {
			return PostDiscordResult{
				Posted: false,
				Day:    already.Day,
				Number: already.Number,
				Title:  already.Title,
				Message: fmt.Sprintf("nothing to post: day %d already went to Discord today (%s)",
					already.Day, *already.PostedAt[queue.DestinationDiscord]),
			}, nil
		}
	}

	entry, ok := queue.SelectNext(q, queue.DestinationDiscord)
	if !ok {
		return PostDiscordResult{Posted: false, Message: "nothing to post: every queued entry is already on Discord"}, nil
	}

	contentPath := filepath.Join(repoRoot, entry.Folder, "POST_DISCORD.md")
	content, err := os.ReadFile(contentPath)
	if err != nil {
		return PostDiscordResult{}, fmt.Errorf("day %d (%s): reading POST_DISCORD.md: %w", entry.Day, entry.Title, err)
	}

	// The hero image is best-effort: a day whose screenshot failed during
	// content generation should still be able to go out as text.
	heroPath := filepath.Join(repoRoot, entry.Folder, "HERO.png")
	if _, statErr := os.Stat(heroPath); statErr != nil {
		heroPath = ""
	}

	if err := discordpost.Post(webhookURL, string(content), heroPath); err != nil {
		return PostDiscordResult{}, fmt.Errorf("day %d (%s): %w", entry.Day, entry.Title, err)
	}

	postedAt := time.Now().UTC()
	if !q.MarkPosted(entry.Number, queue.DestinationDiscord, postedAt) {
		return PostDiscordResult{}, fmt.Errorf("posted day %d but could not find entry %d to mark it", entry.Day, entry.Number)
	}
	if err := q.Save(queuePath); err != nil {
		return PostDiscordResult{}, fmt.Errorf("posted day %d but saving the queue failed: %w", entry.Day, err)
	}

	return PostDiscordResult{
		Posted:   true,
		Day:      entry.Day,
		Number:   entry.Number,
		Title:    entry.Title,
		PostedAt: postedAt.Format(time.RFC3339),
		Message:  fmt.Sprintf("posted day %d to Discord", entry.Day),
	}, nil
}

// postToX is a seam so tests can exercise PostX's queue and file
// handling without reaching the real API. Production always uses the
// real one.
var postToX = xpost.Post

// PostXResult reports what a post-x run did. Posted is false on an idle
// day, when every queued entry has already gone to X. URL is a link to
// the post that just went out, so a workflow log points at the result
// rather than only claiming success.
type PostXResult struct {
	Posted   bool   `json:"posted"`
	Day      int    `json:"day,omitempty"`
	Number   int    `json:"number,omitempty"`
	Title    string `json:"title,omitempty"`
	PostedAt string `json:"postedAt,omitempty"`
	URL      string `json:"url,omitempty"`
	Message  string `json:"message"`
}

// PostX posts the oldest queue entry not yet sent to X, then records the
// timestamp against that entry's x destination.
//
// The shape mirrors PostDiscord deliberately, including the same-day
// guard and the write-queue-only-after-the-network-accepts ordering: one
// post a day is the whole shape of the challenge, and on X specifically
// a retry loop that reposts the same text is not just untidy — near
// duplicate posts in quick succession are what X's platform manipulation
// policy treats as automated abuse.
//
// handle is the account the post lands on, used only to build URL. An
// empty handle just means URL is omitted.
func PostX(repoRoot, queuePath, handle string, creds xpost.Credentials, opts ...PostOption) (PostXResult, error) {
	var cfg postConfig
	for _, opt := range opts {
		opt(&cfg)
	}

	// Credentials are checked before the queue is even read: a run with
	// a missing secret should say so, not look like an idle day.
	if err := creds.Validate(); err != nil {
		return PostXResult{}, err
	}

	q, err := queue.Load(queuePath)
	if err != nil {
		return PostXResult{}, fmt.Errorf("load queue %s: %w", queuePath, err)
	}

	if !cfg.allowSameDay {
		if already, ok := queue.PostedOn(q, queue.DestinationX, time.Now()); ok {
			return PostXResult{
				Posted: false,
				Day:    already.Day,
				Number: already.Number,
				Title:  already.Title,
				Message: fmt.Sprintf("nothing to post: day %d already went to X today (%s)",
					already.Day, *already.PostedAt[queue.DestinationX]),
			}, nil
		}
	}

	entry, ok := queue.SelectNext(q, queue.DestinationX)
	if !ok {
		return PostXResult{Posted: false, Message: "nothing to post: every queued entry is already on X"}, nil
	}

	contentPath := filepath.Join(repoRoot, entry.Folder, "POST_X.md")
	content, err := os.ReadFile(contentPath)
	if err != nil {
		return PostXResult{}, fmt.Errorf("day %d (%s): reading POST_X.md: %w", entry.Day, entry.Title, err)
	}

	// The hero image is best-effort: a day whose screenshot failed during
	// content generation should still be able to go out as text.
	heroPath := filepath.Join(repoRoot, entry.Folder, "HERO.png")
	if _, statErr := os.Stat(heroPath); statErr != nil {
		heroPath = ""
	}

	postID, err := postToX(creds, strings.TrimSpace(string(content)), heroPath)
	if err != nil {
		return PostXResult{}, fmt.Errorf("day %d (%s): %w", entry.Day, entry.Title, err)
	}

	postedAt := time.Now().UTC()
	if !q.MarkPosted(entry.Number, queue.DestinationX, postedAt) {
		return PostXResult{}, fmt.Errorf("posted day %d but could not find entry %d to mark it", entry.Day, entry.Number)
	}
	if err := q.Save(queuePath); err != nil {
		return PostXResult{}, fmt.Errorf("posted day %d but saving the queue failed: %w", entry.Day, err)
	}

	var postURL string
	if handle != "" && postID != "" {
		postURL = fmt.Sprintf("https://x.com/%s/status/%s", strings.TrimPrefix(handle, "@"), postID)
	}

	return PostXResult{
		Posted:   true,
		Day:      entry.Day,
		Number:   entry.Number,
		Title:    entry.Title,
		PostedAt: postedAt.Format(time.RFC3339),
		URL:      postURL,
		Message:  fmt.Sprintf("posted day %d to X", entry.Day),
	}, nil
}

// BatchResult reports what a batch run prepared. MarkCommand is the
// exact follow-up command to run once the days have actually been
// scheduled on the destination.
//
// Warnings names anything about the prepared days that will publish but
// won't look right — an article missing its meta description, a title
// that a search result will truncate. They are reported rather than
// raised because a batch is prepared for a person who is about to look
// at every day in it anyway.
type BatchResult struct {
	Days        []int    `json:"days"`
	Numbers     []int    `json:"numbers"`
	OutPath     string   `json:"outPath,omitempty"`
	MarkCommand string   `json:"markCommand,omitempty"`
	Warnings    []string `json:"warnings,omitempty"`
	Message     string   `json:"message"`
}

// LinkedInBatchResult is the pre-Substack name for BatchResult, kept so
// callers and the JSON shape don't change.
type LinkedInBatchResult = BatchResult

// batchSection is one file to include per day in a batch document,
// under its own heading.
//
// HasMeta marks the file as long-form: its front matter is lifted out
// and printed as its own block, because LinkedIn and Substack both ask
// for the SEO title and description in separate fields at publish time
// rather than reading them off the article.
type batchSection struct {
	Heading  string
	Filename string
	HasMeta  bool
}

// prepareBatch writes the next count unposted days for destination to
// outPath as one paste-ready file.
//
// It deliberately does not mark anything: scheduling by hand is a
// partial process, and claiming days that never got scheduled would skip
// them forever. Marking is the explicit follow-up in MarkCommand,
// listing only what actually went out.
func prepareBatch(repoRoot, queuePath, destination, network, instructions string, sections []batchSection, count int, outPath string) (BatchResult, error) {
	if err := validateDestination(destination); err != nil {
		return BatchResult{}, err
	}
	q, err := queue.Load(queuePath)
	if err != nil {
		return BatchResult{}, fmt.Errorf("load queue %s: %w", queuePath, err)
	}

	pending := q.NextUnposted(destination, count)
	if len(pending) == 0 {
		return BatchResult{Message: fmt.Sprintf("nothing to prepare: every queued entry is already on %s", destination)}, nil
	}

	var doc strings.Builder
	days := make([]int, 0, len(pending))
	numbers := make([]int, 0, len(pending))
	var warnings []string

	fmt.Fprintf(&doc, "# %s batch — %d day(s) for %s\n\n", network, len(pending), destination)
	doc.WriteString(instructions + "\n")

	for _, e := range pending {
		days = append(days, e.Day)
		numbers = append(numbers, e.Number)

		fmt.Fprintf(&doc, "\n\n---\n\n## Day %d — %s (#%d)\n\n", e.Day, e.Title, e.Number)
		heroPath := filepath.Join(e.Folder, "HERO.png")
		if _, err := os.Stat(filepath.Join(repoRoot, heroPath)); err == nil {
			fmt.Fprintf(&doc, "Image to attach: `%s`\n", heroPath)
		} else {
			doc.WriteString("Image to attach: (no HERO.png for this day)\n")
		}

		for _, sec := range sections {
			raw, err := os.ReadFile(filepath.Join(repoRoot, e.Folder, sec.Filename))
			if err != nil {
				return BatchResult{}, fmt.Errorf("day %d (%s): reading %s: %w", e.Day, e.Title, sec.Filename, err)
			}

			body := strings.TrimSpace(string(raw))
			if sec.HasMeta {
				meta, stripped, err := postmeta.Parse(string(raw))
				if err != nil {
					return BatchResult{}, fmt.Errorf("day %d (%s): %s: %w", e.Day, e.Title, sec.Filename, err)
				}
				body = stripped
				fmt.Fprintf(&doc, "\n### %s — publish settings\n\n", sec.Heading)
				fmt.Fprintf(&doc, "- Meta title (%d/%d): %s\n", utf8.RuneCountInString(meta.Title), postmeta.MaxTitleChars, orNotSet(meta.Title))
				fmt.Fprintf(&doc, "- Meta description (%d/%d): %s\n", utf8.RuneCountInString(meta.Description), postmeta.MaxDescriptionChars, orNotSet(meta.Description))
				fmt.Fprintf(&doc, "- Canonical URL: %s\n", orNotSet(meta.Canonical))
				fmt.Fprintf(&doc, "- Tags: %s\n", orNotSet(strings.Join(meta.Tags, ", ")))
				for _, w := range meta.Warnings() {
					warnings = append(warnings, fmt.Sprintf("day %d (%s), %s: %s", e.Day, e.Title, sec.Filename, w))
				}
			}

			fmt.Fprintf(&doc, "\n### %s\n\n%s\n", sec.Heading, body)
		}
	}

	markCmd := fmt.Sprintf(`leetcodectl mark-posted '{"queuePath":%q,"destination":%q,"numbers":%s}'`,
		queuePath, destination, intsToJSON(numbers))
	fmt.Fprintf(&doc, "\n\n---\n\nOnce scheduled, mark them:\n\n    %s\n", markCmd)

	if len(warnings) > 0 {
		doc.WriteString("\nWorth fixing before publishing:\n\n")
		for _, w := range warnings {
			fmt.Fprintf(&doc, "- %s\n", w)
		}
	}

	if err := os.MkdirAll(filepath.Dir(outPath), 0o755); err != nil {
		return BatchResult{}, err
	}
	if err := os.WriteFile(outPath, []byte(doc.String()), 0o644); err != nil {
		return BatchResult{}, fmt.Errorf("writing batch file: %w", err)
	}

	return BatchResult{
		Days:        days,
		Numbers:     numbers,
		OutPath:     outPath,
		MarkCommand: markCmd,
		Warnings:    warnings,
		Message:     fmt.Sprintf("prepared %d day(s) for %s", len(days), destination),
	}, nil
}

func orNotSet(s string) string {
	if strings.TrimSpace(s) == "" {
		return "(not set)"
	}
	return s
}

// LinkedInBatch writes the next count unposted days' LinkedIn content to
// outPath as one paste-ready file, for scheduling by hand in LinkedIn's
// own composer.
func LinkedInBatch(repoRoot, queuePath, destination string, count int, outPath string) (BatchResult, error) {
	return prepareBatch(repoRoot, queuePath, destination, "LinkedIn",
		"Schedule these in LinkedIn's composer, then run the mark-posted command at the bottom for the days you actually scheduled.",
		[]batchSection{
			{Heading: "Short post", Filename: "POST_LINKEDIN.md"},
			{Heading: "Newsletter article", Filename: "POST_LINKEDIN_ARTICLE.md", HasMeta: true},
		}, count, outPath)
}

// SubstackBatch writes the next count unposted days' Substack content to
// outPath as one paste-ready file.
//
// Substack is manual for a harder reason than LinkedIn's: it has no
// public publishing API at all — only inbound RSS and email import — so
// there is nothing to automate against even with an approved app. The
// tooling prepares the post and tracks what went out; the paste is
// yours.
func SubstackBatch(repoRoot, queuePath string, count int, outPath string) (BatchResult, error) {
	return prepareBatch(repoRoot, queuePath, queue.DestinationSubstack, "Substack",
		"Paste each of these into a new Substack post, set the SEO fields from the publish settings block, then run the mark-posted command at the bottom for the days you actually published.",
		[]batchSection{
			{Heading: "Post", Filename: "POST_SUBSTACK.md", HasMeta: true},
		}, count, outPath)
}

// MarkPostedResult reports which numbers were marked and which weren't
// in the queue at all.
type MarkPostedResult struct {
	Marked      []int  `json:"marked"`
	NotFound    []int  `json:"notFound,omitempty"`
	Destination string `json:"destination"`
	Message     string `json:"message"`
}

// MarkPosted stamps the given question numbers as posted to destination.
// It's the manual counterpart to the Discord cron, for destinations
// published by hand.
func MarkPosted(queuePath, destination string, numbers []int) (MarkPostedResult, error) {
	if err := validateDestination(destination); err != nil {
		return MarkPostedResult{}, err
	}
	q, err := queue.Load(queuePath)
	if err != nil {
		return MarkPostedResult{}, fmt.Errorf("load queue %s: %w", queuePath, err)
	}

	now := time.Now().UTC()
	res := MarkPostedResult{Destination: destination}
	for _, n := range numbers {
		if q.MarkPosted(n, destination, now) {
			res.Marked = append(res.Marked, n)
		} else {
			res.NotFound = append(res.NotFound, n)
		}
	}
	if len(res.Marked) > 0 {
		if err := q.Save(queuePath); err != nil {
			return MarkPostedResult{}, fmt.Errorf("saving queue: %w", err)
		}
	}
	res.Message = fmt.Sprintf("marked %d entr(ies) posted to %s", len(res.Marked), destination)
	return res, nil
}

// validateDestination rejects typos rather than silently writing a key
// nothing will ever read.
func validateDestination(destination string) error {
	for _, known := range queue.KnownDestinations {
		if destination == known {
			return nil
		}
	}
	return fmt.Errorf("unknown destination %q (known: %s)", destination, strings.Join(queue.KnownDestinations, ", "))
}

func intsToJSON(ns []int) string {
	parts := make([]string, len(ns))
	for i, n := range ns {
		parts[i] = fmt.Sprint(n)
	}
	return "[" + strings.Join(parts, ",") + "]"
}
