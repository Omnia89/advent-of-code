package main

import (
	"fmt"
	"strings"

	"advent2016/util"
)

func main() {
	data := util.GetDataByRow("day01")
	// data := util.GetTestByRow("day01")

	s := strings.ReplaceAll(data[0], ",", "")
	list := strings.Split(s, " ")

	part1(list)
	part2(list)
}

func part1(data []string) {
	counter := 0

	turns := map[string]map[string]string{
		"up": {
			"R": "right",
			"L": "left",
		},
		"right": {
			"R": "down",
			"L": "up",
		},
		"down": {
			"R": "left",
			"L": "right",
		},
		"left": {
			"R": "up",
			"L": "down",
		},
	}

	direction := "up"
	h := 0
	v := 0

	for _, d := range data {
		direction = turns[direction][string(d[0])]
		switch direction {
		case "up":
			v += util.ToInt(d[1:])
		case "down":
			v -= util.ToInt(d[1:])
		case "right":
			h += util.ToInt(d[1:])
		case "left":
			h -= util.ToInt(d[1:])
		}
	}

	counter = util.IntAbs(h) + util.IntAbs(v)

	fmt.Printf("Part 1: %d\n", counter)
}

// 155 too high
func part2(data []string) {
	counter := 0

	turns := map[string]map[string]string{
		"up": {
			"R": "right",
			"L": "left",
		},
		"right": {
			"R": "down",
			"L": "up",
		},
		"down": {
			"R": "left",
			"L": "right",
		},
		"left": {
			"R": "up",
			"L": "down",
		},
	}

	distancesVisited := map[string]bool{}

	direction := "up"
	h := 0
	v := 0

	for _, d := range data {
		direction = turns[direction][string(d[0])]
		dist := util.ToInt(d[1:])
		for range dist {
			switch direction {
			case "up":
				v++
			case "down":
				v--
			case "right":
				h++
			case "left":
				h--
			}
			k := fmt.Sprintf("%d-%d", h, v)
			if _, ok := distancesVisited[k]; ok {
				counter = util.IntAbs(h) + util.IntAbs(v) - 1
				break
			}
			distancesVisited[k] = true
		}
	}
	fmt.Printf("Part 2: %d\n", counter)
}
