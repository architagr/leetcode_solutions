package cli

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"

	"leetcode_solutions/challenge/internal/gitmap"
	"leetcode_solutions/challenge/internal/leetcode"
	"leetcode_solutions/challenge/internal/reorg"

	"gopkg.in/yaml.v3"
)

// ReorgItem is one folder the sweep looked at.
type ReorgItem struct {
	From       string `json:"from"`
	To         string `json:"to,omitempty"`
	Number     int    `json:"number,omitempty"`
	Difficulty string `json:"difficulty,omitempty"`
	Title      string `json:"title,omitempty"`
	// Via records which signal identified the question: "gitmap",
	// "folder-slug", or "readme-title".
	Via string `json:"via,omitempty"`
	// Note carries the reason for an unresolved or conflicting item.
	Note string `json:"note,omitempty"`
}

// ReorgScanResult is the sweep's plan, and what it applied.
type ReorgScanResult struct {
	Scanned    int         `json:"scanned"`
	Canonical  int         `json:"canonical"`
	Ready      []ReorgItem `json:"ready"`
	Unresolved []ReorgItem `json:"unresolved"`
	Conflicts  []ReorgItem `json:"conflicts"`
	Applied    []ReorgItem `json:"applied,omitempty"`
	DryRun     bool        `json:"dryRun"`
}

var canonicalRe = regexp.MustCompile(`^(easy|medium|hard)_problems/[0-9]+_[0-9]+/[^/]+$`)

// skipDirs are top-level directories that hold no solution folders.
var skipDirs = map[string]bool{
	"challenge": true, "common": true, "docs": true,
	".git": true, ".claude": true, "node_modules": true, ".worktrees": true,
}

// ReorgScan finds every solution folder that is not in the canonical
// {difficulty}_problems/{range}/{slug} layout and works out where each
// one belongs, identifying the question from git history first, then the
// folder name as a slug, then the title in its README.
//
// It is a dry run unless apply is true, so the plan can be reviewed
// before any folder moves.
func ReorgScan(repoRoot, mapPath, cachePath, aliasPath string, apply bool) (ReorgScanResult, error) {
	res := ReorgScanResult{DryRun: !apply}

	folders, err := solutionFolders(repoRoot)
	if err != nil {
		return res, err
	}
	res.Scanned = len(folders)

	byFolder := map[string]int{}
	if m, err := gitmap.Load(mapPath); err == nil {
		for number, folder := range m.Entries {
			byFolder[folder] = number
		}
	}

	aliases := loadAliases(aliasPath)
	cache := loadLookupCache(cachePath)
	client := leetcode.NewClient()
	dirty := false

	for _, from := range folders {
		if canonicalRe.MatchString(from) {
			res.Canonical++
			continue
		}

		item := ReorgItem{From: from}
		meta, ok := identify(client, repoRoot, from, byFolder, aliases, cache, &dirty)
		if !ok {
			item.Note = "could not identify the question from git history, the folder name, or a README title"
			res.Unresolved = append(res.Unresolved, item)
			continue
		}

		item.Number, item.Difficulty, item.Title, item.Via = meta.Number, meta.Difficulty, meta.Title, meta.via
		item.To = reorg.TargetPath(meta.Difficulty, meta.Number, meta.Slug)

		if item.To == from {
			res.Canonical++
			continue
		}
		if entries, err := os.ReadDir(filepath.Join(repoRoot, item.To)); err == nil && len(entries) > 0 {
			item.Note = "target folder already exists and is not empty"
			res.Conflicts = append(res.Conflicts, item)
			continue
		}
		res.Ready = append(res.Ready, item)
	}

	sort.Slice(res.Ready, func(i, j int) bool { return res.Ready[i].Number < res.Ready[j].Number })
	if dirty {
		saveLookupCache(cachePath, cache)
	}

	if !apply {
		return res, nil
	}
	for _, item := range res.Ready {
		slug := strings.TrimPrefix(item.To, filepath.Dir(item.To)+"/")
		if _, err := reorg.Move(repoRoot, item.From, item.Difficulty, item.Number, slug); err != nil {
			item.Note = err.Error()
			res.Conflicts = append(res.Conflicts, item)
			continue
		}
		res.Applied = append(res.Applied, item)
	}
	res.Ready = nil
	return res, nil
}

type identified struct {
	leetcode.Meta
	via string
}

// identify resolves a folder to a question, cheapest signal first.
func identify(c *leetcode.Client, repoRoot, folder string, byFolder map[string]int, aliases map[string]string, cache map[string]*leetcode.Meta, dirty *bool) (identified, bool) {
	base := filepath.Base(folder)

	// A hand-written alias outranks every guess: it is there precisely
	// because the folder name does not match the real slug.
	if slug, ok := aliases[folder]; ok {
		if m, ok := lookupSlug(c, slug, cache, dirty); ok {
			return identified{Meta: m, via: "alias"}, true
		}
	}

	// Git history is authoritative and free, but only carries the
	// number, so the difficulty still has to come from LeetCode.
	if n, ok := byFolder[folder]; ok {
		if m, ok := lookupSlug(c, slugFromFolder(base), cache, dirty); ok {
			m.Number = n
			return identified{Meta: m, via: "gitmap"}, true
		}
	}

	if m, ok := lookupSlug(c, slugFromFolder(base), cache, dirty); ok {
		return identified{Meta: m, via: "folder-slug"}, true
	}

	if title := readmeTitle(filepath.Join(repoRoot, folder, "README.md")); title != "" {
		if m, ok := lookupSlug(c, slugFromTitle(title), cache, dirty); ok {
			return identified{Meta: m, via: "readme-title"}, true
		}
	}
	return identified{}, false
}

func lookupSlug(c *leetcode.Client, slug string, cache map[string]*leetcode.Meta, dirty *bool) (leetcode.Meta, bool) {
	if slug == "" {
		return leetcode.Meta{}, false
	}
	if m, seen := cache[slug]; seen {
		if m == nil {
			return leetcode.Meta{}, false
		}
		return *m, true
	}
	// Spread the requests out; this sweep makes a few hundred of them.
	time.Sleep(250 * time.Millisecond)
	m, found, err := c.FetchQuestionMeta(slug)
	if err != nil {
		return leetcode.Meta{}, false
	}
	*dirty = true
	if !found {
		cache[slug] = nil
		return leetcode.Meta{}, false
	}
	cache[slug] = &m
	return m, true
}

func slugFromFolder(base string) string {
	return strings.ReplaceAll(strings.ToLower(base), "_", "-")
}

var nonSlug = regexp.MustCompile(`[^a-z0-9]+`)

func slugFromTitle(title string) string {
	s := nonSlug.ReplaceAllString(strings.ToLower(title), "-")
	return strings.Trim(s, "-")
}

// readmeTitle reads the first markdown heading out of a README, which is
// the question's real title in the folders this repo already has.
func readmeTitle(path string) string {
	f, err := os.Open(path)
	if err != nil {
		return ""
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if !strings.HasPrefix(line, "#") {
			continue
		}
		title := strings.TrimSpace(strings.TrimLeft(line, "#"))
		// Drop a leading "14. " style number if one is there.
		if i := strings.Index(title, ". "); i > 0 && i < 5 {
			if _, err := fmt.Sscanf(title[:i], "%d", new(int)); err == nil {
				title = strings.TrimSpace(title[i+2:])
			}
		}
		return title
	}
	return ""
}

// solutionFolders lists every directory holding at least one .go file,
// relative to repoRoot.
func solutionFolders(repoRoot string) ([]string, error) {
	seen := map[string]bool{}
	err := filepath.WalkDir(repoRoot, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, relErr := filepath.Rel(repoRoot, path)
		if relErr != nil {
			return nil
		}
		if d.IsDir() {
			if skipDirs[d.Name()] || strings.HasPrefix(d.Name(), ".") && rel != "." {
				return filepath.SkipDir
			}
			return nil
		}
		if strings.HasSuffix(d.Name(), ".go") {
			seen[filepath.Dir(rel)] = true
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	out := make([]string, 0, len(seen))
	for d := range seen {
		if d != "." {
			out = append(out, d)
		}
	}
	sort.Strings(out)
	return out, nil
}

// loadAliases reads the hand-maintained folder -> slug overrides.
func loadAliases(path string) map[string]string {
	out := map[string]string{}
	if path == "" {
		return out
	}
	raw, err := os.ReadFile(path)
	if err != nil {
		return out
	}
	var doc struct {
		Aliases map[string]string `yaml:"aliases"`
	}
	if err := yaml.Unmarshal(raw, &doc); err != nil {
		return out
	}
	for k, v := range doc.Aliases {
		out[filepath.Clean(k)] = v
	}
	return out
}

func loadLookupCache(path string) map[string]*leetcode.Meta {
	cache := map[string]*leetcode.Meta{}
	if path == "" {
		return cache
	}
	raw, err := os.ReadFile(path)
	if err != nil {
		return cache
	}
	_ = json.Unmarshal(raw, &cache)
	return cache
}

func saveLookupCache(path string, cache map[string]*leetcode.Meta) {
	if path == "" {
		return
	}
	raw, err := json.MarshalIndent(cache, "", "  ")
	if err != nil {
		return
	}
	_ = os.WriteFile(path, raw, 0o644)
}
