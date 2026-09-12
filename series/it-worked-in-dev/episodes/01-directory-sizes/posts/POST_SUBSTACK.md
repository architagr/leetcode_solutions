<!-- Substack. Subject line is the H1. Images upload inline at the markers.
     Substack renders markdown tables and fenced code, so this pastes clean. -->

---
meta_title: "Why du crawls on node_modules: depth, not size, is the axis"
meta_description: "801 directories in a chain ran 3x slower than 8,191 in a balanced tree, and the naive version 34x slower than the fix. Both Go versions and the numbers."
hero: ../HERO.png
tags: [golang, algorithms, performance, recursion, dsa]
hashtags: "#Golang #DSA #Algorithms #Performance #SoftwareEngineering #DataStructures #100DaysOfCode"
---

![Why du crawls on node_modules](../HERO.png)

# The benchmark that surprised me: depth, not size

A thing I assumed was true for years turned out to be backwards, and a
benchmark caught it in about four minutes.

This is the first of a new series. Each one takes a problem a working developer
has actually hit, writes the version you would genuinely write first, measures
where it stops being fine, and works out the reframing — with the code and the
benchmark committed so you can run it yourself and get your own numbers.

Not complexity classes. Nanoseconds on a machine I can name.

## The problem

Report the total bytes under every directory in a tree. What `du` prints.

## The version I wrote

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
```

Then call it for every directory. That is it. And I want to be clear that this
is not a strawman — `SizeOf` is correct, it is the clearest possible statement
of what a directory size is, and calling it once per directory is the obvious
way to get every answer.

![SizeOf(/root) walks everything beneath it. Correct, and reasonable.](../images/walkthrough-2.png)

## Where I expected it to hurt

I expected big trees. I benchmarked big trees. Apple M1 Pro, go1.26.4:

| shape | directories | brute force | |
|---|---|---|---|
| balanced, depth 12 | 8,191 | 864,951 ns | 0.86 ms |
| chain, depth 200 | 201 | 157,175 ns | 0.16 ms |
| chain, depth 800 | 801 | 2,633,461 ns | 2.63 ms |

Look at the top row against the bottom. **8,191 directories finish three times
faster than 801 directories.** Ten times the input, a third of the time.

The 8,191 are twelve levels deep. The 801 are eight hundred.

And within the chains, four times the directories cost 16.8 times the work.
Four in, sixteen out. That is a quadratic, and it means each additional level of
nesting costs more than the level before it.

Size was never the axis. Depth is. Which is precisely the shape of a
`node_modules` — not enormous, just nested as far as the dependency graph goes.

## From the symptom to the shape

Not the fix. The fix is two lines and you will have guessed it by the end of
this section. The part worth having is how you get there when nobody has
written you an article - when it is just your own slow code and a hunch.

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

## Stop here and write it

Seriously. You have the rule and it is a small change: no new data structure, no
cache, no second pass.

```bash
git clone https://github.com/architagr/leetcode_solutions
cd leetcode_solutions/series/it-worked-in-dev/episodes/01-directory-sizes
go test ./...
```

The tests will tell you when you have it. Ten minutes, and the rep is the thing
you are actually here for.

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

`out[d.Path]` moved from before the recursion to after it, and the function
returns its total rather than discarding it. The return value is the mechanism.
It is how a finished child answer reaches a parent without the parent walking
down to find it.

![Every file counted once. Same answers, one pass.](../images/walkthrough-7.png)

## Both of them

| shape | directories | brute force | postorder | ratio |
|---|---|---|---|---|
| balanced, depth 3 | 85 | 4.96 µs | 3.91 µs | 1.3x |
| balanced, depth 12 | 8,191 | 865 µs | 631 µs | 1.4x |
| chain, depth 200 | 201 | 157 µs | 12.5 µs | 12.6x |
| chain, depth 800 | 801 | 2,633 µs | 76.3 µs | 34.5x |

Eight thousand directories: 40% slower. You would ship that and never hear
about it.

Eight hundred directories in a line: 34x.

## The honest caveat

At your scale the naive version is probably fine. It falls over on depth, and
most trees are not deep. Allocations are identical between the two, so there is
no memory argument either.

The fast version also costs something real. `SizeOf` as a standalone function
goes away — callers wanting one directory now index into the whole-tree map or
keep the slow function beside it. On a small tree, do not bother.

Knowing which axis to watch is the part that keeps.

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

## Sharing this one

You are already getting these in the inbox, so there is nothing to subscribe
to. But two things are worth having.

If you want to send this to someone, the **Medium** version is the one that
travels — https://medium.com/@architagr. Same piece, and people who do not know you will open a
Medium link when they will not open a forwarded email.

If you want to argue with me about it, **LinkedIn** is where the comments
actually happen — https://www.linkedin.com/in/architagarwal984/.

The code, the tests, the benchmark you can run yourself and all seven
walkthrough diagrams: https://github.com/architagr/leetcode_solutions/tree/main/series/it-worked-in-dev

The daily LeetCode series, all 365: https://github.com/architagr/leetcode_solutions

Next episode: the nested loop that was fine until it wasn't. Two to three of
these a week.

#Golang #DSA #Algorithms #Performance #SoftwareEngineering
