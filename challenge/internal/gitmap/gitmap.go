// Package gitmap builds and maintains a LeetCode question number →
// solution folder map by scanning git commit history, following this
// repo's "Type(Number) Title" commit message convention (e.g.
// "Easy(965) Univalued Binary Tree").
package gitmap

import (
	"fmt"
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
		if exitErr, ok := err.(*exec.ExitError); ok {
			return nil, sinceSHA, fmt.Errorf("git log %s: %w: %s", rangeArg, err, exitErr.Stderr)
		}
		return nil, sinceSHA, fmt.Errorf("git log %s: %w", rangeArg, err)
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

// folderForCommit returns (string, bool) rather than propagating an
// error: it's called once per matching commit inside ScanNewCommits's
// bulk history walk, and a `git show` failure on one commit (e.g. a
// shallow clone missing that object) reasonably degrades to "skip this
// commit" rather than aborting the entire scan. This does mean a real
// git failure here is indistinguishable from "no matching .go file in
// this commit" — accepted tradeoff, not an oversight.
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
