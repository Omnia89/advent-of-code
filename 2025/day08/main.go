package main

import (
	"fmt"
	"math"
	"slices"
	"sort"
	"strings"

	"advent2025/util"
)

func main() {
	data := util.GetDataByRow("day08")
	// data := util.GetTestByRow("day08")

	points := []Point{}
	for _, r := range data {
		s := strings.Split(r, ",")
		points = append(points, Point{
			X: util.ToInt(s[0]),
			Y: util.ToInt(s[1]),
			Z: util.ToInt(s[2]),
		})
	}

	// part1(points, true)

	part1(points, false)
	part2(points)
}

type Point struct {
	X int
	Y int
	Z int
}

func (p Point) toString() string {
	return fmt.Sprintf("%d-%d-%d", p.X, p.Y, p.Z)
}

type CalculatedDistance struct {
	Distance float64
	A        *Point
	B        *Point
}

func distance(a Point, b Point) float64 {
	return math.Sqrt(math.Pow(float64(a.X-b.X), 2) + math.Pow(float64(a.Y-b.Y), 2) + math.Pow(float64(a.Z-b.Z), 2))
}

func part1(points []Point, test bool) {
	counter := 0

	allDistances := []CalculatedDistance{}
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			allDistances = append(allDistances, CalculatedDistance{
				A:        &points[i],
				B:        &points[j],
				Distance: distance(points[i], points[j]),
			})
		}
	}

	// Ordina per distance
	slices.SortFunc(allDistances, func(a CalculatedDistance, b CalculatedDistance) int {
		if a.Distance < b.Distance {
			return -1
		}
		if a.Distance > b.Distance {
			return 1
		}
		return 0
	})

	currentGroup := 1
	circuits := make(map[int][]string)

	size := 1000
	if test {
		size = 10
	}

	for i := range size {
		dist := allDistances[i]

		// search A group
		aGroup := 0
		for k, g := range circuits {
			if slices.Contains(g, dist.A.toString()) {
				aGroup = k
			}
		}

		// search B group
		bGroup := 0
		for k, g := range circuits {
			if slices.Contains(g, dist.B.toString()) {
				bGroup = k
			}
		}

		if aGroup != 0 && bGroup == 0 {
			circuits[aGroup] = append(circuits[aGroup], dist.B.toString())
		} else if aGroup == 0 && bGroup != 0 {
			circuits[bGroup] = append(circuits[bGroup], dist.A.toString())
		} else if aGroup+bGroup == 0 {
			circuits[currentGroup] = []string{dist.A.toString(), dist.B.toString()}
			currentGroup++
		} else if aGroup != bGroup {
			// merge
			circuits[aGroup] = append(circuits[aGroup], circuits[bGroup]...)
			circuits[bGroup] = []string{}
		}
		// If the same do nothing

	}

	circuitSize := []int{}

	for _, g := range circuits {
		circuitSize = append(circuitSize, len(g))
	}

	sort.Ints(circuitSize)

	counter = 1
	counter *= circuitSize[len(circuitSize)-1]
	counter *= circuitSize[len(circuitSize)-2]
	counter *= circuitSize[len(circuitSize)-3]

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(points []Point) {
	counter := 0

	allDistances := []CalculatedDistance{}
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			allDistances = append(allDistances, CalculatedDistance{
				A:        &points[i],
				B:        &points[j],
				Distance: distance(points[i], points[j]),
			})
		}
	}

	// Ordina per distance
	slices.SortFunc(allDistances, func(a CalculatedDistance, b CalculatedDistance) int {
		if a.Distance < b.Distance {
			return -1
		}
		if a.Distance > b.Distance {
			return 1
		}
		return 0
	})

	circuits := make(map[int][]string)
	for i, p := range points {
		circuits[i+1] = []string{p.toString()}
	}
	numCircuits := len(circuits)

	i := 0

	lastX := []int{0, 0}
	fmt.Printf("N.points [%d] - max Distances [%d]", numCircuits, len(allDistances))
	for numCircuits > 1 && i < len(allDistances) {
		dist := allDistances[i]
		i++

		// search A group
		aGroup := 0
		for k, g := range circuits {
			if slices.Contains(g, dist.A.toString()) {
				aGroup = k
			}
		}

		// search B group
		bGroup := 0
		for k, g := range circuits {
			if slices.Contains(g, dist.B.toString()) {
				bGroup = k
			}
		}

		if aGroup != 0 && bGroup == 0 {
			circuits[aGroup] = append(circuits[aGroup], dist.B.toString())
			circuits[bGroup] = []string{}
			lastX[0] = dist.A.X
			lastX[1] = dist.B.X
		} else if aGroup == 0 && bGroup != 0 {
			circuits[bGroup] = append(circuits[bGroup], dist.A.toString())
			circuits[aGroup] = []string{}
			lastX[0] = dist.A.X
			lastX[1] = dist.B.X
		} else if aGroup != bGroup {
			// merge
			circuits[aGroup] = append(circuits[aGroup], circuits[bGroup]...)
			circuits[bGroup] = []string{}
			lastX[0] = dist.A.X
			lastX[1] = dist.B.X
		}
		// If the same do nothing

	}

	counter = lastX[0] * lastX[1]

	fmt.Printf("Part 2: %d\n", counter)
}
