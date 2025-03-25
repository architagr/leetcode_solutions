package courseschedule

func canFinish(numCourses int, prerequisites [][]int) bool {
	result := make([]int, 0, numCourses)

	aj := make(map[int][]int)
	inDegree := make([]int, numCourses)
	for _, pre := range prerequisites {
		inDegree[pre[0]]++
		aj[pre[1]] = append(aj[pre[1]], pre[0])
	}
	q := make([]int, 0, numCourses)
	push := func(c int) {
		q = append(q, c)
	}
	pop := func() int {
		x := q[0]
		q = q[1:]
		return x
	}
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			push(i)
		}
	}
	for len(q) > 0 {
		c := pop()
		result = append(result, c)
		for _, de := range aj[c] {
			inDegree[de]--
			if inDegree[de] == 0 {
				push(de)
			}
		}
	}
	return len(result) == numCourses
}
