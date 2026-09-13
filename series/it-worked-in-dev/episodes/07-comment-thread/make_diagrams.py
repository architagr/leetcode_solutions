import os, sys
sys.path.insert(0, "/Users/architagarwal/code/codestreak-growth/tools")
from diagram import *

# a thread: 1 -> {2 -> 3, 4}
C1 = (505, 95)
C2 = (270, 185)
C3 = (270, 275)
C4 = (760, 185)
EDGES = [(C2, C1), (C3, C2), (C4, C1)]
BOXES = [(C1, "#1"), (C2, "#2"), (C3, "#3"), (C4, "#4")]

def thread(states, subs=None, note=None, order=None):
    subs = subs or {}
    out = []
    for (x, y), (px, py) in EDGES:
        out.append(connect(px + BW/2, py + BH, x + BW/2, y, width=3))
    for i, ((x, y), name) in enumerate(BOXES):
        out += box(x, y, name, states.get(i, "pend"), subs.get(i))
    if order:
        out.append(label(600, 400, order, 28, STROKE, bold=True))
    if note:
        out.append(label(600, 470, note, 26, MUTED))
    return out

P, C, D, R = "pend", "cur", "done", "repeat"

steps = [
 (thread({}, note="parent, then its replies, then the next sibling"),
  "the page renders #1, #2, #3, #4 - in that order"),

 (thread({0: C, 1: C, 2: C, 3: C},
         {0: "00", 1: "00.00", 2: "00.00.00", 3: "00.01"},
         note="one segment per level, zero-padded so 2 sorts before 10"),
  "the sort-key version builds a path for every comment"),

 (thread({0: R, 1: R, 2: R, 3: R},
         note="a comment 60 deep carries a 60-segment string"),
  "building them is quadratic: the keys grow with depth"),

 (thread({1: R, 3: R}, note="two siblings share every segment but the last"),
  "then each comparison walks the shared prefix to decide"),

 (thread({0: D, 1: D, 2: D, 3: D}, order="#1  #2  #3  #4",
         note="the order the tree already had"),
  "sorting reconstructs what the structure encoded"),

 (thread({0: C}, order="#1", note="record the comment, then walk its replies"),
  "instead: emit in the order the thread is already in"),

 (thread({0: D, 1: D, 2: D, 3: D}, order="#1  #2  #3  #4",
         note="no keys, no comparator, no sort"),
  "same order, one walk, nothing compared"),
]
render(steps, os.path.join(os.path.dirname(os.path.abspath(__file__)), "images"))
