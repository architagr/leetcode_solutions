# Intuition

You’re given two non-negative numbers as strings and asked to return their sum as a string.
No built-in big integers. No cheating by converting to `int`.

Basically… become a human calculator 🧮

## 🔍 How Humans Do It

Forget coding for a second.
How do we add numbers?
Example:

```
   456
+   77
------
   533
```

We:

1. Start from the rightmost digit.
1. Add digit by digit.
1. Carry anything above 9 to the next step.

So our code should mimic this same behaviour.

## 🏗 Core Idea

Since we can’t convert strings to integers:

- We traverse both strings from right to left
- Add digits using ASCII math
- Maintain a carry
- Build the result digit by digit

Just like childhood… except now without the shout of:
"Carry over likh lo!"

## ✋ Key Observations

1. Strings are immutable — but bytes are not.
1. Traversing from the end avoids reversing inputs.
1. ASCII trick:
   - `'5' - '0' = 5`
1. Keep looping until:
   - both strings are consumed AND
   - carry is zero
