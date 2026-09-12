<!-- LinkedIn feed post. Short, drives to the article. Attach HERO.png.
     LinkedIn truncates around 210 characters, so the surprise has to land
     in the first two lines. -->

---
meta_title: "801 directories beat 8,191. Depth is the axis."
meta_description: "A Go benchmark where 801 directories ran three times slower than 8,191, and the two-line change that closed the gap."
hero: ../HERO.png
hashtags: "#Golang #DSA #Algorithms #Performance #SoftwareEngineering #DataStructures #100DaysOfCode"
---

![10x smaller. 3x slower.](../HERO.png)

**It worked in dev — Episode 1**

801 directories took three times longer than 8,191 directories. Same code, same
machine.

The smaller tree was 800 levels deep. The bigger one was 12.

I was benchmarking the function everyone writes — total bytes under every
directory, what `du` prints — expecting big trees to be the problem. They are
not. `SizeOf` re-walks the entire subtree on every call, so a file nested 800
deep gets added into a running total 800 separate times. Four times the
directories cost 16.8 times the work.

The fix is two lines, and I would rather you find it than read it. The question
that gets you there is not "which algorithm" — it is **does the answer for a
node depend only on the answers for its children?** When that sentence is true,
every answer can be computed once and handed upward instead of being fetched.
When it is false, you need something else entirely. That distinction is the
skill; the traversal is just what happens after it fires.

Full write-up has both implementations, the benchmark you can run yourself, and
seven walkthrough diagrams. It also says plainly where the naive version is
still the right call — which is more often than this post makes it sound.

Article in comments.

#Golang #DSA #Algorithms #Performance #SoftwareEngineering #DataStructures #100DaysOfCode
