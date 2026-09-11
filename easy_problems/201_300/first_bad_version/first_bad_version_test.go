package first_bad_version

import "testing"

// Cases are the worked example from the problem statement, plus the two
// edges: the first version bad, and only the last one bad.
func TestFirstBadVersion(t *testing.T) {
	tests := []struct {
		name     string
		n, first int
	}{
		{name: "example 1", n: 5, first: 4},
		{name: "example 2", n: 1, first: 1},
		{name: "first version is bad", n: 10, first: 1},
		{name: "only the last is bad", n: 10, first: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := isBadVersion
			isBadVersion = func(v int) bool { return v >= tt.first }
			defer func() { isBadVersion = original }()

			if got := FirstBadVersion(tt.n); got != tt.first {
				t.Errorf("FirstBadVersion(%d) = %v, want %v", tt.n, got, tt.first)
			}
		})
	}
}
