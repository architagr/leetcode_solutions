package leetcode

import (
	"os"
	"testing"
)

// Run with LEETCODE_LIVE=1 to probe the real API.
func TestLiveLookupProbe(t *testing.T) {
	if os.Getenv("LEETCODE_LIVE") == "" {
		t.Skip("set LEETCODE_LIVE=1 to hit the real API")
	}
	c := NewClient()
	for _, s := range []string{
		"two-sum", "roman-to-interger", "largest-common-prefix",
		"merge-sorted-list", "contains-duplicate-2",
		"longest-common-prefix", "implement-strstr",
		"binary-tree-inorder-triversal",
	} {
		m, found, err := c.FetchQuestionMeta(s)
		t.Logf("%-32s found=%-5v n=%-5d diff=%-8s err=%v", s, found, m.Number, m.Difficulty, err)
	}
}
