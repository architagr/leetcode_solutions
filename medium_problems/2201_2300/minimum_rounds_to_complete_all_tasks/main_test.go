package minimumroundstocompletealltasks

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestMinimumRounds(t *testing.T) {
	tests := []struct {
		name  string
		tasks []int
		want  int
	}{
		{name: "example 1", tasks: []int{2, 2, 3, 3, 2, 4, 4, 4, 4, 4}, want: 4},
		{name: "example 2", tasks: []int{2, 3, 3}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumRounds(tt.tasks); got != tt.want {
				t.Errorf("minimumRounds() = %v, want %v", got, tt.want)
			}
		})
	}
}
