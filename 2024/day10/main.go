package main

import (
	"fmt"

	"advent2024/util"
)

func main() {
	// data := util.GetTestByRow("day10")
	data := util.GetDataByRow("day10")

	grid := [][]int{}

	for _, d := range data {
		grid = append(grid, util.StringToIntSlice(d, ""))
	}

	part1(grid)
	part2(grid)
}

type Point struct {
	StartX int
	StartY int
	X      int
	Y      int
	Value  int
}

func (p Point) ToString() string {
	return fmt.Sprintf("%d-%d", p.X, p.Y)
}

func part1(data [][]int) {
	sum := 0

	queue := []Point{}

	// find all zeros
	for r := range data {
		// fmt.Printf("%v\n", data[r])
		for c := range data[0] {
			if data[r][c] == 0 {
				queue = append(queue, Point{c, r, c, r, 0})
			}
		}
	}

	trails := map[string]bool{}

	for len(queue) > 0 {
		var p Point
		p, queue = queue[0], queue[1:]
		// fmt.Printf("[%v] - %v - ", p, queue)

		if p.Value == 9 {
			trails[fmt.Sprintf("%d%d%d%d", p.StartX, p.StartY, p.X, p.Y)] = true
			continue
		}

		// up
		if p.Y > 0 && data[p.Y-1][p.X] == p.Value+1 {
			pt := Point{p.StartX, p.StartY, p.X, p.Y - 1, p.Value + 1}
			queue = append(queue, pt)
			// fmt.Printf("[%v] - ", pt)
		}

		// right
		if p.X < len(data[0])-1 && data[p.Y][p.X+1] == p.Value+1 {
			pt := Point{p.StartX, p.StartY, p.X + 1, p.Y, p.Value + 1}
			queue = append(queue, pt)
			// fmt.Printf("[%v] - ", pt)
		}

		// down
		if p.Y < len(data)-1 && data[p.Y+1][p.X] == p.Value+1 {
			pt := Point{p.StartX, p.StartY, p.X, p.Y + 1, p.Value + 1}
			queue = append(queue, pt)
			// fmt.Printf("[%v] - ", pt)
		}

		// left
		if p.X > 0 && data[p.Y][p.X-1] == p.Value+1 {
			pt := Point{p.StartX, p.StartY, p.X - 1, p.Y, p.Value + 1}
			queue = append(queue, pt)
			// fmt.Printf("[%v] - ", pt)
		}
		// fmt.Printf("\n")

	}
	sum = len(trails)

	fmt.Printf("Part 1: %d\n", sum)
}

func part2(data [][]int) {
	sum := 0

	queue := []Point{}

	// find all zeros
	for r := range data {
		// fmt.Printf("%v\n", data[r])
		for c := range data[0] {
			if data[r][c] == 0 {
				queue = append(queue, Point{c, r, c, r, 0})
			}
		}
	}

	for len(queue) > 0 {
		var p Point
		p, queue = queue[0], queue[1:]
		// fmt.Printf("[%v] - %v - ", p, queue)

		if p.Value == 9 {
			sum++
			continue
		}

		// up
		if p.Y > 0 && data[p.Y-1][p.X] == p.Value+1 {
			pt := Point{p.StartX, p.StartY, p.X, p.Y - 1, p.Value + 1}
			queue = append(queue, pt)
			// fmt.Printf("[%v] - ", pt)
		}

		// right
		if p.X < len(data[0])-1 && data[p.Y][p.X+1] == p.Value+1 {
			pt := Point{p.StartX, p.StartY, p.X + 1, p.Y, p.Value + 1}
			queue = append(queue, pt)
			// fmt.Printf("[%v] - ", pt)
		}

		// down
		if p.Y < len(data)-1 && data[p.Y+1][p.X] == p.Value+1 {
			pt := Point{p.StartX, p.StartY, p.X, p.Y + 1, p.Value + 1}
			queue = append(queue, pt)
			// fmt.Printf("[%v] - ", pt)
		}

		// left
		if p.X > 0 && data[p.Y][p.X-1] == p.Value+1 {
			pt := Point{p.StartX, p.StartY, p.X - 1, p.Y, p.Value + 1}
			queue = append(queue, pt)
			// fmt.Printf("[%v] - ", pt)
		}
		// fmt.Printf("\n")

	}
	fmt.Printf("Part 2: %d\n", sum)
}
