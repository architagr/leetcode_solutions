import os, sys
sys.path.insert(0, "/Users/architagarwal/code/codestreak-growth/tools")
from diagram import *

# region -> zone-a -> {rack-1 -> api-7 -> pod-2, rack-2 -> web-3 -> pod-9}
# region -> zone-b
RG = (505, 75)
ZA, ZB = (250, 175), (760, 175)
R1, R2 = (80, 275), (420, 275)
A7, W3 = (80, 375), (420, 375)
P2, P9 = (80, 470), (420, 470)

BOXES = [(RG, "region"), (ZA, "zone-a"), (ZB, "zone-b"), (R1, "rack-1"),
         (R2, "rack-2"), (A7, "api-7"), (W3, "web-3"), (P2, "pod-2"), (P9, "pod-9")]
EDGES = [(ZA, RG), (ZB, RG), (R1, ZA), (R2, ZA), (A7, R1), (W3, R2), (P2, A7), (P9, W3)]

RG_i, ZA_i, ZB_i, R1_i, R2_i, A7_i, W3_i, P2_i, P9_i = range(9)
Pn, C, D, R = "pend", "cur", "done", "repeat"


def topo(states, subs=None, notes=None):
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
 (topo({}, notes=[(950, 330, "which two", 27, STROKE),
                  (950, 372, "are furthest", 27, STROKE),
                  (950, 414, "apart?", 27, STROKE)]),
  "traffic between two instances climbs to a shared level and back"),

 (topo({P2_i: C, A7_i: D, R1_i: D, ZA_i: D, R2_i: D, W3_i: D, P9_i: D, RG_i: D, ZB_i: D},
       {P2_i: "start", P9_i: "6 hops"},
       notes=[(950, 350, "furthest from", 26, MUTED),
              (950, 392, "pod-2: 6 hops", 26, MUTED)]),
  "the scan: from one instance, measure every other instance"),

 (topo({P9_i: C, A7_i: R, R1_i: R, ZA_i: R, R2_i: R, W3_i: R, P2_i: R, RG_i: R, ZB_i: R},
       {P9_i: "start"},
       notes=[(950, 350, "same walk,", 26, MUTED),
              (950, 392, "new start", 26, MUTED)]),
  "then from the next one. Once per machine in the fleet."),

 (topo({i: R for i in range(9)},
       notes=[(950, 308, "4,197 starts", 27, STROKE),
              (950, 350, "x 4,197 nodes", 27, STROKE),
              (950, 400, "17.6 million", 27, REPEAT),
              (950, 442, "visits", 27, REPEAT)]),
  "every node is walked once per starting point"),

 (topo({ZA_i: C, R1_i: D, A7_i: D, P2_i: D, R2_i: D, W3_i: D, P9_i: D},
       {ZA_i: "bends here"},
       notes=[(950, 350, "3 down one side,", 26, MUTED),
              (950, 392, "3 down the other", 26, MUTED)]),
  "every path has one highest node. This path's is zone-a."),

 (topo({RG_i: C, ZA_i: D, ZB_i: D, R1_i: R, A7_i: R, P2_i: R, R2_i: R, W3_i: R, P9_i: R},
       {RG_i: "4 + 1 = 5"},
       notes=[(950, 350, "true answer:", 26, REPEAT),
              (950, 392, "6 hops", 26, REPEAT)]),
  "the root's two deepest branches miss a path it never sees"),
 (topo({i: D for i in range(9)},
       {P2_i: "0", A7_i: "1", R1_i: "2", ZA_i: "3", RG_i: "4", ZB_i: "0",
        R2_i: "2", W3_i: "1", P9_i: "0"},
       notes=[(950, 308, "each node returns", 25, MUTED),
              (950, 346, "how far down it", 25, MUTED),
              (950, 384, "reaches. zone-a", 25, MUTED),
              (950, 422, "sees 3 and 3.", 25, MUTED)]),
  "one pass, bottom up: a height up, a running best kept"),

]
render(steps, os.path.join(os.path.dirname(os.path.abspath(__file__)), "images"))
