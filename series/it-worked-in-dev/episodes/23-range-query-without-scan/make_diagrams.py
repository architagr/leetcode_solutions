import os, sys
sys.path.insert(0, "/Users/architagarwal/code/codestreak-growth/tools")
from diagram import *

# A 15-product price index, prices in rupees, and the query [45, 60].
BWS = 72
Y = [86, 186, 286, 386]
NODES = {50: (564, 0), 30: (284, 1), 70: (844, 1), 20: (144, 2), 40: (424, 2),
         60: (704, 2), 80: (984, 2), 15: (74, 3), 25: (214, 3), 35: (354, 3),
         45: (494, 3), 55: (634, 3), 65: (774, 3), 75: (914, 3), 85: (1054, 3)}
KIDS = {50: [30, 70], 30: [20, 40], 70: [60, 80], 20: [15, 25], 40: [35, 45],
        60: [55, 65], 80: [75, 85]}
LO, HI = 45, 60
IN = [p for p in NODES if LO <= p <= HI]
Pn, C, D, R = "pend", "cur", "done", "repeat"


def tree(states, notes=()):
    out = []
    for p, ks in KIDS.items():
        px, pl = NODES[p]
        for k in ks:
            kx, kl = NODES[k]
            out.append(connect(px + BWS / 2, Y[pl] + BH, kx + BWS / 2, Y[kl], width=3))
    for p, (x, lvl) in NODES.items():
        out += box(x, Y[lvl], str(p), states.get(p, Pn), w=BWS)
    for y, text, color in notes:
        out.append(label(600, y, text, 26, color))
    return out


def subtree(p):
    out = [p]
    for k in KIDS.get(p, []):
        out += subtree(k)
    return out


PRUNED = [15, 20, 25, 35, 80, 75, 85, 65]  # never visited with pruning
VISITED = [p for p in NODES if p not in PRUNED]

steps = [

 (tree({p: D for p in IN}, notes=[
     (500, "products priced 45 to 60: four of fifteen", MUTED),
     (540, "left is cheaper, right is dearer, at every node", STROKE)]),
  "a price index: every left subtree cheaper, every right dearer"),

 (tree({p: (D if p in IN else R) for p in NODES}, notes=[
     (500, "walk everything in order, keep what is in range", MUTED),
     (540, "15 read for 4. At 200,000 products: 200,000 for 20.", REPEAT)]),
  "what you would write: read every product, keep the ones in range"),

 (tree({30: C, 20: R, 15: R, 25: R}, notes=[
     (500, "30 is below 45, and everything left of 30 is cheaper still", MUTED),
     (540, "the whole left side of 30 was ruled out at 30", STROKE)]),
  "at 30, one comparison already says nothing on its left can match"),

 (tree({70: C, 80: R, 75: R, 85: R, 65: R}, notes=[
     (500, "70 is above 60: everything right of 70 is dearer", MUTED),
     (540, "and at 60, the right side is out too", STROKE)]),
  "the mirror image: above the range, the right side cannot match"),

 (tree({**{p: D for p in VISITED if p not in IN}, **{p: C for p in IN}}, notes=[
     (500, "visit a side only if it can hold something in range", MUTED),
     (540, "7 visited here. 43 of 200,000 for 20 results: 6517x.", STROKE)]),
  "prune both sides: the walk follows only the paths into the range"),
]
render(steps, os.path.join(os.path.dirname(os.path.abspath(__file__)), "images"))
