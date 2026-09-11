package combimationsum

import (
	"fmt"
	"sort"
	"strings"
)

func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0, 150)
	current := make([]int, 0, len(candidates))
	m := make(map[string]bool, 150)
	recursiveSum(candidates, current, target, 0, &result, m)
	return result
}

func arrayToStr(arr []int) string {
	sb := strings.Builder{}
	sb.Grow(3 * len(arr))
	for _, v := range arr {
		sb.WriteString(fmt.Sprintf("%d,", v))
	}
	return strings.TrimRight(sb.String(), ",")
}

func recursiveSum(candidates, current []int, target, currentSum int, output *[][]int, m map[string]bool) {
	if currentSum > target {
		return
	}
	if currentSum == target {
		// add array to output
		x := append([]int{}, current...)
		sort.Ints(x)
		str := arrayToStr(x)
		if _, ok := m[str]; !ok {
			*output = append(*output, x)
			m[str] = true
		}
		return
	}

	for _, score := range candidates {
		currentSum += score
		current = append(current, score)
		recursiveSum(candidates, current, target, currentSum, output, m)
		currentSum -= score
		current = current[:len(current)-1]
	}
}
