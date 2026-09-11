package typeoftriangle

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestTriangleType(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want string
	}{
		{name: "example 1", nums: []int{3, 3, 3}, want: "equilateral"},
		{name: "example 2", nums: []int{3, 4, 5}, want: "scalene"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := triangleType(tt.nums); got != tt.want {
				t.Errorf("triangleType() = %v, want %v", got, tt.want)
			}
		})
	}
}
