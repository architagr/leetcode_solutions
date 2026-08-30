package cli

import (
	"os"
	"path/filepath"
	"testing"

	"leetcode_solutions/challenge/internal/gitmap"
	"leetcode_solutions/challenge/internal/hero"
	"leetcode_solutions/challenge/internal/queue"
)

func TestResolveFindsCanonical(t *testing.T) {
	root := t.TempDir()
	target := filepath.Join(root, "easy_problems", "901_1000", "univalued_binary_tree")
	if err := os.MkdirAll(target, 0o755); err != nil {
		t.Fatal(err)
	}
	mapPath := filepath.Join(root, "challenge", "number_folder_map.yaml")

	result, err := Resolve(root, mapPath, 965, "easy", "univalued-binary-tree")
	if err != nil {
		t.Fatalf("Resolve: %v", err)
	}
	if result.Status != "resolved" || result.Source != "canonical" {
		t.Errorf("result = %+v", result)
	}
}

func TestResolveUnresolved(t *testing.T) {
	root := t.TempDir()
	mapPath := filepath.Join(root, "challenge", "number_folder_map.yaml")

	result, err := Resolve(root, mapPath, 999999, "easy", "does-not-exist")
	if err != nil {
		t.Fatalf("Resolve: %v", err)
	}
	if result.Status != "unresolved" {
		t.Errorf("result = %+v, want unresolved", result)
	}
}

func TestResolveFallsBackToGitmap(t *testing.T) {
	root := t.TempDir()
	// Actual folder lives somewhere resolver's static candidates and
	// FuzzyFind's easy/medium/hard walk never look (outside all three
	// difficulty roots), so only the gitmap entry can find it.
	folder := filepath.Join("archive", "965_univalued_binary_tree")
	if err := os.MkdirAll(filepath.Join(root, folder), 0o755); err != nil {
		t.Fatal(err)
	}
	mapPath := filepath.Join(root, "challenge", "number_folder_map.yaml")
	if err := os.MkdirAll(filepath.Dir(mapPath), 0o755); err != nil {
		t.Fatal(err)
	}
	m := &gitmap.Map{Entries: map[int]string{965: folder}}
	if err := m.Save(mapPath); err != nil {
		t.Fatal(err)
	}

	result, err := Resolve(root, mapPath, 965, "easy", "univalued-binary-tree")
	if err != nil {
		t.Fatalf("Resolve: %v", err)
	}
	if result.Status != "resolved" || result.Source != "gitmap" || result.Path != folder {
		t.Errorf("result = %+v", result)
	}
}

func TestResolveFallsBackToFuzzy(t *testing.T) {
	root := t.TempDir()
	mapPath := filepath.Join(root, "challenge", "number_folder_map.yaml")
	// Folder name predates a LeetCode slug change, but content matches
	// case-insensitively; the range bucket (701_800) doesn't match
	// RangeBucket(155) (101_200), so it can't be found as a static
	// candidate either.
	if err := os.MkdirAll(filepath.Join(root, "medium_problems", "701_800", "Min_Stack"), 0o755); err != nil {
		t.Fatal(err)
	}

	result, err := Resolve(root, mapPath, 155, "medium", "min-stack")
	if err != nil {
		t.Fatalf("Resolve: %v", err)
	}
	wantPath := filepath.Join("medium_problems", "701_800", "Min_Stack")
	if result.Status != "resolved" || result.Source != "fuzzy" || result.Path != wantPath {
		t.Errorf("result = %+v", result)
	}
}

func TestQueueAppendAndHas(t *testing.T) {
	queuePath := filepath.Join(t.TempDir(), "queue.yaml")

	has, err := QueueHas(queuePath, 965)
	if err != nil {
		t.Fatalf("QueueHas: %v", err)
	}
	if has {
		t.Error("QueueHas on empty queue = true, want false")
	}

	day, err := QueueAppend(queuePath, queue.Entry{
		Number: 965, Title: "Univalued Binary Tree", Difficulty: "easy",
		Folder: "easy_problems/901_1000/univalued_binary_tree", Batch: "binary tree easy problems",
	})
	if err != nil {
		t.Fatalf("QueueAppend: %v", err)
	}
	if day != 1 {
		t.Errorf("day = %d, want 1", day)
	}

	has, err = QueueHas(queuePath, 965)
	if err != nil {
		t.Fatalf("QueueHas: %v", err)
	}
	if !has {
		t.Error("QueueHas after append = false, want true")
	}
}

func TestCompaniesLookup(t *testing.T) {
	path := filepath.Join(t.TempDir(), "dataset.json")
	if err := os.WriteFile(path, []byte(`{"965":["Google"]}`), 0o644); err != nil {
		t.Fatal(err)
	}
	got, err := CompaniesLookup(path, 965)
	if err != nil {
		t.Fatalf("CompaniesLookup: %v", err)
	}
	if len(got) != 1 || got[0] != "Google" {
		t.Errorf("got = %v", got)
	}
}

func TestHeroRenderHTML(t *testing.T) {
	tmplPath := filepath.Join(t.TempDir(), "t.html")
	if err := os.WriteFile(tmplPath, []byte(`Day {{.Day}}: {{.Title}}`), 0o644); err != nil {
		t.Fatal(err)
	}
	outPath := filepath.Join(t.TempDir(), "out.html")

	if err := HeroRenderHTML(tmplPath, heroDataFixture(), outPath); err != nil {
		t.Fatalf("HeroRenderHTML: %v", err)
	}
	data, err := os.ReadFile(outPath)
	if err != nil {
		t.Fatalf("read output: %v", err)
	}
	if string(data) != "Day 1: Two Sum" {
		t.Errorf("output = %q", data)
	}
}

func heroDataFixture() hero.Data {
	return hero.Data{Day: 1, Title: "Two Sum"}
}
