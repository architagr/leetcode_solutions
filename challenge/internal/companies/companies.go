// Package companies looks up which companies are known to have asked a
// given LeetCode question, from a vendored, manually-refreshed dataset
// snapshot (LeetCode's own company-tag data is Premium-only with no
// public API).
package companies

import (
	"encoding/json"
	"os"
	"strconv"
)

// Dataset maps a LeetCode question number (as a string key) to the list
// of companies known to have asked it.
type Dataset map[string][]string

// Load reads a Dataset from a JSON file. A missing file is not an error
// — it returns an empty Dataset, so lookups are simply best-effort.
func Load(path string) (Dataset, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return Dataset{}, nil
		}
		return nil, err
	}
	var ds Dataset
	if err := json.Unmarshal(data, &ds); err != nil {
		return nil, err
	}
	return ds, nil
}

// Lookup returns the companies known for a question number, or an empty
// (non-nil) slice if there's no entry.
func (d Dataset) Lookup(number int) []string {
	if companies, ok := d[strconv.Itoa(number)]; ok {
		return companies
	}
	return []string{}
}
