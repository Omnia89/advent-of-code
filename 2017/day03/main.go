package main

import (
	"fmt"

	"advent2017/util"
)

func main() {
	data := util.GetDataByRow("day03")
	// data := util.GetTestByRow("day03")

	num := util.ToInt(data[0])

	part1(num)
	part2(num)
}

func part1(num int) {
	counter := 0

	side := 0

	for i := 1; i < num/2; i += 2 {
		if num <= i*i {
			side = i
			break
		}
	}

	start := (side - 2) * (side - 2)

	middle := 0

	for i := 1; i < 5; i++ {
		if num <= start+(side-1)*i {
			middle = (start + (side-1)*i + start + (side-1)*(i-1)) / 2
			break
		}
	}

	delta := util.IntAbs(num - middle)
	counter = side/2 + delta

	fmt.Printf("Part 1: %d\n", counter)
}

type Point struct {
	x int
	y int
}

func (p *Point) add(d Point) {
	p.x += d.x
	p.y += d.y
}

func nextDirection(dir Point) Point {
	switch dir {
	case Point{1, 0}:
		return Point{0, -1}
	case Point{0, -1}:
		return Point{-1, 0}
	case Point{-1, 0}:
		return Point{0, 1}
	default:
		return Point{1, 0}
	}
}

func sumNeighborns(p Point, grid map[Point]int) int {
	s := 0

	if v, ok := grid[Point{p.x, p.y + 1}]; ok {
		s += v
	}
	if v, ok := grid[Point{p.x, p.y - 1}]; ok {
		s += v
	}
	if v, ok := grid[Point{p.x + 1, p.y}]; ok {
		s += v
	}
	if v, ok := grid[Point{p.x - 1, p.y}]; ok {
		s += v
	}

	if v, ok := grid[Point{p.x + 1, p.y + 1}]; ok {
		s += v
	}
	if v, ok := grid[Point{p.x + 1, p.y - 1}]; ok {
		s += v
	}
	if v, ok := grid[Point{p.x - 1, p.y - 1}]; ok {
		s += v
	}
	if v, ok := grid[Point{p.x - 1, p.y + 1}]; ok {
		s += v
	}
	return s
}

// too high 283758
func part2(num int) {
	counter := 0

	if num == 0 {
		fmt.Printf("Part 2: %d\n", 1)
		return
	}

	side := 3
	current := Point{0, 0}
	direction := Point{0, -1}

	grid := map[Point]int{current: 1}

free:
	for {
		current.add(Point{1, 1})
		for range 4 {
			for range side - 1 {
				current.add(direction)
				v := sumNeighborns(current, grid)
				if v > num {
					counter = v
					break free
				}
				grid[current] = v
			}
			direction = nextDirection(direction)
		}
		side += 2
	}

	fmt.Printf("Part 2: %d\n", counter)
}
