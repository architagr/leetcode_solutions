## Solution walkthrough

`maxVowels(s string, k int) int` in `main.go` keeps a count of vowels in a window of `k`
letters and tracks the largest count. `isVowel` checks a single byte.

We'll trace Example 1: `s = "abciiidef"`, `k = 3`, expected `3`.

1. **Offset and prime.** `k--` makes `k = 2`. The first loop counts vowels in `"ab"`:
   `count = 1`.

   ![Step 1: primed with "ab"](images/walkthrough-1.png)

2. **Add, check, remove.** At `i = 2`, `c` isn't a vowel, so the window `"abc"` has 1.
   `max = 1`. Then `s[i-k] = 'a'` leaves, and since it's a vowel, `count` drops to 0.

   ![Step 2: "abc" has 1](images/walkthrough-2.png)

3. **Slide.** `"bci"` has 1, `"cii"` has 2, so `max = 2`.

   ![Step 3: "cii" has 2](images/walkthrough-3.png)

4. **The best window.** `"iii"` has 3. That's `k` vowels in a window of `k` letters, the
   most any window can hold.

   ![Step 4: "iii" has 3](images/walkthrough-4.png)

5. **The rest can't beat it.** `"iid"`, `"ide"` and `"def"` have 2, 2 and 1. The loop
   checks them anyway; an early `return` once `max == k` would skip them.

   ![Step 5: final answer 3](images/walkthrough-5.png)

**Complexity:** O(n) time, O(1) space.
