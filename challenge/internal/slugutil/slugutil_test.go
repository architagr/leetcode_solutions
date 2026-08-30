package slugutil

import "testing"

func TestFolderName(t *testing.T) {
	cases := map[string]string{
		"univalued-binary-tree": "univalued_binary_tree",
		"two-sum":               "two_sum",
		"already_snake":         "already_snake",
	}
	for slug, want := range cases {
		if got := FolderName(slug); got != want {
			t.Errorf("FolderName(%q) = %q, want %q", slug, got, want)
		}
	}
}

func TestRangeBucket(t *testing.T) {
	cases := map[int]string{
		1:    "1_100",
		55:   "1_100",
		100:  "1_100",
		101:  "101_200",
		965:  "901_1000",
		1001: "1001_1100",
	}
	for number, want := range cases {
		if got := RangeBucket(number); got != want {
			t.Errorf("RangeBucket(%d) = %q, want %q", number, got, want)
		}
	}
}
