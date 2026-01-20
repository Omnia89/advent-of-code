package main

import (
	"fmt"
	"slices"

	"advent2024/util"
)

func main() {
	// data := util.GetTestByRow("day12")
	data := util.GetDataByRow("day12")

	list := parse(data)

	part1(list)
	part2(list, data)
}

type Area struct {
	Key             string
	Area            int
	Peri            int
	PerimeterPoints []Point
	Points          []Point
}

type Point struct {
	X int
	Y int
}

func (p Point) ToString() string {
	return fmt.Sprintf("%d_%d", p.X, p.Y)
}

func (p Point) GetValue(data []string) string {
	return string(data[p.Y][p.X])
}

func parse(data []string) []Area {
	alreadyDone := map[string]bool{}

	areas := []Area{}

	for y := range data {
		for x := range data[0] {
			p := Point{x, y}
			if done := alreadyDone[p.ToString()]; done {
				continue
			}

			areaKey := p.GetValue(data)

			areaPoints := []Point{}
			queue := []Point{p}
			var n Point

			compare := func(p Point) func(t Point) bool {
				return func(t Point) bool {
					return t.Y == p.Y && t.X == p.X
				}
			}

			for len(queue) > 0 {
				n, queue = queue[0], queue[1:]

				if n.GetValue(data) != areaKey || slices.ContainsFunc(areaPoints, compare(n)) {
					continue
				}
				areaPoints = append(areaPoints, n)

				if n.Y > 0 {
					tp := Point{n.X, n.Y - 1}
					if !slices.ContainsFunc(areaPoints, compare(tp)) {
						queue = append(queue, tp)
					}
				}
				if n.X > 0 {
					tp := Point{n.X - 1, n.Y}
					if !slices.ContainsFunc(areaPoints, compare(tp)) {
						queue = append(queue, tp)
					}
				}
				if n.Y < len(data)-1 {
					tp := Point{n.X, n.Y + 1}
					if !slices.ContainsFunc(areaPoints, compare(tp)) {
						queue = append(queue, tp)
					}
				}
				if n.X < len(data[0])-1 {
					tp := Point{n.X + 1, n.Y}
					if !slices.ContainsFunc(areaPoints, compare(tp)) {
						queue = append(queue, tp)
					}
				}
			}

			area := Area{areaKey, 0, 0, []Point{}, areaPoints}
			for _, pp := range areaPoints {
				alreadyDone[pp.ToString()] = true

				area.Area += 1
				peri := 0

				// up
				if pp.Y == 0 || string(data[pp.Y-1][pp.X]) != areaKey {
					peri++
				}

				// down
				if pp.Y == len(data)-1 || string(data[pp.Y+1][pp.X]) != areaKey {
					peri++
				}

				// left
				if pp.X == 0 || string(data[pp.Y][pp.X-1]) != areaKey {
					peri++
				}

				// right
				if pp.X == len(data[0])-1 || string(data[pp.Y][pp.X+1]) != areaKey {
					peri++
				}
				if peri > 0 {
					area.Peri += peri
					area.PerimeterPoints = append(area.PerimeterPoints, pp)
				}
			}
			areas = append(areas, area)
		}
	}
	return areas
}

func part1(areas []Area) {
	sum := 0

	for _, a := range areas {
		sum += a.Peri * a.Area
	}

	fmt.Printf("Part 1: %d\n", sum)
}

func orderPerimeter(points []Point) []Point {
	ordered := make([]Point, 0, len(points))

	directions := []Point{
		{0, -1},
		{1, -1},
		{1, 0},
		{1, 1},
		{0, 1},
		{-1, 1},
		{-1, 0},
		{-1, -1},
	}

	pointMap := map[Point]bool{}

	// Select upper-left
	startIndex := -1
	for i, p := range points {
		pointMap[p] = true
		if startIndex == -1 || p.Y < points[startIndex].Y || (p.Y == points[startIndex].Y && p.X < points[startIndex].X) {
			startIndex = i
		}
	}

	curr := points[startIndex]
	direction := 7

	for {
		ordered = append(ordered, curr)
		found := false

		for i := range 8 {
			nextDirection := (direction + i) % 8
			move := directions[nextDirection]

			nextPoint := Point{curr.X + move.X, curr.Y + move.Y}
			if pointMap[nextPoint] {
				curr = nextPoint
				// Start from the next direction clockwise from the relative position of old point.
				// So the last point will be visited last
				direction = (nextDirection + 5) % 8
				found = true
				break
			}
		}
		if !found || points[startIndex] == curr {
			// Non ci sono punti o sono arrivato all'inizio
			break
		}

		// Safety break
		if len(ordered) > len(points)*2 {
			break
		}
	}

	return ordered
}

func countSides(area Area, height int, width int) int {
	sides := 0

	areaMap := map[Point]bool{}
	for _, a := range area.Points {
		areaMap[a] = true
	}

	// up and down sides
	for y := range height {
		insideSideUp := false
		insideSideDown := false

		for x := range width {
			p := Point{x, y}
			if areaMap[p] {
				// check upper side
				upper := Point{p.X, p.Y - 1}
				if !insideSideUp && (p.Y == 0 || !areaMap[upper]) {
					insideSideUp = true
					sides++
				}
				// Check bottom side
				bottom := Point{p.X, p.Y + 1}
				if !insideSideDown && (p.Y == height || !areaMap[bottom]) {
					insideSideDown = true
					sides++
				}

				if areaMap[upper] {
					insideSideUp = false
				}
				if areaMap[bottom] {
					insideSideDown = false
				}
			} else {
				insideSideDown = false
				insideSideUp = false
			}
		}
	}

	// left and right sides
	for x := range width {
		insideSideLeft := false
		insideSideRight := false

		for y := range height {
			p := Point{x, y}
			if areaMap[p] {
				// check left side
				left := Point{p.X - 1, p.Y}
				if !insideSideLeft && (p.X == 0 || !areaMap[left]) {
					insideSideLeft = true
					sides++
				}
				// Check right side
				right := Point{p.X + 1, p.Y}
				if !insideSideRight && (p.X == width || !areaMap[right]) {
					insideSideRight = true
					sides++
				}

				if areaMap[left] {
					insideSideLeft = false
				}
				if areaMap[right] {
					insideSideRight = false
				}
			} else {
				insideSideLeft = false
				insideSideRight = false
			}
		}
	}

	return sides
}

func part2(areas []Area, raw []string) {
	sum := 0

	for _, a := range areas {
		// ordered := orderPerimeter(a.PerimeterPoints)
		sides := countSides(a, len(raw), len(raw[0]))
		sum += sides * a.Area
	}

	fmt.Printf("Part 2: %d\n", sum)
}
