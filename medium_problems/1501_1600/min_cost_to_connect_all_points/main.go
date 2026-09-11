package mincosttoconnectallpairs

import "container/heap"

type edge struct {
	p1, p2, cost int
}

type minHeap []edge

func (mHeap *minHeap) Len() int {
	return len(*mHeap)
}

func (mHeap *minHeap) Less(i, j int) bool {
	return (*mHeap)[i].cost < (*mHeap)[j].cost
}

func (mHeap *minHeap) Swap(i, j int) {
	(*mHeap)[i], (*mHeap)[j] = (*mHeap)[j], (*mHeap)[i]
}

func (h *minHeap) Push(val any) {
	*h = append(*h, val.(edge))
}
func (h *minHeap) Pop() any {
	old := *h
	n := old.Len()
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func minCostConnectPoints(points [][]int) int {

	n := len(points)
	pq := &minHeap{}
	heap.Init(pq)

	edgeCount := n - 1
	visited := make([]bool, n)
	result := 0
	for i := 1; i < n; i++ {
		heap.Push(pq, edge{p1: 0, p2: i, cost: calcDisctance(points[0], points[i])})
	}
	visited[0] = true
	for edgeCount > 0 && pq.Len() > 0 {
		x := heap.Pop(pq).(edge)
		if !visited[x.p2] {
			result += x.cost
			visited[x.p2] = true
			edgeCount--
			for i := 0; i < n; i++ {
				if !visited[i] {
					heap.Push(pq, edge{p1: x.p2, p2: i, cost: calcDisctance(points[x.p2], points[i])})
				}
			}
		}
	}
	return result
}

func calcDisctance(p1, p2 []int) int {
	return adsValue(p1[0]-p2[0]) + adsValue(p1[1]-p2[1])
}

func adsValue(a int) int {
	if a < 0 {
		a *= -1
	}
	return a
}
