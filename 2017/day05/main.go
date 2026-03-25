package main

import (
	"fmt"
	"slices"

	"advent2017/util"
)

func main() {
	data := util.GetRawData("day05")
	// data := util.GetRawTest("day05")

	list := util.StringToIntSlice(data, "\n")
	list2 := slices.Clone(list)

	part1(list)
	part2(list2)
}

func part1(data []int) {
	counter := 0

	p := 0

	for p >= 0 && p < len(data) {
		data[p]++
		p += data[p] - 1
		counter++
	}

	fmt.Printf("Part 1: %d\n", counter)
}

func part2(data []int) {
	counter := 0

	p := 0

	for p >= 0 && p < len(data) {
		v := data[p]

		if v >= 3 {
			data[p]--
		} else {
			data[p]++
		}

		p += v
		counter++
	}

	fmt.Printf("Part 2: %d\n", counter)
}
