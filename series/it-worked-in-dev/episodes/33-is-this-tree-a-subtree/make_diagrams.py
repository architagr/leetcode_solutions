import os, sys
sys.path.insert(0, "/Users/architagarwal/code/codestreak-growth/tools")
from diagram import *

# A reply chain on the page, c1..c5, and the new render of its last three
# comments, where the bottom one was edited.
Pn, C, D, R = "pend", "cur", "done", "repeat"
CW = 150
PX, SX = 140, 760
PY = [112, 178, 244, 310, 376]
SY = [178, 244, 310]


def chain(x, ys, names, states, subs=None, last=None):
    subs = subs or {}
    out = []
    for i in range(len(ys) - 1):
        out.append(connect(x + CW / 2, ys[i] + BH, x + CW / 2, ys[i + 1], width=3))
    for i, (y, nm) in enumerate(zip(ys, names)):
        out += box(x, y, nm, states.get(nm, Pn), sub=subs.get(nm), w=CW)
    return out


def scene(ps, ss=None, psubs=None, ssubs=None, notes=(), snip=True):
    out = [label(PX + CW / 2, 100, "the page", 22, MUTED)]
    out += chain(PX, PY, ["c1", "c2", "c3", "c4", "c5"], ps, psubs)
    if snip:
        out.append(label(SX + CW / 2, 166, "new render", 22, MUTED))
        out += chain(SX, SY, ["s1", "s2", "s3*"], ss or {}, ssubs)
    for y, text, color in notes:
        out.append(label(600, y, text, 26, color))
    return out


steps = [

 (scene({}, notes=[
     (470, "before mounting a tree, the renderer asks: is it already on the page?", MUTED),
     (510, "same kinds, same children, same order. s3* was edited: it is not.", STROKE)]),
  "is this component tree already somewhere on the page?"),

 (scene({"c1": C, "c2": D, "c3": D, "c4": D}, {"s1": D, "s2": D, "s3*": R}, notes=[
     (470, "every comment is a place the snippet could start", MUTED),
     (510, "from c1: c1=s1, c2=s2, c3 vs s3*: no. Then try c2.", STROKE)]),
  "what you would write: try every anchor, compare in lockstep"),

 (scene({"c2": C, "c3": R, "c4": R, "c5": R}, {"s1": D, "s2": D, "s3*": R},
        psubs={"c3": "read from c1, c2", "c4": "read from c2, c3"}, notes=[
     (470, "each anchor walks down until the snippet runs out", MUTED),
     (510, "a comment is read once for every comment above it", REPEAT),
     (550, "5,000 deep: 13.5 million comparisons, 75.1 ms", REPEAT)]),
  "on a reply chain, every anchor rereads the comments below it"),

 (scene({"c1": D, "c2": C, "c3": C, "c4": C, "c5": C},
        psubs={"c2": "inside c1's subtree"}, snip=False, notes=[
     (470, "the whole subtree under c2 is part of the subtree under c1", MUTED),
     (510, "asking c1 read all of it. Asking c2 reads it again, from scratch.", STROKE)]),
  "the answer for c2 was read while asking about c1"),

 (scene({"c5": D, "c4": D, "c3": D, "c2": D, "c1": D}, {"s1": D, "s2": D, "s3*": D},
        psubs={"c5": "f5 = f(comment, a, b)", "c4": "f4 = f(comment, a, b, f5)",
               "c3": "f3 = f(.., f4)", "c2": "f2", "c1": "f1"},
        ssubs={"s3*": "g3", "s2": "g2", "s1": "g1"}, notes=[
     (470, "a fingerprint of each subtree, from its kind and its kids' prints", MUTED),
     (510, "bottom-up: every node once, the children answer first", STROKE)]),
  "fingerprint every subtree bottom-up, in one pass"),

 (scene({"c1": D, "c2": D, "c3": D, "c4": D, "c5": D}, {"s1": C},
        psubs={"c1": "f1 != g1", "c2": "f2 != g1", "c3": "f3 != g1"}, notes=[
     (470, "compare one number per node against g1, the snippet's print", MUTED),
     (510, "different numbers: different trees. Equal: confirm once.", STROKE),
     (550, "the same reply chain: 163 µs, 461x faster", DONE)]),
  "one number per node: only an equal print gets compared in full"),

 (None,
  "the other one-pass answer: write it out and search"),
]

# step 7 draws only notes; keep the page chain off it
steps[6] = ([label(600, y, t, 26, c) for y, t, c in [
     (200, "card(img(p))  and  card(img, p)", STROKE),
     (260, "kinds in order, as text: card,img,p, and card,img,p,", REPEAT),
     (320, "same text, different trees: Contains says yes", REPEAT),
     (400, "with brackets: (card(img(p))) and (card(img)(p))", DONE),
     (460, "correct, and one pass - but the page becomes a string", MUTED)]],
  "the other one-pass answer: write it out and search")
render(steps, os.path.join(os.path.dirname(os.path.abspath(__file__)), "images"))
