import os, sys
sys.path.insert(0, "/Users/architagarwal/code/codestreak-growth/tools")
from diagram import *

# docs / 2026 / q3 / {report.pdf, notes.md}
ROOT = (505, 110)
Y1   = (505, 205)
Q3   = (505, 300)
F1   = (300, 395)
F2   = (710, 395)
EDGES = [(Y1, ROOT), (Q3, Y1), (F1, Q3), (F2, Q3)]
BOXES = [(ROOT, "docs"), (Y1, "2026"), (Q3, "q3"), (F1, "report"), (F2, "notes")]

def tree(states, subs=None, note=None, up=None):
    subs = subs or {}
    out = []
    for (x, y), (px, py) in EDGES:
        out.append(connect(px + BW/2, py + BH, x + BW/2, y, width=3))
    for i, ((x, y), name) in enumerate(BOXES):
        out += box(x, y, name, states.get(i, "pend"), subs.get(i))
    if up:
        # the upward walk, drawn as one arrow from the file to the root
        fx, fy = BOXES[up][0]
        out.append(connect(fx + BW/2, fy, ROOT[0] + BW/2, ROOT[1] + BH, REPEAT, "up", 5))
    if note:
        out.append(label(600, 505, note, 26, MUTED))
    return out

P, C, D, R = "pend", "cur", "done", "repeat"

steps = [
 (tree({}, note="every row on the page shows its full path"),
  "docs / 2026 / q3 / report.pdf - for every node"),

 (tree({3: C, 0: R, 1: R, 2: R}, up=3,
       note="collected backwards, then reversed"),
  "report walks up: q3, 2026, docs - three segments"),

 (tree({3: D, 4: C, 0: R, 1: R, 2: R}, {3: "done"}, up=4,
       note="the identical prefix, rebuilt from scratch"),
  "notes walks up through q3, 2026, docs again"),

 (tree({0: R, 1: R, 2: R, 3: D, 4: D}, {3: "done", 4: "done"},
       note="and every file in every sibling folder does the same"),
  "one prefix, rebuilt once per file underneath it"),

 (tree({0: D}, {0: "docs"},
       note="a complete path, before any child needs it"),
  "but the parent's path was finished first"),

 (tree({0: D, 1: D, 2: C}, {0: "docs", 1: "docs/2026", 2: "+ /q3"},
       note="each node adds one segment to what it was handed"),
  "carry it down instead: parent's path plus one name"),

 (tree({0: D, 1: D, 2: D, 3: D, 4: D},
       {2: "docs/2026/q3", 3: "+ /report", 4: "+ /notes"},
       note="no collecting, no reversing, no join"),
  "same paths, one pass, nothing rebuilt"),
]
render(steps, os.path.join(os.path.dirname(os.path.abspath(__file__)), "images"))
