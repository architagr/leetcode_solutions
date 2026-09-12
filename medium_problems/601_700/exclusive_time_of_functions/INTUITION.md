# Exclusive Time of Functions — intuition

## Builds on

- [Day 60: Min Stack](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/101_200/min_stack/) — stack entries carrying a second field, so that popping restores the right state without recomputation

## The problem in one line

Given function start and end logs from a single thread, how long did each function spend running *itself*, not counting the functions it called?

## Why a stack

Function calls nest. A function that starts while another is running must finish before that one can resume — that is what a call stack is, and the log is a recording of one.

So the structure is not a design choice here. The input *is* stack activity, and the algorithm is replaying it.

## Exclusive, not total

The word doing the work is "exclusive". If `0` runs from time 0 to 6 but `1` runs from 2 to 5 inside it, function 0's exclusive time is 2 units, not 6. The time it spent waiting on `1` belongs to `1`.

So the running function has to be paused whenever another starts, and resumed when that one ends. Every stack entry carries the timestamp it most recently resumed at, and that field is updated — not just written once.

That is the day 60 idea again: an entry holds a value that makes popping cheap. There it was a minimum; here it is a resume time.

## The three events

**A start, with something already running.** The running function is charged for the interval it just completed, then the new one is pushed.

**An end.** The top is charged and popped. The end timestamp is inclusive — a function that starts and ends at the same instant ran for one unit — so the charge is `time - start + 1`.

**An end, with something underneath.** The newly exposed function resumes at `time + 1`, because the unit at `time` was already charged to the function that just finished.

That last line is the one that makes the time exclusive. Without it, the parent is charged again for the unit its child already claimed.

## The two different arithmetic rules

A paused interval is `time - start`, with no `+1`: the function ran up to but not including the instant its child started.

A completed interval is `time - start + 1`: it ran through the end timestamp inclusive.

Mixing them up is the main way to get this wrong, and both appear in the same short function.

## Complexity

- **Time: O(n)** in the number of logs; each is parsed once and causes one push or pop.
- **Space: O(d)** for the stack, the maximum call depth.
