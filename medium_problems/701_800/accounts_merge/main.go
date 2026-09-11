package accountsmerge

import "sort"

// accountsMerge groups the accounts that belong to one person. Two
// accounts belong together when they share any email, and that
// relationship is transitive: the first and last account here are the
// same person even though they share nothing directly.
//
//	["David","David0@m.co","David1@m.co"]
//	["David","David1@m.co","David2@m.co"]
//
// A single pass that maps each email to the account it first appeared in
// cannot see that, because the link arrives after both have been placed.
// Union-find can, since merging two groups later re-points both.
func accountsMerge(accounts [][]string) [][]string {
	parent := make([]int, len(accounts))
	for i := range parent {
		parent[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		for parent[x] != x {
			parent[x] = parent[parent[x]] // path halving
			x = parent[x]
		}
		return x
	}
	union := func(a, b int) {
		ra, rb := find(a), find(b)
		if ra != rb {
			parent[rb] = ra
		}
	}

	owner := make(map[string]int, len(accounts)*4)
	for i, account := range accounts {
		for _, email := range account[1:] {
			if j, seen := owner[email]; seen {
				union(i, j)
			} else {
				owner[email] = i
			}
		}
	}

	// Collect each root's emails, deduplicated.
	emails := make(map[int]map[string]bool, len(accounts))
	for email, i := range owner {
		root := find(i)
		if emails[root] == nil {
			emails[root] = make(map[string]bool)
		}
		emails[root][email] = true
	}

	result := make([][]string, 0, len(emails))
	for root, set := range emails {
		merged := make([]string, 0, len(set)+1)
		for email := range set {
			merged = append(merged, email)
		}
		sort.Strings(merged)
		result = append(result, append([]string{accounts[root][0]}, merged...))
	}
	return result
}
