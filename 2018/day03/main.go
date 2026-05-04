package main

import (
	"fmt"
	"regexp"
	"strings"

	"advent2018/util"
)

func main() {
	data := util.GetDataByRow("day03")
	// data := util.GetTestByRow("day03")

	list := parse(data)

	part1(list)
	part2(list)
}

type Point struct {
	x int
	y int
}

type Patch struct {
	id     int
	point  Point
	width  int
	height int
}

func parse(data []string) []Patch {
	ps := []Patch{}

	reg := regexp.MustCompile(`#(\d+)\s@\s(\d+),(\d+):\s(\d+)x(\d+)`)

	for _, s := range data {
		pieces := reg.FindStringSubmatch(s)
		ps = append(ps, Patch{
			id: util.ToInt(pieces[1]),
			point: Point{
				x: util.ToInt(pieces[2]),
				y: util.ToInt(pieces[3]),
			},
			width:  util.ToInt(pieces[4]),
			height: util.ToInt(pieces[5]),
		})
	}

	return ps
}

func fill(fabric [][]int, p Patch) {
	for y := p.point.y; y < p.point.y+p.height; y++ {
		for x := p.point.x; x < p.point.x+p.width; x++ {
			fabric[y][x] += 1
		}
	}
}

func printFabric(f [][]int) {
	var sb strings.Builder
	sb.WriteString("-------------------------------\n")
	for _, r := range f {
		for _, v := range r {
			if v > 1 {
				sb.WriteRune('X')
			} else if v == 1 {
				sb.WriteRune('o')
			} else {
				sb.WriteRune('.')
			}
		}
		sb.WriteString("\n")
	}
	sb.WriteString("-------------------------------")
	fmt.Println(sb.String())
}

func part1(data []Patch) {
	counter := 0

	// Contains at least 1000x1000
	size := 1010

	fabric := make([][]int, size)
	for i := range size {
		fabric[i] = make([]int, size)
	}

	for _, p := range data {
		fill(fabric, p)
	}

	for _, r := range fabric {
		for _, v := range r {
			if v > 1 {
				counter++
			}
		}
	}

	// printFabric(fabric)

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(data []Patch) {
	counter := 0

	points := map[Point][]int{}

	okPatches := map[int]bool{}

	for _, p := range data {
		overlap := false
		for y := p.point.y; y < p.point.y+p.height; y++ {
			for x := p.point.x; x < p.point.x+p.width; x++ {
				po := Point{x, y}
				points[po] = append(points[po], p.id)
				if len(points[po]) > 1 {
					overlap = true
					// remove others from okPatches
					for _, i := range points[po] {
						okPatches[i] = false
					}
				}
			}
		}
		if !overlap {
			okPatches[p.id] = true
		}
	}

	for id, ok := range okPatches {
		if ok {
			counter = id
			break
		}
	}

	fmt.Printf("Part 2: %d\n", counter)
}
