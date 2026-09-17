# Measured

```
Apple M1 Pro · go1.26.4 darwin/arm64
go test -bench=. -benchtime=300ms
```

Reproduce with the command above in this directory. Time-based rather than a
fixed iteration count, because the scan at 4,197 nodes runs three times in 300
milliseconds and a fixed count would either take all afternoon or measure one
sample.

## The scan, on its own

| shape | nodes | widest pair | scan from every node |
|---|---:|---:|---:|
| one rack, 24 instances | 25 | 2 hops | 8.03 µs |
| a fleet: 4 zones, 10 racks, 30 each | 1,245 | 6 hops | 12.9 ms |
| a fleet one level deeper | 4,197 | 8 hops | 153 ms |
| a topology that is really a chain | 2,000 | 1,999 hops | 73.3 ms |

Raw ns: 8,027 · 12,862,965 · 152,958,930 · 73,287,100

**3.4x the machines cost 11.9x the time** — 1,245 to 4,197 nodes, 12.9 ms to
153 ms. That is the whole episode in one row pair.

The title and the drafts round that to **12x**, and nothing else is rounded
anywhere: every other figure here is the benchmark's own.

## Both implementations

| shape | nodes | scan | one pass | ratio |
|---|---:|---:|---:|---:|
| one rack | 25 | 8.03 µs | 78 ns | 102x |
| a fleet | 1,245 | 12.9 ms | 3.62 µs | 3,554x |
| a deeper fleet | 4,197 | 153 ms | 13.3 µs | 11,512x |
| a chain | 2,000 | 73.3 ms | 18.6 µs | 3,933x |

Raw ns: 8,027 / 78.38 · 12,862,965 / 3,619 · 152,958,930 / 13,287 ·
73,287,100 / 18,633

Every ratio above compares `WidestByScan` against `WidestByOnePass` on the
same topology.

## Memory

| shape | scan | one pass |
|---|---|---|
| one rack, 25 | 14.9 KB / 206 allocs | 0 B / 0 allocs |
| a fleet, 1,245 | 44.9 MB / 52,437 allocs | 0 B / 0 allocs |
| a deeper fleet, 4,197 | **529 MB** / 273,728 allocs | 0 B / 0 allocs |
| a chain, 2,000 | 44.3 MB / 4,004,022 allocs | 0 B / 0 allocs |

529 megabytes of garbage to return one small integer about 4,197 machines. The
one-pass version allocates nothing at all: it carries two ints per frame and
returns one.

## Removing the garbage does not fix it

`WidestByScanReusing` is the same scan with one visited slice and two frontier
buffers, allocated once and reused for every starting point.

| shape | scan | scan, reusing buffers | | still slower than one pass |
|---|---:|---:|---|---:|
| one rack, 25 | 8.03 µs | 3.44 µs | 2.3x faster | 44x |
| a fleet, 1,245 | 12.9 ms | 3.70 ms | 3.5x faster | 1,022x |
| a deeper fleet, 4,197 | 153 ms | 42.2 ms | 3.6x faster | 3,175x |
| a chain, 2,000 | 73.3 ms | 22.6 ms | 3.2x faster | 1,211x |

Raw ns: 3,439 · 3,697,859 · 42,182,484 · 22,558,383

Allocation was the loudest symptom and not the problem. Deleting all of it
buys **3.6x** at 4,197 nodes; changing the shape of the work buys **11,512x**.

## Naming the pair costs something, and not much

`WidestPair` is the one-pass version carrying the deepest leaf's name up
alongside its height, so the answer reads "api-7 and web-3, 8 hops" rather
than "8".

| shape | one pass | one pass, naming the pair | |
|---|---:|---:|---|
| one rack, 25 | 78 ns | 110 ns | 1.4x |
| a fleet, 1,245 | 3.62 µs | 4.81 µs | 1.3x |
| a deeper fleet, 4,197 | 13.3 µs | 16.6 µs | 1.3x |
| a chain, 2,000 | 18.6 µs | 33.6 µs | 1.8x |

Raw ns: 110.1 · 4,814 · 16,625 · 33,602

Still zero allocations: the strings are already in the tree, and only the two
names that are currently winning are carried.

## The work counted, not claimed

`TestScanVisitsEveryNodePerStart` counts breadth-first visits and asserts they
come to exactly `nodes × nodes` — 40,000 visits on a 200-node chain. That
product is what 153 ms is spent on at 4,197 nodes: 17.6 million visits.

## The shortcut that is wrong, not just slow

`WidestFromRoot` takes the root's two deepest branches. On the topology in
`TestRootOnlyIsWrong` — one zone holding two three-level branches, beside a
zone holding nothing — it returns **5 hops** where the true answer is **6**.
The winning path never touches the root, so a check anchored at the root
cannot see it.
