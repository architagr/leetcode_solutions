# Measured

```
Apple M1 Pro · go1.26.4 darwin/arm64
go test -bench=. -benchtime=100x -run=XXX
```

Reproduce with the command above in this directory.

| shape | nodes | walking up | carrying down | ratio |
|---|---:|---:|---:|---:|
| one folder, 500 files | 501 | 47.9 µs | 36.1 µs | 1.3x |
| a project tree, 5 levels | 1,365 | 275 µs | 93.0 µs | 3.0x |
| 100 deep | 101 | 108 µs | 12.0 µs | 9.0x |
| 600 deep | 601 | 3.66 ms | 275 µs | **13.3x** |

Raw ns: 47,912 / 36,082 · 275,108 / 92,982 · 107,895 / 12,048 · 3,659,173 / 275,371

## The fast version allocates less, which is new

The two episodes before this one both bought speed with memory. This one does
not.

| shape | walking up | carrying down |
|---|---|---|
| one folder, 500 files | 58.0 KB / 1,504 allocs | 61.0 KB / 517 allocs |
| a project tree | 402 KB / 6,738 allocs | 146 KB / 1,386 allocs |
| 600 deep | **8.58 MB / 6,043 allocs** | **1.27 MB / 617 allocs** |

At 600 deep the upward version allocates **6.7x the memory across 9.8x as many allocations**. It builds a slice of segments per node, reverses
it, and joins it - all of which is thrown away immediately.

Carrying the prefix down does none of that. A child's path is its parent's
string plus one segment, so there is no collecting, no reversing and no join.

## Depth is the axis, again

100 deep to 600 deep is six times the depth:

- walking up: 108 µs to 3.66 ms, a factor of **33.9**. Six in, thirty-four out.
- carrying down: 12.0 µs to 275 µs, a factor of **22.9**

## Why even the fast version is not linear

Worth being precise about, because it is the honest limit of this fix.

In a chain of `n` nodes, the node at depth `d` has a path with `d + 1`
segments. The total size of the answer is therefore the sum of every depth -
quadratic in `n` - before anybody writes a line of code.

Carrying down removes the *repeated walking*. It cannot remove the output,
which is why it still grows faster than linearly on a chain. On a wide, shallow
tree - a real file browser - paths are short and both versions stay close to
linear, which is why that row is only 1.3x.

## Segment count

`TestUpwardSegmentCount` asserts the upward version copies `(d+1)(d+2)/2`
segments on a chain of depth `d` - 181,503 at 600 deep. That is what 3.66 ms is
spent on.
