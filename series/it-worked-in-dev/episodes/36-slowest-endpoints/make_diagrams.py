import os, sys
sys.path.insert(0, "/Users/architagarwal/code/codestreak-growth/tools")
from diagram import *

# Twelve requests in arrival order, latency in ms. Drawn with k = 3: the three
# slowest are 92, 88 and 81.
LAT = [42, 35, 81, 30, 55, 92, 38, 47, 88, 33, 61, 40]
X0, STEP, W_ = 60, 92, 70
BASE = 380
FILL = {"pend": "#d9dee3", "cur": CUR, "done": DONE, "repeat": "#f2c4bd", "gone": "#eef1f4"}


def bars(states, lat=LAT, notes=(), bar=None, extra=()):
    out = []
    for i, v in enumerate(lat):
        h = v * 3
        x = X0 + i * STEP
        st = states.get(i, "pend")
        out.append(f'<rect x="{x}" y="{BASE - h}" width="{W_}" height="{h}" rx="6" fill="{FILL[st]}"/>')
        out.append(label(x + W_ / 2, BASE - h - 10, str(v), 22, STROKE if st != "gone" else MUTED))
        out.append(label(x + W_ / 2, BASE + 28, "r%d" % i, 20, MUTED))
    if bar is not None:
        y = BASE - bar * 3
        out.append(f'<line x1="40" y1="{y}" x2="1160" y2="{y}" stroke="{CUR}" stroke-width="3" '
                   f'stroke-dasharray="10 6"/>')
    out += list(extra)
    for y, text, color in notes:
        out.append(label(600, y, text, 26, color))
    return out


TOP = {2: "done", 5: "done", 8: "done"}
SORTED = sorted(LAT)
steps = [

 (bars(TOP, notes=[
     (470, "an hour of requests in arrival order; the page shows the slowest", MUTED),
     (510, "drawn here with the top 3: r5 92 ms, r8 88, r2 81", STROKE)]),
  "the incident page: the ten slowest requests of the hour"),

 (bars({i: ("done" if v >= 81 else "repeat") for i, v in enumerate(SORTED)}, lat=SORTED, notes=[
     (470, "copy the log, sort all of it slowest first, keep the first 10", MUTED),
     (510, "every request gets a place in the order. The page shows 10.", REPEAT)]),
  "what you would write: sort a copy, take the first ten"),

 (bars({**{i: "gone" for i in range(12)}, 0: "done", 1: "done", 2: "cur"}, bar=35, notes=[
     (470, "after the first three, the weakest of them is r1 at 35 ms", MUTED),
     (510, "nothing at or under the weakest kept can ever be shown", STROKE)]),
  "the first k requests set a bar: the weakest one kept"),

 (bars({**{i: "gone" for i in range(12)}, 2: "done", 4: "done", 5: "cur"}, bar=55, notes=[
     (470, "a request has to beat the bar to get in: one comparison", MUTED),
     (510, "an ordinary hour: 142 of 3,600,000 requests ever get in", STROKE)]),
  "everything under the bar is rejected with one comparison"),

 (bars({2: "done", 5: "done", 8: "done", 10: "repeat"}, bar=81, notes=[
     (470, "keep the k with the weakest at the root of a min-heap", MUTED),
     (510, "r10 at 61 loses to 81 and is gone. 2.72 ms, 154x the sort.", STROKE)]),
  "a min-heap of k: the weakest kept is always at the root"),

 (bars({i: "repeat" for i in range(12)}, lat=SORTED, notes=[
     (470, "if every request is slower than the one before, all get in", MUTED),
     (510, "each one costs a heap fix. The sort wins this hour, 4.20x.", REPEAT)]),
  "when every request beats the bar, the heap loses"),
]
render(steps, os.path.join(os.path.dirname(os.path.abspath(__file__)), "images"))
