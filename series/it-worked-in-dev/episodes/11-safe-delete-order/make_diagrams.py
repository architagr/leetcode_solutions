import os, sys
sys.path.insert(0, "/Users/architagarwal/code/codestreak-growth/tools")
from diagram import *

# workspace -> project-a -> {board-1 -> card-9, board-3}, workspace -> board-7
WS = (505, 75)
PA, B7 = (250, 185), (760, 185)
B1, B3 = (80, 295), (420, 295)
C9 = (80, 405)

BOXES = [(WS, "workspace"), (PA, "project-a"), (B7, "board-7"),
         (B1, "board-1"), (B3, "board-3"), (C9, "card-9")]
EDGES = [(PA, WS), (B7, WS), (B1, PA), (B3, PA), (C9, B1)]

WS_i, PA_i, B7_i, B1_i, B3_i, C9_i = range(6)
Pn, C, D, R = "pend", "cur", "done", "repeat"


def rows(states, subs=None, notes=None):
    subs = subs or {}
    res = []
    for (x, y), (px, py) in EDGES:
        res.append(connect(px + BW / 2, py + BH, x + BW / 2, y, width=3))
    for i, ((x, y), name) in enumerate(BOXES):
        res += box(x, y, name, states.get(i, "pend"), subs.get(i))
    for (x, y, t, size, color) in (notes or []):
        res.append(label(x, y, t, size, color))
    return res


steps = [
 (rows({}, notes=[(950, 330, "every child", 27, STROKE),
                  (950, 372, "points at its", 27, STROKE),
                  (950, 414, "parent. What", 27, STROKE),
                  (950, 456, "can go first?", 27, STROKE)]),
  "a parent cannot be deleted while a child still references it"),

 (rows({C9_i: C, B3_i: C, B7_i: C},
       {C9_i: "delete", B3_i: "delete", B7_i: "delete"},
       notes=[(950, 350, "wave 1: three", 26, MUTED),
              (950, 392, "rows, one call", 26, MUTED)]),
  "whatever has nothing pointing at it can go now"),

 (rows({C9_i: D, B3_i: D, B7_i: D, B1_i: C}, {B1_i: "delete"},
       notes=[(950, 350, "wave 2, after", 26, MUTED),
              (950, 392, "a second walk", 26, MUTED)]),
  "then look again. board-1 is empty now, so it can go"),

 (rows({C9_i: D, B3_i: D, B7_i: D, B1_i: D, PA_i: R, WS_i: R},
       notes=[(950, 308, "one full walk", 27, STROKE),
              (950, 350, "per wave", 27, STROKE),
              (950, 400, "2,000 waves", 27, REPEAT),
              (950, 442, "on a thread", 27, REPEAT)]),
  "four waves here. One per nesting level, every time."),

 (rows({C9_i: C, B1_i: R, B3_i: R, B7_i: R},
       {C9_i: "wave 1", B3_i: "waits", B7_i: "waits"},
       notes=[(950, 350, "grouping by", 26, REPEAT),
              (950, 392, "depth, not height", 26, REPEAT)]),
  "by distance from the root, only card-9 goes first. Wrong."),

 (rows({i: D for i in range(6)},
       {C9_i: "0", B1_i: "1", B3_i: "0", B7_i: "0", PA_i: "2", WS_i: "3"},
       notes=[(950, 308, "a row's wave is", 25, MUTED),
              (950, 346, "its height: one", 25, MUTED),
              (950, 384, "past the slowest", 25, MUTED),
              (950, 422, "child it has", 25, MUTED)]),
  "the wave number is height, not depth. board-3 is height 0."),

 (rows({i: D for i in range(6)},
       {C9_i: "wave 1", B1_i: "wave 2", B3_i: "wave 1", B7_i: "wave 1",
        PA_i: "wave 3", WS_i: "wave 4"},
       notes=[(950, 350, "one walk,", 26, MUTED),
              (950, 392, "every wave", 26, MUTED)]),
  "one pass, bottom up: each row filed as its children report"),
]
render(steps, os.path.join(os.path.dirname(os.path.abspath(__file__)), "images"))
