package main

import (
	"fmt"

	"advent2018/util"
)

func main() {
	data := util.GetDataByRow("day25")
	// data := util.GetTestByRow("day25")

	list := parse(data)

	part1(list)
	part2(list)
}

type Point struct {
	x int
	y int
	z int
	t int
}

func (p Point) distance(o Point) int {
	return util.IntAbs(p.x-o.x) + util.IntAbs(p.y-o.y) + util.IntAbs(p.z-o.z) + util.IntAbs(p.t-o.t)
}

func parse(data []string) []Point {
	ps := []Point{}

	for _, s := range data {
		ns := util.StringToIntSlice(s, ",")
		ps = append(ps, Point{
			ns[0],
			ns[1],
			ns[2],
			ns[3],
		})
	}

	return ps
}

type Union struct {
	parent []int
	rank   []int
}

// Get the index of the parent of the group of `x`
func (u *Union) find(x int) int {
	if u.parent[x] != x {
		// path compression
		u.parent[x] = u.find(u.parent[x])
	}
	return u.parent[x]
}

func (u *Union) unite(a, b int) {
	pA := u.find(a)
	pB := u.find(b)

	if pA == pB {
		return
	}

	if u.rank[pA] < u.rank[pB] {
		u.parent[pA] = pB
	} else if u.rank[pA] > u.rank[pB] {
		u.parent[pB] = pA
	} else {
		u.parent[pB] = pA
		u.rank[pA]++
	}
}

// 439 too high
func part1(data []Point) {
	counter := 0

	union := &Union{
		parent: make([]int, len(data)),
		rank:   make([]int, len(data)),
	}

	for i := range union.parent {
		union.parent[i] = i
	}

	for i := range data {
		for j := i + 1; j < len(data); j++ {
			if data[i].distance(data[j]) <= 3 {
				union.unite(i, j)
			}
		}
	}

	constellations := map[int][]Point{}

	for i, p := range data {
		root := union.find(i)
		constellations[root] = append(constellations[root], p)
	}

	counter = len(constellations)

	// for ic, c := range constellations {
	// 	fmt.Printf("----------\n[%d]\n", ic)
	// 	for _, p := range c {
	// 		fmt.Printf("\t[%v]\n", p)
	// 	}
	// }

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(data []Point) {
	counter := 0

	fmt.Printf("Part 2: %d\n", counter)
}
