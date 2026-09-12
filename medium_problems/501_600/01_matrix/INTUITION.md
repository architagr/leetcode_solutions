# 01 Matrix — intuition

## Builds on

- [Day 40: Rotting Oranges](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/901_1000/rotting_oranges/) — the same question, distance from every cell to the nearest source, answered there with a queue and here without one

## The problem in one line

For every cell, how far is the nearest `0`?

## It is yesterday's problem wearing different clothes

Yesterday: how many minutes until this orange rots, given every rotten orange spreads at once. Today: how far is this cell from the nearest zero, given every zero is a source.

Those are the same question. Multi-source BFS solves this one directly, and if you reach for it here you will be right.

This solution does something else, and the reason to read it is that the something else is a genuinely different way of thinking that happens to fit.

## The idea: a shortest path is made of shorter paths

The distance from a cell to the nearest zero is 1 more than the smallest distance among its four neighbours. That is obviously true and, written down naively, obviously circular: every cell depends on its neighbours, and its neighbours depend on it.

The trick is to break the circle by direction.

Any shortest path from a cell to a zero has to make its first step in one of four directions. Split those four into two groups: up-or-left, and down-or-right.

- Sweeping the grid **top-left to bottom-right**, when you reach a cell you have already finalised everything above and to the left of it. So you can correctly compute the best distance for paths that start by going up or left.
- Sweeping **bottom-right to top-left**, you have already finalised everything below and to the right. That covers the other two directions.

Take the minimum of the two passes and every direction has been accounted for. No queue, no visited set.

## Why two passes are enough

Every path from a cell to a zero is a sequence of steps, and in a grid with only four directions the shortest such path never needs to backtrack. The first pass captures all shortest paths whose moves are up and left; the second captures the rest, and can also improve on the first pass's answer because it takes a minimum rather than overwriting.

That is the whole correctness argument, and it is why this is exactly two passes and not three.

## Which one should you write?

BFS, in an interview. It is the obvious framing, it generalises to weighted moves and to more than two dimensions, and it is easier to explain out loud.

The two-pass sweep is worth knowing because it is faster in practice with no queue allocation, and because the pattern (a hard dependency broken by sweeping in a fixed order) shows up constantly once you have seen it.

## Complexity

- **Time: O(m x n).** Two passes over the grid, constant work per cell.
- **Space: O(m x n)** for the output, which the problem asks for anyway. No extra structure.
