package flipgame

import (
	"reflect"
	"testing"
)

// Cases are the worked examples from the problem statement on LeetCode.
func TestGeneratePossibleNextMoves(t *testing.T) {
	tests := []struct {
		name         string
		currentState string
		want         []string
	}{
		{name: "example 1", currentState: "++++", want: []string{"--++", "+--+", "++--"}},
		{name: "example 2", currentState: "+", want: []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generatePossibleNextMoves(tt.currentState)
			if len(got) == 0 && len(tt.want) == 0 {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generatePossibleNextMoves() = %v, want %v", got, tt.want)
			}
		})
	}
}
