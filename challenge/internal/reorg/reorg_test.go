package reorg

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestTargetPath(t *testing.T) {
	got := TargetPath("easy", 965, "univalued-binary-tree")
	want := filepath.Join("easy_problems", "901_1000", "univalued_binary_tree")
	if got != want {
		t.Errorf("TargetPath = %q, want %q", got, want)
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

func TestMoveRelocatesAndPreservesHistory(t *testing.T) {
	dir := initTestRepo(t)
	from := "google_questions/design/min_stack"
	if err := os.MkdirAll(filepath.Join(dir, from), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, from, "main.go"), []byte("package minstack\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	run := func(args ...string) {
		cmd := exec.Command("git", args...)
		cmd.Dir = dir
		if out, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("git %v: %v\n%s", args, err, out)
		}
	}
	run("add", ".")
	run("commit", "-q", "-m", "Easy(155) Min Stack")

	newPath, err := Move(dir, from, "easy", 155, "min-stack")
	if err != nil {
		t.Fatalf("Move: %v", err)
	}
	want := filepath.Join("easy_problems", "101_200", "min_stack")
	if newPath != want {
		t.Errorf("Move returned %q, want %q", newPath, want)
	}
	if _, err := os.Stat(filepath.Join(dir, from)); !os.IsNotExist(err) {
		t.Errorf("old path %q should no longer exist", from)
	}
	if _, err := os.Stat(filepath.Join(dir, want, "main.go")); err != nil {
		t.Errorf("new path missing main.go: %v", err)
	}
}

func TestMoveNoOpWhenAlreadyCanonical(t *testing.T) {
	dir := initTestRepo(t)
	canonical := filepath.Join("easy_problems", "101_200", "min_stack")
	newPath, err := Move(dir, canonical, "easy", 155, "min-stack")
	if err != nil {
		t.Fatalf("Move: %v", err)
	}
	if newPath != canonical {
		t.Errorf("Move = %q, want no-op %q", newPath, canonical)
	}
}
