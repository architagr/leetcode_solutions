package queue

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadMissingFileReturnsEmptyQueue(t *testing.T) {
	q, err := Load(filepath.Join(t.TempDir(), "queue.yaml"))
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if q.NextDay != 1 {
		t.Errorf("NextDay = %d, want 1", q.NextDay)
	}
	if len(q.Entries) != 0 {
		t.Errorf("Entries = %v, want empty", q.Entries)
	}
}

func TestAppendAssignsDayAndIncrements(t *testing.T) {
	q := &Queue{NextDay: 23}
	day := q.Append(Entry{
		Number:     965,
		Title:      "Univalued Binary Tree",
		Difficulty: "easy",
		Folder:     "easy_problems/901_1000/univalued_binary_tree",
		Batch:      "binary tree easy problems",
	})
	if day != 23 {
		t.Errorf("Append returned day %d, want 23", day)
	}
	if q.NextDay != 24 {
		t.Errorf("NextDay after append = %d, want 24", q.NextDay)
	}
	if len(q.Entries) != 1 || q.Entries[0].Status != StatusContentReady {
		t.Fatalf("Entries = %+v", q.Entries)
	}
}

func TestHas(t *testing.T) {
	q := &Queue{Entries: []Entry{{Number: 965}}}
	if !q.Has(965) {
		t.Error("Has(965) = false, want true")
	}
	if q.Has(1) {
		t.Error("Has(1) = true, want false")
	}
}

func TestSaveLoadRoundTrip(t *testing.T) {
	path := filepath.Join(t.TempDir(), "queue.yaml")
	q := &Queue{NextDay: 2}
	q.Append(Entry{Number: 1, Title: "Two Sum", Difficulty: "easy", Folder: "easy_problems/1_100/two_sum", Batch: "arrays"})
	if err := q.Save(path); err != nil {
		t.Fatalf("Save: %v", err)
	}
	if _, err := os.Stat(path); err != nil {
		t.Fatalf("file not written: %v", err)
	}
	loaded, err := Load(path)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if loaded.NextDay != 3 || len(loaded.Entries) != 1 || loaded.Entries[0].Number != 1 {
		t.Fatalf("loaded = %+v", loaded)
	}
}
