package dailytemperatures

import (
	"reflect"
	"testing"
)

// Cases are the worked examples from the problem statement on LeetCode.
func TestDailyTemperatures(t *testing.T) {
	tests := []struct {
		name         string
		temperatures []int
		want         []int
	}{
		{name: "example 1", temperatures: []int{73, 74, 75, 71, 69, 72, 76, 73}, want: []int{1, 1, 4, 2, 1, 1, 0, 0}},
		{name: "example 2", temperatures: []int{30, 40, 50, 60}, want: []int{1, 1, 1, 0}},
		{name: "example 3", temperatures: []int{30, 60, 90}, want: []int{1, 1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dailyTemperatures(tt.temperatures); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dailyTemperatures() = %v, want %v", got, tt.want)
			}
		})
	}
}
