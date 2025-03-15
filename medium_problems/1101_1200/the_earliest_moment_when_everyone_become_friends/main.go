package theearliestmomentwheneveryonebecomefriends

import "sort"

type disJointSet struct {
	count int
	root  []int
	rank  []int
}

func InitDisjoinSet(n int) *disJointSet {

	root := make([]int, n)
	rank := make([]int, n)
	for i := range n {
		root[i] = i
		rank[i] = 1
	}

	return &disJointSet{
		count: n,
		root:  root,
		rank:  rank,
	}
}

func (set *disJointSet) Find(x int) int {
	if x == set.root[x] {
		return x
	}
	set.root[x] = set.Find(set.root[x])
	return set.root[x]
}

func (set *disJointSet) Union(a, b int) {
	rootA, rootB := set.Find(a), set.Find(b)
	if rootA == rootB {
		return
	}
	if set.rank[rootA] > set.rank[rootB] {
		set.root[rootB] = rootA
	} else if set.rank[rootA] < set.rank[rootB] {
		set.root[rootA] = rootB
	} else {
		set.root[rootA] = rootB
		set.rank[rootB]++

	}
	set.count--
}
func (set *disJointSet) ComponentCount() int {
	return set.count
}
func earliestAcq(logs [][]int, n int) int {
	set := InitDisjoinSet(n)

	sort.Slice(logs, func(i, j int) bool {
		return logs[i][0] < logs[j][0]
	})
	timestamp := -1
	for _, e := range logs {
		set.Union(e[1], e[2])
		if set.ComponentCount() == 1 {
			if timestamp > e[0] || timestamp == -1 {
				timestamp = e[0]
			}
		}
	}

	return timestamp
}
