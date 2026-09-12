<!-- X. Two posts, and that is the cap. X is the top of the funnel, not a
     place to spend the argument - the thread version gave away the number,
     the mechanism and the fix, which left nobody a reason to click.

     Post 1 carries the surprise and stops. Post 2 sends them somewhere that
     can actually hold a benchmark, and offers the choice rather than picking
     for them: Medium if that is where they read, the newsletter if they want
     it in the inbox, GitHub if they want to run it.

     A link costs 23 characters on X no matter how long it is. -->

---
meta_title: "801 dirs beat 8,191 dirs. Depth is the axis."
meta_description: "A Go benchmark where the smaller tree ran three times slower. Two posts, the surprise and the link, no spoilers."
hero: ../HERO.png
tags: [golang, algorithms, performance]
hashtags: "#golang #dsa"
---

![10x smaller. 3x slower.](../HERO.png)

**1/2** [attach HERO.png]

801 directories took 3x longer than 8,191.

Same code. Same machine.

The small tree was 800 levels deep. The big one was 12.

I benchmarked the function everyone writes for du and had the wrong mental model of what actually makes it slow.

#golang #dsa

---

**2/2**

The fix is two moved lines. That's not the interesting part.

The interesting part is the question that gets you there on your own — and when it doesn't apply.

Medium: https://medium.com/@architagr
Inbox: https://www.linkedin.com/newsletters/the-weekly-golang-journal-7261403856079597568/
Code: https://github.com/architagr/leetcode_solutions/tree/main/series/it-worked-in-dev
