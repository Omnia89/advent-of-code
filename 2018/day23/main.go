package main

import (
	"fmt"
	"regexp"

	"advent2018/util"
)

func main() {
	// data := util.GetDataByRow("day23")
	data := util.GetTestByRow("day23")

	list := parse(data)

	part1(list)
	part2(list)
}

type Point struct {
	x int
	y int
	z int
}

type Bot struct {
	p      Point
	radius int
}

func parse(data []string) []Bot {
	bs := []Bot{}

	capture := regexp.MustCompile(`<(\-?\d+),(\-?\d+),(\-?\d+)>.*?(\d+)`)

	for _, s := range data {
		matches := capture.FindStringSubmatch(s)
		bs = append(bs, Bot{
			p:      Point{util.ToInt(matches[1]), util.ToInt(matches[2]), util.ToInt(matches[3])},
			radius: util.ToInt(matches[4]),
		})
	}

	return bs
}

func part1(bots []Bot) {
	counter := 0

	var stronger Bot

	for _, b := range bots {
		if b.radius > stronger.radius {
			stronger = b
		}
	}

	for _, b := range bots {
		delta := util.IntAbs(b.p.x - stronger.p.x)
		delta += util.IntAbs(b.p.y - stronger.p.y)
		delta += util.IntAbs(b.p.z - stronger.p.z)

		if delta <= stronger.radius {
			counter++
		}
	}

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(bots []Bot) {
	touched := map[Point]int{}

	for _, b := range bots {
		for dz := range b.radius + 1 {
			newRadius := b.radius - dz

			for dy := range newRadius*2 + 1 {
				y := dy - (newRadius - b.p.y)

				remaining := newRadius - util.IntAbs(b.p.y-y)

				for dx := range remaining*2 + 1 {
					x := b.p.x - remaining + dx
					if dz != 0 {
						touched[Point{x, y, b.p.z - dz}] += 1
						touched[Point{x, y, b.p.z + dz}] += 1
					} else {
						touched[Point{x, y, b.p.z}] += 1
					}
				}
			}
		}
	}

	maxTouch := 0
	var p Point

	for pp, n := range touched {
		if n > maxTouch {
			maxTouch = n
			p = pp
		}
	}

	fmt.Printf("%d,%d,%d\n", p.x, p.y, p.z)
	counter := util.IntAbs(p.x) + util.IntAbs(p.y) + util.IntAbs(p.z)

	fmt.Printf("Part 2: %d\n", counter)
}
