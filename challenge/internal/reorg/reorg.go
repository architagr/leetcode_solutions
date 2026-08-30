// Package reorg relocates an already-solved question's folder from a
// non-canonical location (a root straggler, or one of the old
// topic-organized directories) into the canonical
// {difficulty}_problems/{range}/{slug} layout, preserving git history
// via `git mv`.
package reorg

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"leetcode_solutions/challenge/internal/slugutil"
)

// TargetPath returns the canonical destination path (relative to the
// repo root) a solved question should live at.
func TargetPath(difficulty string, number int, slug string) string {
	return filepath.Join(difficulty+"_problems", slugutil.RangeBucket(number), slugutil.FolderName(slug))
}

// Move git-mv's fromPath (relative to repoRoot) to the canonical
// location for (difficulty, number, slug), creating parent directories
// as needed, and returns the new relative path. If fromPath is already
// the canonical location, it's a no-op.
func Move(repoRoot, fromPath, difficulty string, number int, slug string) (string, error) {
	target := TargetPath(difficulty, number, slug)
	if fromPath == target {
		return target, nil
	}
	if err := os.MkdirAll(filepath.Join(repoRoot, filepath.Dir(target)), 0o755); err != nil {
		return "", err
	}
	cmd := exec.Command("git", "-C", repoRoot, "mv", fromPath, target)
	if out, err := cmd.CombinedOutput(); err != nil {
		return "", fmt.Errorf("git mv %s -> %s: %w: %s", fromPath, target, err, out)
	}
	return target, nil
}
