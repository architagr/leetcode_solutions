# Course Schedule — intuition

## Builds on

- [Day 36: Find if Path Exists in Graph](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/1901_2000/find_if_path_exists_in_graph/) — building an adjacency list from a list of pairs, with the difference that these pairs are directed

## The problem in one line

Given courses and their prerequisites, can you finish all of them?

## The question is about cycles

You can finish everything unless some group of courses depends on itself: A needs B, B needs C, C needs A. Nobody in that group can ever start.

So "can I finish all the courses" is "does this directed graph have a cycle", asked in different words.

## Directed edges, and getting the direction right

Every graph so far in this arc has been undirected, and every adjacency list recorded each edge twice. This one is directed, and the direction carries meaning.

`prerequisites[i] = [a, b]` means you must take `b` before `a`. The useful edge points from the prerequisite to the thing it unlocks: `b -> a`. When `b` is finished, `a` becomes more available.

Build the list the other way and the algorithm runs happily and answers a different question.

## Kahn's algorithm: finish what has nothing left to wait for

Count, for each course, how many prerequisites it still has. That is its in-degree.

- Any course with in-degree 0 can be taken right now. Queue all of them.
- Take one, record it as finished, and decrement the in-degree of everything it unlocks.
- Anything that drops to 0 has just become available. Queue it.
- Repeat.

This is a traversal in the same family as the BFS from day 36, with one change: a node is not enqueued when it is *reached*, it is enqueued when it is *ready*. A course with three prerequisites is reached three times and only enters the queue on the third.

## Cycle detection by absence

Here is the part worth remembering.

If every course comes out of the queue, there was no cycle. If some never do, those courses are exactly the ones in a cycle or downstream of one, because their in-degree could only be reduced by a course that is itself waiting.

So the test is a count:

```go
return len(result) == numCourses
```

Nothing looks for a cycle. The cycle is detected by what fails to appear, which is the same move as day 43 — and both times it comes out shorter than searching.

## Complexity

- **Time: O(V + E).** One pass over the prerequisites to build the graph, then each course dequeued once and each edge relaxed once.
- **Space: O(V + E)** for the adjacency list and the in-degree array.
