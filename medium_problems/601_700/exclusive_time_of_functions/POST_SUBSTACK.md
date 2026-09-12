---
meta_title: "Two interval rules, twenty lines apart"
meta_description: "Exclusive Time replays a call stack from its log. One interval needs a +1 and the other does not, and the line that resumes the parent is what makes it exclusive."
tags: [golang, stack, dsa, interview]
---

![Day 64](HERO.png)

*365 Days of LeetCode Challenge — Day 64/365*

**[636. Exclusive Time of Functions](https://leetcode.com/problems/exclusive-time-of-functions/)** (Medium)

A single-threaded program logs when each function starts and ends. Compute, for each function, how long it spent running *itself* — excluding time spent inside functions it called.

This closes the stack arc, and it is the problem in the set where the stack is not a modelling decision at all.

## The input is stack activity

Every other problem this week chose a stack because the problem had a last-in-first-out shape. This one does not choose anything.

Function calls nest. A function that starts while another is running must finish before that one can resume — that is not an algorithmic insight, it is what a call stack is, and these logs are a recording of one.

So the algorithm is not "use a stack to solve this". It is "reconstruct the stack the log describes, and charge time to whoever was on top".

## The word doing the work is "exclusive"

If function `0` runs from time 0 to time 6, and function `1` runs from 2 to 5 inside it, then function 0's exclusive time is **2**, not 6.

The three units it spent waiting on function 1 belong to function 1. Function 0 was on the stack for all six, and running for two of them.

That distinction is the entire problem, and it forces the design. A function must be **paused** when another starts on top of it, and **resumed** when that one ends. Each stack entry therefore carries the timestamp it most recently resumed at.

This is day 60's idea in a new setting. There, each Min Stack entry carried the minimum at its own height, so popping restored the right answer with no work. Here each entry carries its own resume time, so a parent knows where to pick up. Both are "store something in the entry that makes the pop cheap".

## Walking the events

```go
func exclusiveTime(n int, logs []string) []int {
	st := [][]int{} // [id,startTime]
	out := make([]int, n)
	for i := 0; i < len(logs); i++ {
		id, logType, time := getData(logs[i])
		if logType == "start" {
			if len(st) > 0 {
				idx, sTime := st[len(st)-1][0], st[len(st)-1][1]
				out[idx] += time - sTime
			}
			st = append(st, []int{id, time})
		} else {
			idx, sTime := st[len(st)-1][0], st[len(st)-1][1]
			st = st[:len(st)-1]
			out[idx] += time - sTime + 1
			if len(st) > 0 {
				st[len(st)-1][1] = time + 1
			}
		}
	}
	return out
}
```

![Step 1](images/walkthrough-1.png)

The first start has nothing beneath it, so nothing is charged. Function 0 is pushed with its start time.

![Step 2](images/walkthrough-2.png)

Function 1 starts at 2. Before pushing it, function 0 is charged for the stretch it just completed: `2 - 0 = 2`.

Function 0 is not finished. It is paused, and it stays on the stack underneath.

## Two interval rules, and they are not the same

Here is the detail that makes this a Medium rather than an Easy, and it is arithmetic rather than algorithm.

```go
out[idx] += time - sTime        // a pause
out[idx] += time - sTime + 1    // an end
```

**Pausing has no `+1`.** Function 0 ran at instants 0 and 1. At instant 2, function 1 took over. So it ran for two units, and `2 - 0` is exactly right. The interval is half-open: it includes the start and excludes the moment the child began.

**Ending has a `+1`.** An end timestamp is inclusive — the function was still running *during* that unit. A function that starts at 3 and ends at 3 ran for one unit, not zero, and `3 - 3 + 1` gives that.

![Step 3](images/walkthrough-3.png)

Function 1 ends at 5, so it is charged `5 - 2 + 1 = 4`: instants 2, 3, 4 and 5.

Two different rules, in one twenty-line function, twenty lines apart. Mixing them up is the most common way to get this wrong, and neither mistake produces an obviously absurd answer — you get a number that is off by one per call, which looks plausible until you check it against a hand-worked example.

## The line that makes the answer exclusive

```go
if len(st) > 0 {
	st[len(st)-1][1] = time + 1
}
```

![Step 4](images/walkthrough-4.png)

When function 1 finishes at time 5, function 0 becomes the running function again — but not at time 5. Instant 5 has already been charged to function 1. Function 0 resumes at **6**.

It is worth being concrete about what happens without this line, because the failure is instructive.

The entry underneath still holds its original start time of 0. When function 0 finally ends at 6, it is charged `6 - 0 + 1 = 7` — its entire span, including every instant its child was running.

That is not a crash or a nonsense value. It is the **inclusive** time: a perfectly meaningful quantity, and the answer to a different, easier question. One missing line silently changes which problem you solved.

![Step 5](images/walkthrough-5.png)

Function 0 ends at 6 and is charged `6 - 6 + 1 = 1` more unit, for a total of 3.

## The entry is mutated, and that is the contrast with day 60

```go
st[len(st)-1][1] = time + 1
```

`st` is a `[][]int`, so each entry is a slice and this writes through to it. The same entry can be updated many times over its life — once for every child that starts and ends inside it.

Day 60's entries were the opposite. A Min Stack entry's `min` is computed once at push and never touched again, and that immutability is exactly why popping needs no work: the entry underneath was already correct and nothing could have invalidated it.

Here the entry is a running tally that keeps changing while it sits on the stack. Both designs put information in the entry to make the pop cheap; only one of them can get away with never revisiting it.

## What this code assumes

Two things, both fine for the problem and both worth naming.

The `end` branch reads the top of the stack without checking that it is non-empty, and `getData` discards the errors from `strconv.Atoi`.

The problem guarantees well-formed, balanced logs, so neither can fail. Code parsing a real trace would need to handle an `end` with nothing running, a malformed line, and timestamps that go backwards — none of it hard, but none of it here.

## Complexity

- **Time: O(n)** in the number of log entries. Each is parsed once and causes exactly one push or one pop.
- **Space: O(d)** for the stack, where `d` is the maximum call depth, plus O(n) for the output array.

Note that recursion works without any special handling: the same function id can appear at several stack levels at once, because each entry carries its own resume time rather than the id indexing into shared state.

## Builds on

- [Day 60: Min Stack](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/101_200/min_stack/) — stack entries carrying a second field, so that popping restores the right state without recomputation

Full code and the step-by-step walkthrough:
[exclusive_time_of_functions](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/601_700/exclusive_time_of_functions/SOLUTION.md)

---

*Solution and code by Archit Agarwal. Write-up drafted with AI assistance from the code and problem statement.*
