import os, sys
sys.path.insert(0, "/Users/architagarwal/code/codestreak-growth/tools")
from diagram import *

# home -> {docs (deep), contact (dead end)}
HOME = (505, 90)
DOCS = (250, 180)
CONT = (760, 180)
D1 = (250, 270)
D2 = (250, 360)
EDGES = [(DOCS, HOME), (CONT, HOME), (D1, DOCS), (D2, D1)]
BOXES = [(HOME, "home"), (DOCS, "/docs"), (CONT, "/contact"),
         (D1, "/docs/a"), (D2, "/docs/b")]

def site(states, subs=None, note=None):
    subs = subs or {}
    out = []
    for (x, y), (px, py) in EDGES:
        out.append(connect(px + BW/2, py + BH, x + BW/2, y, width=3))
    for i, ((x, y), name) in enumerate(BOXES):
        out += box(x, y, name, states.get(i, "pend"), subs.get(i))
    if note:
        out.append(label(600, 495, note, 26, MUTED))
    return out

P, C, D, R = "pend", "cur", "done", "repeat"

steps = [
 (site({}, note="how many clicks to a page with nothing below it?"),
  "/contact is a dead end. So is the bottom of /docs."),

 (site({0: D, 1: C}, {0: "1 click"}, note="it must finish this branch before it can compare"),
  "the recursion takes the first child and commits"),

 (site({0: D, 1: R, 3: R, 4: C}, note="500 pages in a real docs section, not 3"),
  "down through /docs to the very bottom"),

 (site({0: D, 1: D, 3: D, 4: D, 2: C}, {1: "3", 4: "1"},
       note="the answer was here the whole time"),
  "only now does it look at /contact, and finds 1"),

 (site({0: D, 1: R, 2: R, 3: R, 4: R}, note="502 pages read to answer '2'"),
  "every page was read. The answer never needed them."),

 (site({0: C}, {0: "level 1"}, note="all the one-click pages, before any two-click page"),
  "instead: walk the site a level at a time"),

 (site({0: D, 1: D, 2: D}, {2: "dead end"},
       note="3 pages read. /docs was never opened."),
  "level 2 holds /contact, which ends it"),
]
render(steps, os.path.join(os.path.dirname(os.path.abspath(__file__)), "images"))
