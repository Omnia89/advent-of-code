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

func part2(data []string) {
	counter := 0
	for _, p := range path {
		if p == end {

	fmt.Printf("Part 2: %d\n", counter)
}
