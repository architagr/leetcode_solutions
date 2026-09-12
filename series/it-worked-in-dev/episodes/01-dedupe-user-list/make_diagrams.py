import os, sys
sys.path.insert(0, "/Users/architagarwal/code/codestreak-growth/tools")
from diagram import *

# Two columns: incoming on the left, existing on the right. The diagrams show
# how many times the right column gets read, which is the whole point.
INX, EXX = 150, 640
ROW = 78
TOP = 150

IN = ["a@x", "b@x", "c@x"]
EX = ["a@x", "q@x", "r@x", "s@x"]

def cols(in_states, ex_states, note=None, arrows=()):
    out = []
    out.append(label(INX + BW/2, TOP - 46, "incoming", 26, MUTED, bold=True))
    out.append(label(EXX + BW/2, TOP - 46, "existing", 26, MUTED, bold=True))
    for i, e in enumerate(IN):
        out += box(INX, TOP + i*ROW, e, in_states.get(i, "pend"))
    for i, e in enumerate(EX):
        out += box(EXX, TOP + i*ROW, e, ex_states.get(i, "pend"))
    for (a, b, colr) in arrows:
        out.append(connect(INX + BW, TOP + a*ROW + BH/2,
                           EXX, TOP + b*ROW + BH/2, colr, None, 3))
    if note:
        out.append(label(600, TOP + 4*ROW + 46, note, 25, MUTED))
    return out

P, C, D, R = "pend", "cur", "done", "repeat"

steps = [
 (cols({}, {}),
  "for each incoming email, is it already in existing?"),

 (cols({0: C}, {0: C}, "found at the first try, but it still had to look",
       arrows=[(0, 0, CUR)]),
  "a@x scans existing from the top and matches"),

 (cols({0: D, 1: C}, {0: R, 1: R, 2: R, 3: R},
       "no match, so every one of them was read",
       arrows=[(1, 0, REPEAT), (1, 1, REPEAT), (1, 2, REPEAT), (1, 3, REPEAT)]),
  "b@x scans the whole list again. Same list, unchanged."),

 (cols({0: D, 1: D, 2: C}, {0: R, 1: R, 2: R, 3: R},
       "and again. existing never changed between any of these",
       arrows=[(2, 0, REPEAT), (2, 1, REPEAT), (2, 2, REPEAT), (2, 3, REPEAT)]),
  "c@x scans it a third time"),

 (cols({0: D, 1: D, 2: D}, {0: R, 1: R, 2: R, 3: R},
       "3 incoming x 4 existing = 12 comparisons, for 3 answers"),
  "every question was answered by re-reading the same list"),

 (cols({}, {0: C, 1: C, 2: C, 3: C},
       "one pass, before the loop starts"),
  "instead: read existing once and remember what is in it"),

 (cols({0: D, 1: D, 2: D}, {0: D, 1: D, 2: D, 3: D},
       "4 to build the set, then 3 lookups. 7 instead of 12."),
  "each incoming email is now one lookup, not a scan"),
]
render(steps, os.path.join(os.path.dirname(os.path.abspath(__file__)), "images"))
