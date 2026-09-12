---
meta_title: "Detecting a cycle by what never shows up"
meta_description: "Course Schedule never searches for a cycle. It counts how many courses came out, and the ones stuck in a cycle are the ones that never do."
tags: [golang, graphs, topological-sort, dsa]
---

![Day 44](HERO.png)

*365 Days of LeetCode Challenge — Day 44/365*

**[207. Course Schedule](https://leetcode.com/problems/course-schedule/)** (Medium)

There are `numCourses` courses. Some have prerequisites. Can you finish all of them?

This is the last day of the graph arc, and it ends on the same idea yesterday ended on: a cycle established without ever searching for one.

## Reading the question properly

You can finish every course unless some group of them depends on itself.

A needs B. B needs C. C needs A. Nobody in that group can ever be first, so nobody in it can ever be taken, and it does not matter how the rest of the schedule looks.

So this problem is cycle detection in a directed graph, wearing a scheduling costume. Nothing else about it matters.

## The first directed graph in the arc

Every graph so far has been undirected, and every adjacency list built in this arc wrote each edge under **both** of its endpoints. Day 36 made a point of it, day 39 repeated it.

This one is directed, and the direction carries the meaning:

```go
inDegree[pre[0]]++
aj[pre[1]] = append(aj[pre[1]], pre[0])
```

`prerequisites[i] = [a, b]` means `b` must be taken before `a`.

There are two edges you could record, and only one of them is useful. Pointing from `a` to `b` would say "a depends on b", which is true and leads nowhere: to make progress you would have to repeatedly ask each course whether all its dependencies happen to be done.

Pointing from `b` to `a` says "finishing b brings a closer to available", and that is a direction you can *push* along. So `aj` is keyed by `pre[1]`, the prerequisite, and holds `pre[0]`, the course it unlocks.

Alongside it, `inDegree[pre[0]]` counts how many prerequisites the course `a` is still waiting on.

Get the direction backwards and the program runs, terminates, and answers a question nobody asked. There is no crash to tell you.

![Step 1](images/walkthrough-1.png)

## Seeding with whatever is already possible

```go
for i := 0; i < numCourses; i++ {
	if inDegree[i] == 0 {
		push(i)
	}
}
```

A course with no prerequisites can be taken immediately. All of them go in before the walk begins, which is day 40's multi-source seeding for the second time.

There is a case worth noticing here: if this loop pushes nothing, then every course has at least one prerequisite, which is only possible if there is a cycle. The main loop never runs, `result` stays empty, and the final comparison returns false. The degenerate case handles itself.

## The line that makes this not a BFS

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
```

![Step 2](images/walkthrough-2.png)

![Step 3](images/walkthrough-3.png)

Put this next to day 36's BFS and the loops look nearly identical. The difference is three lines long and it is the entire algorithm.

In day 36, a node was enqueued the first time it was **reached**. Reaching it was sufficient, because the question was reachability and any arrival answered it.

Here, reaching a node is not enough. Reaching it only decrements its counter. It is enqueued when that counter hits **zero**.

A course with three prerequisites gets reached three times, once as each prerequisite finishes, and it enters the queue only on the third. Until then it has been touched but is not ready.

That `if` is not a guard against duplicate work. It is the scheduling rule, expressed as arithmetic.

![Step 4](images/walkthrough-4.png)

Course 3 is waiting on both 1 and 2. Its in-degree reaches zero only after both have come out of the queue.

## The cycle check that is not a cycle check

```go
return len(result) == numCourses
```

![Step 5](images/walkthrough-5.png)

Read the function again and look for the cycle detection. There is none. No visited set, no recursion stack, no back-edge test, no colouring.

Here is why the count is sufficient.

If every course came out of the queue, then every course reached in-degree zero, which means every course had all its prerequisites satisfied at some point. That is a complete schedule, and a complete schedule cannot contain a cycle.

If some courses never came out, consider one of them. Its in-degree never reached zero, so at least one of its prerequisites never finished. That prerequisite also never came out, for the same reason. Follow that chain backwards through a finite set of courses and it must eventually repeat, and a repeat is a cycle.

So the courses missing from `result` are exactly the courses in a cycle, or downstream of one. The cycle is detected by what fails to appear.

Yesterday inferred a cycle from an edge count. Today infers it from a completion count. In both cases the inference is shorter than the search would have been, and in both cases it comes from knowing a property of the structure rather than writing a cleverer traversal.

If I had one thing to take out of this whole arc, it would be that. Several of these problems have a straightforward search-based solution that works, and a shorter one that follows from a fact. Knowing the fact is the difference.

## Something this function computes and discards

The function uses `result` only for its length. An `int` counter would do.

But look at what is actually in that slice: the courses, in an order where every course appears after all of its prerequisites. That is a topological ordering, and it is a genuinely useful object. It is a valid schedule.

Course Schedule II is the same problem asking for the ordering instead of the yes/no, and this solution already builds it and throws it away. Changing the return type is most of the work.

## Complexity

- **Time: O(V + E).** One pass over the prerequisites to build the graph and the in-degree counts, then each course dequeued exactly once and each edge relaxed exactly once.
- **Space: O(V + E)** for the adjacency list and in-degree array, plus O(V) for the queue and the result.

## Builds on

- [Day 36: Find if Path Exists in Graph](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/1901_2000/find_if_path_exists_in_graph/) — building an adjacency list from a list of pairs, with the difference that these pairs are directed

Full code and the step-by-step walkthrough:
[course_schedule](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/201_300/course_schedule/SOLUTION.md)

---

*Solution and code by Archit Agarwal. Write-up drafted with AI assistance from the code and problem statement.*
