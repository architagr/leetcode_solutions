package compareversionnumbers

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestCompareVersion(t *testing.T) {
	tests := []struct {
		name     string
		version1 string
		version2 string
		want     int
	}{
		{name: "example 1", version1: "1.2", version2: "1.10", want: -1},
		{name: "example 2", version1: "1.01", version2: "1.001", want: 0},
		{name: "example 3", version1: "1.0", version2: "1.0.0.0", want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compareVersion(tt.version1, tt.version2); got != tt.want {
				t.Errorf("compareVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
