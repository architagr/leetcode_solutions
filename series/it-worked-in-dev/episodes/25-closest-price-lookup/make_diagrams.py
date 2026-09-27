import os, sys
sys.path.insert(0, "/Users/architagarwal/code/codestreak-growth/tools")
from diagram import *

# Seven prices and a target of 57. The nearest is 60.
W_ = 100
Y = [86, 186, 286]
POS = {50: (550, 0), 30: (270, 1), 70: (830, 1), 20: (130, 2), 40: (410, 2),
       60: (690, 2), 80: (970, 2)}
KIDS = {50: [30, 70], 30: [20, 40], 70: [60, 80]}
SORTED = [20, 30, 40, 50, 60, 70, 80]
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
    if row:
        if row_title:
            out.append(label(40, 392, row_title, 22, MUTED, anchor="start"))
        for i, (v, st) in enumerate(row):
            out += box(40 + i * 160, 404, str(v), st, w=120)
    for y, text, color in notes:
        out.append(label(600, y, text, 26, color))
    return out


steps = [

 (scene({60: C}, notes=[
     (500, "nearest to 57: 50 is 7 away, 60 is 3 away", MUTED),
     (540, "the answer is one price. The index holds 200,000.", STROKE)]),
  "the product priced nearest a target: here 57, and the answer is 60"),

 (scene({p: D for p in POS},
        row=[(v, C if v in (50, 60) else D) for v in SORTED],
        row_title="Sorted(root), then binary search for 57",
        notes=[(540, "the listing page's sorted list, then search it", STROKE)]),
  "what you would write: take the sorted prices, binary search"),

 (scene({p: R for p in POS},
        row=[(v, R) for v in SORTED], row_title="built per question",
        notes=[(540, "200,000 products: 51.5 MB and 32.8 ms per lookup", REPEAT)]),
  "every lookup rebuilds the whole sorted list to read two entries"),

 (scene({50: C, 60: C},
        row=[(v, C if v in (50, 60) else Pn) for v in SORTED],
        row_title="only these two can be nearest",
        notes=[(500, "the nearest price is a neighbour of 57 in sorted order", MUTED),
               (540, "the last price below it, or the first at or above it", STROKE)]),
  "only two prices can be nearest: the neighbours either side"),

 (scene({50: D, 70: D, 60: C},
        notes=[(500, "50 is below 57: note it, go right. 70 is above: note it, go left.", MUTED),
               (540, "60 is above: note it. Nil. Candidates 50 and 60: 60.", STROKE)]),
  "both neighbours lie on the search path: one descent, no list"),
]
render(steps, os.path.join(os.path.dirname(os.path.abspath(__file__)), "images"))
