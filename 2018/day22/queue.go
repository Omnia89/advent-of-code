package main

type State struct {
	p       Point
	equip   Equipment
	minutes int
}

type StateQueue []*State

func (sq StateQueue) Len() int {
	return len(sq)
}

func (sq StateQueue) Less(i, j int) bool {
	return sq[i].minutes < sq[j].minutes
}

func (sq StateQueue) Swap(i, j int) {
	sq[i], sq[j] = sq[j], sq[i]
}

func (sq *StateQueue) Push(x interface{}) {
	item := x.(*State)
	*sq = append(*sq, item)
}

func (sq *StateQueue) Pop() interface{} {
	old := *sq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*sq = old[0 : n-1]
	return item
}
