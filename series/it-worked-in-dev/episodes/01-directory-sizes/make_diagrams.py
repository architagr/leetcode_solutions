import os, sys
sys.path.insert(0, os.path.join(os.path.dirname(__file__), "..", "..", "..", "..", "tools"))
from diagram import *

# /root -> a -> {a1, a2}, b
ROOT = (505, 120)
A    = (260, 270)
B    = (760, 270)
A1   = (120, 420)
A2   = (400, 420)

def tree(states, subs=None):
    subs = subs or {}
    out = []
    for (x, y), (px, py) in [(A, ROOT), (B, ROOT), (A1, A), (A2, A)]:
        out.append(connect(px + BW/2, py + BH, x + BW/2, y, width=3))
    out += box(*ROOT, "/root", states.get("root", "pend"), subs.get("root"))
    out += box(*A,    "  a",   states.get("a", "pend"),    subs.get("a"))
    out += box(*B,    "  b",   states.get("b", "pend"),    subs.get("b"))
    out += box(*A1,   " a1",   states.get("a1", "pend"),   subs.get("a1"))
    out += box(*A2,   " a2",   states.get("a2", "pend"),   subs.get("a2"))
    return out

P = "pend"; C = "cur"; D = "done"; R = "repeat"
steps = [
 (tree({}, {"root":"2 files","a":"1 file","b":"4 files","a1":"3 files","a2":"5 files"}),
  "each directory holds its own files, and its subdirectories"),

 (tree({"root":C,"a":R,"b":R,"a1":R,"a2":R}),
  "SizeOf(/root) walks everything beneath it. Correct, and reasonable."),

 (tree({"root":D,"a":C,"a1":R,"a2":R,"b":P}, {"root":"15"}),
  "then SizeOf(a) walks a1 and a2 AGAIN - they were just counted"),

 (tree({"root":D,"a":D,"a1":C,"a2":C,"b":D}, {"root":"15","a":"9","b":"4"}),
  "a1 and a2 were summed three times each: by a1/a2, by a, by /root"),

 (tree({"a1":D,"a2":D,"a":P,"b":P,"root":P}, {"a1":"3","a2":"5"}) +
  [label(600, 560, "leaves resolve first", 25, MUTED)],
  "postorder instead: compute children before the parent asks"),

 (tree({"a1":D,"a2":D,"b":D,"a":C,"root":P}, {"a1":"3","a2":"5","b":"4"}) +
  [connect(A1[0]+BW/2, A1[1], A[0]+BW/2, A[1]+BH, DONE, "up", 5),
   connect(A2[0]+BW/2, A2[1], A[0]+BW/2, A[1]+BH, DONE, "up", 5)],
  "a adds 1 + 3 + 5 = 9, reading numbers its children already computed"),

 (tree({"a1":D,"a2":D,"b":D,"a":D,"root":D}, {"a1":"3","a2":"5","b":"4","a":"9","root":"15"}) +
  [connect(A[0]+BW/2, A[1], ROOT[0]+BW/2, ROOT[1]+BH, DONE, "up", 5),
   connect(B[0]+BW/2, B[1], ROOT[0]+BW/2, ROOT[1]+BH, DONE, "up", 5)],
  "every file counted once. Same answers, one pass."),
]
render(steps, os.path.join(os.path.dirname(__file__), "images"))
