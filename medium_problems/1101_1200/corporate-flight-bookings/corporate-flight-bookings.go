package corporateflightbookings

func corpFlightBookings(bookings [][]int, n int) []int {
	temp := make([]int, n+1)
	for _, booking := range bookings {
		temp[booking[0]-1] += booking[2]
		temp[booking[1]] -= booking[2]
	}
	result := make([]int, n)
	result[0] = temp[0]
	for i := 1; i < n; i++ {
		result[i] = result[i-1] + temp[i]
	}
	return result
}
