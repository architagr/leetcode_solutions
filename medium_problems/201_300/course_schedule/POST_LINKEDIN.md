365 Days of LeetCode Challenge — Day 44/365

207. Course Schedule (Medium)
https://leetcode.com/problems/course-schedule/

You can finish every course unless some group depends on itself: A needs B, B needs C, C needs A. Nobody in that group can ever start. So this is cycle detection in a directed graph, phrased as scheduling.

Kahn's algorithm: count how many prerequisites each course still has, queue everything at zero, and each time you finish a course, decrement everything it unlocks.

The line that separates this from an ordinary BFS is small. In a BFS a node is enqueued the first time it is REACHED. Here, reaching a node only decrements its counter - it is enqueued when that counter hits ZERO. A course with three prerequisites is reached three times and queued once. That if is not an optimisation, it is the scheduling rule.

Then the cycle check, which never looks for a cycle:

return len(result) == numCourses

Courses stuck in a cycle never reach in-degree zero, so they never enter the queue and never appear in the result. The cycle is detected by what fails to show up.

That closes the graph arc on the same move yesterday used - inferring a cycle from a count rather than searching for one.

Full breakdown in today's newsletter article ⬇

#DSA #LeetCode #Golang #Graphs #TopologicalSort #CodingInterview #Algorithms
