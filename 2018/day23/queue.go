package main

type Box struct {
	xMin int
	xMax int
	yMin int
	yMax int
	zMin int
	zMax int

	points  int
	minDist int
}

func (b Box) volume() int {
	return (b.xMax - b.xMin + 1) * (b.yMax - b.yMin + 1) * (b.zMax - b.zMin + 1)
}

type BoxQueue []*Box

func (bq BoxQueue) Len() int {
	return len(bq)
}

func (bq BoxQueue) Less(i, j int) bool {
	if bq[i].points == bq[j].points {
		return bq[i].minDist < bq[j].minDist
	}
	return bq[i].points > bq[j].points
}

func (bq BoxQueue) Swap(i, j int) {
	bq[i], bq[j] = bq[j], bq[i]
}

func (bq *BoxQueue) Push(x any) {
	item := x.(*Box)
	*bq = append(*bq, item)
}

func (bq *BoxQueue) Pop() any {
	old := *bq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*bq = old[0 : n-1]
	return item
}
