---
meta_title: "The day to not reach for the heap"
meta_description: "One added sentence about ordering makes sorting the better tool. The heap still works, and its comparator has an inversion you have to reason about every time."
tags: [golang, sorting, heap, dsa]
---

![Day 79](HERO.png)

*365 Days of LeetCode Challenge — Day 79/365*

**[692. Top K Frequent Words](https://leetcode.com/problems/top-k-frequent-words/)** (Medium)

Given an array of strings and an integer `k`, return the `k` most frequent words, sorted by frequency from highest to lowest, with equally frequent words sorted lexicographically.

## One sentence separates this from yesterday

Put the two statements side by side.

Day 78: *return the k most frequent elements. You may return the answer in any order.*

Today: *return the k most frequent strings, sorted by the frequency from highest to lowest. Sort the words with the same frequency by their lexicographical order.*

The counting is identical. The selection is identical. One clause about ordering has been added, and that clause is enough to change which tool is the right one.

I have spent three days building the case for the heap. Today is the day to notice that the technique is not the goal, and that adding a requirement can make a "worse" algorithm the better choice.

## Counting, then flattening

```go
m := make(map[string]int)
for i := 0; i < len(words); i++ {
	m[words[i]]++
}
```

![Step 1](images/walkthrough-1.png)

Go deliberately randomises map iteration order, so nothing useful can be read off the map directly. The counts are flattened into a slice before anything can be arranged.

## The comparator is the solution

```go
sort.Slice(wordCount, func(i, j int) bool {
	if wordCount[i].count == wordCount[j].count {
		return wordCount[i].word < wordCount[j].word
	}
	return wordCount[i].count > wordCount[j].count
})
```

Read that against the problem statement, clause by clause.

"Sorted by the frequency from highest to lowest" — `count > count`. Descending.

"Sort the words with the same frequency by their lexicographical order" — `word < word`. Ascending.

The code says what the requirement says, in the same order, with the same directions. There is nothing to translate.

![Step 2](images/walkthrough-2.png)

`i` and `love` both appear twice, so the tie-break decides between them, and `"i" < "love"`.

![Step 3](images/walkthrough-3.png)

![Step 4](images/walkthrough-4.png)

## Now the heap version, honestly

It works. I want to walk through it properly rather than dismissing it, because the reason to prefer sorting only becomes clear once you have written the alternative.

You would keep a min-heap capped at k, as on day 78. That gives the correct *set* of k words, in ascending frequency, so you reverse at the end to get the required order.

The interesting part is the comparator. A heap evicts its **root**, so the root has to be the entry you would most like to lose — the *worst* one.

For frequency, worst means least frequent. Straightforward.

For a tie, worst means the word that should **lose the tie**. The problem says ties are broken by lexicographic order, so the winner is the earlier word — which makes the loser the **later** one. So on ties the comparison has to point the other way:

```go
if a.count == b.count {
	return a.word > b.word   // inverted
}
return a.count < b.count     // not inverted
```

One comparator, two fields, opposite directions.

That is correct. It is also the kind of line that sits wrong in a codebase for months without anyone noticing, because nothing about it looks unusual, and it only produces wrong output on ties — which are exactly the inputs that casual testing skips.

## The honest comparison

| | sort | size-k heap |
|---|---|---|
| Time | O(d log d) | O(d log k) |
| Output order | already correct | reverse at the end |
| Comparator | reads like the spec | inverts on one field |

On paper the heap wins the complexity column when k is small and d is large.

Then look at the constraints: `words.length` is at most 500. So `d` is at most 500, `log d` is under 9, and the difference between the two is not merely small — it is not measurable.

What *is* measurable is the cost of the comparator. One of them can be checked against the problem statement by reading it. The other requires you to reconstruct an argument about which entry a heap evicts before you can tell whether it is right.

At 500 elements, I take the readable one every time. At 5 million with k of 10, the calculation changes and the heap is worth the comparator. **The point is that it is a calculation, not a reflex.**

## Two things that could be tightened

```go
result := make([]string, 0, k)
for i := 0; i < k; i++ {
	result = append(result, wordCount[i].word)
}
```

If `k` were larger than the number of distinct words, this would index past the end of the slice and panic. The constraints guarantee it cannot happen — `k` is in the range of unique words — so the code is correct as it stands.

Adding `&& i < len(wordCount)` costs nothing and removes the dependence on a constraint that lives in a different document from the code. I would add it.

The other: `sort.Slice` is **not** a stable sort. That is fine here, because the comparator defines a total order — every pair of distinct words is decided either by count or by spelling, and no pair is left to arbitrary choice.

It is worth knowing the shape of the bug it prevents, though. An unstable sort with an *incomplete* comparator produces output that is correct-but-different between runs, and that is a genuinely miserable thing to debug.

## Complexity

- **Time: O(n + d log d).** O(n) to count every word, then sorting the d distinct ones.
- **Space: O(d)** for the map and the flattened slice.

## Builds on

- [Day 78: Top K Frequent Elements](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/301_400/top_k_frequent_elements/) — the same counting problem with an ordering requirement added, which is enough to change the tool

Full code and the step-by-step walkthrough:
[top_k_frequent_words](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/601_700/top_k_frequent_words/SOLUTION.md)

---

*Solution and code by Archit Agarwal. Write-up drafted with AI assistance from the code and problem statement.*
