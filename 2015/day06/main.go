package main

import (
	"fmt"
	"regexp"

	"advent2015/util"
)

func main() {
	data := util.GetDataByRow("day06")
	//data := util.GetTestByRow("day06")

	istr := parse(data)

	part1(istr)
	part2(istr)
}

type Point struct {
	x int
	y int
}

type Istruction struct {
	op string
	p1 Point
	p2 Point
}

func parse(data []string) []Istruction {
	reg := regexp.MustCompile(`(turn on|toggle|turn off)\s(\d{1,3}),(\d{1,3})\sthrough\s(\d{1,3}),(\d{1,3})`)

	istr := []Istruction{}

	for _, s := range data {
		capt := reg.FindStringSubmatch(s)
		istr = append(istr, Istruction{
			op: capt[1],
			p1: Point{util.ToInt(capt[2]), util.ToInt(capt[3])},
			p2: Point{util.ToInt(capt[4]), util.ToInt(capt[5])},
		})
	}
	return istr
}

func exec(op string, v int) int {
	switch op {
	case "turn on":
		return 1
	case "turn off":
		return 0
	case "toggle":
		if v == 1 {
			return 0
		}
		return 1
	}
	return 0
}

func getBright(op string, v int) int {
	t := v
	switch op {
	case "turn on":
		t += 1
	case "turn off":
		t -= 1
	case "toggle":
		t += 2
	}
	if t < 0 {
		return 0
	}
	return t
}

func part1(istruction []Istruction) {
	counter := 0

	grid := [1000][1000]int{}

	for _, ist := range istruction {
		lowX := util.IntMin(ist.p1.x, ist.p2.x)
		lowY := util.IntMin(ist.p1.y, ist.p2.y)
		highX := util.IntMax(ist.p1.x, ist.p2.x)
		highY := util.IntMax(ist.p1.y, ist.p2.y)

		for x := lowX; x <= highX; x++ {
			for y := lowY; y <= highY; y++ {
				grid[y][x] = exec(ist.op, grid[y][x])
			}
		}
	}

	for y := range grid {
		for _, v := range grid[y] {
			counter += v
		}
	}

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(istruction []Istruction) {
	counter := 0

	grid := [1000][1000]int{}

	for _, ist := range istruction {
		lowX := util.IntMin(ist.p1.x, ist.p2.x)
		lowY := util.IntMin(ist.p1.y, ist.p2.y)
		highX := util.IntMax(ist.p1.x, ist.p2.x)
		highY := util.IntMax(ist.p1.y, ist.p2.y)

		//fmt.Printf("[%d,%d] [%d,%d]\n", lowX, lowY, highX, highY)
		for x := lowX; x <= highX; x++ {
			for y := lowY; y <= highY; y++ {
				temp := getBright(ist.op, grid[y][x])
				grid[y][x] = temp
			}
		}
	}

	for y := range grid {
		for _, v := range grid[y] {
			counter += v
		}
	}
	fmt.Printf("Part 2: %d\n", counter)
}
