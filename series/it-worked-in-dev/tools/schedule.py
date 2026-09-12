#!/usr/bin/env python3
"""What can publish, what is blocked, and when the blocked ones unlock.

An episode links back to the LeetCode days that teach its technique. Those
links are only worth anything if the reader can open them, so an episode is not
publishable until every day it depends on has actually been POSTED - written
and committed in the public repo is not enough.

This checks that mechanically instead of from memory.

  python3 tools/schedule.py
"""
import datetime
import os
import sys

try:
    import yaml
except ImportError:
    sys.exit("needs pyyaml:  pip3 install pyyaml")

HERE = os.path.dirname(os.path.abspath(__file__))
SERIES = os.path.dirname(HERE)
QUEUE = os.path.join(SERIES, "queue.yaml")

# The series sits inside the daily-challenge repo now, so the day queue it
# gates on is two levels up rather than in a sibling checkout. That also means
# the gate can never be checked against a stale copy of the other repo.
REPO = os.path.dirname(os.path.dirname(SERIES))
LEETCODE = os.environ.get(
    "LEETCODE_QUEUE", os.path.join(REPO, "challenge", "queue.yaml"))

POSTED_TO = "linkedin_main_account"  # the surface that decides "the audience saw it"


def load_leetcode():
    """day -> (title, posted_bool) from the public repo."""
    if not os.path.exists(LEETCODE):
        return None
    with open(LEETCODE) as f:
        q = yaml.safe_load(f) or {}
    out = {}
    for e in q.get("entries", []):
        at = (e.get("posted_at") or {}).get(POSTED_TO)
        out[e["day"]] = (e.get("title", "?"), bool(at))
    return out


def main():
    with open(QUEUE) as f:
        q = yaml.safe_load(f) or {}
    days = load_leetcode()

    if days is None:
        print(f"cannot read {LEETCODE}")
        print("set LEETCODE_QUEUE to the public repo's challenge/queue.yaml")
        return

    posted = sorted(d for d, (_, p) in days.items() if p)
    print(f"LeetCode days posted to {POSTED_TO}: {len(posted)}"
          + (f" (through day {posted[-1]})" if posted else ""))
    print(f"written but not yet posted: {len(days) - len(posted)}\n")

    ready, blocked = [], []
    for ep in q.get("episodes", []):
        need = ep.get("requires_days") or []
        missing = [d for d in need if not days.get(d, ("", False))[1]]
        (blocked if missing else ready).append((ep, missing))

    print("PUBLISHABLE NOW")
    print("-" * 15)
    any_ready = False
    for ep, _ in ready:
        if ep["status"] == "planned":
            continue
        any_ready = True
        need = ep.get("requires_days") or []
        cites = ", ".join(f"day {d}" for d in need) or "no prerequisites"
        print(f"  ep {ep['episode']}  {ep['title']}")
        print(f"          {ep['technique']}  |  cites {cites}")
    if not any_ready:
        print("  (nothing written and unblocked)")

    print("\nWRITTEN BUT BLOCKED")
    print("-" * 19)
    shown = False
    for ep, missing in blocked:
        if ep["status"] == "planned":
            continue
        shown = True
        names = ", ".join(f"day {d} ({days[d][0][:26]})" for d in missing if d in days)
        print(f"  ep {ep['episode']}  {ep['title']}")
        print(f"          waiting on: {names}")
    if not shown:
        print("  (none)")

    print("\nPLANNED, UNLOCKS AT")
    print("-" * 19)
    for ep, missing in sorted(blocked + ready, key=lambda t: t[0]["episode"]):
        if ep["status"] != "planned":
            continue
        need = ep.get("requires_days") or []
        if not need:
            when = "now - no prerequisites"
        else:
            last = max(need)
            if all(days.get(d, ("", False))[1] for d in need):
                when = "now"
            else:
                ahead = last - (posted[-1] if posted else 0)
                eta = datetime.date.today() + datetime.timedelta(days=ahead)
                when = f"after day {last} posts, about {eta.isoformat()}"
        print(f"  ep {ep['episode']}  {ep['title'][:44]:<44} {when}")

    print("\nThe gate exists so no episode ships a link a reader cannot open.")


if __name__ == "__main__":
    main()
