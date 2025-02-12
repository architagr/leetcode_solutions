package exclusivetimeoffunctions

import (
	"strconv"
	"strings"
)

func exclusiveTime(n int, logs []string) []int {
	st := [][]int{} // [id,startTime]
	out := make([]int, n)
	for i := 0; i < len(logs); i++ {
		id, logType, time := getData(logs[i])
		if logType == "start" {
			if len(st) > 0 {
				idx, sTime := st[len(st)-1][0], st[len(st)-1][1]
				out[idx] += time - sTime
			}
			st = append(st, []int{id, time})
		} else {
			idx, sTime := st[len(st)-1][0], st[len(st)-1][1]
			st = st[:len(st)-1]
			out[idx] += time - sTime + 1
			if len(st) > 0 {
				st[len(st)-1][1] = time + 1
			}
		}
	}
	return out
}

func getData(log string) (id int, operation string, time int) {
	data := strings.Split(log, ":")
	id, _ = strconv.Atoi(data[0])
	time, _ = strconv.Atoi(data[2])
	operation = data[1]
	return
}
