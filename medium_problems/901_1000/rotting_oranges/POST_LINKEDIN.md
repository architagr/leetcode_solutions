365 Days of LeetCode Challenge — Day 40/365

994. Rotting Oranges (Medium)
https://leetcode.com/problems/rotting-oranges/

Four days ago I said BFS and DFS answer reachability equally well, and that the distinction starts mattering when a problem asks about distance. This is that problem.

BFS reaches every cell by a shortest path. DFS reaches it by whatever route it wandered down first. "When does this orange rot?" is a shortest-path question - it rots when the nearest rotten one gets to it - so DFS here gives wrong answers, not slow ones.

The trick is multi-source BFS: every already-rotten orange goes into the queue before the walk starts. A BFS seeded with many sources expands as one combined wavefront, and each cell ends up with its distance to the nearest source. Nothing in the loop needs to know there was more than one start.

One more thing that cannot be skipped: the final scan. BFS only visits what it can reach, so an orange sealed behind empty cells never rots and the queue drains without noticing. The -1 case has to be looked for afterwards.

Full breakdown in today's newsletter article ⬇

#DSA #LeetCode #Golang #Graphs #BFS #CodingInterview #Algorithms
