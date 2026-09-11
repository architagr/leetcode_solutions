package keys_and_rooms

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestCanVisitAllRooms(t *testing.T) {
	tests := []struct {
		name  string
		rooms [][]int
		want  bool
	}{
		{name: "example 1", rooms: [][]int{[]int{1}, []int{2}, []int{3}, []int{}}, want: true},
		{name: "example 2", rooms: [][]int{[]int{1, 3}, []int{3, 0, 1}, []int{2}, []int{0}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canVisitAllRooms(tt.rooms); got != tt.want {
				t.Errorf("canVisitAllRooms() = %v, want %v", got, tt.want)
			}
		})
	}
}
