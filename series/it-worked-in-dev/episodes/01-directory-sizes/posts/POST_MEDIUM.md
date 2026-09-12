<!-- Medium. Import via the story importer pointed at the GitHub README so the
     canonical tag points home, or paste this and set the canonical link by
     hand under Story settings. Medium allows 5 tags; they are in the
     frontmatter. Subtitle goes in the kicker field under the title. -->

---
meta_title: "Depth, not size: the axis nobody benchmarks"
meta_description: "801 directories ran slower than 8,191. A measured Go walkthrough of the du problem, derived from the symptom rather than announced."
canonical: "https://github.com/architagr/leetcode_solutions/blob/main/series/it-worked-in-dev/episodes/01-directory-sizes/README.md"
hero: ../HERO.png
tags: [golang, algorithms, programming, software-engineering, performance]
hashtags: "#Golang #Algorithms #Performance #SoftwareEngineering #DataStructures"
---

![10x smaller. 3x slower.](../HERO.png)

# Depth, not size: the axis nobody benchmarks

## 801 directories ran three times slower than 8,191 directories. Same code, same machine.

Most writing about algorithmic complexity is true and does nothing. O(n²)
versus O(n) is a fact you can recite and still not act on, because it does not
tell you the one thing you actually want to know: *at what point does this start
costing me anything?*

So this is the first of a series that answers that with a number instead. A real
problem, the code a competent person writes first, a committed benchmark, and
the reframing — with the machine and the Go version named, because a number
without them is decoration.

## The problem

Report the total bytes under every directory in a tree. What `du` prints. A
directory's size is its own files plus everything nested under it.

## The version you would write

```go
func SizeOf(dir *Dir) int64 {
	var total int64
	for _, f := range dir.Files {
		total += f.Size
	}
	for _, child := range dir.Children {
		total += SizeOf(child)
	}
	return total
}

func AllSizesBrute(root *Dir) map[string]int64 {
	out := make(map[string]int64)
	var walk func(*Dir)
	walk = func(d *Dir) {
		out[d.Path] = SizeOf(d)
		for _, child := range d.Children {
			walk(child)
		}
	}
	walk(root)
	return out
}
```

I want to defend this for a moment, because the exercise is worthless if the
slow version is a strawman. `SizeOf` is the clearest possible statement of what
a directory's size is. Calling it once per directory is the obvious way to get
an answer for all of them. This passes review anywhere.

![SizeOf(/root) walks everything beneath it. Correct, and reasonable.](../images/walkthrough-2.png)

## Measure it before touching it

Apple M1 Pro, go1.26.4, `go test -bench=BenchmarkBrute -benchtime=200x`. The
brute force alone, no comparison yet:

| shape | directories | brute force | |
| --- | --- | --- | --- |
| balanced, depth 12 | 8,191 | 864,951 ns | 0.86 ms |
| chain, depth 200 | 201 | 157,175 ns | 0.16 ms |
| chain, depth 800 | 801 | 2,633,461 ns | 2.63 ms |

Two things fall out of that table.

Four times the directories, in the bottom pair, cost **16.8 times the work**.
Four in, sixteen out is the fingerprint of a quadratic, which means every extra
level of nesting is more expensive than the level before it.

And 8,191 directories finish in **a third** of the time 801 directories take.
Ten times the input, a third of the cost. The 8,191 are twelve levels deep. The
801 are eight hundred.

Size was never the axis. Depth is. Which is exactly what a `node_modules` is:
not enormous, just nested as far as the dependency graph goes.

## From the symptom to the shape

This is the section I care about, and it is the one most articles skip by going
straight from the slow code to the fast code. Next time there will be no
article. There will be slow code that belongs to you and a hunch.

**The issue, said plainly: the same numbers are added more than once.**

`a1` holds 3 bytes. That 3 gets added in once when the walk reaches `a1`, again
when `a` is computed, and again when `/root` is computed. Three times, for a
directory two levels down. A directory `d` levels deep is summed `d + 1` times,
once for itself and once per ancestor. That is the quadratic, and it is why
depth is the axis rather than size.

**Why is it allowed to happen?** Because `SizeOf(a)` has no memory of
`SizeOf(a1)`. Every call rebuilds its subtree total from the leaves. Nothing is
wrong with `SizeOf` in isolation, which is exactly why it passes review - the
problem is only visible in the walk that calls it.

![Then SizeOf(a) walks a1 and a2 again, and they were just counted](../images/walkthrough-3.png)

**And we already had the answer.** When the walk reaches `a1` it computes 3 and
writes it into `out`. The answer is sitting in a map we are holding. Three
frames later `/root` derives that same 3 from scratch and never looks at the
map. The information was not missing, it was produced too late to use. The
results are coming out in the wrong order.

Notice that nothing so far has named an algorithm.

**So what shape is this data?** Directories nest. Every directory has exactly
one parent. Nothing loops back on itself. Follow a path down and it ends. One
parent, no cycles, a single root - that is a **tree**, and it is worth saying
out loud, because the next step is only true for trees.

![Directories nest, every one has a single parent, nothing loops back. A tree.](../images/walkthrough-5.png)

**On a tree, write down what you want as an equation:**

```
size(d) = sum of d's own files
        + size(child) for each child of d
```

Every term on the right is either a number sitting in `d`, or the answer for a
node strictly below `d`. Nothing refers to `d`'s parent, `d`'s siblings, or the
path taken to reach `d`. That is the property to go looking for in your own
problems:

**Does the answer for a node depend only on the answers for its children?**

**Then the order is forced.** If `size(d)` needs `size(child)`, every child has
to be finished before `d` can be. Not preferred, required. Children first, then
the parent, applied to the whole tree. That order has a name: **postorder**.

**One piece left - how does a child's answer reach its parent?** It returns it.
That is why the fast version's signature changes from `func(*Dir)` to
`func(*Dir) int64`. The return value is the mechanism.

**And when does none of this apply?** Go back to the equation. It closed because
the right-hand side mentioned only children. The moment a node needs its parent,
its siblings, or the path taken to reach it, the equation no longer closes and
postorder gives you nothing. Depth-of-each-node needs the path. Distance between
two nodes needs a common ancestor. Both are trees. Neither is this.

Knowing which one you are holding is the skill. The traversal is what happens
after you know.

## Stop and write it

You have the rule, and the change is small: no new data structure, no cache, no
second pass.

```bash
git clone https://github.com/architagr/leetcode_solutions
cd leetcode_solutions/series/it-worked-in-dev/episodes/01-directory-sizes
go test ./...
```

The tests will tell you when you have it. The rep is the thing you are here for,
and reading the next section costs it.

## The version that scales

```go
func AllSizesPostorder(root *Dir) map[string]int64 {
	out := make(map[string]int64)
	var visit func(*Dir) int64
	visit = func(d *Dir) int64 {
		var total int64
		for _, f := range d.Files {
			total += f.Size
		}
		// Recurse first. Each child hands back a number it computed once.
		for _, child := range d.Children {
			total += visit(child)
		}
		// Only now is this directory's own answer known.
		out[d.Path] = total
		return total
	}
	visit(root)
	return out
}
```

Two lines moved. `out[d.Path]` went from before the recursion to after it, and
the function returns its total rather than discarding it.

That return value is the mechanism. It is how a finished child answer reaches a
parent without the parent walking down to find it.

![Every file counted once. Same answers, one pass.](../images/walkthrough-7.png)

## Now both of them

| shape | directories | brute force | postorder | ratio |
| --- | --- | --- | --- | --- |
| balanced, depth 3 | 85 | 4.96 µs | 3.91 µs | 1.3x |
| balanced, depth 12 | 8,191 | 865 µs | 631 µs | 1.4x |
| chain, depth 200 | 201 | 157 µs | 12.5 µs | 12.6x |
| chain, depth 800 | 801 | 2,633 µs | 76.3 µs | 34.5x |

Eight thousand directories, forty percent slower. You would ship that and never
hear about it.

Eight hundred directories in a line, thirty-four times slower.

Going from 200-deep to 800-deep, the postorder version costs a factor of 6.1
against 4x the input — linear, plus the map growing. The brute force costs 16.8.
That gap is two lines of code.

## What the fast version costs

Almost nothing, here. Both allocate identically — same `B/op`, same `allocs/op`
at every shape — because the output map is the same size either way.

The real cost is structural. `SizeOf` as a standalone function disappears.
Callers who want one directory's size now either keep the slow function
alongside or call the whole-tree version and index into it. That is a genuine
API trade, and on a small tree it is not worth making.

**At your scale, the naive version is probably fine.** It falls over on depth,
not on size, and most trees are not deep. The value is knowing which axis to
watch before it matters.

## Where the technique comes from

This is three days of my 365-day LeetCode series, applied to something with
bytes in it. Each one contributes a different piece of the derivation above:

- **[Day 1 — Maximum Depth of Binary Tree](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/101_200/maximum_depth_of_binary_tree/SOLUTION.md)** — a node's depth is 1 + the
  deeper of its children. The equation with only children on the right-hand
  side, and a recursion that returns a number up the tree.
- **[Day 4 — Sum of Left Leaves](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/401_500/sum_of_left_leaves/SOLUTION.md)** — a parent reads what its children
  returned instead of going down to look. The mechanism.
- **[Day 9 — Binary Tree Postorder Traversal](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/101_200/binary_tree_postorder_traversal/SOLUTION.md)** — children before the
  parent, stated on its own. The order.

All of it, and the other 78 days so far, is at
[github.com/architagr/leetcode_solutions](https://github.com/architagr/leetcode_solutions).

That is what the daily grind is actually building. Not trivia. A set of shapes
you recognise when the slow code is yours.

## The one line to keep

When a recursive function is called at every node *and* re-walks the subtree
underneath it, the subtree's answer wants to be a return value, not a second
walk.

---

## Where to keep reading

These run two to three times a week, and there are two ways to not miss one.

**The Weekly Golang Journal** — https://www.linkedin.com/newsletters/the-weekly-golang-journal-7261403856079597568/ — is my newsletter on
LinkedIn. Every episode goes out in it, along with the Go writing that does not
fit an article. If LinkedIn is somewhere you already are, that is the lowest
friction option by a distance.

**Substack** — https://substack.com/@architagr — if you would rather it arrived by email and stayed
out of a feed.

If you would rather keep reading here, follow on Medium and the next episode
turns up in your feed. Same content either way, so pick one.

If you want to argue with me about any of it, the comments that actually happen
are on LinkedIn: https://www.linkedin.com/in/architagarwal984/.

The code, the tests, the benchmark you can run yourself and all seven
walkthrough diagrams: https://github.com/architagr/leetcode_solutions/tree/main/series/it-worked-in-dev

The daily LeetCode series, all 365: https://github.com/architagr/leetcode_solutions

#Golang #Algorithms #Performance #SoftwareEngineering #DataStructures
