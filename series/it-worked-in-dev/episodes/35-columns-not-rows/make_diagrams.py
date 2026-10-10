import os, sys
sys.path.insert(0, "/Users/architagarwal/code/codestreak-growth/tools")
from diagram import *

# A depth-3 decision tree, ids in level order. Columns: -2 [3], -1 [1],
# 0 [0, 4, 5], 1 [2], 2 [6]. 4 and 5 share column 0 and row 2: 4 is left.
Pn, C, D, R = "pend", "cur", "done", "repeat"
COL = {0: 0, 1: -1, 2: 1, 3: -2, 4: 0, 5: 0, 6: 2}
DEP = {0: 0, 1: 1, 2: 1, 3: 2, 4: 2, 5: 2, 6: 2}
KIDS = {0: [1, 2], 1: [3, 4], 2: [5, 6]}
NW = 80


def pos(i):
    # 4 and 5 share a column; nudge them apart so both are visible
    nudge = {4: -48, 5: 48}.get(i, 0)
    return 560 + COL[i] * 170 + nudge, 110 + DEP[i] * 105


def tree(states, notes=(), guides=True, subs=None):
    subs = subs or {}
    out = []
    if guides:
        for c in range(-2, 3):
            x = 600 + c * 170
            out.append(f'<line x1="{x}" y1="92" x2="{x}" y2="400" stroke="#dfe5ea" '
                       f'stroke-width="2" stroke-dasharray="6 6"/>')
            out.append(label(x, 425, "col %d" % c, 20, MUTED))
    for p, ks in KIDS.items():
        px, py = pos(p)
        for k in ks:
            kx, ky = pos(k)
            out.append(connect(px + NW / 2, py + BH, kx + NW / 2, ky, width=3))
    for i in COL:
        x, y = pos(i)
        out += box(x, y, str(i), states.get(i, Pn), w=NW)
    for y, text, color in notes:
        out.append(label(600, y, text, 26, color))
    return out


ALL = {i: D for i in COL}
steps = [

 (tree(ALL, notes=[
     (480, "columns left to right; top down; a shared row left to right", MUTED),
     (520, "[3]  [1]  [0 4 5]  [2]  [6]", STROKE)]),
  "draw the tree in columns: a node's column is its x position"),

 (tree(ALL, notes=[
     (480, "walk it, note (column, depth): 0 1 3 4 2 5 6 in walk order", MUTED),
     (520, "then sort by column, then depth, and cut into columns", STROKE)]),
  "what you would write: note column and depth, then sort"),

 (tree({**ALL, 4: R, 5: R}, notes=[
     (480, "4 and 5: same column, same depth. The sort calls them equal.", MUTED),
     (520, "sort.Slice may swap them. On 12 nodes or fewer it never does.", REPEAT)]),
  "sort.Slice: right on every tree small enough to check by eye"),

 (tree({**ALL, 4: C, 5: C}, notes=[
     (480, "sort.SliceStable keeps the walk's order for ties: correct", MUTED),
     (520, "1M nodes: 48 swaps per node, 580 ms. 8.11x the wrong one.", REPEAT)]),
  "the stable sort is right, and 8.11x slower"),

 (tree({0: D, 1: D, 2: D, 3: C, 4: C, 5: C, 6: C}, notes=[
     (480, "the sort key is (row, left to right) inside each column", MUTED),
     (520, "reading the tree row by row, left to right, IS that order", STROKE)]),
  "the order the sort rebuilds is the order rows are read in"),

 (tree(ALL, subs={}, notes=[
     (480, "row by row with a queue; append each node to its column", MUTED),
     (520, "every column comes out sorted. No sort: 30.7 ms, 18.9x.", STROKE)]),
  "walk the rows, append on arrival: the columns are already sorted"),
]
render(steps, os.path.join(os.path.dirname(os.path.abspath(__file__)), "images"))
