package cli

import (
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
