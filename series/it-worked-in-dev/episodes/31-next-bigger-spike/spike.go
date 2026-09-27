package nextbiggerspike

// Requests per period, oldest first. For capacity planning the dashboard
// shows, for every period, how long until traffic next went higher than it:
// how long a peak "held". 0 means it was never beaten.

// WaitByScan is what I would write: from each period, look forward until
// something is higher.
func WaitByScan(load []int) []int {
	out := make([]int, len(load))
	for i := range load {
		for j := i + 1; j < len(load); j++ {
			if load[j] > load[i] {
				out[i] = j - i
				break
			}
		}
	}
	return out
}

// WaitByStack walks backwards, keeping the periods to the right that could
// still be somebody's answer. Anything not higher than the current period is
// hidden behind it for everyone further left, so it is dropped for good.
func WaitByStack(load []int) []int {
	out := make([]int, len(load))
	stack := make([]int, 0, 64) // indices; their loads rise from top to bottom
	for i := len(load) - 1; i >= 0; i-- {
		// Not higher than load[i] means nobody left of i will ever pick it:
		// i is nearer and at least as high.
		for len(stack) > 0 && load[stack[len(stack)-1]] <= load[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			out[i] = stack[len(stack)-1] - i
		}
		stack = append(stack, i)
	}
	return out
}
