import os, sys
sys.path.insert(0, "/Users/architagarwal/code/codestreak-growth/tools")
from diagram import *

# Eight hours of traffic (thousands of requests). Answers: 1 1 4 2 1 1 0 0.
LOAD = [73, 74, 75, 71, 69, 72, 76, 73]
WAIT = [1, 1, 4, 2, 1, 1, 0, 0]
Pn, C, D, R = "pend", "cur", "done", "repeat"
X0, STEP, W_ = 60, 135, 110
BASE = 360


def bars(states, labels=None, notes=(), extra=()):
    out = []
    for i, v in enumerate(LOAD):
        h = (v - 60) * 14
        x = X0 + i * STEP
        fill = {"done": DONE, "cur": CUR, "pend": "#d9dee3", "repeat": REPEAT}[states.get(i, Pn)]
        out.append(f'<rect x="{x}" y="{BASE - h}" width="{W_}" height="{h}" rx="6" fill="{fill}"/>')
        out.append(label(x + W_ / 2, BASE - h - 12, str(v), 24, STROKE))
        out.append(label(x + W_ / 2, BASE + 32, "h%d" % i, 22, MUTED))
        if labels and labels.get(i) is not None:
            out.append(label(x + W_ / 2, BASE + 72, labels[i], 26, DONE, bold=True))
    out += list(extra)
    for y, text, color in notes:
        out.append(label(600, y, text, 26, color))
    return out


steps = [

 (bars({}, labels={i: str(w) for i, w in enumerate(WAIT)}, notes=[
     (490, "for every hour: how many hours until traffic went higher", MUTED),
     (530, "0: it was never beaten. How long each peak held.", STROKE)]),
  "for every hour, how long until traffic next went higher"),

 (bars({2: C, 3: R, 4: R, 5: R, 6: D}, labels={2: "4"}, notes=[
     (490, "from each hour, look forward until something is higher", MUTED),
     (530, "h2 reads h3, h4, h5, h6. h3 will read h4, h5 again.", REPEAT)]),
  "what you would write: from each hour, scan forward"),

 (bars({5: C, 6: D}, notes=[
     (490, "a peak that is never beaten scans to the end of the data", MUTED),
     (530, "after a launch, every day's peak is one. 4.02x slower.", REPEAT)]),
  "falling traffic: many peaks are never beaten, and scan to the end"),

 (bars({3: D, 4: R}, notes=[
     (490, "h3 is 71: higher than h4, and nearer to every hour on its left", MUTED),
     (530, "so nobody left of h3 will ever stop at h4. Drop it for good.", STROKE)]),
  "a lower hour behind a higher, nearer one is never anyone's answer"),

 (bars({6: D, 2: C, 3: R, 5: R}, labels={2: "4", 3: "2", 4: "1", 5: "1", 6: "0", 7: "0"}, notes=[
     (490, "right to left. At h2 the stack is h6, h5, h3: pop h3, pop h5.", MUTED),
     (530, "the top, h6, is the answer: 4 hours. Nothing is popped twice. 17.9x.", STROKE)]),
  "a monotonic stack: each hour pushed once and popped at most once"),
]
render(steps, os.path.join(os.path.dirname(os.path.abspath(__file__)), "images"))
