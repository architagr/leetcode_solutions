package numberofconnectedcomponentsinanundirectedgraph

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
	if set.rank[root1] > set.rank[root2] {
		set.root[root2] = root1
	} else if set.rank[root1] < set.rank[root2] {
		set.root[root1] = root2
	} else {
		set.root[root1] = root2
		set.rank[root2]++
	}
	set.count--
}

func (set *DisJointSet) CountComponents() int {
	return set.count
}
func countComponents(n int, edges [][]int) int {
	set := InitDisjoinSet(n)

	for _, e := range edges {
		set.Union(e[0], e[1])
	}
	return set.CountComponents()
}
