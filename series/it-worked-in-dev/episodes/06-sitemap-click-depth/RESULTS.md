# Measured

```
Apple M1 Pro · go1.26.4 darwin/arm64
go test -bench=. -benchtime=200x
```

Reproduce with the command above in this directory.

| shape | pages | answer | recursion | by level | winner |
|---|---:|---:|---:|---:|---|
| marketing site + docs | 5,464 | 2 clicks | 11.6 µs | 117 ns | **level, 99.2x** |
| uniform, every leaf at depth 6 | 1,365 | 6 clicks | 2.85 µs | 7.11 µs | **recursion, 2.5x** |
| one long route, no shortcuts | 400 | 400 clicks | 3.48 µs | 5.83 µs | **recursion, 1.7x** |
| deep docs, shallow page last | 2,002 | 2 clicks | 21.8 µs | 58.8 ns | **level, 370.9x** |

Raw ns: 11,619 / 117.1 · 2,854 / 7,109 · 3,483 / 5,834 · 21,789 / 58.75

## It goes both ways, and the dividing line is not size

The two shapes the level walk wins are the two where **the answer is near the
top**. It stops as soon as it meets a dead end, so 5,464 pages cost the same as
the four it actually looked at.

The two it loses are the ones where nothing can be skipped. Every leaf in the
uniform tree is at depth 6, and the chain has a single route, so the level walk
reaches the bottom anyway - having paid for a queue the whole way down.

**What decides it is not how big the site is. It is how far the answer is from
the homepage**, which the recursion cannot know in advance and the level walk
does not need to.

## Allocations

| shape | recursion | by level |
|---|---|---|
| marketing site + docs | 0 B, 0 allocs | 104 B, 4 allocs |
| uniform, depth 6 | 0 B, 0 allocs | 41.3 KB, 16 allocs |
| one long route | 0 B, 0 allocs | 3.13 KB, 400 allocs |

**The recursion allocates nothing at all**, at any shape. It carries integers on
the call stack.

The level walk allocates a slice per level. On the winning shapes that is 104
bytes, because it only ever builds two of them. On the uniform tree it is
41.3 KB, because the widest level holds 1,024 pages - and it pays that to lose.

## Pages visited

`TestLevelWalkStopsEarly` counts them on a site with a 500-page docs section and
one shallow page: the recursion visits all 502, the level walk visits **3**.
That is the whole result, without a timer.
