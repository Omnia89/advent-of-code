package main

import (
	"fmt"
	"strings"

	"advent2017/util"
)

func main() {
	data := util.GetRawData("day11")
	// data := util.GetRawTest("day11")

	list := strings.Split(data, ",")

	part1(list)
	part2(list)
}

// even if its hexagonal, use the x,y system (doubled coordinates)
type Point struct {
	x int
	y int
}

func part1(data []string) {
	counter := 0

	dirs := map[string]int{}

	for _, s := range data {
		t := strings.TrimSpace(s)
		dirs[t]++
	}

	// Nullify opposite directions
	if dirs["s"] > dirs["n"] {
		dirs["s"] -= dirs["n"]
		dirs["n"] = 0
	} else {
		dirs["n"] -= dirs["s"]
		dirs["s"] = 0
	}
	if dirs["sw"] > dirs["ne"] {
		dirs["sw"] -= dirs["ne"]
		dirs["ne"] = 0
	} else {
		dirs["ne"] -= dirs["sw"]
		dirs["sw"] = 0
	}
	if dirs["se"] > dirs["nw"] {
		dirs["se"] -= dirs["nw"]
		dirs["nw"] = 0
	} else {
		dirs["nw"] -= dirs["se"]
		dirs["se"] = 0
	}

	// simplify direction
	type couple struct {
		left   string
		right  string
		target string
	}

	couples := []couple{
		{"nw", "ne", "n"},
		{"n", "se", "ne"},
		{"ne", "s", "se"},
		{"se", "sw", "s"},
		{"s", "nw", "sw"},
		{"sw", "n", "nw"},
	}

	for _, c := range couples {
		if dirs[c.left] > 0 && dirs[c.right] > 0 {
			larger := c.right
			lesser := c.left
			if dirs[c.left] > dirs[c.right] {
				larger = c.left
				lesser = c.right
			}
			dirs[larger] -= dirs[lesser]
			dirs[c.target] += dirs[lesser]
			dirs[lesser] = 0
		}
	}

	for _, n := range dirs {
		counter += n
	}

	fmt.Printf("Part 1: %d\n", counter)
}

func getPoint(p Point, direction string) Point {
	switch direction {
	case "n":
		return Point{p.x, p.y - 2}
	case "ne":
		return Point{p.x + 1, p.y - 1}
	case "se":
		return Point{p.x + 1, p.y + 1}
	case "s":
		return Point{p.x, p.y + 2}
	case "sw":
		return Point{p.x - 1, p.y + 1}
	case "nw":
		return Point{p.x - 1, p.y - 1}
	}
	return Point{}
}

func originDist(p Point) int {
	aX := util.IntAbs(p.x)
	aY := util.IntAbs(p.y)

	s := 0
	if aX <= aY {
		s = (aY - aX) / 2
	}

	return aX + s
}

func part2(data []string) {
	counter := 0

	current := Point{0, 0}
	maxDist := 0

	for n, s := range data {
		t := strings.TrimSpace(s)
		current = getPoint(current, t)
		d := originDist(current)
		if n < 100 {
			fmt.Printf("p[%v] dir[%s] [%d]\n", current, s, d)
		}
		if d > maxDist {
			maxDist = d
		}
	}
	counter = maxDist

	fmt.Printf("Part 2: %d\n", counter)
}
