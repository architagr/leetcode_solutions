365 Days of LeetCode Challenge — Day 62/365

150. Evaluate Reverse Polish Notation (Medium)
https://leetcode.com/problems/evaluate-reverse-polish-notation/

Why postfix needs no parentheses: (2 + 1) * 3 needs brackets because + sits between its operands, so precedence rules have to decide what binds first. Written postfix it is 2 1 + 3 *, and each operator applies to the two values immediately before it. Evaluation order is already in the token order, so there is nothing for parentheses to do.

That is the whole basis of the algorithm. Numbers wait on a stack; an operator consumes the two most recent and leaves its result in their place.

The actual trap is smaller and easier to miss. Popping returns the operands in REVERSE - the first pop is the right operand, the second is the left.

For + and * that makes no difference, which is exactly why the bug survives casual testing. For - and / it picks between the right answer and a plausible wrong one. ["5","3","-"] means 5 - 3 = 2; get the order backwards and you get -2.

Full breakdown in today's newsletter article ⬇

#DSA #LeetCode #Golang #Stack #Parsing #CodingInterview #Algorithms
