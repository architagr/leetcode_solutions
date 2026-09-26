import os, sys
sys.path.insert(0, "/Users/architagarwal/code/codestreak-growth/tools")
from diagram import *

# Five components and the config's links in file order. Line 5 is the extra
# one that closes a loop; the first four alone are a valid tree.
W_ = 130
POS = {"A": (250, 100), "B": (70, 220), "D": (430, 220), "C": (70, 340), "E": (430, 340)}
LINES = [("A", "B"), ("B", "C"), ("A", "D"), ("D", "E"), ("C", "E")]
Pn, C, D, R = "pend", "cur", "done", "repeat"
PX = 640


def link(a, b, color=EDGE):
    ax, ay = POS[a]
    bx, by = POS[b]
    if ay == by:
        return connect(ax + W_, ay + BH / 2, bx, by + BH / 2, color=color, width=4)
    return connect(ax + W_ / 2, ay + BH, bx + W_ / 2, by, color=color, width=4)


def cfg(n_lines, states=None, panel=(), title=None, notes=(), hot=None):
    states = states or {}
    out = []
    for i, (a, b) in enumerate(LINES[:n_lines]):
        out.append(link(a, b, REPEAT if hot == i else EDGE))
    for n, (x, y) in POS.items():
        out += box(x, y, n, states.get(n, Pn), w=W_)
    if title:
        out.append(label(PX, 76, title, 22, MUTED, anchor="start"))
    for i, (text, state) in enumerate(panel):
        out += box(PX, 96 + i * 68, text, state, w=520)
    for y, text, color in notes:
        out.append(label(600, y, text, 26, color))
    return out


ALL = {n: D for n in POS}

steps = [

 (cfg(4, ALL, title="the config file",
      panel=[("line 1: A - B", D), ("line 2: B - C", D), ("line 3: A - D", D),
             ("line 4: D - E", D)],
      notes=[(480, "valid only if the links form one tree:", MUTED),
             (520, "everything connected, and no loops", STROKE)]),
  "a config is valid only if its links form one tree"),

 (cfg(2, {"A": C, "B": D, "C": D, "D": R},
      title="line 3: is A already connected to D?",
      panel=[("PathExists over lines 1-2", D), ("builds the adjacency list", D),
             ("walks A, B, C: no D. Accept.", D)],
      notes=[(480, "the tooling's tested search, one call per line", MUTED),
             (520, "a yes means this line closes a loop", STROKE)]),
  "what you would write: before each line, search for a path"),

 (cfg(4, {n: R for n in POS}, title="the cost",
      panel=[("each line: rebuild + search", R), ("then a search per component", R),
             ("5,000 lines: 9,998 searches", R), ("2.74 GB allocated, 1.95 s", R)],
      notes=[(520, "1,000 lines took 68.2 ms. 5x the lines, 28.7x the time.", REPEAT)]),
  "every line reruns a whole search over everything read so far"),

 (cfg(2, {"A": D, "B": D, "C": D, "D": Pn, "E": Pn}, title="groups after line 2",
      panel=[("{A, B, C}", D), ("{D}", Pn), ("{E}", Pn),
             ("groups only ever merge", C)],
      notes=[(480, "line 3's search rediscovers {A, B, C}", MUTED),
             (520, "which line 2's search had already found", STROKE)]),
  "the searches keep rediscovering groups that only ever merge"),

 (cfg(5, ALL, hot=4, title="5 components, 5 links",
      panel=[("a tree on n needs n - 1 links", D), ("5 is not 4: not a tree", R),
             ("decided in 2.42 ns", D), ("but which line is wrong?", C)],
      notes=[(520, "connected and n - 1 links means no room for a loop", STROKE)]),
  "count first: a tree on n components has exactly n - 1 links"),

 (cfg(5, {"A": D, "B": D, "C": C, "D": D, "E": C}, hot=4,
      title="line 5: C - E",
      panel=[("every group has one root", D), ("root of C: A", D), ("root of E: A", D),
             ("same root: line 5 closes a loop", R)],
      notes=[(520, "42.2 µs for 5,000 lines, one allocation", STROKE)]),
  "union-find: one root per group, and the line that loops is named"),
]
render(steps, os.path.join(os.path.dirname(os.path.abspath(__file__)), "images"))
