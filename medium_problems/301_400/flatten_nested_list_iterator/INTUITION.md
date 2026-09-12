# Flatten Nested List Iterator — intuition

## Builds on

- [Day 55: Binary Search Tree Iterator](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/101_200/binary_search_tree_iterator/) — the same design: a stack holding the path to the cursor, advanced only when the caller asks

## The problem in one line

Given a nested list of integers, iterate the integers in order.

## The solution that is not an iterator

Walk the whole structure in the constructor, copy every integer into a slice, and have `Next` return the next element.

It produces correct values. It is not an iterator, and the distinction is the point of the problem.

A flatten-then-index does all the work up front, whether or not the caller consumes any of it, and it holds a copy of every integer for as long as the object lives. That is fine on a small input and wrong in the situations iterators exist for: a structure too large to copy, or a caller that stops after three elements.

## What an iterator has to do instead

Do nothing until asked. Keep only enough state to resume.

For a nested list, "enough state to resume" is the path from the outer list down to wherever the cursor currently is: which list you are in, and how far into it, for each level of nesting.

That is a stack of cursors, one frame per level. And it is exactly what a BST iterator holds — the path from the root to the current node — which is why day 55 and today are the same design wearing different data.

## Where the work goes

`HasNext` does the advancing, which is unusual and worth stating.

Most of the time `HasNext` sounds like a passive question. Here it is the method that moves the stack: it pops finished lists, descends into nested ones, and stops as soon as the cursor is sitting on an integer. Once it has stopped, `Next` only has to read and step past.

That split is why `HasNext` can be called repeatedly without consuming anything. If the cursor is already on an integer, it does nothing and returns true.

## The line that keeps it terminating

When `HasNext` meets a nested list, it steps the parent's cursor past it **before** descending.

Descend first and the parent's cursor still points at the list you just entered, so when that list finishes and is popped, the parent walks straight back into it. Forever.

This is day 37's "sink the cell before recursing" and day 42's "record the copy before cloning its neighbours", in a third costume. The pattern is the same: mark the thing as handled before you go into it.

## Complexity

- **Time: O(1)** construction. Across a full iteration, every element is visited once, so `Next` is amortised O(1).
- **Space: O(d)**, the nesting depth — not the element count.
