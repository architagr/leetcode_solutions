# Measured

```
Apple M1 Pro · go1.26.4 darwin/arm64
go test -bench=. -benchtime=100x      (three runs, median reported)
go test -run TestPasses -v
```

Reproduce with the commands above in this directory.

Every shape is 5,000 packages in layers, each package importing one to three
packages from the layer below. `deps_first` lists libraries before the apps
that use them; `apps_first` is the same kind of repo listed the other way
round, which is what an alphabetical `apps/ ... libs/` tree gives you. The
number is how many layers deep the imports go.

`TestPasses`:

```
deps_first_10       5000 packages, built  5000, passes    1
apps_first_10       5000 packages, built  5000, passes   10
apps_first_100      5000 packages, built  5000, passes  100
apps_first_1000     5000 packages, built  5000, passes 1000
chain_5000          5000 packages, built  5000, passes 5000
cycle_10            5000 packages, built  3665, passes   11  cycle: 2 packages
```

The pass loop makes one pass per layer when the list is in the wrong order,
and one pass in total when it is in the right order.

## Time

| shape | layers | pass loop | Kahn, a slice per package | Kahn, flat arrays |
|---|---:|---:|---:|---:|
| `deps_first_10` | 10 | 55.1 µs | 281 µs | 115 µs |
| `apps_first_10` | 10 | 99.3 µs | 276 µs | 112 µs |
| `apps_first_100` | 100 | 474 µs | 312 µs | 144 µs |
| `apps_first_1000` | 1,000 | 3,735 µs | 331 µs | 129 µs |
| `chain_5000` | 5,000 | 18,757 µs | 276 µs | 128 µs |
| `cycle_10` | 10 | 74.2 µs | 273 µs | 104 µs |

Raw ns, in column order:

```
55,067 / 281,410 / 115,234
99,291 / 275,580 / 111,889
474,244 / 311,581 / 143,699
3,735,157 / 330,958 / 128,926
18,757,135 / 276,224 / 127,555
74,175 / 272,832 / 103,862
```

Pass loop against flat Kahn, same repo: **0.478x**, **0.887x**, **3.3x**,
**29x**, **147x**, **0.714x**. Below 1, the pass loop is the faster one.

Kahn with a slice per package against the pass loop, same repo: **5.11x**,
**2.78x** slower on the 10-layer repos, **3.68x** on `cycle_10`. It wins from
100 layers: the pass loop is **1.52x**, **11.3x**, **67.9x** slower there.

Kahn with a slice per package against flat Kahn: **2.44x**, **2.46x**,
**2.17x**, **2.57x**, **2.17x**, **2.63x**. Same algorithm; the difference is
allocation.

`apps_first_10` against `chain_5000`: the pass loop takes **189x** as long;
flat Kahn **1.14x**.

## Allocated, same runs

| shape | pass loop | Kahn, a slice per package | Kahn, flat arrays |
|---|---:|---:|---:|
| 10 layers | 45.3 KB, 2 allocs | 331 KB, 8,205 allocs | 232 KB, 5 |
| 100 layers and deeper | 45.3 KB, 2 allocs | 347 KB, 9,139 to 9,995 allocs | 240 KB, 5 |

Raw bytes: 46,336 / 339,462 / 237,568 on `apps_first_10`;
46,336 / 355,194 / 245,760 on `apps_first_100`.

The pass loop allocates its `built` flags and its output. Kahn allocates those
plus a count per package and the reverse edges - who imports whom - which is
what it needs to know which packages a build can unblock.

## The cycle

`cycle_10` adds one import from a bottom-layer package back to a top-layer app
that depends on it. Both versions stop with 3,665 of 5,000 built: the cycle and
everything that depends on it. `Cycle` follows unbuilt imports from the first
unbuilt package until it revisits one, and names the loop: 2 packages here.
