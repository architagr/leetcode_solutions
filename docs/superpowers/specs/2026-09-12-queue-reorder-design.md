# Queue reorder — braided topic arcs (days 12–81)

Status: approved design, not yet implemented.
Date: 2026-09-12.
Supersedes nothing; extends the ordering rules in `.claude/skills/leetcode-content/SKILL.md`.

## Problem

`challenge/queue.yaml` holds 51 days, every one of them binary tree. Days 1–10 are posted,
day 11 (112 Path Sum) posts next and is treated as fixed. That leaves **40 consecutive
binary-tree days** ahead.

A reader who followed for a month has seen one topic and one traversal idea rehearsed forty
ways. The publishing system is healthy — cron posting, hero rendering, per-destination
tracking all work — but the content shape asks the audience to stay interested in a single
data structure for six weeks. That is the failure this reorder addresses.

It is worth being precise about what this does and does not fix. Reordering improves
**retention of people already following**. It does not improve **acquisition**, and it is
not a growth lever. It is worth doing because it is cheap and because publishing 40 straight
tree days would actively cost followers, not because it moves any follower target.

## Goals

- No unbroken run of one topic longer than ~13 days.
- Every topic switch is *earned* — each arc opens on a question whose technique the previous
  arc taught, so the series stays one chain rather than becoming a playlist.
- Every question drawn from problems already solved in this repo.
- No `builds_on` forward references, and no renumbering of a posted day.

## Non-goals

- Growing the follower count. See above.
- Reshaping days 82+. The horizon is deliberately ~2.3 months.
- Deduplicating the whole repo. Two related defects are noted below and fixed only where
  they touch a day in this plan.

## The braid

Forty remaining tree questions split into four themed arcs, interleaved with four new-topic
arcs drawn from already-solved problems. Longest unbroken tree run drops from **41 days to
13**.

| Days | Arc | Content | Hinge into it |
|---|---|---|---|
| 12–18 | Tree — recursion & bottom-up | ready | continues day 11 Path Sum |
| 19–26 | Linked List | new | day 10 was 114 Flatten Binary Tree to Linked List — the handoff is literal |
| 27–35 | Tree — BFS / level order | ready | 637 Average of Levels opens gently |
| 36–44 | Graph BFS/DFS | new | day 29 Level Order *is* BFS; day 42 lands 261 Graph Valid Tree as "the tree was a graph all along" |
| 45–57 | Tree — BST | ready | 700 Search in a BST |
| 58–64 | Stack & Iterator | new | day 55 BST Iterator → day 63 Flatten Nested List Iterator, same design problem |
| 65–74 | Tree — structure & compare | ready | 965/993, light after the BST run |
| 75–81 | Heap | 1 ready + 6 new | 703 Kth Largest in a Stream, already written, opens it |

The closing hinge is deliberate: **day 80 is 23 Merge k Sorted Lists**, which needs the heap
arc and day 22's Merge Two Sorted Lists. Two arcs 58 days apart converge on one hard problem.

## Difficulty mix — where this plan breaks its own rule

`SKILL.md` mandates "3-4 easy, then 1-2 medium or hard, repeating." Three arcs bend it and
one breaks it outright:

| Arc | Mix | Verdict |
|---|---|---|
| Graph BFS/DFS | **E1 M8** | breaks the rule |
| Tree — BFS | E3 M6 | bends it |
| Stack & Iterator | E2 M5 | bends it |

The graph arc is not fixable by reordering: `1971 Find if Path Exists in Graph` is the only
easy graph question solved in this repo. Every other one is medium or hard.

Decision: **accept it.** By day 36 the reader has had 35 days of ramp, and the arc opens on
the one easy. The alternative — solving two or three easy graph questions first — is a
reasonable choice if the day 36–44 stretch proves too steep, and can be made later without
disturbing anything else in this plan.

## Day map

Day 11 (112 Path Sum) is unchanged and posts as scheduled. `builds_on` lists LeetCode
numbers; every one was verified to land on an earlier day than the entry citing it,
including against the eleven frozen days.

### Days 12–18 · Tree — recursion & bottom-up

7 days · E3 M3 H1 · 0 need content generated

| Day | # | Title | Diff | Content | `builds_on` | Folder (new arcs only) |
|---|---|---|---|---|---|---|
| 12 | 1022 | Sum of Root To Leaf Binary Numbers | E | ready | 112,257 | — |
| 13 | 543 | Diameter of Binary Tree | E | ready | 104 | — |
| 14 | 563 | Binary Tree Tilt | E | ready | 543 | — |
| 15 | 2265 | Count Nodes Equal to Average of Subtree | M | ready | 543,563 | — |
| 16 | 366 | Find Leaves of Binary Tree | M | ready | 104,543 | — |
| 17 | 606 | Construct String from Binary Tree | M | ready | 257,144 | — |
| 18 | 124 | Binary Tree Maximum Path Sum | H | ready | 543 | — |

### Days 19–26 · Linked List

8 days · E5 M3 H0 · 8 need content generated

| Day | # | Title | Diff | Content | `builds_on` | Folder (new arcs only) |
|---|---|---|---|---|---|---|
| 19 | 206 | Reverse Linked List | E | **new** | 114 | `easy_problems/201_300/reverse_linked_list` |
| 20 | 876 | Middle of the Linked List | E | **new** | — | `easy_problems/801_900/middle_of_the_linked_list` |
| 21 | 141 | Linked List Cycle | E | **new** | 876 | `easy_problems/101_200/linked_list_cycle` |
| 22 | 21 | Merge Two Sorted Lists | E | **new** | — | `easy_problems/1_100/merge_two_sorted_lists` |
| 23 | 234 | Palindrome Linked List | E | **new** | 206,876 | `easy_problems/201_300/palindrome_linked_list` |
| 24 | 19 | Remove Nth Node From End of List | M | **new** | 876 | `medium_problems/1_100/remove_nth_node_from_end_of_list` |
| 25 | 92 | Reverse Linked List II | M | **new** | 206 | `medium_problems/1_100/reverse_linked_list_ii` |
| 26 | 143 | Reorder List | M | **new** | 206,876,21 | `medium_problems/101_200/reorder_list` |

### Days 27–35 · Tree — BFS / level order

9 days · E3 M6 H0 · 0 need content generated

| Day | # | Title | Diff | Content | `builds_on` | Folder (new arcs only) |
|---|---|---|---|---|---|---|
| 27 | 637 | Average of Levels in Binary Tree | E | ready | — | — |
| 28 | 993 | Cousins in Binary Tree | E | ready | 637 | — |
| 29 | 102 | Binary Tree Level Order Traversal | M | ready | 637 | — |
| 30 | 103 | Binary Tree Zigzag Level Order Traversal | M | ready | 102,104 | — |
| 31 | 199 | Binary Tree Right Side View | M | ready | 102 | — |
| 32 | 872 | Leaf-Similar Trees | E | ready | 257 | — |
| 33 | 1161 | Maximum Level Sum of a Binary Tree | M | ready | 102,637,103 | — |
| 34 | 1302 | Deepest Leaves Sum | M | ready | 104,637 | — |
| 35 | 116 | Populating Next Right Pointers in Each Node | M | ready | 102,103 | — |

### Days 36–44 · Graph BFS/DFS

9 days · E1 M8 H0 · 9 need content generated

| Day | # | Title | Diff | Content | `builds_on` | Folder (new arcs only) |
|---|---|---|---|---|---|---|
| 36 | 1971 | Find if Path Exists in Graph | E | **new** | — | `easy_problems/1901_2000/find_if_path_exists_in_graph` |
| 37 | 200 | Number of Islands | M | **new** | 102 | `medium_problems/201_300/number_of_islands` |
| 38 | 695 | Max Area of Island | M | **new** | 200 | `medium_problems/601_700/max_area_of_island` |
| 39 | 547 | Number of Provinces | M | **new** | 200 | `medium_problems/501_600/number_of_provinces` |
| 40 | 994 | Rotting Oranges | M | **new** | 102,200 | `medium_problems/901_1000/rotting_oranges` |
| 41 | 542 | 01 Matrix | M | **new** | 994 | `medium_problems/501_600/01_matrix` |
| 42 | 261 | Graph Valid Tree | M | **new** | 104,200 | `medium_problems/201_300/graph_valid_tree` |
| 43 | 133 | Clone Graph | M | **new** | 200 | `medium_problems/101_200/clone_graph` |
| 44 | 207 | Course Schedule | M | **new** | 133 | `medium_problems/201_300/course_schedule` |

### Days 45–57 · Tree — BST

13 days · E6 M6 H1 · 0 need content generated

| Day | # | Title | Diff | Content | `builds_on` | Folder (new arcs only) |
|---|---|---|---|---|---|---|
| 45 | 700 | Search in a Binary Search Tree | E | ready | 108 | — |
| 46 | 501 | Find Mode in Binary Search Tree | E | ready | — | — |
| 47 | 530 | Minimum Absolute Difference in BST | E | ready | 501 | — |
| 48 | 235 | Lowest Common Ancestor of a Binary Search Tree | M | ready | 700 | — |
| 49 | 783 | Minimum Distance Between BST Nodes | E | ready | 530 | — |
| 50 | 236 | Lowest Common Ancestor of a Binary Tree | M | ready | 235 | — |
| 51 | 653 | Two Sum IV - Input is a BST | E | ready | 700 | — |
| 52 | 98 | Validate Binary Search Tree | M | ready | 501,530 | — |
| 53 | 897 | Increasing Order Search Tree | E | ready | 501 | — |
| 54 | 1038 | Binary Search Tree to Greater Sum Tree | M | ready | 530,783 | — |
| 55 | 173 | Binary Search Tree Iterator | M | ready | 501,530 | — |
| 56 | 450 | Delete Node in a BST | M | ready | 700,98 | — |
| 57 | 272 | Closest Binary Search Tree Value II | H | ready | 173 | — |

### Days 58–64 · Stack & Iterator

7 days · E2 M5 H0 · 7 need content generated

| Day | # | Title | Diff | Content | `builds_on` | Folder (new arcs only) |
|---|---|---|---|---|---|---|
| 58 | 20 | Valid Parentheses | E | **new** | — | `easy_problems/1_100/valid_parentheses` |
| 59 | 496 | Next Greater Element I | E | **new** | 20 | `easy_problems/401_500/next_greater_element_i` |
| 60 | 155 | Min Stack | M | **new** | 20 | `medium_problems/101_200/min_stack` |
| 61 | 739 | Daily Temperatures | M | **new** | 496 | `medium_problems/701_800/daily_temperatures` |
| 62 | 150 | Evaluate Reverse Polish Notation | M | **new** | 20 | `medium_problems/101_200/evaluate_reverse_polish_notation` |
| 63 | 341 | Flatten Nested List Iterator | M | **new** | 173 | `medium_problems/301_400/flatten_nested_list_iterator` |
| 64 | 636 | Exclusive Time of Functions | M | **new** | 155 | `medium_problems/601_700/exclusive_time_of_functions` |

### Days 65–74 · Tree — structure & compare

10 days · E6 M4 H0 · 0 need content generated

| Day | # | Title | Diff | Content | `builds_on` | Folder (new arcs only) |
|---|---|---|---|---|---|---|
| 65 | 965 | Univalued Binary Tree | E | ready | — | — |
| 66 | 617 | Merge Two Binary Trees | E | ready | — | — |
| 67 | 1379 | Find a Corresponding Node of a Binary Tree in a Clone of That Tree | E | ready | 617 | — |
| 68 | 572 | Subtree of Another Tree | E | ready | 617 | — |
| 69 | 2331 | Evaluate Boolean Binary Tree | E | ready | 145 | — |
| 70 | 671 | Second Minimum Node In a Binary Tree | E | ready | — | — |
| 71 | 814 | Binary Tree Pruning | M | ready | 145,2331 | — |
| 72 | 2415 | Reverse Odd Levels of Binary Tree | M | ready | 102,103 | — |
| 73 | 156 | Binary Tree Upside Down | M | ready | 897 | — |
| 74 | 314 | Binary Tree Vertical Order Traversal | M | ready | 102 | — |

### Days 75–81 · Heap

7 days · E2 M4 H1 · 6 need content generated

| Day | # | Title | Diff | Content | `builds_on` | Folder (new arcs only) |
|---|---|---|---|---|---|---|
| 75 | 703 | Kth Largest Element in a Stream | E | ready | 700 | — |
| 76 | 1046 | Last Stone Weight | E | **new** | 703 | `easy_problems/1001_1100/last_stone_weight` |
| 77 | 215 | Kth Largest Element in an Array | M | **new** | 703 | `medium_problems/201_300/kth_largest_element_in_an_array` |
| 78 | 347 | Top K Frequent Elements | M | **new** | 215 | `medium_problems/301_400/top_k_frequent_elements` |
| 79 | 692 | Top K Frequent Words | M | **new** | 347 | `medium_problems/601_700/top_k_frequent_words` |
| 80 | 23 | Merge k Sorted Lists | H | **new** | 21,703 | `hard_problems/1_100/merge_k_sorted_lists` |
| 81 | 264 | Ugly Number II | M | **new** | 703 | `medium_problems/201_300/ugly_number_ii` |
## Content generation schedule

Thirty questions need content written. The queue tolerates a day reserved without content —
`Queue.NextUnposted` gates on `HasContent()` — but it **skips past** such a day rather than
stopping at it, so an unwritten day 20 means day 21 publishes first and the reader sees
"Day 21/365" before "Day 20/365".

Content is therefore generated **one full arc ahead**, in a single sitting per arc:

| Batch | Questions | Must be done before |
|---|---|---|
| Linked List | 8 | day 19 |
| Graph | 9 | day 36 |
| Stack & Iterator | 7 | day 58 |
| Heap | 6 | day 76 |

The linked-list batch is the tight one: eight questions, roughly eight days out from day 11.

## `queue-renumber`

A new `leetcodectl` subcommand. There is no renumber operation today — `queue-append` is the
only queue mutation — and the queue will need reshaping again over 365 days, so this is built
as a tool rather than a one-off script.

The day number is not private to `queue.yaml`. It is copied into published-facing artifacts,
so moving day N to day M means all of:

- rewriting `Day <N>/365` and `![Day <N>](HERO.png)` across `POST_LINKEDIN.md`,
  `POST_DISCORD.md`, `POST_X.md`, `POST_SUBSTACK.md` and `POST_LINKEDIN_ARTICLE.md`
- re-rendering `HERO.png`, since the day number is printed on the card and the background
  palette is derived from the day
- repairing `[Day <N>: <title>](...)` cross-references, which live in the *referring*
  question's files rather than the moved one

### Input

JSON: the full intended day → entry mapping. Entries already in the queue are matched by
LeetCode number; entries not yet in the queue are inserted with their folder and
`builds_on`. Insertion and renumbering are one operation, so the queue is never left in a
half-reordered state.

### Validation, before anything is written

The command refuses the whole run — writing nothing — if any of these hold:

1. A day whose `posted_at` carries any non-null timestamp would move. Those day numbers are
   public on LinkedIn and Discord.
2. Any `builds_on` entry would land on the same day or later than the entry citing it.
3. A target folder does not exist on disk.
4. The mapping has a duplicate day or a duplicate question number.

### Execution order

File work first, `queue.yaml` last, so a failure partway through leaves the queue describing
the state the files were in before the run rather than a state that never existed.

1. rewrite the five `POST_*.md` files for each moved day
2. repair cross-references repo-wide, resolved by target folder rather than by day number
3. re-render `HERO.png` for each moved day
4. write `queue.yaml`

### Dry run

`"dryRun": true` reports the full diff — days moved, files that would change, heroes that
would re-render — and touches nothing.

### Tests

- refuses to move a posted day
- refuses a mapping that creates a forward `builds_on` reference
- refuses a missing folder
- repairs a cross-reference in a folder other than the moved one
- re-running the same mapping is a no-op
- dry run writes nothing

## Two repo defects found while planning

Both touch folders this plan schedules, so both are fixed as part of the work.

**`number_folder_map.yaml` maps question 150 to the wrong folder.** It records
`150: medium_problems/701_800/daily_temperature`. Question 150 is Evaluate Reverse Polish
Notation, whose folder is `medium_problems/101_200/evaluate_reverse_polish_notation`;
Daily Temperatures is 739. Both questions are scheduled (days 62 and 61), so this must be
corrected before content generation resolves either one.

**Seven duplicate solution folders**, both copies holding real code:

| Pair | Note |
|---|---|
| `701_800/daily_temperature` / `daily_temperatures` | 739; scheduled day 61 |
| `1_100/3_sum` / `3sum` | not scheduled |
| `101_200/number_of_islands` / `201_300/number_of_islands` | 200; scheduled day 37 |
| `601_700/search_in_a_binary_search_tree` / `701_800/...` | 700; **601_700 already holds the generated content** for day 45, 701_800 is stale |
| `1_100/pow_x_n` / `powx_n` | not scheduled |
| `1_100/sqrt_x` / `sqrtx` | not scheduled |
| `2501_2600/merge_two_2_d_arrays...` / `...2d...` | not scheduled |

Only the three scheduled pairs need resolving before this plan runs. The rest are noted so
the next person does not rediscover them.

## Out of scope, recorded here so they are not lost

These came up in the same conversation and were deliberately deferred. Each needs its own
design:

- **Instrumentation.** Nothing in this repo records post reach, impressions or follower
  delta. Every content decision is currently made blind, including this one. This is the
  highest-value missing piece.
- **LLD/HLD content track.** The stated goal is to be known for LLD/HLD and problem solving.
  That is also the material that converts to a paid course, and it has no plan yet.
- **Discord real-world-problem series** — brute force, then the same problem with the right
  data structure, measured. Strong format; deferred as LinkedIn posts first, because a
  Discord server at current audience size is likely to read as empty.
- **Weekend problem-solving marathon.** Deferred for the same reason, more strongly.
- **Agent fleet and management UI.** Deferred indefinitely. It is infrastructure for a
  content operation whose constraint is distribution, not tooling.
