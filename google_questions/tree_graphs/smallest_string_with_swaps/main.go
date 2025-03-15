package smalleststringwithswaps

import (
	"sort"
)

func smallestStringWithSwaps(s string, pairs [][]int) string {

	set := InitDisjoinSet(len(s))
	for _, p := range pairs {
		set.Union(p[0], p[1])
	}
	rootToComponent := make(map[int][]int)
	for i := range len(s) {
		r := set.Find(i)
		rootToComponent[r] = append(rootToComponent[r], i)
	}
	sb := make([]byte, len(s))

	for _, vertexList := range rootToComponent {
		characters := make([]byte, len(vertexList))
		for i, v := range vertexList {
			characters[i] = s[v]
		}
		sort.Slice(characters, func(i, j int) bool {
			return characters[i] < characters[j]
		})

		for index := 0; index < len(characters); index++ {
			sb[vertexList[index]] = characters[index]
		}
	}

	return string(sb)
}

type DisJointSet struct {
	root  []int
	rank  []int
	count int
}

func InitDisjoinSet(n int) *DisJointSet {
	root := make([]int, n)
	rank := make([]int, n)

	for i := range n {
		root[i] = i
		rank[i] = 1
	}
	return &DisJointSet{
		count: n,
		root:  root,
		rank:  rank,
	}

}

func (set *DisJointSet) Find(node int) int {
	if node == set.root[node] {
		return node
	}
	set.root[node] = set.Find(set.root[node])
	return set.root[node]
}
func (set *DisJointSet) Union(node1, node2 int) {
	root1 := set.Find(node1)
	root2 := set.Find(node2)

	if root1 == root2 {
		return
	}
	if set.rank[root1] >= set.rank[root2] {
		set.root[root2] = root1
		set.rank[root1] += set.rank[root2]
	} else {
		set.root[root1] = root2
		set.rank[root2] += set.rank[root1]
	}
	set.count--
}
