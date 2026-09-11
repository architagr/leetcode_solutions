package rabbitsinforest

func numRabbits(answers []int) int {
	rabbits := make(map[int]int)
	res := 0
	for _, ans := range answers {
		rabbits[ans] += 1
		//for example, if 3 rabbits telling us that there's 1 more same rabbit
		if rabbits[ans] == ans+1 {
			res += ans + 1
			delete(rabbits, ans)
		}
	}
	for ans, _ := range rabbits {
		res += ans + 1
	}
	return res
}
