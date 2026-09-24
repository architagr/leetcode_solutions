# It worked in dev — 2-Day Go Performance Workshop

**Sat 10 + Sun 11 October 2026, 3–6 PM IST.** Hands-on, Go.

*Benchmark it, profile it, then pick the data structure.*

This directory is the workshop's starter code. Right now it holds the setup
check; the three labs land here before Day 2.

---

## Pre-work — do this before Saturday

From the repository root:

```bash
make verify
```

Ten minutes, most of it downloading `benchstat` once. It checks your Go
version, builds and tests this module, installs `benchstat`, and then runs four
benchmarks.

If anything fails, bring the error to the setup clinic on **Friday 9 October**
rather than to Saturday morning.

You need **Go 1.21 or newer** and a laptop you can type on.

---

## What `make verify` shows you

Four benchmarks, in two pairs. Each pair calls the same function with the same
argument — once throwing the result away, once keeping it. Same work, so each
pair should cost the same.

One pair does. The other does not:

```
                   │ sec/op        │
MixDiscarded-8       0.3184n ± 1%
MixKept-8             1.054n ± 1%   <- 3.3x apart
SumToDiscarded-8      324.9n ± 0%
SumToKept-8           324.4n ± 1%   <- the same
```

Ten runs each through `benchstat`, not one, because one run is not evidence —
measured on a loaded laptop a single run reported the SumTo pair 1.3x apart
when they are the same. The `± n%` is the variance, and it is how you tell a
real difference from a busy machine. Those are M1 Pro numbers; yours will
differ.

`Mix` discarded never ran. **0.315 ns on a 3.2 GHz machine is about one clock
cycle, and five operations do not happen in one cycle** — the number is its own
tell. Any benchmark result near a single cycle is a benchmark of nothing.

The obvious explanation is that `Mix` got inlined and `SumTo` did not. That
explanation is wrong, and you can check it yourself:

```bash
cd series/it-worked-in-dev/workshop
go test -gcflags='-m' -bench=XXX -run=XXX . 2>&1 | grep 'inlining call'
```

Both are inlined. What separates them is elsewhere, and it is hour one of
Day 1. **Bring the question, not the answer.**

Your numbers will differ — different CPU, different Go version. The ratio is
what matters. If your two pairs both agree, say so on the day; that is a real
result about your toolchain and worth the room's time.

---

## The harness

```bash
make bench
```

Ten runs at a fixed iteration count, piped through `benchstat`. This is the
thing you keep and point at your own code on Monday.

Fixed iterations rather than a fixed duration (`-benchtime=200x`, not `-benchtime=1s`)
so two runs on the same machine are comparable — which is the entire reason for
running ten of them and handing the output to `benchstat`. One run is not
evidence. `benchstat` reports variance, and a difference it marks `~` is a
difference you cannot claim.

If `benchstat` is not found after `make verify`:

```bash
export PATH="$(go env GOPATH)/bin:$PATH"
```

---

## Prep reading — free, and it spoils nothing

The [episodes](../episodes) are the best preparation. Each one takes a
production-shaped problem, benchmarks the obvious implementation, shows where
it collapses, and derives the fix.

| Episode | Why it is worth reading first |
|---|---|
| [01 — dedupe user list](../episodes/01-dedupe-user-list) | Free at 100 rows, 516 ms at 20,000 |
| [03 — pricing tier lookup](../episodes/03-pricing-tier-lookup) | A plain loop beats binary search by 1.7x at six tiers, loses by 22.9x at five thousand |
| [12 — pipeline never finishes](../episodes/12-pipeline-never-finishes) | A loop check that allocates 4.7 MB to return a boolean |

**The labs are new problems, not these.** Reading ahead helps and gives nothing
away.

---

## Layout

```
warmup.go        Mix and SumTo - the setup check's two subjects
warmup_test.go   the four benchmarks, and the lesson in the comment
```

The labs are added here before Day 2. Their solutions drop in Discord over the
week afterwards — the 13th, 15th and 17th — so try them yourself first.
