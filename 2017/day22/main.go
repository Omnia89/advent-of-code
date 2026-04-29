package main

import (
	"fmt"

	"advent2017/util"
)

func main() {
	data := util.GetDataByRow("day22")
	// data := util.GetTestByRow("day22")

	layer1, center1 := parse(data)
	layer2, center2 := parse(data)

	part1(layer1, center1)
	part2(layer2, center2)
}

const (
	statusClean    = 0
	statusWeakened = 2
	statusInfected = 1
	statusFlagged  = 3
)

type Point struct {
	x int
	y int
}

func (p *Point) add(o Point) {
	p.x += o.x
	p.y += o.y
}

func parse(data []string) (layer map[Point]int, center Point) {
	center = Point{
		x: len(data[0]) / 2,
		y: len(data) / 2,
	}

	layer = map[Point]int{}

	for y, s := range data {
		for x, c := range s {
			val := statusClean
			if c == '#' {
				val = statusInfected
			}
			layer[Point{x, y}] = val
		}
	}
	return
}

var leftTurns = map[Point]Point{
	{0, -1}: {-1, 0},
	{-1, 0}: {0, 1},
	{0, 1}:  {1, 0},
	{1, 0}:  {0, -1},
}

var rightTurns = map[Point]Point{
	{0, -1}: {1, 0},
	{1, 0}:  {0, 1},
	{0, 1}:  {-1, 0},
	{-1, 0}: {0, -1},
}

func turn1(direction Point, value int) Point {
	if value == statusClean {
		return leftTurns[direction]
	}
	return rightTurns[direction]
}

func turn2(direction Point, value int) Point {
	switch value {
	case statusClean:
		return leftTurns[direction]
	case statusInfected:
		return rightTurns[direction]
	case statusFlagged:
		return Point{-direction.x, -direction.y}
	}

	return direction
}

func part1(layer map[Point]int, center Point) {
	counter := 0

	direction := Point{0, -1}
	current := center

	for range 10000 {
		val := layer[current]
		direction = turn1(direction, val)

		newVal := statusClean
		if val == statusClean {
			newVal = statusInfected
		}
		counter += newVal
		layer[current] = newVal

		current.add(direction)
	}

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(layer map[Point]int, center Point) {
	counter := 0

	direction := Point{0, -1}
	current := center

	for range 10000000 {
		val := layer[current]
		direction = turn2(direction, val)

		var newVal int

		switch val {
		case statusClean:
			newVal = statusWeakened
		case statusWeakened:
			newVal = statusInfected
		case statusInfected:
			newVal = statusFlagged
		case statusFlagged:
			newVal = statusClean
		}

		if newVal == statusInfected {
			counter++
		}
		layer[current] = newVal

		current.add(direction)
	}
	fmt.Printf("Part 2: %d\n", counter)
}
