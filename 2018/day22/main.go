package main

import (
	"container/heap"
	"fmt"
	"strings"

	"advent2018/util"
)

func main() {
	data := util.GetDataByRow("day22")
	// data := util.GetTestByRow("day22")

	target, depth := parse(data)

	part1(target, depth)
	part2(target, depth)
}

type Point struct {
	x int
	y int
}

type Region = int

const (
	Rocky Region = iota
	Wet
	Narrow
)

func parse(data []string) (target Point, depth int) {
	_, d, _ := strings.Cut(data[0], ":")
	depth = util.ToInt(d)
	_, p, _ := strings.Cut(data[1], ":")
	parts := util.StringToIntSlice(p, ",")
	target = Point{
		parts[0],
		parts[1],
	}
	return
}

func getRegionMap(target Point, depth int, xBound int, yBound int) map[Point]int {
	erosion := map[Point]int{}
	regions := map[Point]Region{}

	for y := range yBound {
		for x := range xBound {
			var geoIndex int
			if y+x == 0 {
				geoIndex = 0
			} else if target.x == x && target.y == y {
				geoIndex = 0
			} else if x == 0 {
				geoIndex = y * 48271
			} else if y == 0 {
				geoIndex = x * 16807
			} else {
				geoIndex = erosion[Point{x, y - 1}] * erosion[Point{x - 1, y}]
			}

			p := Point{x, y}
			erosion[p] = (geoIndex + depth) % 20183
			regions[p] = erosion[p] % 3
		}
	}
	return regions
}

func part1(target Point, depth int) {
	counter := 0

	regions := getRegionMap(target, depth, target.x+1, target.y+1)

	for _, r := range regions {
		counter += r
	}

	fmt.Printf("Part 1: %d\n", counter)
}

type Equipment = int

const (
	None Equipment = iota
	Torch
	Climbing
)

func getNear(p Point) []Point {
	ps := []Point{}

	if p.x > 0 {
		ps = append(ps, Point{p.x - 1, p.y})
	}
	if p.y > 0 {
		ps = append(ps, Point{p.x, p.y - 1})
	}
	ps = append(ps, Point{p.x + 1, p.y})
	ps = append(ps, Point{p.x, p.y + 1})

	return ps
}

func canEnter(r Region, e Equipment) bool {
	switch r {
	case Rocky:
		return e != None
	case Wet:
		return e != Torch
	case Narrow:
		return e != Climbing
	}
	return false
}

func part2(target Point, depth int) {
	regions := getRegionMap(target, depth, target.x*3, target.y*3)

	type StateKey struct {
		p     Point
		equip Equipment
	}

	targetKey := StateKey{target, Torch}
	distances := map[StateKey]int{}

	queue := make(StateQueue, 0)
	heap.Init(&queue)

	push := func(p Point, e Equipment, minutes int) {
		k := StateKey{p, e}
		if d, ok := distances[k]; !ok || minutes < d {
			distances[k] = minutes
			heap.Push(&queue, &State{p: p, equip: e, minutes: minutes})
		}
	}

	push(Point{0, 0}, Torch, 0)

	for queue.Len() > 0 {
		s := heap.Pop(&queue).(*State)

		sk := StateKey{s.p, s.equip}
		if distances[sk] < s.minutes {
			continue
		}
		if sk == targetKey {
			break
		}

		currentRegion := regions[s.p]

		for _, p := range getNear(s.p) {
			r, exists := regions[p]
			if !exists || !canEnter(r, s.equip) {
				continue
			}
			push(p, s.equip, s.minutes+1)
		}

		for _, e := range []Equipment{None, Torch, Climbing} {
			if e != s.equip && canEnter(currentRegion, e) {
				push(s.p, e, s.minutes+7)
			}
		}
	}

	fmt.Printf("Part 2: %d\n", distances[targetKey])
}
