package insert_interval

func InsertInterval(intervals [][]int, newInterval []int) [][]int {
	result := make([][]int, 0)
	n := len(intervals)
	i := 0
	// Add all the elements to the list which are less then the new interval
	for i < n && intervals[i][1] < newInterval[0] {
		result = append(result, intervals[i])
		i++
	}
	// merge new interval with overlapping interlapping interval and then add it to the list
	// The reason of using intervals[i][0]<newInterval[1] is, that we already know that intervals[i][0] and newInterval[0] overlaps thats why the above while loop is exited and so
	// now the while loop should run for all the intervals which overlaps and so checking the start of interval is less that the end of newinterval.
	for i < n && intervals[i][0] <= newInterval[1] {
		newInterval[0] = minValue(intervals[i][0], newInterval[0])
		newInterval[1] = maxValue(intervals[i][1], newInterval[1])
		i++
	}
	result = append(result, newInterval)

	// add all elements to the list which are greater the new interval
	for i < n {
		result = append(result, intervals[i])
		i++
	}

	return result
}

func minValue(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
