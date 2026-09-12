**365 Days of LeetCode Challenge — Day 79/365**
**Top K Frequent Words** (Medium)
🔗 https://leetcode.com/problems/top-k-frequent-words/

One sentence separates this from yesterday.

Day 78: return the k most frequent elements, in any order.
Today: return them sorted by frequency, ties broken lexicographically.

That is enough to change which tool is right.

```go
sort.Slice(wordCount, func(i, j int) bool {
	if wordCount[i].count == wordCount[j].count {
		return wordCount[i].word < wordCount[j].word
	}
	return wordCount[i].count > wordCount[j].count
})
```

Read it against the spec: "highest to lowest" is `count >`, "ties lexicographically" is `word <`. The code says what the requirement says.

Yesterday's heap still works - but a heap evicts its ROOT, which must be the WORST entry. For counts, worst is least frequent. For a tie, worst is the word that should LOSE, the lexicographically later one. So the fields point opposite ways:

```go
if a.count == b.count { return a.word > b.word }  // inverted
return a.count < b.count                          // not inverted
```

Correct, and exactly the kind of line that sits wrong in a codebase for months - it only misbehaves on ties.

O(d log d) vs O(d log k), with input capped at 500 words. Unmeasurable. Days 76-78 built the case for the heap; this is the day to notice the technique is not the goal.

Full walkthrough: https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/601_700/top_k_frequent_words/SOLUTION.md
