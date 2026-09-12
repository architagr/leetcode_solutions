# Top K Frequent Words — intuition

## Builds on

- [Day 78: Top K Frequent Elements](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/301_400/top_k_frequent_elements/) — the same counting problem with an ordering requirement added, which is enough to change the tool

## The problem in one line

Return the `k` most frequent words, ordered by frequency descending, ties broken by spelling.

## One sentence separates this from yesterday

Yesterday: *return the k most frequent elements, in any order*.

Today: *return them sorted by frequency, and sort words with the same frequency lexicographically*.

That is the whole difference, and it is enough to make sorting the better tool.

## Why the ordering requirement matters

Yesterday's size-k min-heap produces the right *set* in ascending frequency. To satisfy today's spec you would still have to reverse it, and handle the tie-break inside the comparator.

And the comparator is where it gets unpleasant. In a min-heap that evicts the least frequent, ties have to break the *other* way — the word that should win must be treated as larger so it is not the one thrown out. So the comparison inverts on one field and not the other:

```go
if a.count == b.count {
    return a.word > b.word   // inverted
}
return a.count < b.count     // not inverted
```

That is correct and it is the kind of line that is wrong in production for months.

## What sorting gives instead

The spec is a total order over the words, stated plainly:

```go
if wordCount[i].count == wordCount[j].count {
	return wordCount[i].word < wordCount[j].word
}
return wordCount[i].count > wordCount[j].count
```

The comparator reads like the problem statement, both fields point the same way as the requirement, and the result is already in output order with no reversal.

## The honest trade

Sorting is O(d log d) against the heap's O(d log k). When k is tiny and d is huge, the heap wins on paper.

With `words.length` capped at 500, d is at most 500 and the difference is not measurable. The comparator being obviously correct is worth more than an asymptotic improvement that cannot be observed.

**The technique is not the goal.** Days 76 to 78 built up to the heap; this is the day to notice when not to reach for it.

## Complexity

- **Time: O(n + d log d)** — counting, then sorting the distinct words.
- **Space: O(d)** for the counts and the slice.
