package designauthenticationmanager

import "container/heap"

type token struct {
	id             string
	expirationTime int
}
type minHeap []token

func (mHeap *minHeap) Len() int {
	return len(*mHeap)
}

func (mHeap *minHeap) Less(i, j int) bool {
	return (*mHeap)[i].expirationTime < (*mHeap)[j].expirationTime
}

func (mHeap *minHeap) Swap(i, j int) {
	(*mHeap)[i], (*mHeap)[j] = (*mHeap)[j], (*mHeap)[i]
}

func (h *minHeap) Push(val any) {
	*h = append(*h, val.(token))
}
func (h *minHeap) Pop() any {
	old := *h
	n := old.Len()
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type AuthenticationManager struct {
	tokenMap            map[string]int
	tokenExpriationHeap *minHeap
	ttl                 int
}

func Constructor(timeToLive int) AuthenticationManager {
	minHeapObj := &minHeap{}
	heap.Init(minHeapObj)
	return AuthenticationManager{
		tokenMap:            make(map[string]int),
		tokenExpriationHeap: minHeapObj,
		ttl:                 timeToLive,
	}
}

func (this *AuthenticationManager) clearOldToken(currentTime int) {
	for this.tokenExpriationHeap.Len() > 0 {
		x := heap.Pop(this.tokenExpriationHeap).(token)

		if x.expirationTime > currentTime {
			heap.Push(this.tokenExpriationHeap, x)
			break
		}

		if this.tokenMap[x.id] == x.expirationTime {
			delete(this.tokenMap, x.id)
		}
	}
}
func (this *AuthenticationManager) add(tokenId string, currentTime int) {
	heap.Push(this.tokenExpriationHeap, token{
		id:             tokenId,
		expirationTime: currentTime + this.ttl,
	})
	this.tokenMap[tokenId] = currentTime + this.ttl
}
func (this *AuthenticationManager) Generate(tokenId string, currentTime int) {
	this.clearOldToken(currentTime)
	this.add(tokenId, currentTime)
}

func (this *AuthenticationManager) Renew(tokenId string, currentTime int) {
	this.clearOldToken(currentTime)
	if _, ok := this.tokenMap[tokenId]; ok {
		this.add(tokenId, currentTime)
	}
}

func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	this.clearOldToken(currentTime)
	return len(this.tokenMap)
}

/**
 * Your AuthenticationManager object will be instantiated and called as such:
 * obj := Constructor(timeToLive);
 * obj.Generate(tokenId,currentTime);
 * obj.Renew(tokenId,currentTime);
 * param_3 := obj.CountUnexpiredTokens(currentTime);
 */
