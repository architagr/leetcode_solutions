package networkdelaytime

import (
	"container/heap"
	"math"
)

type networkCost struct {
	p1, p2, cost int
}

type minHeap []networkCost

func (h *minHeap) Len() int {
	return len(*h)
}
func (mHeap *minHeap) Less(i, j int) bool {
	return (*mHeap)[i].cost < (*mHeap)[j].cost
}
func (h *minHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *minHeap) Push(x any) {
	*h = append(*h, x.(networkCost))
}
func (h *minHeap) Pop() any {
	old := *h
	n := old.Len()
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func networkDelayTime(times [][]int, n int, k int) int {
	visited := make(map[int]bool, n)
	nodeCost := make(map[int]int, n)

	for i := 1; i <= n; i++ {
		nodeCost[i] = math.MaxInt
	}
	nodeCost[k] = 0

	aj := getAj(times, n)
	mHeap := &minHeap{}
	heap.Init(mHeap)
	processNodeNextPathCost(visited, nodeCost, k, aj, mHeap)

	for mHeap.Len() > 0 {
		networkNode := heap.Pop(mHeap).(networkCost)
		processNodeNextPathCost(visited, nodeCost, networkNode.p2, aj, mHeap)
	}
	if len(visited) < n {
		return -1
	}
	ans := 0
	for _, c := range nodeCost {
		ans = maxValue(ans, c)
	}
	return ans
}
func minValue(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func processNodeNextPathCost(visited map[int]bool, nodeCost map[int]int, k int, aj map[int][]networkCost, mHeap *minHeap) {
	if _, found := visited[k]; found {
		return
	}
	currentMinCost := nodeCost[k]
	for _, n := range aj[k] {
		if _, found := visited[n.p2]; !found {
			nodeCost[n.p2] = minValue(nodeCost[n.p2], currentMinCost+n.cost)
			heap.Push(mHeap, networkCost{
				p1:   n.p1,
				p2:   n.p2,
				cost: currentMinCost + n.cost,
			})
		}
	}
	visited[k] = true
}

func getAj(times [][]int, n int) map[int][]networkCost {
	aj := make(map[int][]networkCost, n)
	for _, t := range times {
		aj[t[0]] = append(aj[t[0]], networkCost{
			p1:   t[0],
			p2:   t[1],
			cost: t[2],
		})
	}
	return aj
}
