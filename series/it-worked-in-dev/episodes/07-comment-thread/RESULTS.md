# Measured

```
Apple M1 Pro · go1.26.4 darwin/arm64
go test -bench=. -benchtime=300ms -run=XXX
```

Reproduce with the command above in this directory.

**The `-benchtime=300ms` matters.** An earlier run at `-benchtime=200x` reported
9,370 ns for a 61-comment walk that actually takes under a microsecond - eleven
times out, and in the direction that would have made the write-up wrong. Fixed
iteration counts are too few to be stable at these sizes. Let Go pick.

| shape | comments | sort key | walking | ratio |
|---|---:|---:|---:|---:|
| a discussion, 2 deep | 73 | 16.9 µs | 912 ns | 18.6x |
| a busy thread, 3 deep | 259 | 51.6 µs | 4.14 µs | 12.4x |
| an argument, 60 deep | 61 | 47.4 µs | 967 ns | 49.0x |
| a long argument, 400 deep | 401 | 1.50 ms | 7.48 µs | **200.5x** |

Raw ns: 16,942 / 912.1 · 51,557 / 4,144 · 47,406 / 966.6 · 1,500,670 / 7,484

## Depth costs the sort key twice

Row three is the one to look at. **61 comments cost nearly as much as 259** - 47.4 µs against 51.6 µs - because
those 61 are sixty levels deep and the 259 are three.

Depth hits the sort-key version on two axes at once. Each key is one segment per
level, so a comment sixty deep carries a string sixty segments long - and every
comparison during the sort walks those strings to decide. Longer keys, and more
of the key read per comparison.

From 60 deep to 400 deep is 6.7 times the depth for **31.7 times the work**.

## Allocations

| shape | sort key | walking |
|---|---|---|
| a discussion, 73 | 19.0 KB / 301 allocs | 3.75 KB / 4 allocs |
| a busy thread, 259 | 67.2 KB / 1,046 allocs | 15.8 KB / 6 allocs |
| a long argument, 401 | **2.03 MB / 1,263 allocs** | **15.8 KB / 6 allocs** |

At 400 deep the sort-key version allocates **132 times the memory**. Every
comment gets a path built for it, and the path grows with depth.

The walk allocates six times, total - it appends into one slice that doubles as
it grows, and nothing else.

## Key characters

`TestSortKeyCharacterCount` asserts the total key length on a chain is the sum
of the depths, which is quadratic. That is the string-building half of the cost,
before a single comparison happens.
