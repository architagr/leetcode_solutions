package cli

import (
	"os"
	"path/filepath"
	"strings"

	"leetcode_solutions/challenge/internal/leetcode"
	"leetcode_solutions/challenge/internal/reorg"
	"leetcode_solutions/challenge/internal/resolver"
)

// langExt maps LeetCode's lang slugs to the file a solution folder
// holds. The repo is Go-only, so anything else is pulled down under its
// own extension and flagged rather than dropped into main.go.
var langExt = map[string]string{
	"golang": ".go",
	"go":     ".go",
}

// ScaffoldResult reports what ScaffoldFromSubmission did (or why it
// declined to do anything).
type ScaffoldResult struct {
	// Status is "scaffolded", "not-solved", or "exists".
	Status string `json:"status"`
	// Source, on an "exists" result, says where the already-present
	// solution was found, matching Resolve's source values.
	Source string `json:"source,omitempty"`
	// Path is the canonical folder, set for "scaffolded" and "exists".
	Path string `json:"path,omitempty"`
	// File is the solution file written, relative to the repo root.
	File string `json:"file,omitempty"`
	// Lang is the LeetCode lang slug of the submission that was pulled.
	Lang string `json:"lang,omitempty"`
	// SubmittedAt is the accepted submission's date (RFC3339, UTC).
	SubmittedAt string `json:"submittedAt,omitempty"`
	// NonGo is true when the accepted submission was not Go, in which
	// case the code was written under its own extension and the folder
	// still needs a Go port by hand.
	NonGo bool `json:"nonGo,omitempty"`
}

// ScaffoldFromSubmission creates the canonical folder for a question the
// user has solved on LeetCode but never committed here, writing their
// own most recent accepted submission into it.
//
// It never overwrites: if the canonical folder already holds a solution
// file, it reports "exists" and leaves the tree untouched.
func ScaffoldFromSubmission(repoRoot, sessionPath string, number int, difficulty, slug, title string) (ScaffoldResult, error) {
	// Runs before authentication on purpose: a question that is already
	// in the repo needs neither a cookie nor a network round trip.
	// Guard against duplicating a solution that is already here under a
	// non-canonical path. The skill resolves before it scaffolds, so this
	// should not fire in the normal flow, but scaffolding a second copy
	// is destructive enough to be worth checking independently of call
	// order.
	if existing, err := resolver.FindAnywhere(repoRoot, number, difficulty, slug); err == nil && existing.Path != "" {
		return ScaffoldResult{Status: "exists", Path: existing.Path, Source: string(existing.Source)}, nil
	}

	sess, err := leetcode.LoadSession(sessionPath)
	if err != nil {
		return ScaffoldResult{}, err
	}
	client := leetcode.NewClient()
	client.Session = &sess

	sub, found, err := client.FetchAcceptedSubmission(slug, "golang")
	if err != nil {
		return ScaffoldResult{}, err
	}
	if !found {
		return ScaffoldResult{Status: "not-solved"}, nil
	}

	rel := reorg.TargetPath(difficulty, number, slug)
	abs := filepath.Join(repoRoot, rel)

	ext, isGo := langExt[sub.Lang]
	if !isGo {
		ext = "." + sub.Lang
	}
	name := "main" + ext

	// Final guard before writing. Solution folders do not all use
	// main.go — older ones are named after the slug — so treat any
	// non-empty canonical folder as already solved rather than only
	// looking for the file this call would write.
	if entries, err := os.ReadDir(abs); err == nil && len(entries) > 0 {
		return ScaffoldResult{Status: "exists", Path: rel, Source: "canonical"}, nil
	}

	if err := os.MkdirAll(abs, 0o755); err != nil {
		return ScaffoldResult{}, err
	}

	body := sub.Code
	if isGo {
		body = withPackageClause(body, slug)
	}
	if !strings.HasSuffix(body, "\n") {
		body += "\n"
	}
	if err := os.WriteFile(filepath.Join(abs, name), []byte(body), 0o644); err != nil {
		return ScaffoldResult{}, err
	}

	res := ScaffoldResult{
		Status: "scaffolded",
		Path:   rel,
		File:   filepath.Join(rel, name),
		Lang:   sub.Lang,
		NonGo:  !isGo,
	}
	if !sub.Timestamp.IsZero() {
		res.SubmittedAt = sub.Timestamp.Format("2006-01-02")
	}
	return res, nil
}

// withPackageClause prepends the package clause LeetCode's editor omits.
// Submitted Go code is a bare set of declarations, which does not
// compile as a file on its own.
func withPackageClause(code, slug string) string {
	trimmed := strings.TrimLeft(code, " \t\r\n")
	if strings.HasPrefix(trimmed, "package ") {
		return code
	}
	return "package " + packageName(slug) + "\n\n" + trimmed
}

// packageName turns a question slug into the all-lowercase, punctuation
// -free identifier the existing solution folders use.
func packageName(slug string) string {
	var b strings.Builder
	for _, r := range strings.ToLower(slug) {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			b.WriteRune(r)
		}
	}
	out := b.String()
	if out == "" || (out[0] >= '0' && out[0] <= '9') {
		out = "p" + out
	}
	return out
}
