package main

type PointQueue []*Point

func (pq PointQueue) Len() int {
	return len(pq)
}

func (pq PointQueue) Less(i, j int) bool {
	// Higher y
	if pq[i].y != pq[j].y {
		return pq[i].y > pq[j].y
	}
	// Lesser x
	return pq[i].x < pq[j].x
}

func (pq PointQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PointQueue) Push(x interface{}) {
	// TODO: check if using the Point and make the pointer
	item := x.(*Point)
	*pq = append(*pq, item)
}

func (pq *PointQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}
