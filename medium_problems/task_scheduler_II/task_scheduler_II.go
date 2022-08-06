package task_scheduler_II

func TaskSchedulerII(tasks []int, space int) int64 {
	mapTask := make(map[int]int64)
	dayCount := int64(0)

	for i := 0; i < len(tasks); i++ {
		val, ok := mapTask[tasks[i]]
		if ok && dayCount-val < int64(space) {
			dayCount += int64(space) - (dayCount - val)
		}
		dayCount++
		mapTask[tasks[i]] = dayCount
	}
	return dayCount
}
