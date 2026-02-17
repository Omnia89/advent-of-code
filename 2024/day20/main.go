package main

import (
	"fmt"
	"strings"

	"advent2024/util"
)

func main() {
	// data := util.GetTestByRow("day20")
	data := util.GetDataByRow("day20") // p2: 9432 too low

	part1(data)
	part2(data)
}

// TODO: inizio logica pt2
//

type Point struct {
	X int
	Y int
}

func (p Point) toString() string {
	return fmt.Sprintf("%d,%d", p.X, p.Y)
}

func locate(data []string, char string) Point {
	for y, r := range data {
		if x := strings.Index(r, char); x != -1 {
			return Point{x, y}
		}
	}
	return Point{}
}

func isPath(r byte) bool {
	return r == '.' || r == 'E'
}

func track(data []string, start Point, end Point) (map[Point]int, []Point) {
	count := 0

	pointMap := map[Point]int{}
	pointMap[start] = 0
	points := []Point{}

	p := start
	for {
		points = append(points, p)
		count++

		if p == end {
			break
		}

		// up
		temp := Point{p.X, p.Y - 1}
		if _, ok := pointMap[temp]; !ok && isPath(data[temp.Y][temp.X]) {
			pointMap[temp] = count
			p = temp
			continue
		}

		// down
		temp = Point{p.X, p.Y + 1}
		if _, ok := pointMap[temp]; !ok && isPath(data[temp.Y][temp.X]) {
			pointMap[temp] = count
			p = temp
			continue
		}

		// left
		temp = Point{p.X - 1, p.Y}
		if _, ok := pointMap[temp]; !ok && isPath(data[temp.Y][temp.X]) {
			pointMap[temp] = count
			p = temp
			continue
		}

		// right
		temp = Point{p.X + 1, p.Y}
		if _, ok := pointMap[temp]; !ok && isPath(data[temp.Y][temp.X]) {
			pointMap[temp] = count
			p = temp
			continue
		}
	}
	return pointMap, points
}

func getGain(begin Point, end Point, track map[Point]int) int {
	valB, okB := track[begin]
	valE, okE := track[end]

	if !okB || !okE || valE < valB {
		return 0
	}

	// return valE - valB - 1 - 2
	return valE - valB - 2
}

func part1(data []string) {
	counter := 0

	limit := 100

	start := locate(data, "S")
	end := locate(data, "E")
	// fmt.Printf("  start [%s]\n", start.toString())
	// fmt.Printf("  end   [%s]\n", end.toString())

	points, path := track(data, start, end)

	// fmt.Printf("  points [%v]\n", points)
	// fmt.Printf("  path [%v]\n", path)

	for _, p := range path {
		if p == end {
			continue
		}

		// fmt.Printf("   p[%s]", p.toString())

		// up
		temp := Point{p.X, p.Y - 2}
		middle := Point{p.X, p.Y - 1}
		if _, ok := points[temp]; ok && data[middle.Y][middle.X] == '#' {
			gain := getGain(p, temp, points)
			// fmt.Printf(" - UP - [%d]", gain)
			if gain >= limit {
				counter++
			}
		}

		// down
		temp = Point{p.X, p.Y + 2}
		middle = Point{p.X, p.Y + 1}
		if _, ok := points[temp]; ok && data[middle.Y][middle.X] == '#' {
			gain := getGain(p, temp, points)
			// fmt.Printf(" - DOWN - [%d]", gain)
			if gain >= limit {
				counter++
			}
		}

		// left
		temp = Point{p.X - 2, p.Y}
		middle = Point{p.X - 1, p.Y}
		if _, ok := points[temp]; ok && data[middle.Y][middle.X] == '#' {
			gain := getGain(p, temp, points)
			// fmt.Printf(" - LEFT - [%d]", gain)
			if gain >= limit {
				counter++
			}
		}

		// right
		temp = Point{p.X + 2, p.Y}
		middle = Point{p.X + 1, p.Y}
		if _, ok := points[temp]; ok && data[middle.Y][middle.X] == '#' {
			gain := getGain(p, temp, points)
			// fmt.Printf(" - RIGHT - [%d]", gain)
			if gain >= limit {
				counter++
			}
		}
		// fmt.Println()
	}

	fmt.Printf("Part 1: %d\n", counter)
}

func getNeigh(grid []string, start Point, direction string) (Point, bool) {
	var p Point
	if direction == "up" {
		p = Point{start.X, start.Y - 1}
	} else if direction == "down" {
		p = Point{start.X, start.Y + 1}
	} else if direction == "left" {
		p = Point{start.X - 1, start.Y}
	} else if direction == "right" {
		p = Point{start.X + 1, start.Y}
	} else {
		return p, false
	}

	if p.X < 0 || p.Y < 0 || p.X >= len(grid[0]) || p.Y >= len(grid) {
		return p, false
	}
	return p, true
}

// TODO: new function to use
func getGridPoint(data []string, start Point, track map[Point]int) (map[Point]int, Point, int) {
	queue := []Point{start}

	points := map[Point]int{start: 0}

	directions := []string{"up", "down", "left", "right"}

	distance := 1
	var p Point

	var bestPoint Point
	var bestScore int

	for distance < 20 {
		newQ := []Point{}
		for len(queue) > 0 {
			p, queue = queue[0], queue[1:]
			for _, dir := range directions {
				temp, ok := getNeigh(data, p, dir)
				if !ok {
					continue
				}
				if data[temp.Y][temp.X] == '#' {
					if _, okk := points[temp]; !okk {
						newQ = append(newQ, temp)
						points[temp] = distance
					}
				} else if v, okk := track[temp]; okk && track[temp] > track[start] {
					score := v - distance // v - 1 - distance
					if score > bestScore {
						// fmt.Printf("old point and score[%s][%d] new point and score[%s][%d]\n", bestPoint.toString(), bestScore, temp.toString(), score)
						bestScore = score
						bestPoint = temp
					}
				}
			}
		}
		queue = newQ

		distance++
	}
	return points, bestPoint, bestScore
}

// ---- new ----

func getPointsInRange(data []string, start Point, step int) []Point {
	ps := []Point{}

	minY := util.IntMax(0, start.Y-step)
	maxY := util.IntMin(len(data)-1, start.Y+step)

	for y := minY; y <= maxY; y++ {
		// for every row, calculate how much narrow/wide is the x range

		// dy = how far is the row from the origin. This value is used to narrow the width
		dy := start.Y - y
		if dy < 0 {
			dy *= -1
		}
		dx := step - dy // width

		minX := util.IntMax(0, start.X-dx)
		maxX := util.IntMin(len(data[0])-1, start.X+dx)

		for x := minX; x <= maxX; x++ {
			temp := Point{x, y}
			if temp == start {
				continue
			}
			ps = append(ps, temp)
		}
	}
	return ps
}

func distance(a, b Point) int {
	dx := util.IntAbs(a.X - b.X)
	dy := util.IntAbs(a.Y - b.Y)
	return dx + dy
}

func part2(data []string) {
	counter := 0

	limit := 100

	start := locate(data, "S")
	end := locate(data, "E")
	// fmt.Printf("  start [%s]\n", start.toString())
	// fmt.Printf("  end   [%s]\n", end.toString())

	points, path := track(data, start, end)

	// fmt.Printf("  points [%v]\n", points)
	// fmt.Printf("  path [%v]\n", path)

	for _, p := range path {
		if p == end {
			continue
		}

		pointsInReach := getPointsInRange(data, p, 20)
		for _, pp := range pointsInReach {
			endVal, ok := points[pp]
			if !ok {
				continue
			}

			startVal := points[p]

			if endVal > startVal {
				gain := endVal - startVal - distance(p, pp)
				if gain >= limit {
					counter++
				}
			}
		}

	}

	fmt.Printf("Part 2: %d\n", counter)
}
