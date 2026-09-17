import os, sys
sys.path.insert(0, "/Users/architagarwal/code/codestreak-growth/tools")
from diagram import *

# fetch -> validate -> enrich -> notify -> retry, and retry -> validate
Y = 190
XS = [20, 250, 480, 710, 940]
NAMES = ["fetch", "validate", "enrich", "notify", "retry"]
CX = [x + BW / 2 for x in XS]
BACK_Y = 372

Pn, C, D, R = "pend", "cur", "done", "repeat"


def pipeline(states, marks=None, notes=None, back=True, backcolor=EDGE):
    res = []
    for i in range(len(XS) - 1):
        res.append(connect(XS[i] + BW, Y + BH / 2, XS[i + 1], Y + BH / 2,
                           marker="a", width=3))
    if back:
        # retry -> validate, routed under the row
        res.append(connect(CX[4], Y + BH, CX[4], BACK_Y, color=backcolor, width=3))
        res.append(connect(CX[4], BACK_Y, CX[1], BACK_Y, color=backcolor, width=3))
        res.append(connect(CX[1], BACK_Y, CX[1], Y + BH, color=backcolor, width=3,
                           marker="up" if backcolor == DONE else "a"))
    for i, x in enumerate(XS):
        res += box(x, Y, NAMES[i], states.get(i, "pend"))
    for i, text in (marks or {}).items():
        res.append(label(CX[i], Y - 22, text, 25, STROKE, bold=True))
    for (x, y, t, size, color) in (notes or []):
        res.append(label(x, y, t, size, color))
    return res


steps = [
 (pipeline({}, notes=[(600, 462, "one \"next\" pointing backwards, and the runner never stops", 26, MUTED)]),
  "a pipeline is a chain: each stage names the one that follows"),

 (pipeline({0: D, 1: D, 2: C}, marks={2: "here"},
           notes=[(600, 462, "seen: fetch, validate", 26, MUTED)]),
  "the obvious check: remember every stage you have walked past"),

 (pipeline({0: D, 1: R, 2: D, 3: D, 4: D}, marks={1: "again"},
           notes=[(600, 448, "arrived somewhere already visited, so it loops", 26, MUTED),
                  (600, 492, "100,000 stages: 4.73 MB of map, to return true", 26, REPEAT)]),
  "been here before. Correct, and it pays one map entry per stage."),

 (pipeline({i: Pn for i in range(5)}, back=False,
           notes=[(600, 448, "the version most HTTP clients ship: stop after N hops", 26, MUTED),
                  (600, 492, "a 150-stage pipeline with no loop is reported as one", 26, REPEAT)]),
  "giving up after 100 hops is not the same question"),

 (pipeline({1: C, 2: C}, marks={1: "slow", 2: "fast"},
           notes=[(600, 462, "one stage at a time, and two at a time", 26, MUTED)]),
  "instead: two cursors, moving at different speeds"),

 (pipeline({3: C, 2: C}, marks={3: "slow", 2: "fast"},
           notes=[(600, 448, "turn 3: fast came round the back edge, behind slow", 26, MUTED),
                  (600, 492, "inside a loop it closes the gap by one stage a turn", 26, MUTED)]),
  "\"ahead\" stops meaning anything once the chain wraps"),

 (pipeline({4: D}, marks={4: "both"}, backcolor=DONE,
           notes=[(600, 448, "they land on the same stage, never step over it", 26, MUTED),
                  (600, 492, "two variables. No map, no allocation, no limit.", 26, STROKE)]),
  "the cursors meet, so the pipeline loops"),
]
render(steps, os.path.join(os.path.dirname(os.path.abspath(__file__)), "images"))
