import os, sys
sys.path.insert(0, "/Users/architagarwal/code/codestreak-growth/tools")
from diagram import *

# shop -> {books, electronics}, with electronics running deep.
SHOP = (505, 100)
BOOKS = (250, 190)
ELEC = (760, 190)
E1 = (760, 280)
E2 = (760, 370)
EDGES = [(BOOKS, SHOP), (ELEC, SHOP), (E1, ELEC), (E2, E1)]
BOXES = [(SHOP, "shop"), (BOOKS, "books"), (ELEC, "electr"), (E1, "phones"), (E2, "cases")]

def tree(states, subs=None, note=None):
    subs = subs or {}
    out = []
    for (x, y), (px, py) in EDGES:
        out.append(connect(px + BW/2, py + BH, x + BW/2, y, width=3))
    for i, ((x, y), name) in enumerate(BOXES):
        out += box(x, y, name, states.get(i, "pend"), subs.get(i))
    if note:
        out.append(label(600, 505, note, 26, MUTED))
    return out

P, C, D, R = "pend", "cur", "done", "repeat"

steps = [
 (tree({}, note="which categories have one branch far deeper than the rest?"),
  "electronics runs three deep; books is one"),

 (tree({0: C, 1: R, 2: R, 3: R, 4: R},
       note="depth() walks the whole tree to judge one node"),
  "to judge shop, measure both children: 1 and 3"),

 (tree({0: D, 2: C, 3: R, 4: R}, {0: "lopsided"},
       note="the same subtree, walked a second time"),
  "then descend into electronics and measure again"),

 (tree({0: D, 2: D, 3: C, 4: R}, {0: "lopsided"},
       note="a chain of n costs n(n+1)/2 calls into depth()"),
  "and into phones, measuring cases a third time"),

 (tree({0: D, 1: D, 2: D, 3: D, 4: D}, {0: "lopsided"},
       note="every depth was computed, compared, and thrown away"),
  "each node re-measured everything beneath it"),

 (tree({4: C}, {4: "returns 1"},
       note="the depth a parent needs is what the child just returned"),
  "instead: let the recursion hand back the depth it found"),

 (tree({0: D, 1: D, 2: D, 3: D, 4: D},
       {0: "lopsided", 2: "3", 3: "2", 4: "1", 1: "1"},
       note="one pass. Depth and verdict come back together."),
  "same answer, nothing measured twice"),
]
render(steps, os.path.join(os.path.dirname(os.path.abspath(__file__)), "images"))
