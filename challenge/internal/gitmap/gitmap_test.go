package gitmap

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestParseCommitSubject(t *testing.T) {
	cases := map[string]struct {
		number int
		ok     bool
	}{
		"Easy(965) Univalued Binary Tree":      {965, true},
		"medium(3652)":                         {3652, true},
		"Easy (168)":                           {168, true},
		"Binary Tree Vertical Order Traversal": {0, false},
		"upgrade go and testify":               {0, false},
	}
	for subject, want := range cases {
		number, ok := ParseCommitSubject(subject)
		if ok != want.ok || (ok && number != want.number) {
			t.Errorf("ParseCommitSubject(%q) = (%d, %v), want (%d, %v)", subject, number, ok, want.number, want.ok)
		}
	}
}

func initTestRepo(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	run := func(args ...string) {
		cmd := exec.Command("git", args...)
		cmd.Dir = dir
		if out, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("git %v: %v\n%s", args, err, out)
		}
	}
	run("init", "-q")
	run("config", "user.email", "test@example.com")
	run("config", "user.name", "Test")
	return dir
}

func commitFile(t *testing.T, dir, relPath, subject string) string {
	t.Helper()
	full := filepath.Join(dir, relPath)
	if err := os.MkdirAll(filepath.Dir(full), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(full, []byte("package p\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command("git", "-C", dir, "add", relPath)
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("git add: %v\n%s", err, out)
	}
	cmd = exec.Command("git", "-C", dir, "commit", "-q", "-m", subject)
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("git commit: %v\n%s", err, out)
	}
	cmd = exec.Command("git", "-C", dir, "rev-parse", "HEAD")
	out, err := cmd.Output()
	if err != nil {
		t.Fatal(err)
	}
	return string(out[:len(out)-1]) // trim trailing newline
}

func TestScanNewCommitsFromScratch(t *testing.T) {
	dir := initTestRepo(t)
	commitFile(t, dir, "easy_problems/1_100/two_sum/main.go", "Easy(1) Two Sum")
	newest := commitFile(t, dir, "easy_problems/901_1000/univalued_binary_tree/main.go", "Easy(965) Univalued Binary Tree")

	entries, newestSHA, err := ScanNewCommits(dir, "")
	if err != nil {
		t.Fatalf("ScanNewCommits: %v", err)
	}
	if newestSHA != newest {
		t.Errorf("newestSHA = %q, want %q", newestSHA, newest)
	}
	if entries[1] != "easy_problems/1_100/two_sum" {
		t.Errorf("entries[1] = %q", entries[1])
	}
	if entries[965] != "easy_problems/901_1000/univalued_binary_tree" {
		t.Errorf("entries[965] = %q", entries[965])
	}
}

func TestScanNewCommitsIncremental(t *testing.T) {
	dir := initTestRepo(t)
	first := commitFile(t, dir, "easy_problems/1_100/two_sum/main.go", "Easy(1) Two Sum")
	commitFile(t, dir, "easy_problems/901_1000/univalued_binary_tree/main.go", "Easy(965) Univalued Binary Tree")

	entries, _, err := ScanNewCommits(dir, first)
	if err != nil {
		t.Fatalf("ScanNewCommits: %v", err)
	}
	if _, ok := entries[1]; ok {
		t.Errorf("entries should not include commit 1 (already scanned): %v", entries)
	}
	if entries[965] != "easy_problems/901_1000/univalued_binary_tree" {
		t.Errorf("entries[965] = %q", entries[965])
	}
}

func TestSaveLoadRoundTrip(t *testing.T) {
	path := filepath.Join(t.TempDir(), "number_folder_map.yaml")
	m := &Map{LastScannedCommit: "abc123", Entries: map[int]string{1: "easy_problems/1_100/two_sum"}}
	if err := m.Save(path); err != nil {
		t.Fatalf("Save: %v", err)
	}
	loaded, err := Load(path)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if loaded.LastScannedCommit != "abc123" || loaded.Entries[1] != "easy_problems/1_100/two_sum" {
		t.Fatalf("loaded = %+v", loaded)
	}
}

func TestLoadMissingFile(t *testing.T) {
	m, err := Load(filepath.Join(t.TempDir(), "nope.yaml"))
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if m.Entries == nil || len(m.Entries) != 0 {
		t.Errorf("Entries = %v, want empty map", m.Entries)
	}
}
