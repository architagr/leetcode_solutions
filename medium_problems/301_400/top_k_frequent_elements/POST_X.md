Day 78/365 · Top K Frequent Elements (Medium)

To keep the k MOST frequent, use a MIN-heap.

The root must be the weakest entry you kept, because that is the one an arrival displaces. Backwards, and you silently return the k least frequent.

https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/301_400/top_k_frequent_elements/SOLUTION.md

#golang
