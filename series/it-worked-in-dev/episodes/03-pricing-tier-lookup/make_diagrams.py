import os, sys
sys.path.insert(0, "/Users/architagarwal/code/codestreak-growth/tools")
from diagram import *

# Eight tiers in a row, drawn small enough to fit across the canvas. The
# argument is about how many of them a single lookup has to read.
N = 8
BOXW = 128
GAP = 12
LEFT = (W - (N*BOXW + (N-1)*GAP)) // 2
Y = 250
MINS = [0, 1, 5, 10, 25, 50, 100, 250]

def row(states, marks=None, note=None):
    marks = marks or {}
    out = [label(W/2, 150, "usage = 62k units. which tier?", 30, STROKE, bold=True)]
    for i in range(N):
        x = LEFT + i*(BOXW + GAP)
        # No sub-labels: box() draws them to the RIGHT, which lands on the
        # next box in a horizontal row. The detail goes in the note instead.
        out += box(x, Y, "%dk" % MINS[i], states.get(i, "pend"), w=BOXW)
    if note:
        out.append(label(W/2, Y + BH + 70, note, 26, MUTED))
    return out

P, C, D, R = "pend", "cur", "done", "repeat"

steps = [
 (row({}), "eight tiers, sorted by where each one starts"),

 (row({0: R, 1: R, 2: R, 3: R},
      note="each one only confirms what the ordering already promised"),
  "the scan reads tier 0, then 1, then 2, then 3 ..."),

 (row({0: R, 1: R, 2: R, 3: R, 4: R, 5: C}, note="six reads to rule out five tiers nothing could match"),
  "... until 100k is too high. The answer was tier 5."),

 (row({i: D for i in range(6)}, note="and the next lookup starts over at tier 0"),
  "every lookup re-reads tiers it already ruled out"),

 (row({5: C}, note="50k &lt;= 62k, so one comparison removes half the list"),
  "but compare the MIDDLE first: is 50k &lt;= 62k? yes"),

 (row({0: P, 1: P, 2: P, 3: P, 4: P, 5: D, 6: C}, note="100k &gt; 62k, so the answer is tier 5, 6 or 7 - then 5"),
  "everything below the middle is eliminated, unread"),

 (row({5: D}, note="3 comparisons instead of 6, and 12 instead of 5000"),
  "same answer. It never read the tiers it ruled out."),
]
render(steps, os.path.join(os.path.dirname(os.path.abspath(__file__)), "images"))
