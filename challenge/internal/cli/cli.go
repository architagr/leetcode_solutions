// Package cli holds the orchestration logic behind each leetcodectl
// subcommand, kept separate from cmd/leetcodectl/main.go so it's
// testable without spawning a subprocess.
package cli

import (
	"fmt"
	"os"
	"path/filepath"

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
