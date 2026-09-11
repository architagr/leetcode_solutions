package averagesalaryexcludingtheminimumandmaximumsalary

import (
	"math"
	"testing"
)

// Cases are the worked examples from the problem statement on LeetCode.
func TestAverage(t *testing.T) {
	tests := []struct {
		name   string
		salary []int
		want   float64
	}{
		{name: "example 1", salary: []int{4000, 3000, 1000, 2000}, want: 2500.0},
		{name: "example 2", salary: []int{1000, 2000, 3000}, want: 2000.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := average(tt.salary); math.Abs(got-tt.want) > 1e-5 {
				t.Errorf("average() = %v, want %v", got, tt.want)
			}
		})
	}
}
