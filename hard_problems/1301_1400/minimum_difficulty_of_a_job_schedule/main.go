package minimumdifficultyofajobschedule

import "math"

func minDifficulty(jobDifficulty []int, d int) int {
	n := len(jobDifficulty)
	if n < d {
		return -1
	}
	dpTable := make([][]int, n)
	for i := 0; i < n; i++ {
		dpTable[i] = make([]int, d+1)
		for j := 0; j <= d; j++ {
			dpTable[i][j] = -1
		}
	}

	return dp(0, d, jobDifficulty, &dpTable)
}

func dp(i, daysRemainining int, jobDifficulty []int, dpTable *[][]int) int {
	if (*dpTable)[i][daysRemainining] != -1 {
		return (*dpTable)[i][daysRemainining]
	}
	if daysRemainining == 1 {
		x := 0
		for j := i; j < len(jobDifficulty); j++ {
			x = maxValue(x, jobDifficulty[j])
		}
		return x
	}
	res := math.MaxInt
	dailyMaxJobDiff := 0
	for j := i; j < len(jobDifficulty)-daysRemainining+1; j++ {
		dailyMaxJobDiff = maxValue(dailyMaxJobDiff, jobDifficulty[j])
		res = minValue(res, dailyMaxJobDiff+dp(j+1, daysRemainining-1, jobDifficulty, dpTable))
	}
	(*dpTable)[i][daysRemainining] = res
	return (*dpTable)[i][daysRemainining]
}

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minValue(a, b int) int {
	if a < b {
		return a
	}
	return b
}
