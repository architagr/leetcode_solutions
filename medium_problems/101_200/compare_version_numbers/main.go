package compareversionnumbers

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	v1, v2 := parseVersion(version1), parseVersion(version2)

	if len(v1) < len(v2) {
		v1 = append(v1, make([]int, len(v2)-len(v1))...)
	} else if len(v1) > len(v2) {
		v2 = append(v2, make([]int, len(v1)-len(v2))...)
	}

	for i := 0; i < len(v1); i++ {
		if v1[i] > v2[i] {
			return 1
		} else if v1[i] < v2[i] {
			return -1
		}
	}
	return 0
}

func parseVersion(version string) []int {
	strArr := strings.Split(version, ".")
	result := make([]int, len(strArr))
	for i, str := range strArr {
		num, _ := strconv.Atoi(str)
		result[i] = num
	}
	return result
}
