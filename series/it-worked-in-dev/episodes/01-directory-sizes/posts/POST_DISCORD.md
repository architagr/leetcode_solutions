<!-- Discord. ~70 members, so this is the one surface where the "try it before
     you read the answer" beat can actually be interactive instead of
     rhetorical. Posted in two parts with a real gap between them:

     Part 1 goes up first and withholds the fix entirely. It is a question the
     server can answer.
     Part 2 goes up the next day, after people have replied.

     No images - Discord embeds the GitHub link's preview on its own, and
     attachments push the code block below the fold on mobile.

     Set the thread title to the episode title so it is findable later. -->

---
meta_title: "It worked in dev #1 - the du benchmark"
meta_description: "Discord drop in two parts: the benchmark that surprised me, then the reveal a day later once people have posted their answers."
hero: ../HERO.png
hashtags: "#golang #dsa"
---

![10x smaller. 3x slower.](../HERO.png)

## Part 1 — post this first, hold the answer back

**It worked in dev #1 — a benchmark I got wrong**

Starting a new thing alongside the daily LeetCode posts. Real problem, the code you would actually write, a benchmark, and then the version that scales. First one is up, but I want to try something here before I post the write-up.

The job: report the total bytes under every directory in a tree. What `du` prints.

Here is the version I wrote. I will defend it — `SizeOf` is the clearest possible statement of what a directory's size is, and calling it once per directory is the obvious way to get all of them.

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

Benchmarked on an M1 Pro, go1.26.4:

```
8,191 directories, 12 levels deep   0.86 ms
  201 directories, 200 levels deep  0.16 ms
  801 directories, 800 levels deep  2.63 ms
```

801 directories took three times longer than 8,191. Same code, same machine.

**Two questions, and I would rather have your answers than post mine:**

1. Why is the small tree slower than the big one? React with 🌲 when you see it.
2. Fix it without adding a cache, a second pass, or any new data structure. Drop your version in the thread.

The repo has both implementations and the tests, so you can check yourself instead of waiting for me:

```
git clone https://github.com/architagr/leetcode_solutions
cd series/it-worked-in-dev/episodes/01-directory-sizes
go test ./...
```

Reveal tomorrow. Genuinely worth ten minutes first — reading the answer costs you the rep, and the whole point of this series is the reps.

---

## Part 2 — post the next day, after replies

**The reveal — and the part I actually care about**

<@ everyone who answered> got it, and the answers in here were better than mine. Pulling the reasoning together, because the fix is the boring half.

The mechanism: the walk visits each directory once, but `SizeOf` re-walks that directory's **entire subtree** every call. A file 800 levels deep gets added into a running total 800 separate times, once per ancestor asking how big it is. Four times the directories cost 16.8 times the work. Four in, sixteen out is a quadratic.

So it was never size. It is depth. Which is what a `node_modules` actually is — not enormous, just nested as far as the dependency graph goes.

Now the part worth keeping, because next time there is no Discord thread, only your own slow code and a hunch. Walking it the way you would have to:

**The issue, plainly:** the same numbers get added more than once. `a1` holds 3 bytes, and that 3 is added in once by `a1`, again by `a`, again by `/root`. A directory `d` levels deep is summed `d + 1` times, once for itself and once per ancestor.

**Why is that allowed?** `SizeOf(a)` has no memory of `SizeOf(a1)` — every call rebuilds its subtree from the leaves. Nothing is wrong with `SizeOf` on its own, which is why it survives review. The problem is only visible in the walk that calls it.

**But we already had the answer.** When the walk reached `a1` it computed 3 and wrote it into `out`. It is sitting in a map we are holding. Three frames later `/root` derives that same 3 from scratch and never checks. The information was not missing, it was produced too late to use. Wrong order, not wrong arithmetic.

Nothing above names an algorithm yet. Here is where the shape decides it.

**What shape is this data?** Directories nest. Every one has exactly one parent. Nothing loops back on itself. One parent, no cycles, a single root — a **tree**. Worth saying out loud, because what follows is only true for trees.

**On a tree, write the thing you want as an equation:**
```
size(d) = d's own files + size(child) for each child
```
Every term on the right is either a number sitting in `d`, or the answer for a node strictly *below* `d`. Nothing refers to `d`'s parent, siblings, or the path taken to reach it. That is the property to hunt for in your own problems: **does the answer for a node depend only on the answers for its children?**

**Then the order is forced.** If `size(d)` needs `size(child)`, children finish first. Not preferred — required. Children before parent, over a whole tree, is **postorder**.

**Last piece: how does a child's answer reach its parent?** It returns it. Which is why the signature goes from `func(*Dir)` to `func(*Dir) int64`.

**And when does none of this work?** When the equation stops closing — a node needing its parent, its siblings, or the path taken to reach it. Depth-of-each-node needs the path. Distance between two nodes needs a common ancestor. Both trees, neither postorder. Knowing which one you are holding is the actual skill.

```go
visit = func(d *Dir) int64 {
	var total int64
	for _, f := range d.Files {
		total += f.Size
	}
	for _, child := range d.Children {
		total += visit(child)   // child hands back a number it computed once
	}
	out[d.Path] = total         // only now is this directory's answer known
	return total
}
```

Two lines moved, and the function returns its total instead of discarding it. That return value is the whole mechanism.

```
                brute force   postorder   ratio
balanced 8,191       865 µs      631 µs    1.4x
chain      801     2,633 µs     76.3 µs   34.5x
```

Honest caveat, same as in the write-up: at your scale the naive one is probably fine. It falls over on depth and most trees are not deep, and both versions allocate identically. The value is knowing which axis to watch.

This is days 1, 4 and 9 of the daily series with bytes in it: [max depth](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/101_200/maximum_depth_of_binary_tree/SOLUTION.md) gives the equation shape, [sum of left leaves](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/401_500/sum_of_left_leaves/SOLUTION.md) gives the mechanism, [postorder](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/101_200/binary_tree_postorder_traversal/SOLUTION.md) gives the order. All 81 days are at <https://github.com/architagr/leetcode_solutions>.

Full write-up, seven walkthrough diagrams, and the benchmark you can run: https://github.com/architagr/leetcode_solutions/tree/main/series/it-worked-in-dev

**For next week:** a nested loop that is fine until it isn't. Same deal — I will post the benchmark first and hold the fix for a day. If you want a specific slow-code problem benchmarked, say so in here and I will take it.
