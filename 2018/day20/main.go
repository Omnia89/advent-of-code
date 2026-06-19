package main

import (
	"fmt"
	"strings"

	"advent2018/util"
)

func main() {
	data := util.GetDataByRow("day20")
	// data := util.GetTestByRow("day20")

	area := buildArea(data[0])

	part1(area)
	part2(area)
}

type Point struct {
	x int
	y int
}

func (p Point) add(pp Point) Point {
	p.x += pp.x
	p.y += pp.y
	return p
}

type Cell int

const (
	Wall Cell = iota
	Door
	Room
)

var directions map[rune]Point = map[rune]Point{
	'N': {0, -1},
	'S': {0, 1},
	'E': {1, 0},
	'W': {-1, 0},
}

func buildArea(data string) map[Point]Cell {
	area := map[Point]Cell{
		{0, 0}: Room,
	}

	type frame struct {
		starts map[Point]bool
		ends   map[Point]bool
	}

	stack := []frame{}
	current := map[Point]bool{{0, 0}: true}

	for _, c := range data {
		switch c {
		case '^', '$':
		case 'N', 'S', 'E', 'W':
			next := make(map[Point]bool)
			for p := range current {
				d := p.add(directions[c])
				area[d] = Door
				r := d.add(directions[c])
				area[r] = Room
				next[r] = true
			}
			current = next
		case '(':
			starts := make(map[Point]bool)
			for p := range current {
				starts[p] = true
			}
			f := frame{starts, make(map[Point]bool)}
			stack = append(stack, f)
		case '|':
			f := &stack[len(stack)-1]
			for p := range current {
				f.ends[p] = true
			}
			current = make(map[Point]bool, len(f.starts))
			for p := range f.starts {
				current[p] = true
			}
		case ')':
			f := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			for p := range f.ends {
				current[p] = true
			}
		}
	}

	return area
}

func drawArea(dirs []string) map[Point]Cell {
	area := map[Point]Cell{
		{0, 0}: Room,
	}

	for _, ds := range dirs {
		p := Point{0, 0}
		for _, d := range ds {
			p = p.add(directions[d])
			area[p] = Door
			p = p.add(directions[d])
			area[p] = Room
		}
	}
	return area
}

func printArea(area map[Point]Cell) {
	minX, minY, maxX, maxY := 0, 0, 0, 0
	for p := range area {
		if minX > p.x {
			minX = p.x
		}
		if maxX < p.x {
			maxX = p.x
		}
		if minY > p.y {
			minY = p.y
		}
		if maxY < p.y {
			maxY = p.y
		}
	}
	// borders
	minX--
	minY--
	maxX++
	maxY++

	var sb strings.Builder

	sb.WriteString(strings.Repeat("=", 25))
	sb.WriteString("\n")

	for iy := range maxY - minY + 1 {
		y := iy + minY
		for ix := range maxX - minX + 1 {
			x := ix + minX
			c := area[Point{x, y}]
			switch c {
			case Wall:
				sb.WriteString("#")
			case Room:
				sb.WriteString(" ")
			case Door:
				s := "|"
				if x%2 == 0 {
					s = "-"
				}
				sb.WriteString(s)
			}
		}
		sb.WriteString("\n")
	}
	sb.WriteString(strings.Repeat("=", 25))
	sb.WriteString("\n")

	fmt.Print(sb.String())
}

func part1(area map[Point]Cell) {
	counter := 0

	visitedRooms := map[Point]bool{}

	type state struct {
		Point
		hops int
	}

	maxHops := 0

	queue := []state{{Point{0, 0}, 0}}
	var p state
	for len(queue) > 0 {
		p, queue = queue[0], queue[1:]

		if p.hops > maxHops {
			maxHops = p.hops
		}

		for _, d := range directions {
			t := p.add(d)
			if area[t] == Door {
				t = t.add(d)
				if !visitedRooms[t] {
					queue = append(queue, state{
						t,
						p.hops + 1,
					})
					visitedRooms[t] = true
				}
			}
		}
	}

	printArea(area)

	counter = maxHops

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(area map[Point]Cell) {
	counter := 0

	type state struct {
		Point
		hops int
	}
	origin := Point{0, 0}

	for start, t := range area {
		if t != Room {
			continue
		}
		if start == origin {
			continue
		}
		doors := 0
		visitedRooms := map[Point]bool{}

		queue := []state{{start, 0}}
		var p state
		for len(queue) > 0 {
			p, queue = queue[0], queue[1:]

			for _, d := range directions {
				t := p.add(d)
				if area[t] == Door {
					t = t.add(d)
					if t == origin {
						doors = p.hops + 1
						break
					}

					if !visitedRooms[t] {
						queue = append(queue, state{
							t,
							p.hops + 1,
						})
						visitedRooms[t] = true
					}
				}
			}
		}
		if doors >= 1000 {
			counter++
		}
	}

	fmt.Printf("Part 2: %d\n", counter)
}
