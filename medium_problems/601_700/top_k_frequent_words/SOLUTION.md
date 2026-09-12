# Top K Frequent Words — solution walkthrough

From `main.go`:

```go
func topKFrequent(words []string, k int) []string {
	m := make(map[string]int)
	for i := 0; i < len(words); i++ {
		m[words[i]]++
	}

	type WordCount struct {
		word  string
		count int
	}
	wordCount := make([]WordCount, 0)
	for key, val := range m {
		wordCount = append(wordCount, WordCount{word: key, count: val})
	}

	sort.Slice(wordCount, func(i, j int) bool {
		if wordCount[i].count == wordCount[j].count {
			return wordCount[i].word < wordCount[j].word
		}
		return wordCount[i].count > wordCount[j].count
	})

	result := make([]string, 0, k)
	for i := 0; i < k; i++ {
		result = append(result, wordCount[i].word)
	}
	return result
}
```

## One sentence separates this from yesterday

Day 78: *return the k most frequent elements, in any order.*

Today: *return them sorted by frequency from highest to lowest, and sort words with the same frequency by lexicographical order.*

That addition is the entire difference between the two problems, and it is enough to change which tool is right.

## Counting, then flattening

```go
m := make(map[string]int)
for i := 0; i < len(words); i++ {
	m[words[i]]++
}
```

![Step 1](images/walkthrough-1.png)

A Go map has no order, and iterating one gives a deliberately randomised sequence. So the counts are flattened into a slice before anything can be arranged.

## The comparator is the solution

```go
sort.Slice(wordCount, func(i, j int) bool {
	if wordCount[i].count == wordCount[j].count {
		return wordCount[i].word < wordCount[j].word
	}
	return wordCount[i].count > wordCount[j].count
})
```

Read it against the problem statement. "Sorted by frequency from highest to lowest" is `count > count`. "Ties by lexicographical order" is `word < word`. The code says what the requirement says.

![Step 2](images/walkthrough-2.png)

`i` and `love` both appear twice, so the tie-break decides, and `"i" < "love"`.

![Step 3](images/walkthrough-3.png)

![Step 4](images/walkthrough-4.png)

## Why not yesterday's heap

Yesterday's size-k min-heap gives the right *set* in ascending frequency, so today it would need reversing — and the tie-break has to live inside the comparator.

That comparator is where it turns unpleasant. The heap evicts its root, which must be the *worst* entry. For counts, worst means least frequent. For a tie, worst means the word that should lose, which is the lexicographically **later** one.

So the two fields point in opposite directions:

```go
if a.count == b.count {
	return a.word > b.word   // inverted for the tie-break
}
return a.count < b.count     // not inverted
```

That is correct, and it is exactly the kind of line that sits wrong in a codebase for a long time. Nothing about it looks unusual, and it only misbehaves on ties.

## The honest comparison

Sorting is O(d log d). The heap is O(d log k). On paper the heap wins when k is small and d is large.

Here `words.length` is capped at 500, so d is at most 500 and the difference is unmeasurable. What is measurable is that one comparator reads like the specification and the other has an inversion in it that has to be reasoned about every time it is read.

Days 76 to 78 built the case for the heap. This is the day to notice that the technique is not the goal.

## Two things that could be tightened

```go
result := make([]string, 0, k)
for i := 0; i < k; i++ {
```

If `k` exceeded the number of distinct words this would index past the end. The constraints forbid it, so it cannot happen — but `for i := 0; i < k && i < len(wordCount); i++` costs nothing and removes the dependence.

`sort.Slice` is not stable, which does not matter here: the comparator is a total order, with no pair left to arbitrary choice. Worth noticing, since an unstable sort with an incomplete comparator is a real source of irreproducible output.

## Complexity

- **Time: O(n + d log d)** — O(n) to count, then sorting d distinct words.
- **Space: O(d)** for the map and the slice.

## Test

`main_test.go` covers both worked examples. The first is the important one: `i` and `love` tie on frequency, so it fails if the tie-break is missing or the wrong way round.
