package countdayswithoutmeetings

import "sort"

func countDays(days int, meetings [][]int) int {
	freeDays, latestEnd := 0, 0

	sort.Slice(meetings, func(i, j int) bool {
		if meetings[i][0] == meetings[j][0] {
			return meetings[i][1] < meetings[j][1]
		}
		return meetings[i][0] < meetings[j][0]
	})

	for _, meeting := range meetings {
		start, end := meeting[0], meeting[1]

		// Add current range of days without a meeting
		if start > latestEnd+1 {
			freeDays += start - latestEnd - 1
		}

		// Update latest meeting end
		latestEnd = maxValue(latestEnd, end)
	}

	// Add all days after the last day of meetings
	freeDays += days - latestEnd

	return freeDays
}

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
