import os, sys
sys.path.insert(0, "/Users/architagarwal/code/codestreak-growth/tools")
from diagram import *

# Six packages, three layers, listed apps first: web, api, auth, db, log, cfg.
W_ = 130
POS = {"web": (40, 96), "api": (380, 96), "auth": (210, 236), "db": (500, 236),
       "log": (40, 376), "cfg": (380, 376)}
IMPORTS = {"web": ["auth", "log"], "api": ["auth", "db"], "auth": ["log", "cfg"],
           "db": ["cfg"], "log": [], "cfg": []}
Pn, C, D, R = "pend", "cur", "done", "repeat"
PX = 680


def graph(states=None, subs=None, panel=(), title=None, notes=(), reverse=False):
    states, subs = states or {}, subs or {}
    out = []
    for p, ds in IMPORTS.items():
        px, py = POS[p]
        for d in ds:
            dx, dy = POS[d]
            if reverse:
                out.append(connect(dx + W_ / 2, dy, px + W_ / 2, py + BH,
                                   color=DONE, marker="up", width=3))
            else:
                out.append(connect(px + W_ / 2, py + BH, dx + W_ / 2, dy,
                                   marker="a", width=3))
    for n, (x, y) in POS.items():
        out += box(x, y, n, states.get(n, Pn), sub=subs.get(n), w=W_)
    if title:
        out.append(label(PX, 76, title, 22, MUTED, anchor="start"))
    for i, (text, state) in enumerate(panel):
        out += box(PX, 96 + i * 68, text, state, w=480)
    for y, text, color in notes:
        out.append(label(600, y, text, 26, color))
    return out


ALL = {n: D for n in POS}

steps = [

 (graph(ALL, title="listed apps first",
        panel=[("web api auth db log cfg", Pn), ("an arrow: imports", Pn),
               ("a valid build order:", D), ("log cfg auth db web api", D)],
        notes=[(500, "every import built before the package that imports it", STROKE)]),
  "a build order puts every import before the package using it"),

 (graph({"web": R, "api": R, "auth": R, "db": R, "log": D, "cfg": D},
        title="pass 1, in list order",
        panel=[("web: auth not built", R), ("api: auth not built", R),
               ("auth: log not built", R), ("db: cfg not built", R),
               ("log, cfg: built", D)],
        notes=[(540, "repeat until a pass builds nothing", STROKE)]),
  "what you would write: pass over the list, build what is ready"),

 (graph({"web": R, "api": R, "auth": D, "db": D, "log": D, "cfg": D},
        title="pass 2, then pass 3",
        panel=[("pass 2: web, api again", R), ("pass 2: auth, db built", D),
               ("pass 3: web, api built", D), ("3 layers, 3 passes", C)],
        notes=[(500, "one pass per layer when the list is in the wrong order", MUTED),
              (540, "10 layers: 99.3 µs. A 5,000-long chain: 18.8 ms.", REPEAT)]),
  "web and api are re-checked every pass until auth is finally built"),

 (graph({"log": D, "cfg": C, "auth": C, "db": C}, reverse=True,
        title="when cfg is built",
        panel=[("who imports cfg?", D), ("auth and db. Nobody else.", C),
               ("only they can be ready now", C)],
        notes=[(540, "the build already knew who it could unblock", STROKE)]),
  "building a package can only unblock the packages that import it"),

 (graph({"log": C, "cfg": C}, subs={"web": "2", "api": "2", "auth": "2", "db": "1",
                                     "log": "0", "cfg": "0"},
        title="count what each is waiting for",
        panel=[("number: imports not yet built", Pn), ("0: ready now", C),
               ("queue: log, cfg", C)],
        notes=[(540, "reverse the arrows once, and count", STROKE)]),
  "count each package's unbuilt imports, and queue the zeros"),

 (graph({"log": D, "cfg": D, "auth": D, "db": D, "web": C, "api": C},
        subs={"web": "0", "api": "0", "auth": "0", "db": "0"},
        title="each build counts importers down",
        panel=[("log: web 1, auth 1", D), ("cfg: auth 0, db 0", D),
               ("auth: web 0, api 1", D), ("db: api 0", D)],
        notes=[(540, "each arrow looked at once, whatever the order", STROKE)]),
  "Kahn: queue a package the moment its count reaches zero"),
]
render(steps, os.path.join(os.path.dirname(os.path.abspath(__file__)), "images"))
