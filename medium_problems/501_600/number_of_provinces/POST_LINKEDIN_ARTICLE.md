---
meta_title: "Three graph inputs, one adjacency list"
meta_description: "An edge list, a grid and a matrix all describe graphs. Number of Provinces is the third in four days, and the trap is that the matrix is not a grid."
---

![Day 39](HERO.png)

## 365 Days of LeetCode Challenge — Day 39/365

**[547. Number of Provinces](https://leetcode.com/problems/number-of-provinces/)** (Medium)

`isConnected[i][j] == 1` means cities `i` and `j` are directly joined. Count the connected groups.

## The counting is day 37's

```
for every node:
    if not visited:
        count++
        traverse everything reachable
```

Identical to Number of Islands. What changes is only how you find a node's neighbours, and that is the interesting part, because this is the **third graph representation in four days**.

Day 36 gave an edge list. Day 37 gave a grid where edges were implied by position. Today gives a matrix. All three become the same adjacency list before anything interesting happens.

## The trap: the matrix is not a grid

`isConnected` looks exactly like day 37's input and means something completely different.

In day 37, cell `(i, j)` was a **place**. The grid's shape was the map.

Here, cell `(i, j)` is a **relationship** between city `i` and city `j`. An `n x n` matrix describes `n` nodes, not `n²`.

So you do not traverse the matrix. You traverse the cities and consult the matrix to find out who is adjacent to whom. Writing a day-37-style nested scan over the cells here produces something that type-checks, runs, and answers a question nobody asked.

## Reading only above the diagonal

```go
for j := i + 1; j < len(isConnected[i]); j++ {
	if isConnected[i][j] == 1 {
		ajList[i] = append(ajList[i], j)
		ajList[j] = append(ajList[j], i)
	}
}
```

The matrix is symmetric, so every relationship is stored twice. The upper triangle sees each one once. Starting at `i + 1` also skips the diagonal, where `isConnected[i][i]` is always 1 and says only that a city is connected to itself.

The body then writes **both** directions into the adjacency list, which is day 36's undirected-edge rule turning up again.

![Step 1](images/walkthrough-1.png)

![Step 2](images/walkthrough-2.png)

![Step 3](images/walkthrough-3.png)

![Step 4](images/walkthrough-4.png)

![Step 5](images/walkthrough-5.png)

## This one marks on the way out

Day 36 made a point of marking nodes visited as they were **enqueued**. This marks them as they are **dequeued**:

```go
x := pop()
if _, alreadyVisited := visited[x]; !alreadyVisited {
	push(ajList[x]...)
	visited[x] = true
}
```

Both are correct. The dequeue version costs queue size: a city joined to five already-visited cities gets pushed five times, and four of those pops do nothing.

With `n` capped at 200 it is irrelevant here. It is worth seeing the two next to each other, because the difference never shows up in the output, only in how much work the queue does.

## Why convert at all

You could BFS straight off the matrix, asking `isConnected[x][k]` for every `k`. That is O(n) per node no matter how few neighbours it has, so the traversal is O(n²) even on a sparse graph.

Converting costs one O(n²) pass and then lets the traversal be O(V + E).

The total is O(n²) either way, because the matrix has to be read at least once and that read is the floor. What changes is whether the traversal itself scales with the graph or with the matrix.

## Complexity

- **Time: O(n²)**, dominated by reading the matrix.
- **Space: O(n + E)** for the adjacency list, plus O(n) for visited and the queue.

## Builds on

- [Day 36: Find if Path Exists in Graph](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/1901_2000/find_if_path_exists_in_graph/) — building an adjacency list, then BFS from a start node
- [Day 37: Number of Islands](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/101_200/number_of_islands/) — counting components by counting how many times the scan had to start

Full code and the step-by-step walkthrough:
[number_of_provinces](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/501_600/number_of_provinces/SOLUTION.md)

#DSA #LeetCode #Golang #Graphs #BFS #CodingInterview #Algorithms

---

*Solution and code by Archit Agarwal. Write-up drafted with AI assistance from the code and problem statement.*
