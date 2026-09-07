package main

import (
	"fmt"
	"maps"
	"slices"
	"sort"

	"advent2019/util"
)

func main() {
	data := util.GetDataByRow("day10")
	// data := util.GetTestByRow("day10")

	list := parse(data)

	part1(list)
	part2(list)
}

type Point struct {
	x int
	y int
}

func parse(data []string) []Point {
	ps := []Point{}

	for y, r := range data {
		for x, c := range r {
			if c == '#' {
				ps = append(ps, Point{x, y})
			}
		}
	}
	return ps
}

func getDivisors(n int) []int {
	d := []int{n}

	for i := n / 2; i > 1; i-- {
		if n%i == 0 {
			d = append(d, i)
		}
	}
	return d
}

func getMiddlePoints(a, b Point) []Point {
	if a == b {
		return nil
	}
	ps := []Point{}

	deltaX := util.IntAbs(a.x - b.x)
	deltaY := util.IntAbs(a.y - b.y)

	// vertical
	if a.x == b.x {
		var l, h Point

		if a.y < b.y {
			l = a
			h = b
		} else {
			l = b
			h = a
		}

		for i := 1; i < h.y-l.y; i++ {
			ps = append(ps, Point{l.x, l.y + i})
		}
		return ps
	}

	// horizontal
	if a.y == b.y {
		var l, h Point

		if a.x < b.x {
			l = a
			h = b
		} else {
			l = b
			h = a
		}

		for i := 1; i < h.x-l.x; i++ {
			ps = append(ps, Point{l.x + i, l.y})
		}
		return ps
	}

	// diagonal
	if deltaY == deltaX {
		var l, h Point

		if a.x < b.x {
			l = a
			h = b
		} else {
			l = b
			h = a
		}

		var dY int
		if l.y < h.y {
			dY = 1
		} else {
			dY = -1
		}

		for i, j := 1, dY; i < h.x-l.x; i, j = i+1, j+dY {
			ps = append(ps, Point{l.x + i, l.y + j})
		}
		return ps
	}

	// other
	// if its diagonal, irregoral and less than 2, it cannot be present a "middle" point
	if deltaX < 2 || deltaY < 2 {
		return ps
	}

	// select the lower value (lesser operation)
	lower := deltaX
	higher := deltaY
	if deltaY < deltaX {
		lower = deltaY
		higher = deltaX
	}

	// get the divisors
	divs := getDivisors(lower)

	// check if the divisors are in common with the other delta
	leftPoint := a
	rightPoint := b
	if b.x < a.x {
		leftPoint = b
		rightPoint = a
	}
	dY := 1
	if rightPoint.y < leftPoint.y {
		dY = -1
	}
	for _, d := range divs {
		if higher%d == 0 {
			// if the divisor is in common, i should check all the possible intersections (ie: 0 and 12 -> 3, 6, 9)

			stepX := deltaX / d
			stepY := (deltaY / d) * dY

			for sX, sY := stepX, stepY; sX < deltaX; sX, sY = sX+stepX, sY+stepY {
				ps = append(ps, Point{leftPoint.x + sX, leftPoint.y + sY})
			}
		}
	}
	return ps
}

func arrayContains(a []Point, b []Point) bool {
	for _, p := range a {
		if slices.Contains(b, p) {
			return true
		}
	}
	return false
}

func getVisibilityMap(data []Point) map[Point]int {
	values := map[Point]int{}

	for i := range len(data) - 1 {
		a := data[i]
		for j := i + 1; j < len(data); j++ {
			b := data[j]

			middlePs := getMiddlePoints(a, b)
			if !arrayContains(middlePs, data) {
				values[a] += 1
				values[b] += 1
			}
		}
	}

	return values
}

func getSeenPoints(o Point, data []Point) []Point {
	seen := map[Point]bool{}

	for _, d := range data {
		if d == o {
			continue
		}

		middlePs := getMiddlePoints(o, d)
		if !arrayContains(middlePs, data) {
			seen[d] = true
		}
	}
	return slices.Collect(maps.Keys(seen))
}

func part1(data []Point) {
	counter := 0

	values := getVisibilityMap(data)
	for _, n := range values {
		if n > counter {
			counter = n
		}
	}

	fmt.Printf("Part 1: %d\n", counter)
}

func removeFromArray(orig []Point, toRemove []Point) []Point {
	newPs := make([]Point, 0, len(orig)-len(toRemove))
	for _, p := range orig {
		if !slices.Contains(toRemove, p) {
			newPs = append(newPs, p)
		}
	}
	return newPs
}

func half(x, y int) int {
	if y > 0 || (y == 0 && x >= 0) {
		return 0
	}
	return 1
}

func sortClockwise(center Point, points []Point) {
	sort.Slice(points, func(i, j int) bool {
		ax, ay := -(points[i].y - center.y), points[i].x-center.x
		bx, by := -(points[j].y - center.y), points[j].x-center.x

		ha, hb := half(ax, ay), half(bx, by)
		if ha != hb {
			return ha < hb
		}

		cross := ax*by - ay*bx
		if cross != 0 {
			return cross > 0
		}

		return (ax*ax+ay*ay)-(bx*bx+by*by) < 0
	})
}

func part2(data []Point) {
	values := getVisibilityMap(data)
	var p Point
	c := 0
	for k, n := range values {
		if n > c {
			c = n
			p = k
		}
	}

	deleted := 0
	counter := 0

	for deleted < 200 {
		seen := getSeenPoints(p, data)
		if deleted+len(seen) < 200 {
			data = removeFromArray(data, seen)
			deleted += len(seen)
		} else {
			// TODO: find the 200th clockwise
			sortClockwise(p, seen)
			index := 200 - deleted - 1
			counter = seen[index].x*100 + seen[index].y
			break
		}
	}

	fmt.Printf("Part 2: %d\n", counter)
}
