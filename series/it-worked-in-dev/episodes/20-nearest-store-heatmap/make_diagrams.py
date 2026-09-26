import os, sys
sys.path.insert(0, "/Users/architagarwal/code/codestreak-growth/tools")
from diagram import *

# A 5 x 7 city with two stores. FINAL is the heatmap; PASS1 is the grid after
# the top-left sweep, where 99 stands for "not reached yet".
R, CO = 5, 7
STORES = [(0, 1), (4, 5)]
FINAL = [[1, 0, 1, 2, 3, 4, 5], [2, 1, 2, 3, 4, 3, 4], [3, 2, 3, 4, 3, 2, 3],
         [4, 3, 4, 3, 2, 1, 2], [5, 4, 3, 2, 1, 0, 1]]
PASS1 = [[99, 0, 1, 2, 3, 4, 5], [99, 1, 2, 3, 4, 5, 6], [99, 2, 3, 4, 5, 6, 7],
         [99, 3, 4, 5, 6, 7, 8], [99, 4, 5, 6, 7, 0, 1]]
CW, GAP = 62, 8
X0, Y0 = 40, 86
PX = 600
Pn, C, D, Rp = "pend", "cur", "done", "repeat"


def cell(r, c):
    return X0 + c * (CW + GAP), Y0 + r * (BH + GAP)


def grid(values, states, panel=(), title=None, notes=()):
    out = []
    for r in range(R):
        for c in range(CO):
            v = values[r][c] if values else None
            text = "" if v is None else ("-" if v == 99 else str(v))
            x, y = cell(r, c)
            out += box(x, y, text or " ", states(r, c), w=CW)
    if title:
        out.append(label(PX, 76, title, 22, MUTED, anchor="start"))
    for i, (text, state) in enumerate(panel):
        out += box(PX, 96 + i * 72, text, state, w=560)
    for y, text, color in notes:
        out.append(label(600, y, text, 26, color))
    return out


def stores_orange(r, c):
    return C if (r, c) in STORES else D


steps = [

 (grid(FINAL, stores_orange, title="the heatmap",
       panel=[("0 = a store", C), ("n = blocks to the nearest one", D)],
       notes=[(490, "walk along streets: up, down, left, right", MUTED),
              (530, "250,000 blocks in the real city", STROKE)]),
  "every block: how many blocks away the nearest store is"),

 (grid(None, lambda r, c: C if (r, c) == (2, 5) else (D if (r, c) in STORES else Pn),
       title="this block, scanning every store",
       panel=[("store at (0,1): 2 + 4 = 6", Rp), ("store at (4,5): 2 + 0 = 2", D),
              ("keep 2", D)],
       notes=[(490, "|dRow| + |dCol| to each store, keep the smallest", MUTED),
              (530, "exactly right, and no data structure at all", STROKE)]),
  "what you would write: every block checks every store"),

 (grid(FINAL, lambda r, c: Rp,
       title="the cost",
       panel=[("1 store: 1.03 ms", D), ("1,000 stores: 219 ms", Rp),
              ("250,000,000 distances", Rp)],
       notes=[(490, "blocks times stores, and one distance kept per block", MUTED),
              (530, "the chain opened more stores. 213x.", Rp)]),
  "every block measures every store, and keeps one number"),

 (grid(FINAL, lambda r, c: C if (r, c) == (2, 4) else
       (D if (r, c) in [(1, 4), (3, 4), (2, 3), (2, 5)] else Pn),
       title="the block in orange",
       panel=[("up 4, down 2, left 4, right 2", D), ("1 + smallest = 3", C)],
       notes=[(490, "a neighbour already knows its nearest store", MUTED),
              (530, "but every neighbour's answer depends on this one", STROKE)]),
  "the answer is one step from a neighbour's, and that is circular"),

 (grid(PASS1, lambda r, c: C if (r, c) in STORES else
       (D if PASS1[r][c] == FINAL[r][c] else Rp),
       title="after sweep 1: top-left to bottom-right",
       panel=[("read only up and left", D), ("both already final", D),
              ("red: store is down or right", Rp)],
       notes=[(530, "no queue: each block read in order, once", STROKE)]),
  "sweep 1 settles every block whose nearest store is up or left"),

 (grid(FINAL, lambda r, c: C if (r, c) in STORES else
       (D if PASS1[r][c] == FINAL[r][c] else C),
       title="after sweep 2: bottom-right to top-left",
       panel=[("read only down and right", D), ("keep the smaller of the two", D),
              ("orange: improved by sweep 2", C)],
       notes=[(530, "1.48 ms at 1,000 stores, and no 3.82 MB queue", STROKE)]),
  "sweep 2 runs backwards and keeps the minimum: done"),
]
render(steps, os.path.join(os.path.dirname(os.path.abspath(__file__)), "images"))
