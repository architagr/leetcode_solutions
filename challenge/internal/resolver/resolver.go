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

// Source identifies which known repo layout a Location came from.
type Source string

const (
	SourceCanonical         Source = "canonical"
	SourceStraggler         Source = "straggler"
	SourceGoogleQuestions   Source = "google_questions"
	SourceLinkedInQuestions Source = "linkedin_questions"
	SourceGitmap            Source = "gitmap"
	SourceFuzzy             Source = "fuzzy"
)

// topicRootSources maps a topic-root directory name to its Source
// constant, for the topic roots CandidatePaths globs.
var topicRootSources = map[string]Source{
	"google_questions":   SourceGoogleQuestions,
	"linkedin_questions": SourceLinkedInQuestions,
}

// Location is one place a solved question might live in the repo, with
// Path relative to the repo root.
type Location struct {
	Path      string
	Canonical bool
	Source    Source
}

// CandidatePaths returns the static, filesystem-derivable places a
// solved question might live, in priority order. It does not consult
// git history or do fuzzy name matching — see FuzzyFind for that.
func CandidatePaths(repoRoot string, number int, difficulty, slug string) []Location {
	folder := slugutil.FolderName(slug)
	diffRoot := difficulty + "_problems"

	locs := []Location{
		{Path: filepath.Join(diffRoot, slugutil.RangeBucket(number), folder), Canonical: true, Source: SourceCanonical},
		{Path: filepath.Join(diffRoot, folder), Canonical: false, Source: SourceStraggler},
	}
	for _, topicRoot := range []string{"google_questions", "linkedin_questions"} {
		matches, _ := filepath.Glob(filepath.Join(repoRoot, topicRoot, "*", folder))
		for _, m := range matches {
			rel, err := filepath.Rel(repoRoot, m)
			if err != nil {
				continue
			}
			locs = append(locs, Location{Path: rel, Canonical: false, Source: topicRootSources[topicRoot]})
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
		_ = filepath.WalkDir(start, func(path string, d fs.DirEntry, err error) error {
			if err != nil || ok || !d.IsDir() {
				return nil
			}
			if strings.ToLower(d.Name()) == target {
				rel, relErr := filepath.Rel(repoRoot, path)
				if relErr == nil {
					found = Location{Path: rel, Canonical: false, Source: SourceFuzzy}
					ok = true
				}
				return fs.SkipAll
			}
			return nil
		})
		if ok {
			break
		}
	}
	return found, ok
}

// FindAnywhere reports whether a question already has a folder anywhere
// in the repo — canonical, straggler, topic directory, or under a
// differently-cased name. It is the static half of cli.Resolve, without
// the git-history map, and exists so a caller that is about to create a
// folder can cheaply avoid duplicating an existing solution.
//
// A zero Location with a nil error means nothing was found.
func FindAnywhere(repoRoot string, number int, difficulty, slug string) (Location, error) {
	if loc, ok := Find(repoRoot, CandidatePaths(repoRoot, number, difficulty, slug)); ok {
		return loc, nil
	}
	if loc, ok := FuzzyFind(repoRoot, slug); ok {
		return loc, nil
	}
	return Location{}, nil
}
