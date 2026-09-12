**365 Days of LeetCode Challenge — Day 44/365**
**Course Schedule** (Medium)
🔗 https://leetcode.com/problems/course-schedule/

You can finish everything unless some group of courses depends on itself. So this is cycle detection in a directed graph, phrased as scheduling.

Kahn's algorithm: count each course's remaining prerequisites (its in-degree), queue everything at zero, and decrement as courses finish.

```go
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
```

The line that separates this from an ordinary BFS is small. In a BFS a node is enqueued the first time it is REACHED. Here reaching a node only decrements its counter - it is enqueued when the counter hits ZERO. A course with three prerequisites is reached three times and queued once. That `if` is the scheduling rule, not an optimisation.

And the cycle check never looks for a cycle. Courses in a cycle never reach in-degree zero, so they never enter the queue and never appear in `result`. Detected by what fails to show up - the same move as day 43's edge count.

Also: `result` is a valid topological order. This computes it and throws it away. Course Schedule II asks for it.

O(V + E) time and space.

Full walkthrough: https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/201_300/course_schedule/SOLUTION.md
