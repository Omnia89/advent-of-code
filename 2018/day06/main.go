package main

import (
	"fmt"
	"maps"
	"math"
	"slices"
	"strings"

	"advent2018/util"
)

func main() {
	data := util.GetDataByRow("day06")
	//data := util.GetTestByRow("day06")

	list := parse(data)

	part1(list)
	part2(list)
}

type Point struct {
	x int
	y int
}

type Area struct {
	id     string
	center Point
	finite bool
	points []Point
	border []Point
}

func parse(data []string) []Point {
	ps := []Point{}

	for _, s := range data {
		xS, yS, _ := strings.Cut(s, ", ")
		ps = append(ps, Point{
			util.ToInt(xS),
			util.ToInt(yS),
		})
	}

	return ps
}

var (
	Up    = Point{0, -1}
	Down  = Point{0, 1}
	Left  = Point{-1, 0}
	Right = Point{1, 0}
)

func checkBoundaries(p Point, allPoints []Point) bool {
	oks := map[Point]bool{
		Up:    false,
		Down:  false,
		Left:  false,
		Right: false,
	}

	allOk := false
	check := func() bool {
		k := true
		for _, b := range oks {
			k = k && b
		}
		return k
	}

	for _, o := range allPoints {
		if allOk {
			break
		}
		if o == p {
			continue
		}

		dx := util.IntAbs(p.x - o.x)
		dy := util.IntAbs(p.y - o.y)

		if dx < dy {
			if p.y > o.y {
				// blocks Up
				oks[Up] = true
			} else {
				// blocks Down
				oks[Down] = true
			}
		} else if dx > dy {
			if p.x > o.x {
				// blocks Left
				oks[Left] = true
			} else {
				// blocks Right
				oks[Right] = true
			}
		} else {
			// dx == dy

			if o.x < p.x && o.y < p.y {
				// Up-Left
				oks[Up] = true
				oks[Left] = true
			} else if o.x > p.x && o.y < p.y {
				// Up-Right
				oks[Up] = true
				oks[Right] = true
			} else if o.x > p.x && o.y > p.y {
				// Down-Right
				oks[Down] = true
				oks[Right] = true
			} else if o.x < p.x && o.y > p.y {
				// Down-Left
				oks[Down] = true
				oks[Left] = true
			}
		}
		allOk = check()
	}
	return allOk
}

func getNear(p Point) []Point {
	return []Point{
		{p.x, p.y + 1},
		{p.x, p.y - 1},
		{p.x + 1, p.y},
		{p.x - 1, p.y},
	}
}

// Returns the list of finite areas that expanded
func expand(areas map[Point]Area, stalledPoints []Point) (finitedCenters []Point, stalledPoint []Point) {
	finArea := map[Point]struct{}{}

	occupiedPoint := map[Point]struct{}{}
	for _, a := range areas {
		for _, p := range a.points {
			occupiedPoint[p] = struct{}{}
		}
	}
	for _, p := range stalledPoints {
		occupiedPoint[p] = struct{}{}
	}

	// center/id -> new points
	expandedPoints := map[Point]map[Point]struct{}{}

	for c, a := range areas {
		if len(a.border) == 0 {
			continue
		}
		for _, b := range a.border {
			for _, np := range getNear(b) {
				if _, ok := occupiedPoint[np]; !ok {
					if _, k := expandedPoints[np]; !k {
						expandedPoints[np] = map[Point]struct{}{}
					}
					expandedPoints[np][c] = struct{}{}
				}
			}
		}
	}

	newBorders := map[Point][]Point{}

	for p, cs := range expandedPoints {
		if len(cs) > 1 {
			stalledPoint = append(stalledPoint, p)
			continue
		}
		var c Point
		for b := range cs {
			c = b
		}
		a := areas[c]

		a.points = append(a.points, p)
		newBorders[a.center] = append(newBorders[a.center], p)

		areas[c] = a
	}

	for c := range areas {
		a := areas[c]
		a.border = newBorders[c]
		areas[c] = a

		if a.finite && len(a.border) > 0 {
			finArea[a.center] = struct{}{}
		}
	}

	finiteAreas := slices.Collect(maps.Keys(finArea))

	return finiteAreas, stalledPoint
}

type idGen struct {
	curr rune
}

func newIdGen(start rune) *idGen {
	return &idGen{curr: start}
}

func (g *idGen) Next() string {
	char := string(g.curr)
	g.curr++

	return char
}

func draw(areas map[Point]Area) {
	minX, minY := math.MaxInt, math.MaxInt
	maxX, maxY := 0, 0

	for _, a := range areas {
		for _, p := range a.points {
			if p.x < minX {
				minX = p.x
			}
			if p.y < minY {
				minY = p.y
			}
			if p.x > maxX {
				maxX = p.x
			}
			if p.y > maxY {
				maxY = p.y
			}
		}
	}
	grid := make([][]string, maxY-minY+1)
	for y := 0; y < maxY-minY+1; y++ {
		grid[y] = make([]string, maxX-minX+1)
		for x := 0; x < maxX-minX+1; x++ {
			grid[y][x] = "."
		}
	}

	for _, a := range areas {
		for _, p := range a.points {
			grid[p.y-minY][p.x-minX] = a.id
		}
	}

	fmt.Println("-----------------------------------------------")
	for _, r := range grid {
		fmt.Println(strings.Join(r, ""))
	}
	fmt.Println("-----------------------------------------------")
}

func part1(data []Point) {
	counter := 0

	areas := map[Point]Area{}
	finiteAreas := []Point{}
	stalledPoint := []Point{}

	ids := newIdGen('a')

	for _, p := range data {
		a := Area{
			id:     ids.Next(),
			center: p,
		}

		a.finite = checkBoundaries(p, data)
		a.points = []Point{p}
		a.border = []Point{p}
		areas[p] = a
		if a.finite {
			finiteAreas = append(finiteAreas, p)
		}
	}
	for len(finiteAreas) > 0 {
		finiteAreas, stalledPoint = expand(areas, stalledPoint)
		// draw(areas)
	}

	for _, a := range areas {
		if !a.finite {
			continue
		}
		ar := len(a.points)
		if ar > counter {
			counter = ar
		}
	}

	fmt.Printf("Part 1: %d\n", counter)
}

func checkDistance(p Point, all []Point, limit int) bool {
	s := 0
	for _, o := range all {
		s += util.IntAbs(p.x-o.x) + util.IntAbs(p.y-o.y)
	}
	return s < limit
}

func part2(points []Point) {
	counter := 0

	minX, minY := math.MaxInt, math.MaxInt
	maxX, maxY := 0, 0

	for _, p := range points {
		if p.x < minX {
			minX = p.x
		}
		if p.y < minY {
			minY = p.y
		}
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
	}

	queue := []Point{
		{(maxX + minX) / 2, (maxY + minY) / 2},
	}
	var current Point

	if !checkDistance(queue[0], points, 10000) {
		panic("find another start point")
	}
	counter++

	done := map[Point]struct{}{queue[0]: {}}

	for len(queue) > 0 {
		current, queue = queue[0], queue[1:]

		for _, np := range getNear(current) {
			if _, ok := done[np]; !ok {
				done[np] = struct{}{}
				if checkDistance(np, points, 10000) {
					counter++
					queue = append(queue, np)
				}
			}
		}
	}

	fmt.Printf("Part 2: %d\n", counter)
}
