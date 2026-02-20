package main

import (
	"fmt"
	"slices"

	"advent2015/util"
)

func main() {
	data := util.GetDataByRow("day02")
	// data := util.GetTestByRow("day02")

	list := parse(data)

	part1(list)
	part2(list)
}

func parse(data []string) [][]int {
	list := make([][]int, 0)

	for _, s := range data {
		t := util.StringToIntSlice(s, "x")
		list = append(list, t)
	}
	return list
}

func part1(data [][]int) {
	counter := 0

	for _, l := range data {
		a1 := l[0] * l[1]
		a2 := l[0] * l[2]
		a3 := l[1] * l[2]

		smaller := util.IntMin(a1, a2)
		smaller = util.IntMin(smaller, a3)

		counter += smaller
		counter += a1*2 + a2*2 + a3*2
	}

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(data [][]int) {
	counter := 0

	for _, l := range data {
		slices.Sort(l)
		counter += l[0]*2 + l[1]*2 + l[0]*l[1]*l[2]
	}

	fmt.Printf("Part 2: %d\n", counter)
}
