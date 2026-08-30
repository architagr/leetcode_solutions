package resolver

import (
	"os"
	"path/filepath"
	"testing"
)

func mkdirs(t *testing.T, root string, dirs ...string) {
	t.Helper()
	for _, d := range dirs {
		if err := os.MkdirAll(filepath.Join(root, d), 0o755); err != nil {
			t.Fatal(err)
		}
	}
}

func TestCandidatePathsOrderAndSources(t *testing.T) {
	root := t.TempDir()
	mkdirs(t, root, "google_questions/design/min_stack")

	candidates := CandidatePaths(root, 155, "easy", "min-stack")
	if len(candidates) < 3 {
		t.Fatalf("expected at least 3 static candidates, got %d: %+v", len(candidates), candidates)
	}
	if candidates[0].Path != filepath.Join("easy_problems", "101_200", "min_stack") || candidates[0].Source != "canonical" {
		t.Errorf("candidates[0] = %+v", candidates[0])
	}
	if candidates[1].Path != filepath.Join("easy_problems", "min_stack") || candidates[1].Source != "straggler" {
		t.Errorf("candidates[1] = %+v", candidates[1])
	}
	found := false
	for _, c := range candidates {
		if c.Path == filepath.Join("google_questions", "design", "min_stack") && c.Source == "google_questions" {
			found = true
		}
	}
	if !found {
		t.Errorf("expected a google_questions candidate, got %+v", candidates)
	}
}

func TestFindPicksFirstExisting(t *testing.T) {
	root := t.TempDir()
	mkdirs(t, root, "easy_problems/min_stack") // only the straggler location exists

	candidates := CandidatePaths(root, 155, "easy", "min-stack")
	loc, ok := Find(root, candidates)
	if !ok {
		t.Fatal("Find: expected a match")
	}
	if loc.Source != "straggler" {
		t.Errorf("Find matched %+v, want straggler", loc)
	}
}

func TestFindNoMatch(t *testing.T) {
	root := t.TempDir()
	candidates := CandidatePaths(root, 155, "easy", "min-stack")
	if _, ok := Find(root, candidates); ok {
		t.Error("Find: expected no match on empty repo")
	}
}

func TestFuzzyFind(t *testing.T) {
	root := t.TempDir()
	// Folder name predates a LeetCode slug change, but content matches.
	mkdirs(t, root, "medium_problems/701_800/Min_Stack")

	loc, ok := FuzzyFind(root, "min-stack")
	if !ok {
		t.Fatal("FuzzyFind: expected a match")
	}
	if loc.Path != filepath.Join("medium_problems", "701_800", "Min_Stack") || loc.Source != "fuzzy" {
		t.Errorf("FuzzyFind matched %+v", loc)
	}
}
