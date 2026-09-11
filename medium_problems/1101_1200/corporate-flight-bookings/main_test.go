package corporateflightbookings

import (
	"reflect"
	"testing"
)

// Cases are the worked examples from the problem statement on LeetCode.
func TestCorpFlightBookings(t *testing.T) {
	tests := []struct {
		name     string
		bookings [][]int
		n        int
		want     []int
	}{
		{name: "example 1", bookings: [][]int{[]int{1, 2, 10}, []int{2, 3, 20}, []int{2, 5, 25}}, n: 5, want: []int{10, 55, 45, 25, 25}},
		{name: "example 2", bookings: [][]int{[]int{1, 2, 10}, []int{2, 2, 15}}, n: 2, want: []int{10, 25}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := corpFlightBookings(tt.bookings, tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("corpFlightBookings() = %v, want %v", got, tt.want)
			}
		})
	}
}
