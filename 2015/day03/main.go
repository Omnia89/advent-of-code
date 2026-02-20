package main

import (
	"fmt"

	"advent2015/util"
)

func main() {
	data := util.GetDataByRow("day03")
	// data := util.GetTestByRow("day03")

	part1(data)
	part2(data)
}

type Point struct {
	x int
	y int
}

func (p Point) move(direction rune) Point {
	n := Point{p.x, p.y}

	switch direction {
	case '^':
		n.y--
	case '>':
		n.x++
	case 'v':
		n.y++
	case '<':
		n.x--
	}
	return n
}

func part1(data []string) {
	counter := 0

	commands := data[0]

	p := Point{0, 0}
	places := map[Point]bool{p: true}

	for _, c := range commands {
		p = p.move(c)
		places[p] = true
	}
	counter = len(places)

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(data []string) {
	counter := 0

	commands := data[0]

	santa := Point{0, 0}
	robo := Point{0, 0}
	places := map[Point]bool{santa: true, robo: true}

	for i, c := range commands {
		if i%2 == 0 {
			santa = santa.move(c)
			places[santa] = true
		} else {
			robo = robo.move(c)
			places[robo] = true

		}
	}
	counter = len(places)
	fmt.Printf("Part 2: %d\n", counter)
}
