# Measured

```
Apple M1 Pro · go1.26.4 darwin/arm64
go test -bench=. -benchtime=50x -run=XXX
```

Reproduce with the command above in this directory.

Both lists are the same length. Half the incoming emails already exist, which
is the realistic case and also the fair one: a full miss would make the nested
version scan the entire existing list every single time, which flatters the
comparison.

| records each side | nested loop | set | ratio |
|---|---:|---:|---:|
| 100 | 13.7 µs | 5.31 µs | 2.6x |
| 1,000 | 1.12 ms | 40.3 µs | 27.8x |
| 5,000 | 30.3 ms | 282 µs | 108x |
| 20,000 | **516 ms** | 1.30 ms | **398x** |

Raw ns: 13,662 / 5,311 · 1,121,888 / 40,286 · 30,334,418 / 281,760 · 516,400,558 / 1,296,928

## Where it stops being fine

At 100 records the nested version takes 13.7 microseconds. Nothing in a request
is measured in microseconds, so this is free and always will be.

At 1,000 it is 1.12 ms. Still nothing.

At 20,000 it is **516 milliseconds**. Half a second, of one function, doing
nothing but comparing strings. That is past every latency budget worth having,
and 20,000 records is not a large number - it is a mid-size customer's user
table, or one nightly import.

## The growth rate

From 5,000 to 20,000 is four times the input:

- nested: 30.3 ms to 516 ms, a factor of **17.0**. Four in, sixteen out is
  quadratic, and it means the next doubling costs four times again.
- set: 282 µs to 1.30 ms, a factor of **4.6** against 4x the input. Linear,
  plus the map growing.

## The set version uses more memory, and that is the trade

| records | nested B/op | set B/op |
|---|---:|---:|
| 100 | 7,504 | 11,000 |
| 1,000 | 64,848 | 119,458 |
| 5,000 | 474,448 | 692,880 |
| 20,000 | 2,325,947 | 3,199,674 |

At 20,000 the set version allocates about **37% more** - 3.2 MB against 2.3 MB -
because the map has to hold every existing email. Both grow linearly in the
input, so the ratio stays roughly flat rather than getting worse.

That is the whole cost, and it is worth stating plainly rather than pretending
the fast version is free. Half a second of CPU for 0.9 MB is a trade almost
anybody would take, but it IS a trade.

## Comparison count

`TestNestedComparisonCount` asserts the nested version does exactly `n * n`
string comparisons when nothing matches - 400 million of them at n=20,000.
That is the number the benchmark is measuring.
