import os, sys
sys.path.insert(0, "/Users/architagarwal/code/codestreak-growth/tools")
from diagram import *

# Seven cache nodes in a line, the invalidation starting at both ends. Small
# enough that every per-node number fits under its node.
NODES = ["A", "B", "C", "D", "E", "F", "G"]
X = [30 + i * 165 for i in range(7)]
BW_ = 130
Y0 = 120
FROM_A = [0, 1, 2, 3, 4, 5, 6]
FROM_G = [6, 5, 4, 3, 2, 1, 0]
BEST = [0, 1, 2, 3, 2, 1, 0]
Pn, C, D, R = "pend", "cur", "done", "repeat"


def row(y, name, values, colors):
    out = [label(X[0] - 8, y, name, 22, MUTED, anchor="start")] if name else []
    for i, v in enumerate(values):
        if v is None:
            continue
        out.append(label(X[i] + BW_ / 2, y + 36, str(v), 30, colors[i], bold=True))
    return out


def line(states, rows=(), notes=(), origins=("A", "G")):
    out = [connect(X[i] + BW_, Y0 + BH / 2, X[i + 1], Y0 + BH / 2, width=4)
           for i in range(6)]
    for i, n in enumerate(NODES):
        out += box(X[i], Y0, n, states.get(n, Pn), w=BW_)
    for n in origins:
        i = NODES.index(n)
        out.append(label(X[i] + BW_ / 2, Y0 - 16, "origin", 20, CUR))
    y = 230
    for name, values, colors in rows:
        out += row(y, name, values, colors)
        y += 78
    for yy, text, color in notes:
        out.append(label(600, yy, text, 26, color))
    return out


ALL = {n: D for n in NODES}
BLUE = [DONE] * 7
keep_a = [DONE if a <= g else REPEAT for a, g in zip(FROM_A, FROM_G)]
keep_g = [DONE if g < a else REPEAT for a, g in zip(FROM_A, FROM_G)]

steps = [

 (line(ALL, rows=[("rounds until invalidated", BEST, BLUE)],
       notes=[(470, "one hop per round, from whichever origin is nearer", MUTED),
              (510, "every node has it after 3 rounds", STROKE)]),
  "the invalidation starts at A and G and spreads one hop a round"),

 (line({n: (C if n == "A" else D) for n in NODES}, origins=("A",),
       rows=[("HopsFrom(A)", FROM_A, BLUE)],
       notes=[(470, "the dashboard's function: rounds from one origin", MUTED),
              (510, "run it for each origin, keep the smallest per node", STROKE)]),
  "what you would write: one search per origin, then the minimum"),

 (line(ALL, rows=[("HopsFrom(A)", FROM_A, keep_a), ("HopsFrom(G)", FROM_G, keep_g),
                  ("minimum", BEST, BLUE)],
       notes=[(510, "14 numbers computed, 7 kept. 1,000 origins: 10,000,000", REPEAT)]),
  "every search covers the whole cluster, and most of it is discarded"),

 (line({"A": D, "G": D, "B": C, "F": C},
       rows=[("round 1", [None, 1, None, None, None, 1, None], [CUR] * 7)],
       notes=[(470, "the smaller number is the wave that arrived first", MUTED),
              (510, "and a search already visits nodes in arrival order", STROKE)]),
  "the minimum is just the first wave to arrive"),

 (line({"A": D, "G": D, "B": D, "F": D, "C": C, "E": C},
       rows=[("queue", ["A", "G", "B", "F", "C", "E", "D"], [DONE] * 4 + [CUR] * 2 + [MUTED]),
             ("round", [0, 0, 1, 1, 2, 2, 3], [DONE] * 4 + [CUR] * 2 + [MUTED])],
       notes=[(510, "both origins go in at round 0, before anything expands", STROKE)]),
  "one queue, seeded with every origin: the waves merge by themselves"),

 (line({n: (C if n == "D" else D) for n in NODES},
       rows=[("rounds", BEST, BLUE)],
       notes=[(470, "FIFO hands rounds out in order: 0 0 1 1 2 2 3", MUTED),
              (510, "the last one assigned is the answer. No comparison.", STROKE)]),
  "the last node the queue hands out holds the answer: 3 rounds"),
]
render(steps, os.path.join(os.path.dirname(os.path.abspath(__file__)), "images"))
