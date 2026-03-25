package main

import (
	"fmt"
	"slices"

	"advent2017/util"
)

func main() {
	data := util.GetDataByRow("day02")
	// data := util.GetTestByRow("day02")

	list := parse(data)

	part1(list)
	part2(list)
}

func parse(data []string) [][]int {
	r := [][]int{}

	for _, s := range data {
		r = append(r, util.StringToIntSlice(s, "\t"))
	}
	return r
}

func part1(data [][]int) {
	counter := 0

	for _, l := range data {
		minor := slices.Min(l)
		major := slices.Max(l)

		counter += major - minor
	}

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(data [][]int) {
	counter := 0

	for _, l := range data {
	arrayLoop:
		for i := range len(l) - 1 {
			for j := i + 1; j < len(l); j++ {
				minor := l[i]
				major := l[j]
				if major < minor {
					minor, major = major, minor
				}
				if major%minor == 0 {
					counter += major / minor
					break arrayLoop
				}
			}
		}
	}

	fmt.Printf("Part 2: %d\n", counter)
}
