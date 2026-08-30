// Package slugutil converts LeetCode question slugs and numbers into this
// repo's folder-naming conventions (snake_case folder names, 100-wide
// number-range buckets like "901_1000").
package slugutil

import (
	"fmt"
	"strings"
)

// FolderName converts a LeetCode URL slug (dash-case, e.g.
// "univalued-binary-tree") into this repo's folder naming convention
// (snake_case, e.g. "univalued_binary_tree").
func FolderName(slug string) string {
	return strings.ReplaceAll(slug, "-", "_")
}

// RangeBucket returns the 100-wide folder bucket name for a LeetCode
// question number, matching the existing repo convention
// (easy_problems/901_1000, easy_problems/1_100, ...).
func RangeBucket(number int) string {
	start := ((number-1)/100)*100 + 1
	end := start + 99
	return fmt.Sprintf("%d_%d", start, end)
}
