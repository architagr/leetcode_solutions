package wallsandgates

import (
	"reflect"
	"testing"
)

// wallsAndGates works in place, so the case compares rooms afterwards.
// Cases are the worked examples from the problem statement on LeetCode.
func TestWallsAndGates(t *testing.T) {
	const inf = 2147483647
	tests := []struct {
		name  string
		rooms [][]int
		want  [][]int
	}{
		{
			name: "example 1",
			rooms: [][]int{
				{inf, -1, 0, inf},
				{inf, inf, inf, -1},
				{inf, -1, inf, -1},
				{0, -1, inf, inf},
			},
			want: [][]int{
				{3, -1, 0, 1},
				{2, 2, 1, -1},
				{1, -1, 2, -1},
				{0, -1, 3, 4},
			},
		},
		{name: "example 2", rooms: [][]int{{-1}}, want: [][]int{{-1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := make([][]int, len(tt.rooms))
			for i := range tt.rooms {
				got[i] = append([]int{}, tt.rooms[i]...)
			}
			wallsAndGates(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("wallsAndGates() = %v, want %v", got, tt.want)
			}
		})
	}
}
