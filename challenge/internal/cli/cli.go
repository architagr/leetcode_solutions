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

	"leetcode_solutions/challenge/internal/companies"
	"leetcode_solutions/challenge/internal/discordpost"
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
		if already, ok := discordpost.PostedOn(q, queue.DestinationDiscord, time.Now()); ok {
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

	entry, ok := discordpost.SelectNext(q, queue.DestinationDiscord)
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

// LinkedInBatchResult reports what a linkedin-batch run prepared.
// MarkCommand is the exact follow-up command to run once the days have
// actually been scheduled in LinkedIn.
type LinkedInBatchResult struct {
	Days        []int  `json:"days"`
	Numbers     []int  `json:"numbers"`
	OutPath     string `json:"outPath,omitempty"`
	MarkCommand string `json:"markCommand,omitempty"`
	Message     string `json:"message"`
}

// LinkedInBatch writes the next count unposted days' LinkedIn content to
// outPath as one paste-ready file, for scheduling by hand in LinkedIn's
// own composer.
//
// It deliberately does not mark anything: LinkedIn's scheduler is a
// manual, partial process, and claiming days that never got scheduled
// would skip them forever. Marking is the explicit follow-up in
// MarkCommand, listing only what actually went out.
func LinkedInBatch(repoRoot, queuePath, destination string, count int, outPath string) (LinkedInBatchResult, error) {
	if err := validateDestination(destination); err != nil {
		return LinkedInBatchResult{}, err
	}
	q, err := queue.Load(queuePath)
	if err != nil {
		return LinkedInBatchResult{}, fmt.Errorf("load queue %s: %w", queuePath, err)
	}

	pending := q.NextUnposted(destination, count)
	if len(pending) == 0 {
		return LinkedInBatchResult{Message: fmt.Sprintf("nothing to prepare: every queued entry is already on %s", destination)}, nil
	}

	var doc strings.Builder
	days := make([]int, 0, len(pending))
	numbers := make([]int, 0, len(pending))

	fmt.Fprintf(&doc, "# LinkedIn batch — %d day(s) for %s\n\n", len(pending), destination)
	doc.WriteString("Schedule these in LinkedIn's composer, then run the mark-posted command at the bottom for the days you actually scheduled.\n")

	for _, e := range pending {
		teaser, err := os.ReadFile(filepath.Join(repoRoot, e.Folder, "POST_LINKEDIN.md"))
		if err != nil {
			return LinkedInBatchResult{}, fmt.Errorf("day %d (%s): reading POST_LINKEDIN.md: %w", e.Day, e.Title, err)
		}
		article, err := os.ReadFile(filepath.Join(repoRoot, e.Folder, "POST_LINKEDIN_ARTICLE.md"))
		if err != nil {
			return LinkedInBatchResult{}, fmt.Errorf("day %d (%s): reading POST_LINKEDIN_ARTICLE.md: %w", e.Day, e.Title, err)
		}

		days = append(days, e.Day)
		numbers = append(numbers, e.Number)

		fmt.Fprintf(&doc, "\n\n---\n\n## Day %d — %s (#%d)\n\n", e.Day, e.Title, e.Number)
		heroPath := filepath.Join(e.Folder, "HERO.png")
		if _, err := os.Stat(filepath.Join(repoRoot, heroPath)); err == nil {
			fmt.Fprintf(&doc, "Image to attach: `%s`\n", heroPath)
		} else {
			doc.WriteString("Image to attach: (no HERO.png for this day)\n")
		}
		fmt.Fprintf(&doc, "\n### Short post\n\n%s\n\n### Newsletter article\n\n%s\n", strings.TrimSpace(string(teaser)), strings.TrimSpace(string(article)))
	}

	markCmd := fmt.Sprintf(`leetcodectl mark-posted '{"queuePath":%q,"destination":%q,"numbers":%s}'`,
		queuePath, destination, intsToJSON(numbers))
	fmt.Fprintf(&doc, "\n\n---\n\nOnce scheduled, mark them:\n\n    %s\n", markCmd)

	if err := os.MkdirAll(filepath.Dir(outPath), 0o755); err != nil {
		return LinkedInBatchResult{}, err
	}
	if err := os.WriteFile(outPath, []byte(doc.String()), 0o644); err != nil {
		return LinkedInBatchResult{}, fmt.Errorf("writing batch file: %w", err)
	}

	return LinkedInBatchResult{
		Days:        days,
		Numbers:     numbers,
		OutPath:     outPath,
		MarkCommand: markCmd,
		Message:     fmt.Sprintf("prepared %d day(s) for %s", len(days), destination),
	}, nil
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
