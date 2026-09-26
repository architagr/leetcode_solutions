import os, sys
sys.path.insert(0, "/Users/architagarwal/code/codestreak-growth/tools")
from diagram import *

# A 12-person company. Attendees: ana, bo and cy. Their lowest common manager
# is vp1. In/Out are the pre-order numbers from Number.
W_ = 120
Y = [86, 196, 306, 416]
POS = {"ceo": (530, 0), "vp1": (250, 1), "vp2": (810, 1),
       "m1": (110, 2), "m2": (390, 2), "m3": (670, 2), "m4": (950, 2),
       "ana": (20, 3), "bo": (200, 3), "cy": (390, 3), "dev": (670, 3), "eli": (950, 3)}
KIDS = {"ceo": ["vp1", "vp2"], "vp1": ["m1", "m2"], "vp2": ["m3", "m4"],
        "m1": ["ana", "bo"], "m2": ["cy"], "m3": ["dev"], "m4": ["eli"]}
ATT = ["ana", "bo", "cy"]
NUM = {"ceo": "0-11", "vp1": "1-6", "m1": "2-4", "ana": "3", "bo": "4", "m2": "5-6",
       "cy": "6", "vp2": "7-11", "m3": "8-9", "dev": "9", "m4": "10-11", "eli": "11"}
Pn, C, D, R = "pend", "cur", "done", "repeat"


def org(states, subs=None, notes=()):
    subs = subs or {}
    out = []
    for p, ks in KIDS.items():
        px, pl = POS[p]
        for k in ks:
            kx, kl = POS[k]
            out.append(connect(px + W_ / 2, Y[pl] + BH, kx + W_ / 2, Y[kl], width=3))
    for n, (x, l) in POS.items():
        out += box(x, Y[l], n, states.get(n, Pn), sub=subs.get(n), w=W_)
    for y, text, color in notes:
        out.append(label(600, y, text, 26, color))
    return out


steps = [

 (org({**{a: C for a in ATT}, "vp1": D}, notes=[
     (518, "a meeting with ana, bo and cy: who is the lowest", MUTED),
     (554, "manager above all three? vp1.", STROKE)]),
  "the lowest manager above everyone on the invite"),

 (org({"ceo": D, "vp1": D, "m1": D, "ana": C}, notes=[
     (518, "PathTo(ana): search from the top, return the chain", MUTED),
     (554, "ceo, vp1, m1, ana. Then the same for bo, then cy.", STROKE)]),
  "what you would write: a breadcrumb per attendee, then the prefix"),

 (org({"ceo": R, "vp1": R, "m1": R, "m2": R, "ana": C, "bo": C, "cy": C}, notes=[
     (518, "every breadcrumb searches from the ceo down again", MUTED),
     (554, "2,000 attendees: 2,000 searches of 50,000 people", REPEAT)]),
  "every attendee's search walks the same managers from the top"),

 (org({"ceo": D, "vp1": D, "m1": D, "ana": C, "bo": C, "m2": D, "cy": C}, notes=[
     (518, "the first search walked right past bo and cy", MUTED),
     (554, "one walk sees every attendee. It just was not counting.", STROKE)]),
  "one walk already passes every attendee on the way"),

 (org({"ana": C, "bo": C, "cy": C, "m1": D, "m2": D, "vp1": D},
      subs={"ana": "1", "bo": "1", "cy": "1", "m1": "2", "m2": "1", "vp1": "3",
            "dev": "0", "eli": "0", "m3": "0", "m4": "0"},
      notes=[(518, "each person reports how many attendees are below them", MUTED),
             (554, "vp1 is the first to reach 3: stop there. 265x.", STROKE)]),
  "count bottom-up: the first manager to see everyone is the answer"),

 (org({"ceo": D, "vp1": C, "ana": D, "bo": D, "cy": D},
      subs={n: NUM[n] for n in ["ceo", "vp1", "vp2", "m1", "m2", "ana", "bo", "cy"]},
      notes=[(518, "number once: a reporting line is a range. Attendees: 3 to 6.", MUTED),
             (554, "vp1 covers 3-6, none of its reports do. 2,000 in 24.3 µs", STROKE)]),
  "numbered once, the chart answers like a search tree"),
]
render(steps, os.path.join(os.path.dirname(os.path.abspath(__file__)), "images"))
