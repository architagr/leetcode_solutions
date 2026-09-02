## Intuition

Re-sorting (or re-scanning) the entire stream every time `add` is called would be
correct but wasteful — with up to `10^4` calls, doing O(n log n) work per call adds up
fast. The key realization is that we don't need to know the full sorted order of every
score seen so far — we only ever need to answer one question: **what is the kth
largest value right now?**

That means we only need to track the **top k scores**, not all of them. And among just
those top k scores, the kth largest is exactly the **smallest** one — the "weakest
link" of the top-k group. So the problem reduces to: efficiently maintain a set of (at
most) k values, and be able to read/replace its minimum quickly. That's precisely what
a **min-heap** is for.

The strategy:
- Keep a min-heap that holds at most `k` elements — the k largest scores seen so far.
- When a new value arrives:
  - If the heap has fewer than `k` elements, it's automatically part of the current
    top-k, so just push it in.
  - Otherwise, compare the new value against the heap's root (the smallest of the
    current top-k). If the new value is larger, it belongs in the top-k more than the
    current smallest does — pop the smallest out and push the new value in. If the new
    value is smaller than or equal to the current smallest, it doesn't crack the top-k
    at all, and the heap is left untouched.
- After each `add`, the root of the min-heap — its smallest element — is by
  construction the kth largest value seen so far, since the heap holds exactly the
  top-k values and the smallest among them occupies the kth position when everything is
  sorted descending.

The constructor does the same thing, just seeded from the initial `nums` array instead
of one value at a time: fill the heap with the first `k` elements, then run every
remaining element through the same "does it beat the current minimum of the top-k?"
check.

**Complexity:**
- Time: O(log k) per `add` call (a heap push/pop on a heap bounded to size k), and
  O(n log k) for the constructor over an initial array of length n.
- Space: O(k) — the heap never grows past size k, regardless of how many scores the
  stream sees.
