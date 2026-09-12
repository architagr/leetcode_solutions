import os, sys
sys.path.insert(0, "/Users/architagarwal/code/codestreak-growth/tools")
from diagram import *

# A five-deep reporting line, drawn as a column. The whole argument is about
# how many times the same upward path gets walked.
X = 505
TOP = 100
ROW = 80
NAMES = ["CEO", "VP", "Dir", "Mgr", "Eng"]

def chain(states, subs=None, arrows=(), note=None):
    subs = subs or {}
    out = []
    for i in range(len(NAMES) - 1):
        out.append(connect(X + BW/2, TOP + i*ROW + BH, X + BW/2, TOP + (i+1)*ROW, width=3))
    for i, n in enumerate(NAMES):
        out += box(X, TOP + i*ROW, n, states.get(i, "pend"), subs.get(i))
    for (frm, to, colr) in arrows:
        # the upward walk, drawn to the left of the column
        out.append(connect(X - 14, TOP + frm*ROW + BH/2, X - 14, TOP + to*ROW + BH/2, colr, "up", 4))
    if note:
        out.append(label(X + BW/2, TOP + 4*ROW + BH + 42, note, 25, MUTED))
    return out

P, C, D, R = "pend", "cur", "done", "repeat"

steps = [
 (chain({}, {i: "level ?" for i in range(5)}),
  "every row needs its level before the page can render"),

 (chain({4: C, 0: R, 1: R, 2: R, 3: R}, {4: "4 hops"},
        arrows=[(4, 0, REPEAT)],
        note="Eng climbs all the way to the top"),
  "walking up from Eng: 4 hops to reach the CEO"),

 (chain({4: D, 3: C, 0: R, 1: R, 2: R}, {4: "4", 3: "3 hops"},
        arrows=[(3, 0, REPEAT)],
        note="over the exact path Eng just walked"),
  "then Mgr climbs the same path again, minus one"),

 (chain({4: D, 3: D, 2: C, 0: R, 1: R}, {4: "4", 3: "3", 2: "2 hops"},
        arrows=[(2, 0, REPEAT)],
        note="4 + 3 + 2 + 1 = 10 hops for 5 answers"),
  "and Dir again. The same journey, over and over."),

 (chain({0: D, 1: D, 2: D, 3: D, 4: D}, {0: "0", 1: "1", 2: "2", 3: "3", 4: "4"},
        note="each answer is its manager's answer plus one"),
  "but Mgr's level was known one hop before Eng needed it"),

 (chain({0: C}, {0: "level 0"},
        note="start at the top and hand the number down"),
  "instead: one pass down, labelling a whole row at a time"),

 (chain({0: D, 1: D, 2: D, 3: D, 4: D}, {0: "0", 1: "1", 2: "2", 3: "3", 4: "4"},
        note="4 steps instead of 10. Nobody walked upward."),
  "same answers, one pass, no journey repeated"),
]
render(steps, os.path.join(os.path.dirname(os.path.abspath(__file__)), "images"))
