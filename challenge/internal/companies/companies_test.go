package companies

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadMissingFileReturnsEmptyDataset(t *testing.T) {
	ds, err := Load(filepath.Join(t.TempDir(), "nope.json"))
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if len(ds.Lookup(1)) != 0 {
		t.Errorf("Lookup on empty dataset = %v, want empty", ds.Lookup(1))
	}
}

func TestLoadAndLookup(t *testing.T) {
	path := filepath.Join(t.TempDir(), "dataset.json")
	if err := os.WriteFile(path, []byte(`{"965": ["Google", "Amazon"]}`), 0o644); err != nil {
		t.Fatal(err)
	}
	ds, err := Load(path)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	got := ds.Lookup(965)
	if len(got) != 2 || got[0] != "Google" || got[1] != "Amazon" {
		t.Errorf("Lookup(965) = %v", got)
	}
	if len(ds.Lookup(1)) != 0 {
		t.Errorf("Lookup(1) = %v, want empty", ds.Lookup(1))
	}
}
