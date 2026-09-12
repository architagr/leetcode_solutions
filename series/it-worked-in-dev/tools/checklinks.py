#!/usr/bin/env python3
"""Refuse to call an episode ready while a cross-link is unset.

A backlink that 404s is worse than an absent one: it costs the reader's trust
and it is invisible to the person who wrote it.

Two checks. First, that every spoke in LINKS.yaml has a URL. Second, that no
draft still carries an unresolved {{placeholder}} - drafts are written with
{{github_growth}} and friends so the URL lives in one place, and pasting one
into LinkedIn with the braces still in it is the exact failure this prevents.
"""
import os
import re
import sys

try:
    import yaml
except ImportError:
    sys.exit("needs pyyaml:  pip3 install pyyaml")

SERIES = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
CONTENT = os.path.join(SERIES, "episodes")
links = yaml.safe_load(open(os.path.join(SERIES, "LINKS.yaml")))

spokes = links.get("spokes") or {}
hubs = links.get("hub") or {}
known = dict(hubs)
known.update(spokes)

missing = [k for k, v in spokes.items() if not v]
have = [k for k, v in spokes.items() if v]

print("configured:", ", ".join(have) or "(none)")
if missing:
    print("MISSING   :", ", ".join(missing))

# Drafts carry the real URLs now rather than {{placeholders}}, because a draft
# you cannot read the links in is a draft you cannot check. LINKS.yaml stays
# the record of what the canonical URL for each surface is, so the job here is
# drift: a draft pointing somewhere LINKS.yaml does not know about.
known_urls = {v.rstrip("/") for v in known.values() if v}
leftover, unknown, used = [], [], {}

URL = re.compile(r"https?://[^\s)\]>\"']+")
for root, _, files in os.walk(CONTENT):
    for f in files:
        if not f.endswith(".md"):
            continue
        path = os.path.join(root, f)
        rel = os.path.relpath(path, SERIES)
        text = open(path).read()
        for key in set(re.findall(r"\{\{(\w+)\}\}", text)):
            leftover.append("%s: {{%s}} was never substituted" % (rel, key))
        for url in set(URL.findall(text)):
            base = url.rstrip("/.,")
            if any(base.startswith(k) for k in known_urls):
                for key, value in known.items():
                    if value and base.startswith(value.rstrip("/")):
                        used.setdefault(key, set()).add(rel)
                continue
            if "leetcode.com" in base or "go.dev" in base:
                continue
            unknown.append("%s: %s is not a URL LINKS.yaml knows about" % (rel, base))

print()
print("links in drafts, by surface:")
for key in sorted(used):
    print("  %-18s %d file(s)" % (key, len(used[key])))
for key in sorted(set(known) - set(used)):
    print("  %-18s not linked from any draft" % key)

for line in leftover + unknown:
    print("  FAIL", line)
if leftover or unknown:
    missing = missing or ["(see failures above)"]

if missing or unknown:
    print("\nEpisodes can be written and committed, but must not be published to a")
    print("surface whose URL is unset - and cannot cross-link to one either.")
    sys.exit(1)
print("\nall spokes configured, every draft link accounted for")
