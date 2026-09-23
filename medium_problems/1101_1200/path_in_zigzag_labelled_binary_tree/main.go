package pathinzigzaglabelledbinarytree

import "math"

func pathInZigZagTree(label int) []int {
	// step 1: find the level this label will fit at
	level := findLevel(label)

	// step 2: find the index at which this label will be in the level
	levelStartVal := levelStart(level)
	// step 2.1: label - (2^ (level-1) ) if level is even then index is (2^ (level-1) - 1) - index
	index := label - levelStartVal
	if !isOdd(level) {
		index = levelStartVal - 1 - index
	}
	res := make([]int, level)
	res[level-1] = label
	level--
	// step 3: find the remaining path
	// step 3.1: index/=2
	// step 3.2: if level is even than temp index (2^ (level-1)-1)- index
	// step 3.2: path will be (2^ (level-1)) + temp index

	for level > 0 {
		index /= 2
		tempIndex := index
		if !isOdd(level) {
			tempIndex = pow2(level-1) - 1 - tempIndex
		}
		res[level-1] = pow2(level-1) + tempIndex
		level--
	}
	// repeate step 3 for all level in reverse order
	return res
}
func levelStart(level int) int {
	return pow2(level - 1)
}
func isOdd(x int) bool {
	return x%2 == 1
}
func findLevel(label int) int {
	level := 0

	for x := 1; x <= label; {
		level++
		x = pow2(level)
	}
	return level
}

func pow2(x int) int {
	return int(math.Pow(float64(2), float64(x)))
}
