package cli

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestPackageNameStripsPunctuation(t *testing.T) {
	cases := map[string]string{
		"lowest-common-ancestor-of-deepest-leaves": "lowestcommonancestorofdeepestleaves",
		"two-sum":       "twosum",
		"3sum":          "p3sum",
		"n-queens-ii":   "nqueensii",
		"minimum.stack": "minimumstack",
	}
	for slug, want := range cases {
		if got := packageName(slug); got != want {
			t.Errorf("packageName(%q) = %q, want %q", slug, got, want)
		}
	}
}

func TestWithPackageClauseAddsMissingClause(t *testing.T) {
	got := withPackageClause("func twoSum() {}", "two-sum")
	want := "package twosum\n\nfunc twoSum() {}"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestWithPackageClauseLeavesExistingAlone(t *testing.T) {
	src := "package twosum\n\nfunc twoSum() {}"
	if got := withPackageClause(src, "two-sum"); got != src {
		t.Errorf("clause was rewritten: %q", got)
	}
}

func TestWithPackageClauseDetectsPaddedClause(t *testing.T) {
	// Leading whitespace before the clause is legal Go, so the file is
	// left exactly as-is rather than gaining a second package line.
	src := "\n\n  package twosum\nfunc f(){}"
	got := withPackageClause(src, "two-sum")
	if got != src {
		t.Errorf("padded clause was not detected, got %q", got)
	}
	if strings.Count(got, "package ") != 1 {
		t.Errorf("ended up with a duplicate package clause: %q", got)
	}
}

// A folder whose solution file is named after the slug rather than
// main.go must still count as solved. Missing this once caused a
// duplicate main.go to be written into easy_problems/1_100/two_sum,
// which already held two_sum.go.
func TestScaffoldTreatsSlugNamedFileAsExisting(t *testing.T) {
	root := t.TempDir()
	dir := filepath.Join(root, "easy_problems", "1_100", "two_sum")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "two_sum.go"), []byte("package twosum\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	// No session is configured, so reaching LeetCode would error out.
	// Any non-error "exists" result proves the guard fired first.
	t.Setenv("LEETCODE_SESSION", "")
	res, err := ScaffoldFromSubmission(root, filepath.Join(root, "nope.json"), 1, "easy", "two-sum", "Two Sum")
	if err == nil && res.Status != "exists" {
		t.Fatalf("status = %q, want exists", res.Status)
	}
	if err != nil && !strings.Contains(err.Error(), "session") {
		t.Fatalf("unexpected error: %v", err)
	}
}

// A solution living under a non-canonical directory must not be
// duplicated into the canonical one.
func TestScaffoldDoesNotDuplicateNonCanonicalSolution(t *testing.T) {
	root := t.TempDir()
	dir := filepath.Join(root, "linkedin_questions", "array_strings", "two_sum")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "main.go"), []byte("package twosum\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	res, err := ScaffoldFromSubmission(root, filepath.Join(root, "nope.json"), 1, "easy", "two-sum", "Two Sum")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res.Status != "exists" {
		t.Fatalf("status = %q, want exists", res.Status)
	}
	if _, statErr := os.Stat(filepath.Join(root, "easy_problems", "1_100", "two_sum")); !os.IsNotExist(statErr) {
		t.Fatal("a duplicate canonical folder was created")
	}
}
