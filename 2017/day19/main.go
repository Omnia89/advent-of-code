package main

import (
	"fmt"
	"regexp"
	"strings"

	"advent2017/util"
)

func main() {
	data := util.GetDataByRow("day19")
	// data := util.GetTestByRow("day19")

	part1(data)
	part2(data)
}

type Point struct {
	x int
	y int
}

func (p *Point) add(o Point) {
	p.x += o.x
	p.y += o.y
}

var alpha = regexp.MustCompile(`[A-Z]`)

func isLetter(s string) bool {
	return alpha.MatchString(s)
}

func getDirections(d Point) []Point {
	if d.x != 0 {
		return []Point{
			{0, 1},
			{0, -1},
		}
	}
	return []Point{
		{1, 0},
		{-1, 0},
	}
}

func part1(data []string) {
	startX := strings.Index(data[0], "|")

	current := Point{startX, 0}
	direction := Point{0, 1}

	letters := []string{}

	for {
		c := string(data[current.y][current.x])
		if c == " " {
			break
		}

		if isLetter(c) {
			letters = append(letters, c)
		} else if c == "+" {
			// change direction
			for _, d := range getDirections(direction) {
				y := current.y + d.y
				x := current.x + d.x
				if y < 0 || x < 0 || y >= len(data) || x >= len(data[0]) {
					continue
				}
				if data[current.y+d.y][current.x+d.x] != ' ' {
					direction = d
					break
				}
			}
		}
		current.add(direction)
	}

	solution := strings.Join(letters, "")

	fmt.Printf("Part 1: %s\n", solution)
}

func part2(data []string) {
	counter := 0

	startX := strings.Index(data[0], "|")

	current := Point{startX, 0}
	direction := Point{0, 1}

	for {
		c := string(data[current.y][current.x])
		if c == " " {
			break
		}

		if c == "+" {
			// change direction
			for _, d := range getDirections(direction) {
				y := current.y + d.y
				x := current.x + d.x
				if y < 0 || x < 0 || y >= len(data) || x >= len(data[0]) {
					continue
				}
				if data[current.y+d.y][current.x+d.x] != ' ' {
					direction = d
					break
				}
			}
		}
		current.add(direction)
		counter++
	}
	fmt.Printf("Part 2: %d\n", counter)
}
