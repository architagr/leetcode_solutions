# Measured

```
Apple M1 Pro · go1.26.4 darwin/arm64
go test -bench=. -benchtime=100x
```

Reproduce with the command above in this directory.

| shape | people | levels | walking up | level order | ratio |
|---|---:|---:|---:|---:|---:|
| flat | 1,000 | 1 | 65.1 µs | 61.6 µs | 1.1x |
| a real company | 5,461 | 6 | 1.09 ms | 572 µs | 1.9x |
| deep thread | 200 | 200 | 503 µs | 19.8 µs | 25.5x |
| deeper thread | 2,000 | 2,000 | 64.0 ms | 255 µs | **251x** |

Raw ns: 65,143 / 61,641 · 1,089,758 / 571,889 · 502,970 / 19,755 · 63,977,230 / 255,102

## The headline is that a real org chart is fine

**5,461 people across six levels is 1.9x.** Half a millisecond against one. On a
page that also hits a database, nobody will ever find that.

A thousand people in one flat level is **1.1x** — indistinguishable, because
walking up from a direct report is a single hop.

That is the honest answer for an org chart, and it is worth saying before the
rest: if this is what your tree looks like, the naive version is correct code
and you should leave it alone.

## Depth is the axis, again

The two deep shapes are the same code on a different shape:

- 200 deep: **25.5x**
- 2,000 deep: **251x**, which is 64 ms of one function

From 200 to 2,000 is ten times the input for **127 times the work**. Ten in, a
hundred out is quadratic. The level-order version over the same change costs
12.9x against 10x the input.

It is not the number of people. It is how far the average person is from the
top, because that is exactly how many hops the upward walk makes.

## Where a tree is actually that deep

Not org charts. Threaded comments, nested categories, a bill of materials, a
file tree, a dependency chain - anything where the structure grows by nesting
rather than by breadth.

## Allocations

| shape | walking up | level order |
|---|---:|---:|
| a real company | 809 KB / 34 allocs | 1.12 MB / 4,152 allocs |
| deeper thread | 404 KB / 18 allocs | 370 KB / 4,019 allocs |

The level-order version allocates **more, and far more often** - a slice per
level, plus the reports index. On the shallow shapes it is carrying that cost
for a speed-up nobody needed, which is the other half of why the naive version
wins there.

## Hop count

`TestUpwardHopCount` asserts the upward walk makes `n(n-1)/2` hops on a chain of
`n` - 1,999,000 of them at 2,000 deep. That is what 64 ms is spent on.
