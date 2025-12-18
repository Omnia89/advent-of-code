package main

import (
	"fmt"
	"math"
	"strings"

	"advent2025/util"
)

func main() {
	rows := util.GetDataByRow("day09")
	// rows := util.GetTestByRow("day09")

	// 1574681964 è troppo bassa come risposta

	points := []Point{}
	for _, r := range rows {
		s := strings.Split(r, ",")
		points = append(points, Point{
			X: util.ToInt(s[0]),
			Y: util.ToInt(s[1]),
		})
	}

	part1(points)
	part2(points)
}

type Point struct {
	X int
	Y int
}

type Segment struct {
	Min int
	Max int
}

func (p Point) toString() string {
	return fmt.Sprintf("%d-%d", p.X, p.Y)
}

func (p Point) areaWith(other Point) int {
	w := util.IntAbs(p.X-other.X) + 1
	h := util.IntAbs(p.Y-other.Y) + 1

	return h * w
}

func (p Point) validAreaWith(other Point, segments map[int]Segment) (area int, ok bool) {
	minX := util.IntMin(p.X, other.X)
	maxX := util.IntMax(p.X, other.X)
	minY := util.IntMin(p.Y, other.Y)
	maxY := util.IntMax(p.Y, other.Y)

	for y := minY; y <= maxY; y++ {
		segment := segments[y]
		if minX < segment.Min || maxX > segment.Max {
			return 0, false
		}
	}

	return p.areaWith(other), true
}

func part1(points []Point) {
	maxArea := 0

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			area := points[i].areaWith(points[j])
			if area > maxArea {
				maxArea = area
			}
		}
	}

	fmt.Printf("Part 1: %d\n", maxArea)
}

func part2(points []Point) {
	maxArea := 0

	pointsInX := map[int][]Point{}
	pointsInY := map[int][]Point{}

	for _, p := range points {
		pointsInX[p.X] = append(pointsInX[p.X], p)
		pointsInY[p.Y] = append(pointsInX[p.X], p)
	}

	perimeterByY := map[int][]Point{}

	j := len(points) - 1
	for i := 0; i < len(points); i++ {
		minY := util.IntMin(points[i].Y, points[j].Y)
		minX := util.IntMin(points[i].X, points[j].X)
		maxY := util.IntMax(points[i].Y, points[j].Y)
		maxX := util.IntMax(points[i].X, points[j].X)

		for x := minX; x <= maxX; x++ {
			for y := minY; y <= maxY; y++ {
				p := Point{x, y}
				perimeterByY[y] = append(perimeterByY[y], p)
			}
		}
		j = i
	}

	horizontalSegments := map[int]Segment{}
	for y, pp := range perimeterByY {
		min := math.MaxInt
		max := -1

		for _, p := range pp {
			min = util.IntMin(min, p.X)
			max = util.IntMax(max, p.X)
		}
		horizontalSegments[y] = Segment{min, max}
	}

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			area, ok := points[i].validAreaWith(points[j], horizontalSegments)
			if ok && area > maxArea {
				maxArea = area
			}
		}
	}

	fmt.Printf("Part 2: %d\n", maxArea)
}
