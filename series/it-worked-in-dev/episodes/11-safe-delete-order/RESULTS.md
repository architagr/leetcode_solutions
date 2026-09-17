# Measured

```
Apple M1 Pro · go1.26.4 darwin/arm64
go test -bench=. -benchtime=300ms
```

Reproduce with the command above in this directory.

## The repeated scan, on its own

| shape | records | waves | repeated scan |
|---|---:|---:|---:|
| one board of cards | 21 | 2 | 2.19 µs |
| a workspace: 3 projects, 5 boards, 20 cards | 319 | 4 | 31.1 µs |
| one tenant, wider again | 1,245 | 4 | 123 µs |
| categories nested 10 deep | 2,047 | 11 | 386 µs |
| a reply to a reply, 2,000 deep | 2,000 | 2,000 | 75.8 ms |

Raw ns: 2,192 · 31,081 · 122,765 · 386,446 · 75,837,800

**1.6x the records cost 618x the time**, comparing the 1,245-record tenant
against the 2,000-record thread. The two differ in nesting, not in size.

## Both implementations

| shape | records | repeated scan | one pass | ratio |
|---|---:|---:|---:|---:|
| one board | 21 | 2.19 µs | 471 ns | 4.7x |
| a workspace | 319 | 31.1 µs | 3.46 µs | 9.0x |
| one tenant | 1,245 | 123 µs | 16.0 µs | 7.7x |
| categories, 10 deep | 2,047 | 386 µs | 28.2 µs | 13.7x |
| a 2,000-deep thread | 2,000 | 75.8 ms | 95.9 µs | 790x |

Raw ns: 2,192 / 471.2 · 31,081 / 3,461 · 122,765 / 15,984 · 386,446 / 28,206 ·
75,837,800 / 95,936

Every ratio compares `WavesByRepeatedScan` against `WavesByOnePass` on the same
hierarchy. Both produce the same plan — the same waves, with the same records
in each.

## Memory

| shape | repeated scan | one pass |
|---|---|---|
| a workspace, 319 | 38.6 KB / 53 allocs | 10.1 KB / 21 allocs |
| one tenant, 1,245 | 167 KB / 69 allocs | 62.2 KB / 26 allocs |
| a 2,000-deep thread | 354 KB / 4,042 allocs | 190 KB / 2,013 allocs |

Memory is not where this one hurts. It is under 2x everywhere, because both
versions build the same plan and the scan's extra cost is time, not space.

## The visits counted, not claimed

`TestScanWalksEverythingPerWave` counts the walk and asserts it comes to
`(waves + 1) × records` — one full walk per wave, plus the last walk that finds
nothing left to delete. On the 2,000-deep thread that is **4 million record
visits** to produce a plan with 2,000 waves in it.

## Grouping by depth is safe, and answers something else

`WavesByDepth` groups by distance from the root and deletes the deepest level
first. `TestEveryPlanIsSafeToApply` confirms it never deletes a parent before
its children.

On the hierarchy in `TestDepthAnswersADifferentQuestion` — a workspace holding
one three-level project and two empty boards — its first wave contains **one
record** where three are deletable immediately. Both plans take 4 waves here,
so the cost is not round trips; it is that "what can I delete first" comes back
with a third of the answer.
