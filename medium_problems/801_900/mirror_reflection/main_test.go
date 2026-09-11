package mirror_reflection

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestMirrorReflection(t *testing.T) {
	tests := []struct {
		name string
		p    int
		q    int
		want int
	}{
		{name: "example 1", p: 2, q: 1, want: 2},
		{name: "example 2", p: 3, q: 1, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MirrorReflection(tt.p, tt.q); got != tt.want {
				t.Errorf("MirrorReflection() = %v, want %v", got, tt.want)
			}
		})
	}
}
