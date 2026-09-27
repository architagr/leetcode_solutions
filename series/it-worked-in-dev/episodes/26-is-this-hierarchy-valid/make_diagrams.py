import os, sys
sys.path.insert(0, "/Users/architagarwal/code/codestreak-growth/tools")
from diagram import *

# Seven prices. The snapshot was patched: the node that held 40 now holds 55.
W_ = 100
Y = [86, 196, 306]
POS = {"50": (550, 0), "30": (270, 1), "70": (830, 1), "20": (130, 2), "55": (410, 2),
       "60": (690, 2), "80": (970, 2)}
KIDS = {"50": ["30", "70"], "30": ["20", "55"], "70": ["60", "80"]}
Pn, C, D, R = "pend", "cur", "done", "repeat"


def tree(states, subs=None, notes=()):
    subs = subs or {}
    out = []
    for p, ks in KIDS.items():
        px, pl = POS[p]
        for k in ks:
            kx, kl = POS[k]
            out.append(connect(px + W_ / 2, Y[pl] + BH, kx + W_ / 2, Y[kl], width=3))
    for p, (x, l) in POS.items():
        out += box(x, Y[l], p, states.get(p, Pn), sub=subs.get(p), w=W_)
    for y, text, color in notes:
        out.append(label(600, y, text, 26, color))
    return out


ALL = {p: D for p in POS}

steps = [

 (tree({**ALL, "55": C}, notes=[
     (450, "a snapshot patched by hand: the price that was 40 is now 55", MUTED),
     (490, "searching for 55 goes right at 50 and never finds it", STROKE),
     (530, "every lookup from episodes 23 and 25 now misses it", REPEAT)]),
  "a patched snapshot: 55 sits on the cheap side of 50"),

 (tree(ALL, subs={"30": "20 &lt; 30 &lt; 55", "50": "30 &lt; 50 &lt; 70", "70": "60 &lt; 70 &lt; 80"}, notes=[
     (450, "each node against its two children", MUTED),
     (490, "every check passes, and the check says valid", STROKE)]),
  "what looks right: every node against its own children"),

 (tree({"50": C, "30": D, "55": R}, subs={"55": "must be under 50"}, notes=[
     (450, "55 is fine against 30, its parent", MUTED),
     (490, "the rule is about the whole left subtree of 50", STROKE),
     (530, "and 55 is in it", REPEAT)]),
  "the rule is about subtrees, not children"),

 (tree({"50": C, "30": R, "20": R, "55": R, "70": R, "60": R, "80": R}, notes=[
     (450, "state it exactly: at every node, scan both whole subtrees", MUTED),
     (490, "correct, and every node is re-read once per ancestor", STROKE),
     (530, "a 5,000-long chain: 91.9 ms", REPEAT)]),
  "what you would write next: every node against its whole subtrees"),

 (tree({"50": D, "30": D, "20": D, "55": R},
       subs={"50": "(-inf, inf)", "30": "(-inf, 50)", "20": "(-inf, 30)", "55": "(30, 50)"},
       notes=[(450, "carry down the range every node must be strictly inside", MUTED),
              (490, "left: the node is the new upper bound. right: the new lower.", STROKE),
              (530, "55 is not inside (30, 50). One check, every ancestor.", REPEAT)]),
  "bounds carried down: each node checked once, against every ancestor"),
]
render(steps, os.path.join(os.path.dirname(os.path.abspath(__file__)), "images"))
