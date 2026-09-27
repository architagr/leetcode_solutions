# Measured

```
Apple M1 Pro · go1.26.4 darwin/arm64
go test -bench=. -benchtime=200x      (three runs, median reported)
go test -run 'TestPasses|TestCountingPassesCrossedBrackets' -v
```

Reproduce with the commands above in this directory.

`TestPasses`:

```
rules_100          4690 bytes, deepest nesting     5
rules_10k        561730 bytes, deepest nesting     5
nested_500         2000 bytes, deepest nesting   500
nested_5000       20000 bytes, deepest nesting  5000
crossed_10k      562074 bytes, deepest nesting     5
```

`rules_*` are files of generated pricing rules, up to five brackets deep.
`nested_*` is one expression wrapped in groups N deep, the shape a rule builder
produces when it wraps every added condition. `crossed_10k` is `rules_10k` with
two adjacent closers swapped halfway through: every count still balances.

## The check that is wrong

`TestCountingPassesCrossedBrackets`:

```
"([)]": counting says balanced, the stack says not
```

It also passes `crossed_10k`, and a crossed `nested_*`.

## Time

| shape | counting | erasing pairs | stack |
|---|---:|---:|---:|
| `rules_100` | 15.8 µs | 24.4 µs | 9.46 µs |
| `rules_10k` | 1.32 ms | 3.42 ms | 1.25 ms |
| `nested_500` | 3.83 µs | 333 µs | 3.62 µs |
| `nested_5000` | 37.7 µs | 21.7 ms | 38.2 µs |
| `crossed_10k` | 1.29 ms, says valid | 3.46 ms | 622 µs |

Raw ns, in column order:

```
15,825 / 24,352 / 9,461
1,323,785 / 3,420,445 / 1,245,787
3,835 / 333,150 / 3,620
37,732 / 21,693,196 / 38,182
1,294,181 / 3,464,133 / 621,862
```

Erasing against the stack, same file: **2.57x**, **2.75x**, **92x**, **568x**,
**5.57x**.

`nested_500` against `nested_5000`: 10x the depth, and erasing takes **65.1x**
as long. The stack takes **10.5x**.

Counting against the stack: **1.67x**, **1.06x**, **1.06x**, **0.988x**,
**2.08x**. It is not faster than the check that is right. On `crossed_10k` the
stack stops at the crossing, halfway through the file; counting reads to the
end and says valid.

## Allocated, same runs

| shape | counting | erasing pairs | stack |
|---|---:|---:|---:|
| `rules_100` | 0 B | 2.27 KB, 19 allocs | 0 B |
| `rules_10k` | 0 B | 349 KB, 35 allocs | 0 B |
| `nested_500` | 0 B | 261 KB, 508 allocs | 896 B, 3 |
| `nested_5000` | 0 B | 25.4 MB, 5,016 allocs | 17.4 KB, 9 |

Raw bytes: 2,328 · 357,313 · 267,450 / 896 · 26,618,476 / 17,792

The stack starts with room for 64 open brackets, so the rules files, five deep,
never grow it.
