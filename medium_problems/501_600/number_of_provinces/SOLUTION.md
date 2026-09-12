# Number of Provinces — solution walkthrough

From `main.go`:

```go
func findCircleNum(isConnected [][]int) int {
	ajList := getAdjeList(isConnected)
	visited := make(map[int]bool, len(isConnected))
	count := 0
	for i := 0; i < len(isConnected); i++ {
		if _, alreadyVisited := visited[i]; !alreadyVisited {
			count++
			findProvince(i, visited, ajList)
		}
	}

	return count
}
```

Same shape as day 37: scan every node, and every time you find one no earlier traversal reached, add one and consume everything it can reach.

## The matrix is not a grid

`isConnected` looks like day 37's input and means something entirely different, which is the one genuine trap in this problem.

In day 37, cell `(i, j)` was a *place*, and the grid's shape was the map. Here, cell `(i, j)` is a *relationship* between city `i` and city `j`. An `n x n` matrix describes `n` nodes, not `n²` of them.

So the outer loop runs over cities, `0` to `n-1`, not over matrix cells.

## Building the adjacency list

```go
func getAdjeList(isConnected [][]int) map[int][]int {
	ajList := make(map[int][]int, len(isConnected))
	for i := 0; i < len(isConnected); i++ {
		for j := i + 1; j < len(isConnected[i]); j++ {
			if isConnected[i][j] == 1 {
				ajList[i] = append(ajList[i], j)
				ajList[j] = append(ajList[j], i)
			}
		}
	}
	return ajList
}
```

![Step 1](images/walkthrough-1.png)

Two details in the loop bounds.

`j := i + 1` reads only above the diagonal. The matrix is symmetric, so every relationship is stored twice, and the upper triangle sees each one exactly once. The body then writes both directions into the list, which is day 36's "record an undirected edge under both endpoints" showing up again.

Starting at `i + 1` also skips the diagonal, where `isConnected[i][i]` is always `1` and only says that a city is connected to itself.

This is the third graph representation in four days: day 36's edge list, day 37's implicit grid, and now a matrix. All three end up as the same adjacency list before anything interesting happens.

## The scan

![Step 2](images/walkthrough-2.png)

```go
if _, alreadyVisited := visited[i]; !alreadyVisited {
	count++
	findProvince(i, visited, ajList)
}
```

![Step 3](images/walkthrough-3.png)

City 0 starts a province, and the BFS marks everything reachable from it. City 1 is now visited, so when the scan reaches it, it is skipped.

![Step 4](images/walkthrough-4.png)

![Step 5](images/walkthrough-5.png)

Two starts, two provinces.

## The BFS

```go
func findProvince(start int, visited map[int]bool, ajList map[int][]int) {
	queue := make([]int, 0, len(ajList))
	push := func(node ...int) { queue = append(queue, node...) }
	pop := func() int {
		x := queue[0]
		queue = queue[1:]
		return x
	}
	push(start)
	for len(queue) > 0 {
		x := pop()
		if _, alreadyVisited := visited[x]; !alreadyVisited {
			push(ajList[x]...)
			visited[x] = true
		}
	}
}
```

The queue is a slice with `q = q[1:]` to pop, which is Go's usual idiom. Day 36 hand-rolled a linked list for the same job; both work, and this one is shorter.

## This version marks on the way out

Day 36 was careful to mark nodes visited as they were *enqueued*. This marks them when they are *dequeued*:

```go
x := pop()
if _, alreadyVisited := visited[x]; !alreadyVisited {
	push(ajList[x]...)
	visited[x] = true
}
```

The result is correct, because the check on dequeue skips any node that was already processed. What it costs is queue size: a city joined to five visited cities gets pushed five times, and four of those entries do nothing but get popped and discarded.

For this problem, where `n` is at most 200, it does not matter. It is worth seeing the two versions side by side, because the difference is invisible in the output and shows up only in how much work the queue does.

## Why not traverse the matrix directly

You can skip `getAdjeList` and ask `isConnected[x][k]` for every `k` during the BFS. That is O(n) per node however few neighbours it has, so the traversal becomes O(n²) even on a sparse graph.

Converting first costs one O(n²) pass and then lets the traversal be O(V + E). On a dense graph that is no saving. On a sparse one it is the difference between checking every city and walking only the joined ones.

Either way the total is O(n²), because the matrix has to be read at least once, and that read is the floor.

## Complexity

- **Time: O(n²).** Dominated by reading the matrix.
- **Space: O(n + E)** for the adjacency list, plus O(n) for the visited map and the queue.

## Test

`main_test.go` covers both worked examples: the matrix where cities 0 and 1 are joined and city 2 is alone, giving 2, and the matrix where nothing is joined, giving 3.
