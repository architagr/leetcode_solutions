import os, sys
sys.path.insert(0, "/Users/architagarwal/code/codestreak-growth/tools")
from diagram import *

# Seven prices, two to a page.
W_ = 100
Y = [86, 186, 286]
POS = {40: (550, 0), 20: (270, 1), 60: (830, 1), 10: (130, 2), 30: (410, 2),
       50: (690, 2), 70: (970, 2)}
KIDS = {40: [20, 60], 20: [10, 30], 60: [50, 70]}
Pn, C, D, R = "pend", "cur", "done", "repeat"


def scene(states, row=None, row_title=None, notes=()):
    out = []
    for p, ks in KIDS.items():
        px, pl = POS[p]
        for k in ks:
            kx, kl = POS[k]
            out.append(connect(px + W_ / 2, Y[pl] + BH, kx + W_ / 2, Y[kl], width=3))
    for p, (x, l) in POS.items():
        out += box(x, Y[l], str(p), states.get(p, Pn), w=W_)
    if row is not None:
        if row_title:
            out.append(label(40, 392, row_title, 22, MUTED, anchor="start"))
        for i, (v, st) in enumerate(row):
            out += box(40 + i * 160, 404, str(v), st, w=120)
    for y, text, color in notes:
        out.append(label(600, y, text, 26, color))
    return out


ORDER = [10, 20, 30, 40, 50, 60, 70]

steps = [

 (scene({p: D for p in POS}, row=[(v, C if v in (10, 20) else D) for v in ORDER],
        row_title="sorted by price, two to a page: page 1 is 10, 20",
        notes=[(540, "the real catalogue: 200,000 products, 10,000 pages", STROKE)]),
  "the catalogue sorted by price, one page at a time"),

 (scene({p: R for p in POS}, row=[(v, C if v in (10, 20) else R) for v in ORDER],
        row_title="NewIterator: the whole in-order list, then page 1",
        notes=[(540, "7 prices listed to show 2. Real: 200,000 for 20.", REPEAT)]),
  "what you would write: build the sorted list, slice out the page"),

 (scene({40: C, 20: C, 10: C}, row=[(v, C) for v in (40, 20, 10)],
        row_title="stack: the left spine, nothing else",
        notes=[(500, "pop the top: 10. Pop again: 20, then push 30's spine.", MUTED),
               (540, "never more than the depth: 31 nodes for 200,000", STROKE)]),
  "a stack of the unfinished path: the next price is always on top"),

 (scene({10: R, 20: R, 30: R, 40: R, 50: C, 60: C},
        row=[(v, R) for v in (10, 20, 30, 40)] + [(v, C) for v in (50, 60)],
        row_title="page 3 by number: step past 4 prices first",
        notes=[(540, "skipping costs every product before the page", REPEAT)]),
  "a page number still means stepping past everything before it"),

 (scene({40: D, 20: Pn, 30: C}, row=[(v, C) for v in (40, 30)],
        row_title="page after 20: descend once, keep what is still to come",
        notes=[(500, "40 is after 20: keep it, go left. 20 is not: go right.", MUTED),
               (540, "30 is after 20: keep it. Next: 30, then 40. No skipping.", STROKE)]),
  "a cursor instead of a page number: one descent, then read the page"),
]
render(steps, os.path.join(os.path.dirname(os.path.abspath(__file__)), "images"))
