# Measured

```
Apple M1 Pro · go1.26.4 darwin/arm64
go test -bench=. -benchtime=100x
```

Reproduce with the command above in this directory.

| shape | nodes | asking twice | one pass | ratio |
|---|---:|---:|---:|---:|
| nav tree, 3 levels | 156 | 1.48 µs | 637 ns | 2.3x |
| nav tree, 4 levels | 781 | 9.51 µs | 4.05 µs | 2.4x |
| one branch, 200 deep | 201 | 146 µs | 2.50 µs | 58.6x |
| one branch, 1,000 deep | 1,001 | 3.95 ms | 13.2 µs | **298.7x** |

Raw ns: 1,482 / 637 · 9,514 / 4,047 · 146,488 / 2,498 · 3,954,433 / 13,239

## A real navigation tree is only 2.4x

781 categories across four levels: 9.51 µs against 4.05 µs. Both are
microseconds and neither is a problem. If your menu looks like this, the
readable version is the right code.

What changes the answer is not the number of categories. It is one branch
growing deep while the others stay shallow — which is exactly the condition the
check exists to find.

## The cross-shape figure, with both sides named

The write-up says the deep tree takes 415 times longer than the nav tree. That
is **asking-twice on 1,001 deep categories against asking-twice on 781 nav
categories** - 3.95 ms against 9.51 µs, or **415.6x**. One implementation, two
shapes.

It is a different comparison from the 298.7x in the table, which is the two
implementations on the same shape. Both are real and they must not be run
together.

## The growth rate

200 deep to 1,000 deep is five times the depth:

- asking twice: 146 µs to 3.95 ms, a factor of **27.0**. Five in, twenty-seven
  out is quadratic.
- one pass: 2.50 µs to 13.2 µs, a factor of **5.3** against 5x the input. Linear.

## Neither allocates

Both report `0 B/op` and `0 allocs/op` at every shape. The recursion carries
integers; the only allocation either makes is the result slice, which is the
same size for both and empty on a balanced tree.

So there is no memory trade here, the same as episode 3. The cost is elsewhere -
see the write-up: the two functions report the same categories in **different
orders**, and the one-pass version fuses two concerns that the readable version
keeps apart.

## Call count

`TestDepthCallCount` asserts `depth()` is entered `n(n+1)/2` times on a chain of
`n` below the root — 500,500 calls at 1,000 deep. That is what 3.95 ms is spent
on.
