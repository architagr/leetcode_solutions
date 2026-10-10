import os, sys
sys.path.insert(0, "/Users/architagarwal/code/codestreak-growth/tools")
from diagram import *

# root: payments (timeout_ms read, legacy_gw not), old_banner (not read),
# pricing: a rule chain whose only read key is the fallback at the bottom.
# The when/price keys of each rule are left out of the drawing.
Pn, C, D, R = "pend", "cur", "done", "repeat"
BW_ = 200
POS = {"root": (500, 92), "payments": (145, 172), "old_banner": (500, 172),
       "pricing": (840, 172), "timeout_ms": (30, 272), "legacy_gw": (255, 272),
       "else1": (840, 252), "else2": (840, 332), "fallback": (840, 412)}
KIDS = {"root": ["payments", "old_banner", "pricing"], "payments": ["timeout_ms", "legacy_gw"],
        "pricing": ["else1"], "else1": ["else2"], "else2": ["fallback"]}
READ = {"timeout_ms", "fallback"}


def tree(states, notes=(), hide=(), marks=True):
    out = []
    for p, ks in KIDS.items():
        if p in hide:
            continue
        px, py = POS[p]
        for k in ks:
            if k in hide:
                continue
            kx, ky = POS[k]
            out.append(connect(px + BW_ / 2, py + BH, kx + BW_ / 2, ky, width=3))
    for k, (x, y) in POS.items():
        if k in hide:
            continue
        out += box(x, y, k + ("*" if k in READ else ""), states.get(k, Pn), w=BW_)
    for y, text, color in notes:
        out.append(label(600, y, text, 26, color))
    return out


steps = [

 (tree({}, notes=[
     (500, "drop every key nothing reads, and every section left empty", MUTED),
     (540, "* is read by the code. Keep payments and the pricing chain.", STROKE)]),
  "drop the config branches nothing in the code reads"),

 (tree({"payments": R, "old_banner": R, "pricing": R, "root": D, "timeout_ms": R,
       "legacy_gw": R, "else1": R, "else2": R, "fallback": R}, notes=[
     (500, "first try: drop any key nobody reads, deciding on the way down", MUTED),
     (540, "payments is not read - and timeout_ms goes with it. So does fallback.", REPEAT)]),
  "deciding on the way down drops sections you still need"),

 (tree({"root": D, "payments": C, "timeout_ms": D, "old_banner": R}, notes=[
     (500, "at each section ask: is anything under here read?", MUTED),
     (540, "payments: yes, keep it and go on. old_banner: no, drop it.", STROKE)]),
  "what you would write: ask each section about everything below it"),

 (tree({"pricing": C, "else1": R, "else2": R, "fallback": D}, notes=[
     (500, "pricing asks: searches down to fallback. Then else1 asks, else2...", MUTED),
     (540, "5,000 rules deep: 37.5 million reads for 15,004 keys, 104 ms", REPEAT)]),
  "a rule chain: every section searches the same keys again"),

 (tree({"pricing": D, "else1": C, "else2": D, "fallback": D}, notes=[
     (500, "asking for pricing walked through else1 and else2 to fallback", MUTED),
     (540, "so their answers were found already - and then asked for again", STROKE)]),
  "the answer for else1 was found while asking about pricing"),

 (tree({k: D for k in POS if k not in ("legacy_gw", "old_banner")}, notes=[
     (500, "children first: each comes back kept or gone, and says so", MUTED),
     (540, "a section stays if it is read or a child stayed. 92.4 µs.", STROKE)]),
  "bottom-up: every key read once, the children answer first"),
]
render(steps, os.path.join(os.path.dirname(os.path.abspath(__file__)), "images"))
